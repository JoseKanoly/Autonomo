package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {
	createListaPreciosTable := `
	CREATE TABLE IF NOT EXISTS lista_precios (
		id SERIAL PRIMARY KEY,
		tipo_residuo TEXT NOT NULL,
		precio_por_kg DECIMAL(10, 2) NOT NULL
	);`

	createResiduosTable := `
	CREATE TABLE IF NOT EXISTS residuos (
		id SERIAL PRIMARY KEY,
		tipo TEXT NOT NULL,
		peso DECIMAL(10, 2) NOT NULL,
		fecha_recoleccion TIMESTAMP NOT NULL,
		empresa_id INTEGER NOT NULL
	);`

	_, err := db.Exec(createListaPreciosTable)
	if err != nil {
		log.Printf("Error creating lista_precios table: %v", err)
		return err
	}

	_, err = db.Exec(createResiduosTable)
	if err != nil {
		log.Printf("Error creating residuos table: %v", err)
		return err
	}

	log.Println("Tables created successfully")
	return nil
}

