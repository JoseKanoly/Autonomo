package resolvers

import (
	"go_project/internal/models"
	"go_project/internal/services"
	"time"

	"github.com/graphql-go/graphql"
)

var residuoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Residuo",
	Fields: graphql.Fields{
		"id":                &graphql.Field{Type: graphql.Int},
		"tipo":              &graphql.Field{Type: graphql.String},
		"peso":              &graphql.Field{Type: graphql.Float},
		"fecha_recoleccion": &graphql.Field{Type: graphql.DateTime},
		"empresa_id":        &graphql.Field{Type: graphql.Int},
	},
})

func GetResiduosQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(residuoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("residuosService").(*services.ResiduosService)
			return service.GetAll()
		},
	}
}

func CreateResiduoMutation() *graphql.Field {
	return &graphql.Field{
		Type: residuoType,
		Args: graphql.FieldConfigArgument{
			"tipo":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"peso":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
			"fecha_recoleccion": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.DateTime)},
			"empresa_id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("residuosService").(*services.ResiduosService)
			residuo := &models.Residuo{
				Tipo:             p.Args["tipo"].(string),
				Peso:             p.Args["peso"].(float64),
				FechaRecoleccion: p.Args["fecha_recoleccion"].(time.Time),
				EmpresaID:        p.Args["empresa_id"].(int),
			}
			err := service.Create(residuo)
			return residuo, err
		},
	}
}

func UpdateResiduoMutation() *graphql.Field {
	return &graphql.Field{
		Type: residuoType,
		Args: graphql.FieldConfigArgument{
			"id":                &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"tipo":              &graphql.ArgumentConfig{Type: graphql.String},
			"peso":              &graphql.ArgumentConfig{Type: graphql.Float},
			"fecha_recoleccion": &graphql.ArgumentConfig{Type: graphql.DateTime},
			"empresa_id":        &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("residuosService").(*services.ResiduosService)
			id := p.Args["id"].(int)
			residuo, err := service.GetByID(id)
			if err != nil {
				return nil, err
			}
			if tipo, ok := p.Args["tipo"].(string); ok {
				residuo.Tipo = tipo
			}
			if peso, ok := p.Args["peso"].(float64); ok {
				residuo.Peso = peso
			}
			if fechaRecoleccion, ok := p.Args["fecha_recoleccion"].(time.Time); ok {
				residuo.FechaRecoleccion = fechaRecoleccion
			}
			if empresaID, ok := p.Args["empresa_id"].(int); ok {
				residuo.EmpresaID = empresaID
			}
			err = service.Update(residuo)
			return residuo, err
		},
	}
}

func DeleteResiduoMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("residuosService").(*services.ResiduosService)
			id := p.Args["id"].(int)
			err := service.Delete(id)
			return err == nil, err
		},
	}
}
