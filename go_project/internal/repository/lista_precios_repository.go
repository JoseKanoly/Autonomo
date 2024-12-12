package repository

import (
	"database/sql"

	"go_project/internal/models"
)

type ListaPreciosRepository struct {
	db *sql.DB
}

func NewListaPreciosRepository(db *sql.DB) *ListaPreciosRepository {
	return &ListaPreciosRepository{db: db}
}

func (r *ListaPreciosRepository) GetAll() ([]models.ListaPrecios, error) {
	rows, err := r.db.Query("SELECT id, tipo_residuo, precio_por_kg FROM lista_precios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listaPrecios []models.ListaPrecios
	for rows.Next() {
		var lp models.ListaPrecios
		if err := rows.Scan(&lp.ID, &lp.TipoResiduo, &lp.PrecioPorKg); err != nil {
			return nil, err
		}
		listaPrecios = append(listaPrecios, lp)
	}

	return listaPrecios, nil
}

func (r *ListaPreciosRepository) GetByID(id int) (*models.ListaPrecios, error) {
	var lp models.ListaPrecios
	err := r.db.QueryRow("SELECT id, tipo_residuo, precio_por_kg FROM lista_precios WHERE id = $1", id).
		Scan(&lp.ID, &lp.TipoResiduo, &lp.PrecioPorKg)
	if err != nil {
		return nil, err
	}
	return &lp, nil
}

func (r *ListaPreciosRepository) Create(lp *models.ListaPrecios) error {
	err := r.db.QueryRow("INSERT INTO lista_precios (tipo_residuo, precio_por_kg) VALUES ($1, $2) RETURNING id",
		lp.TipoResiduo, lp.PrecioPorKg).Scan(&lp.ID)
	return err
}

func (r *ListaPreciosRepository) Update(lp *models.ListaPrecios) error {
	_, err := r.db.Exec("UPDATE lista_precios SET tipo_residuo = $1, precio_por_kg = $2 WHERE id = $3",
		lp.TipoResiduo, lp.PrecioPorKg, lp.ID)
	return err
}

func (r *ListaPreciosRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM lista_precios WHERE id = $1", id)
	return err
}
