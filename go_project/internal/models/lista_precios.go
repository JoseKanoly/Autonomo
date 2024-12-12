package models

type ListaPrecios struct {
	ID          int     `json:"id"`
	TipoResiduo string  `json:"tipo_residuo" validate:"required"`
	PrecioPorKg float64 `json:"precio_por_kg" validate:"required,gt=0"`
}
