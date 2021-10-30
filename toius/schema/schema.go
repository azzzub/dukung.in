package schema

import (
	"toius/resolver"
	"toius/types"

	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"CreateRole": &graphql.Field{
			Type:        types.ObjectRole,
			Description: "Create new role",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.CreateRole,
		},
		"DeleteRole": &graphql.Field{
			Type:        types.ObjectCode,
			Description: "Delete a role",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.DeleteRole,
		},
		"UpdateRole": &graphql.Field{
			Type:        types.ObjectCode,
			Description: "Update a role",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.UpdateRole,
		},
		"CreateAccount": &graphql.Field{
			Type:        types.ObjectAccount,
			Description: "Create new account",
			Args: graphql.FieldConfigArgument{
				"value": &graphql.ArgumentConfig{
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name: "AccountDetail",
							Fields: graphql.InputObjectConfigFieldMap{
								"email": &graphql.InputObjectFieldConfig{
									Type: graphql.NewNonNull(graphql.String),
								},
								"username": &graphql.InputObjectFieldConfig{
									Type: graphql.String,
								},
								"password": &graphql.InputObjectFieldConfig{
									Type: graphql.String,
								},
								"role_id": &graphql.InputObjectFieldConfig{
									Type: graphql.Int,
								},
							},
						},
					),
				},
			},
			Resolve: resolver.CreateAccount,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"GetAllRoles": &graphql.Field{
			Type:        graphql.NewList(types.ObjectRole),
			Description: "Get all roles",
			Resolve:     resolver.GetAllRoles,
		},
		"GetAllAccounts": &graphql.Field{
			Type:        graphql.NewList(types.ObjectAccount),
			Description: "Get all accounts",
			Resolve:     resolver.GetAllAccounts,
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
