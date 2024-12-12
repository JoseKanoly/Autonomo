from flask import request, jsonify
from models.models import User, Company
from database.database import db

class UserController:
    def get_all(self):
        users = User.query.all()
        return jsonify([user.to_dict() for user in users])

    def get_by_id(self, id):
        user = User.query.get(id)
        return jsonify(user.to_dict()) if user else ('', 404)

    def create(self):
        data = request.json
        # Validación de entradas
        if not data.get('username') or not data.get('password'):
            return jsonify({'error': 'Username and password are required'}), 400
        new_user = User(username=data['username'], password=data['password'])
        db.session.add(new_user)
        db.session.commit()
        return jsonify(new_user.to_dict()), 201

    def update(self, id):
        user = User.query.get(id)
        if not user:
            return ('', 404)
        data = request.json
        # Validación de entradas
        if not data.get('username') or not data.get('password'):
            return jsonify({'error': 'Username and password are required'}), 400
        user.username = data['username']
        user.password = data['password']
        db.session.commit()
        return jsonify(user.to_dict())

    def delete(self, id):
        user = User.query.get(id)
        if not user:
            return ('', 404)
        db.session.delete(user)
        db.session.commit()
        return ('', 204)

class CompanyController:
    def get_all(self):
        companies = Company.query.all()
        return jsonify([company.to_dict() for company in companies])

    def get_by_id(self, id):
        company = Company.query.get(id)
        return jsonify(company.to_dict()) if company else ('', 404)

    def create(self):
        data = request.json
        # Validación de entradas
        if not data.get('name'):
            return jsonify({'error': 'Company name is required'}), 400
        new_company = Company(name=data['name'], address=data.get('address'))
        db.session.add(new_company)
        db.session.commit()
        return jsonify(new_company.to_dict()), 201

    def update(self, id):
        company = Company.query.get(id)
        if not company:
            return ('', 404)
        data = request.json
        # Validación de entradas
        if not data.get('name'):
            return jsonify({'error': 'Company name is required'}), 400
        company.name = data['name']
        company.address = data.get('address')
        db.session.commit()
        return jsonify(company.to_dict())

    def delete(self, id):
        company = Company.query.get(id)
        if not company:
            return ('', 404)
        db.session.delete(company)
        db.session.commit()
        return ('', 204)
