package config

import (
	"github.com/karan-khu/go-online-shop/internal/infra/database"
)

type DbConn struct {
	dbs map[string]database.Database
}

func NewDbConn(env *Env) *DbConn {
	mainDb := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	return &DbConn{dbs: map[string]database.Database{
		"main": mainDb,
	}}
}
