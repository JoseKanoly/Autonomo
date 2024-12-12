package graphql

import (
	"go_project/internal/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

// Define root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"listaPrecios": resolvers.GetListaPreciosQuery(),
		"residuos":     resolvers.GetResiduosQuery(),
	},
})

// Define root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createListaPrecio": resolvers.CreateListaPrecioMutation(),
		"updateListaPrecio": resolvers.UpdateListaPrecioMutation(),
		"deleteListaPrecio": resolvers.DeleteListaPrecioMutation(),
		"createResiduo":     resolvers.CreateResiduoMutation(),
		"updateResiduo":     resolvers.UpdateResiduoMutation(),
		"deleteResiduo":     resolvers.DeleteResiduoMutation(),
	},
})

// Create and export the schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
