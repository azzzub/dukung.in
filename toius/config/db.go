package config

import (
	"fmt"
	"toius/tools"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = tools.GetEnv("DB_HOST")
	port     = tools.GetEnv("DB_PORT")
	user     = tools.GetEnv("DB_USER")
	password = tools.GetEnv("DB_PASS")
	dbname   = tools.GetEnv("DB_NAME")
)

func DbConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
