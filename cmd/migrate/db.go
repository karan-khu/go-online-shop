package main

import (
	"fmt"
	"log"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
)

func main() {
	env := config.NewEnv()

	db := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	tx := db.Connect().Begin()

	tx.AutoMigrate(&models.UserRecord{})
	tx.AutoMigrate(&models.UserBalanceRecord{})
	tx.AutoMigrate(&models.ItemRecord{})
	tx.AutoMigrate(&models.InventoryRecord{})
	tx.AutoMigrate(&models.PurchaseHistoryRecord{})

	tx.Migrator().CreateConstraint(&models.UserRecord{}, "Balances")
	tx.Migrator().CreateConstraint(&models.UserRecord{}, "Items")
	tx.Migrator().CreateConstraint(&models.UserRecord{}, "PurchaseHistories")
	tx.Migrator().CreateConstraint(&models.UserRecord{}, "Inventories")
	tx.Migrator().CreateConstraint(&models.ItemRecord{}, "Admin")
	tx.Migrator().CreateConstraint(&models.InventoryRecord{}, "User")
	tx.Migrator().CreateConstraint(&models.InventoryRecord{}, "Item")
	tx.Migrator().CreateConstraint(&models.PurchaseHistoryRecord{}, "Buyer")
	tx.Migrator().CreateConstraint(&models.PurchaseHistoryRecord{}, "Item")
	tx.Migrator().CreateConstraint(&models.UserBalanceRecord{}, "User")

	if err := tx.Commit().Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Fatal("Error rolling back database", err)
		}
		log.Fatal("Error migrating database", err)
	}

	fmt.Println("Database migrated successfully")
}
