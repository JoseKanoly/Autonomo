package models

import "time"

type Residuo struct {
	ID               int       `json:"id"`
	Tipo             string    `json:"tipo" validate:"required"`
	Peso             float64   `json:"peso" validate:"required,gt=0"`
	FechaRecoleccion time.Time `json:"fecha_recoleccion" validate:"required"`
	EmpresaID        int       `json:"empresa_id" validate:"required"`
}
