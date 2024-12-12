from flask import Flask
from flask_graphql import GraphQLView
from flask_cors import CORS
from dotenv import load_dotenv
import os
from database.database import db
from routes.routes import register_routes
from graphql_schema.graphql_schema import schema

load_dotenv()

app = Flask(__name__)
CORS(app)

app.config['SQLALCHEMY_DATABASE_URI'] = os.getenv("DATABASE_URL")
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
db.init_app(app)

with app.app_context():
    db.create_all()

# Register routes
register_routes(app)

# GraphQL endpoint
app.add_url_rule('/graphql', view_func=GraphQLView.as_view('graphql', schema=schema, graphiql=True))

if __name__ == '__main__':
    app.run(port=8082)  # Cambiado a puerto 8082
