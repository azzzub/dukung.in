package resolver

import (
	"errors"
	"fmt"
	"toius/config"
	"toius/types"

	"github.com/graphql-go/graphql"
)

func CreateRole(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	newRole := &types.Role{
		Id:   id,
		Name: name,
	}

	err = db.Create(newRole).Error
	if err != nil {
		return nil, err
	}

	return newRole, nil
}

func DeleteRole(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	id := p.Args["id"].(int)

	err = db.Delete(&types.Role{}, id).Error
	if err != nil {
		return nil, err
	}

	return &types.StaticResp{
		Code:    200,
		Message: fmt.Sprintf("successfully delete role with id: %d", id),
	}, nil
}

func UpdateRole(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	if name == "" {
		return nil, errors.New("role name must be filled")
	}

	err = db.Model(&types.Role{}).Where("id", id).Update("name", name).Error
	if err != nil {
		return nil, err
	}

	return &types.StaticResp{
		Code:    200,
		Message: fmt.Sprintf("successfully update role with id: %d", id),
	}, nil
}

func GetAllRoles(p graphql.ResolveParams) (interface{}, error) {
	db, err := config.DbConnect()
	if err != nil {
		return nil, err
	}

	var roles []types.Role
	err = db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
