package config

import (
	"log"

	"github.com/karan-khu/go-online-shop/pkg/database"
)

type Config struct {
	Env          *Env
	database     *DbConn
	Oauth2Config *Oauth2Config
}

func NewConfig() *Config {
	env := NewEnv()
	db := NewDbConn(env)
	googleOAuth2Config := NewGoogleOAuth2Config(env)

	return &Config{
		Env:          env,
		database:     db,
		Oauth2Config: googleOAuth2Config,
	}
}

func (c *Config) GetDb(name string) database.Database {
	if db, ok := c.database.dbs[name]; ok {
		return db
	}
	log.Fatal("Error", "database", "database with name "+name+" not found")
	return nil
}
