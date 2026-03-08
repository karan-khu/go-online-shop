package inventory

import "time"

type InventoryEntity struct {
	InventoryId  int       `gorm:"column:InventoryId;primaryKey;autoIncrement"`
	UserId       string    `gorm:"column:UserId"`
	ItemId       int       `gorm:"column:ItemId"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
}

func (InventoryEntity) TableName() string {
	return "GOST_Inventories"
}

type QueryInventoryItemEntity struct {
	InventoryEntity
	ItemName        string `gorm:"column:Name"`
	ItemDescription string `gorm:"column:Description"`
	ItemPicture     string `gorm:"column:Picture"`
}
