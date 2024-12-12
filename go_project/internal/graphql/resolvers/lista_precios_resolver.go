package resolvers

import (
	"github.com/graphql-go/graphql"

	"go_project/internal/models"
	"go_project/internal/services"
)

var listaPreciosType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListaPrecios",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"tipo_residuo":  &graphql.Field{Type: graphql.String},
		"precio_por_kg": &graphql.Field{Type: graphql.Float},
	},
})

func GetListaPreciosQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(listaPreciosType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("listaPreciosService").(*services.ListaPreciosService)
			return service.GetAll()
		},
	}
}

func CreateListaPrecioMutation() *graphql.Field {
	return &graphql.Field{
		Type: listaPreciosType,
		Args: graphql.FieldConfigArgument{
			"tipo_residuo":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"precio_por_kg": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("listaPreciosService").(*services.ListaPreciosService)
			listaPrecio := &models.ListaPrecios{
				TipoResiduo: p.Args["tipo_residuo"].(string),
				PrecioPorKg: p.Args["precio_por_kg"].(float64),
			}
			err := service.Create(listaPrecio)
			return listaPrecio, err
		},
	}
}

func UpdateListaPrecioMutation() *graphql.Field {
	return &graphql.Field{
		Type: listaPreciosType,
		Args: graphql.FieldConfigArgument{
			"id":            &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"tipo_residuo":  &graphql.ArgumentConfig{Type: graphql.String},
			"precio_por_kg": &graphql.ArgumentConfig{Type: graphql.Float},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("listaPreciosService").(*services.ListaPreciosService)
			id := p.Args["id"].(int)
			listaPrecio, err := service.GetByID(id)
			if err != nil {
				return nil, err
			}
			if tipoResiduo, ok := p.Args["tipo_residuo"].(string); ok {
				listaPrecio.TipoResiduo = tipoResiduo
			}
			if precioPorKg, ok := p.Args["precio_por_kg"].(float64); ok {
				listaPrecio.PrecioPorKg = precioPorKg
			}
			err = service.Update(listaPrecio)
			return listaPrecio, err
		},
	}
}

func DeleteListaPrecioMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			service := p.Context.Value("listaPreciosService").(*services.ListaPreciosService)
			id := p.Args["id"].(int)
			err := service.Delete(id)
			return err == nil, err
		},
	}
}
