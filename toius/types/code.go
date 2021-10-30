package types

import "github.com/graphql-go/graphql"

type StaticResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ObjectCode = graphql.NewObject(graphql.ObjectConfig{
	Name: "Code",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Type: graphql.Int,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
