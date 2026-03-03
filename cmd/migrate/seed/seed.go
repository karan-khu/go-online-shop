package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/pkg/database"
)

func main() {

	env := config.NewEnv()

	db := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	tx := db.Begin()

	createItemSeed(tx)

	if err := tx.Commit().Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Fatal("Error rolling back database", err)
		}
		log.Fatal("Error seeding database", err)
	}
}

func createItemSeed(tx *gorm.DB) {
	// ข้อมูลจาก Arena of Valor Wiki: https://arena-of-valor.fandom.com/th/wiki/ไอเทม
	items := []item.ItemEntity{
		{
			Name:        "Short Sword",
			Description: "+20 พลังโจมตี",
			Price:       250,
			Picture:     "uploads/sword.png",
		},
		{
			Name:        "Dagger",
			Description: "+10% ความเร็วโจมตี",
			Price:       290,
			Picture:     "uploads/dragger.png",
		},
		{
			Name:        "Astral Spear",
			Description: "+50 พลังโจมตี, เจาะเกราะ +60",
			Price:       830,
			Picture:     "uploads/astral-spear.png",
		},
		{
			Name:        "Shuriken",
			Description: "+20% ความเร็วโจมตี, Heel: การโจมตีปกติสร้างความเสียหายกายภาพเพิ่มขึ้น",
			Price:       750,
			Picture:     "uploads/shuriken.png",
		},
		{
			Name:        "The Beast",
			Description: "+100 พลังโจมตี, ดูดเลือด +25%",
			Price:       1740,
			Picture:     "uploads/the-beast.png",
		},
		{
			Name:        "Muramasa",
			Description: "+75 พลังโจมตี, +10% ลดคูลดาวน์, เจาะเกราะ +45%",
			Price:       2020,
			Picture:     "uploads/muramasa.png",
		},
		{
			Name:        "Broken Spears",
			Description: "+110 พลังโจมตี, Destroyer: เพิ่มเจาะเกราะ (110-250)",
			Price:       2020,
			Picture:     "uploads/broken-spear.png",
		},
		{
			Name:        "The Morning Star",
			Description: "+50 พลังโจมตี, +30% ความเร็วโจมตี, +10% อัตราคริติคอล, Dawning star + Stab",
			Price:       2980,
			Picture:     "uploads/morning-star.png",
		},
	}

	tx.Create(&items)
}
