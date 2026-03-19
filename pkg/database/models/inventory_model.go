package models

import "time"

type InventoryRecord struct {
	InventoryId  int       `gorm:"column:InventoryId;primaryKey;autoIncrement"`
	UserId       string    `gorm:"column:UserId;size:255"`
	ItemId       int       `gorm:"column:ItemId"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
}

func (InventoryRecord) TableName() string {
	return "GOST_Inventories"
}
