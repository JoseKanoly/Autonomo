import graphene
from models.models import User, Company
from database.database import db

# Definición de los tipos para los datos (Usuarios y Empresas)
class UserType(graphene.ObjectType):
    id = graphene.Int()
    username = graphene.String()

class CompanyType(graphene.ObjectType):
    id = graphene.Int()
    name = graphene.String()
    address = graphene.String()

# Consultas (Query - GET)
class Query(graphene.ObjectType):
    all_users = graphene.List(UserType)
    all_companies = graphene.List(CompanyType)

    def resolve_all_users(self, info):
        return User.query.all()

    def resolve_all_companies(self, info):
        return Company.query.all()

# Mutaciones (Mutation - POST, PUT, DELETE)
class CreateUser(graphene.Mutation):
    class Arguments:
        username = graphene.String()
        password = graphene.String()

    user = graphene.Field(UserType)

    def mutate(self, info, username, password):
        new_user = User(username=username, password=password)
        db.session.add(new_user)
        db.session.commit()
        return CreateUser(user=new_user)

class UpdateUser(graphene.Mutation):
    class Arguments:
        id = graphene.Int()
        username = graphene.String()
        password = graphene.String()

    user = graphene.Field(UserType)

    def mutate(self, info, id, username, password):
        user = User.query.get(id)
        if user:
            user.username = username
            user.password = password
            db.session.commit()
            return UpdateUser(user=user)
        return UpdateUser(user=None)

class DeleteUser(graphene.Mutation):
    class Arguments:
        id = graphene.Int()

    success = graphene.Boolean()

    def mutate(self, info, id):
        user = User.query.get(id)
        if user:
            db.session.delete(user)
            db.session.commit()
            return DeleteUser(success=True)
        return DeleteUser(success=False)

class CreateCompany(graphene.Mutation):
    class Arguments:
        name = graphene.String()
        address = graphene.String()

    company = graphene.Field(CompanyType)

    def mutate(self, info, name, address):
        new_company = Company(name=name, address=address)
        db.session.add(new_company)
        db.session.commit()
        return CreateCompany(company=new_company)

class UpdateCompany(graphene.Mutation):
    class Arguments:
        id = graphene.Int()
        name = graphene.String()
        address = graphene.String()

    company = graphene.Field(CompanyType)

    def mutate(self, info, id, name, address):
        company = Company.query.get(id)
        if company:
            company.name = name
            company.address = address
            db.session.commit()
            return UpdateCompany(company=company)
        return UpdateCompany(company=None)

class DeleteCompany(graphene.Mutation):
    class Arguments:
        id = graphene.Int()

    success = graphene.Boolean()

    def mutate(self, info, id):
        company = Company.query.get(id)
        if company:
            db.session.delete(company)
            db.session.commit()
            return DeleteCompany(success=True)
        return DeleteCompany(success=False)

# Definición de la clase de Mutaciones que incluye todas las mutaciones para los usuarios y las empresas
class Mutation(graphene.ObjectType):
    create_user = CreateUser.Field()
    update_user = UpdateUser.Field()
    delete_user = DeleteUser.Field()
    create_company = CreateCompany.Field()
    update_company = UpdateCompany.Field()
    delete_company = DeleteCompany.Field()

# Esquema completo que incluye tanto las consultas (Query) como las mutaciones (Mutation)
schema = graphene.Schema(query=Query, mutation=Mutation)
