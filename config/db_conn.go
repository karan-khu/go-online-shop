package config

import (
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/pkg/database"
)

type DbConn struct {
	dbs map[string]*gorm.DB
}

func NewDbConn(env *Env) *DbConn {
	mainDb := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	return &DbConn{dbs: map[string]*gorm.DB{
		"main": mainDb.DB,
	}}
}
