package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Account struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Id        int64          `json:"id" gorm:"primaryKey"`
	Username  *string        `json:"username" gorm:"unique"`
	Email     string         `json:"email" gorm:"unique,not null"`
	Password  *string        `json:"password"`
	RoleId    *int           `json:"role_id"`
	Role      Role           `json:"role" gorm:"foreignKey:role_id"`
}

var ObjectAccount = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"role_id": &graphql.Field{
			Type: graphql.Int,
		},
		"Role": &graphql.Field{
			Type: ObjectRole,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				src, _ := p.Source.(Account)
				return src.Role, nil
			},
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
