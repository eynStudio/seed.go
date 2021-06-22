package schema

import (
	"github.com/eynstudio/seed.go/domains/sys"
	"github.com/graphql-go/graphql"
)

var accounts []sys.SysAccount

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: graphql.Fields{},
})

// root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"accounts": &graphql.Field{
			Type:        graphql.NewList(sys.SysAccountType),
			Description: "List of system accounts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return accounts, nil
			},
		},
	},
})

var RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
