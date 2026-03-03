package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/app/balance"
	"github.com/karan-khu/go-online-shop/internal/app/inventory"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/internal/app/purchase"
	"github.com/karan-khu/go-online-shop/internal/app/user"
	"github.com/karan-khu/go-online-shop/pkg/database"
)

func main() {
	env := config.NewEnv()

	db := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	tx := db.Begin()

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
	tx.AutoMigrate(&balance.UserBalanceEntity{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.AutoMigrate(&inventory.InventoryEntity{})
}

func itemMigration(tx *gorm.DB) {
	tx.AutoMigrate(&item.ItemEntity{})
}

func purchaseMigration(tx *gorm.DB) {
	tx.AutoMigrate(&purchase.PurchaseHistoryEntity{})
}

func userMigration(tx *gorm.DB) {
	tx.AutoMigrate(&user.UserEntity{})
}
