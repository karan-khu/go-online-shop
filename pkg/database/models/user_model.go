package models

import "time"

type UserRecord struct {
	UserId       string    `gorm:"column:UserId;primaryKey"`
	FullName     string    `gorm:"column:FullName;size:255"`
	Email        string    `gorm:"column:Email;size:255"`
	Picture      string    `gorm:"column:Picture;size:255"`
	Role         string    `gorm:"column:Role;size:10;comment:ADMIN/USER;default:USER"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`

	Balances          []UserBalanceRecord     `gorm:"foreignKey:UserId;references:UserId"`
	Items             []ItemRecord            `gorm:"foreignKey:ItemId;references:ItemId"`
	PurchaseHistories []PurchaseHistoryRecord `gorm:"foreignKey:BuyerId;references:BuyerId"`
	Inventories       []InventoryRecord       `gorm:"foreignKey:UserId;references:UserId"`
}

func (UserRecord) TableName() string {
	return "GOST_Users"
}
