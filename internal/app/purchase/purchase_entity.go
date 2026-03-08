package purchase

import "time"

type PurchaseHistoryEntity struct {
	PurchaseId      int       `gorm:"column:PurchaseId;primaryKey;autoIncrement"`
	BuyerId         string    `gorm:"column:BuyerId"`
	ItemId          int       `gorm:"column:ItemId"`
	ItemName        string    `gorm:"column:ItemName;size:255"`
	ItemDescription string    `gorm:"column:ItemDescription;size:max"`
	ItemPrice       int       `gorm:"column:ItemPrice"`
	Quantity        int       `gorm:"column:Quantity"`
	Type            string    `gorm:"column:Type;size:10;comment:BUY/SELL"`
	ActiveStatus    string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt       time.Time `gorm:"column:CreatedAt;autoCreateTime"`
}

func (PurchaseHistoryEntity) TableName() string {
	return "GOST_PurchaseHistories"
}
