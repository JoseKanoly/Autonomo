from flask import Blueprint
from controllers.controllers import UserController, CompanyController

def register_routes(app):
    user_controller = UserController()
    company_controller = CompanyController()

    # Rutas para usuarios
    app.add_url_rule('/api/users', view_func=user_controller.get_all, methods=['GET'], endpoint='user_get_all')
    app.add_url_rule('/api/users/<int:id>', view_func=user_controller.get_by_id, methods=['GET'], endpoint='user_get_by_id')
    app.add_url_rule('/api/users', view_func=user_controller.create, methods=['POST'], endpoint='user_create')
    app.add_url_rule('/api/users/<int:id>', view_func=user_controller.update, methods=['PUT'], endpoint='user_update')
    app.add_url_rule('/api/users/<int:id>', view_func=user_controller.delete, methods=['DELETE'], endpoint='user_delete')

    # Rutas para empresas
    app.add_url_rule('/api/companies', view_func=company_controller.get_all, methods=['GET'], endpoint='company_get_all')
    app.add_url_rule('/api/companies/<int:id>', view_func=company_controller.get_by_id, methods=['GET'], endpoint='company_get_by_id')
    app.add_url_rule('/api/companies', view_func=company_controller.create, methods=['POST'], endpoint='company_create')
    app.add_url_rule('/api/companies/<int:id>', view_func=company_controller.update, methods=['PUT'], endpoint='company_update')
    app.add_url_rule('/api/companies/<int:id>', view_func=company_controller.delete, methods=['DELETE'], endpoint='company_delete')
