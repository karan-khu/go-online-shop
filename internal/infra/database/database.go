package database

import "gorm.io/gorm"

type Database interface {
	Connect() *gorm.DB
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}
