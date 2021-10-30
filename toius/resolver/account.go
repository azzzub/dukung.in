package resolver

import (
	"toius/config"
	"toius/tools"
	"toius/types"

	"github.com/graphql-go/graphql"
)

func CreateAccount(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	value := p.Args["value"].(map[string]interface{})

	username := value["value"]
	email := value["email"].(string)
	password := value["password"]
	roleId := value["role_id"]

	var hashedPassword string

	newAccount := &types.Account{}

	newAccount.Email = email

	if password != nil {
		hashedPassword, err = tools.HashPassword(password.(string))
		if err != nil {
			return nil, err
		}

		newAccount.Password = &hashedPassword
	}
	if username == nil {
		newAccount.Username = nil
	}

	if roleId == nil {
		// default role id
		id := 1
		newAccount.RoleId = &id
	}

	db.Create(newAccount)

	return newAccount, nil
}

func GetAllAccounts(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	var accounts []types.Account

	err = db.Preload("Role").Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
