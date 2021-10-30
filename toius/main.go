package main

import (
	"net/http"
	"toius/config"
	"toius/schema"
	"toius/types"

	"github.com/graphql-go/handler"
)

func main() {
	db, err := config.DbConnect()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.Role{})
	db.AutoMigrate(&types.Account{})

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/ui", h)
	http.ListenAndServe(":9090", nil)
}
