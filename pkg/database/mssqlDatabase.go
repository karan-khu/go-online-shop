package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database interface {
	Connect() *gorm.DB
}

type sqlServerDatabase struct {
	*gorm.DB
}

var (
	instance *sqlServerDatabase
	once     sync.Once
)

func NewSqlServerDatabase(host, user, pass, db string) Database {
	once.Do(func() {
		connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", host, user, pass, db)

		conn, err := gorm.Open(sqlserver.Open(connStr))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Connected to SQL Server database: %s", db)

		instance = &sqlServerDatabase{conn}
	})

	return instance
}

func (d *sqlServerDatabase) Connect() *gorm.DB {
	return instance.DB
}
