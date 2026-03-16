package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
)

func main() {
	env := config.NewEnv()

	db := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	tx := db.Connect().Begin()

	balanceMigration(tx)
	inventoryMigration(tx)
	itemMigration(tx)
	purchaseMigration(tx)
	userMigration(tx)

	if err := tx.Commit().Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Fatal("Error rolling back database", err)
		}
		log.Fatal("Error migrating database", err)
	}

	fmt.Println("Database migrated successfully")
}

func balanceMigration(tx *gorm.DB) {
	tx.AutoMigrate(&models.UserBalanceRecord{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.AutoMigrate(&models.InventoryRecord{})
}

func itemMigration(tx *gorm.DB) {
	tx.AutoMigrate(&models.ItemRecord{})
}

func purchaseMigration(tx *gorm.DB) {
	tx.AutoMigrate(&models.PurchaseHistoryRecord{})
}

func userMigration(tx *gorm.DB) {
	tx.AutoMigrate(&models.UserRecord{})
}
