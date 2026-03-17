package models

import "time"

type UserRecord struct {
	UserId       string    `gorm:"column:UserId;primaryKey;size:255"` // ต้องระบุ size เพื่อไม่ให้เป็น nvarchar(MAX) - SQL Server ไม่รองรับ FK บน MAX
	FullName     string    `gorm:"column:FullName;size:255"`
	Email        string    `gorm:"column:Email;size:255"`
	Picture      string    `gorm:"column:Picture;size:255"`
	Role         string    `gorm:"column:Role;size:10;comment:ADMIN/USER;default:USER"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`

	// Has Many - ระบุ foreignKey (คอลัมน์ใน child) และ references (คอลัมน์ใน parent)
	Balances          []UserBalanceRecord     `gorm:"foreignKey:UserId;references:UserId"`
	Items             []ItemRecord            `gorm:"foreignKey:AdminId;references:UserId"`
	PurchaseHistories []PurchaseHistoryRecord `gorm:"foreignKey:BuyerId;references:UserId"`
	Inventories       []InventoryRecord       `gorm:"foreignKey:UserId;references:UserId"`
}

func (UserRecord) TableName() string {
	return "GOST_Users"
}
