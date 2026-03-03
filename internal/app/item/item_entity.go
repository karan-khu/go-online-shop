package item

import (
	"strings"
	"time"
)

type ItemEntity struct {
	ItemId       int       `gorm:"column:ItemId;primaryKey;autoIncrement"`
	AdminId      int       `gorm:"column:AdminId"`
	Name         string    `gorm:"column:Name;size:255"`
	Description  string    `gorm:"column:Description;size:max"`
	Picture      string    `gorm:"column:Picture"`
	Price        int       `gorm:"column:Price"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`
}

func (ItemEntity) TableName() string {
	return "GOST_Items"
}

func (e *ItemEntity) ToModel(host string) *ItemModel {
	picture := e.Picture
	if host != "" && picture != "" {
		baseURL := strings.TrimSuffix(host, "/")
		picture = baseURL + "/" + strings.TrimPrefix(picture, "/")
	}
	return &ItemModel{
		ItemId:      e.ItemId,
		AdminId:     e.AdminId,
		Name:        e.Name,
		Description: e.Description,
		Picture:     picture,
		Price:       e.Price,
	}
}
