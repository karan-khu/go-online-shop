package models

import "time"

type ItemRecord struct {
	ItemId       int       `gorm:"column:ItemId;primaryKey;autoIncrement"`
	AdminId      string    `gorm:"column:AdminId;size:255"` // FK ไป UserRecord.UserId
	Name         string    `gorm:"column:Name;size:255"`
	Description  string    `gorm:"column:Description;size:max"`
	Picture      string    `gorm:"column:Picture"`
	Price        int       `gorm:"column:Price"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`

	Admin UserRecord `gorm:"foreignKey:AdminId;references:UserId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (ItemRecord) TableName() string {
	return "GOST_Items"
}
