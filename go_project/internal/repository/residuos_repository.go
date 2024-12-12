package repository

import (
	"database/sql"

	"go_project/internal/models"
)

type ResiduosRepository struct {
	db *sql.DB
}

func NewResiduosRepository(db *sql.DB) *ResiduosRepository {
	return &ResiduosRepository{db: db}
}

func (r *ResiduosRepository) GetAll() ([]models.Residuo, error) {
	rows, err := r.db.Query("SELECT id, tipo, peso, fecha_recoleccion, empresa_id FROM residuos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var residuos []models.Residuo
	for rows.Next() {
		var res models.Residuo
		if err := rows.Scan(&res.ID, &res.Tipo, &res.Peso, &res.FechaRecoleccion, &res.EmpresaID); err != nil {
			return nil, err
		}
		residuos = append(residuos, res)
	}

	return residuos, nil
}

func (r *ResiduosRepository) GetByID(id int) (*models.Residuo, error) {
	var res models.Residuo
	err := r.db.QueryRow("SELECT id, tipo, peso, fecha_recoleccion, empresa_id FROM residuos WHERE id = $1", id).
		Scan(&res.ID, &res.Tipo, &res.Peso, &res.FechaRecoleccion, &res.EmpresaID)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ResiduosRepository) Create(res *models.Residuo) error {
	err := r.db.QueryRow("INSERT INTO residuos (tipo, peso, fecha_recoleccion, empresa_id) VALUES ($1, $2, $3, $4) RETURNING id",
		res.Tipo, res.Peso, res.FechaRecoleccion, res.EmpresaID).Scan(&res.ID)
	return err
}

func (r *ResiduosRepository) Update(res *models.Residuo) error {
	_, err := r.db.Exec("UPDATE residuos SET tipo = $1, peso = $2, fecha_recoleccion = $3, empresa_id = $4 WHERE id = $5",
		res.Tipo, res.Peso, res.FechaRecoleccion, res.EmpresaID, res.ID)
	return err
}

func (r *ResiduosRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM residuos WHERE id = $1", id)
	return err
}
