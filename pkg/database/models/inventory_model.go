package models

import "time"

type InventoryRecord struct {
	InventoryId  int       `gorm:"column:InventoryId;primaryKey;autoIncrement"`
	UserId       string    `gorm:"column:UserId;size:255"` // FK ไป UserRecord.UserId
	ItemId       int       `gorm:"column:ItemId"`          // FK ไป ItemRecord.ItemId
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`

	User UserRecord `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Item ItemRecord `gorm:"foreignKey:ItemId;references:ItemId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (InventoryRecord) TableName() string {
	return "GOST_Inventories"
}
