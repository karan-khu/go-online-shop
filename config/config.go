package config

import (
	"log"

	"github.com/karan-khu/go-online-shop/pkg/database"
)

type Config struct {
	Env      *Env
	database *DbConn
}

func NewConfig() *Config {
	env := NewEnv()
	db := NewDbConn(env)

	return &Config{
		Env:      env,
		database: db,
	}
}

func (c *Config) GetDb(name string) database.Database {
	if db, ok := c.database.dbs[name]; ok {
		return db
	}
	log.Fatal("Error", "database", "database with name "+name+" not found")
	return nil
}
