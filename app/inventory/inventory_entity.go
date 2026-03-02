package inventory

import "time"

type InventoryEntity struct {
	InventoryId  int       `gorm:"column:InventoryId;primaryKey;autoIncrement"`
	UserId       int       `gorm:"column:UserId"`
	ItemId       int       `gorm:"column:ItemId"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
}

func (InventoryEntity) TableName() string {
	return "GOST_Inventories"
}
