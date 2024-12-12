import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    DATABASE_URL = os.getenv("DATABASE_URL")
    PORT = 8082  # Cambiado a 8082
    JWT_SECRET = os.getenv("JWT_SECRET")
