package models

import "time"

type UserBalanceRecord struct {
	BalanceId int       `gorm:"column:BalanceId;primaryKey;autoIncrement"`
	UserId    string    `gorm:"column:UserId;size:255"` // ต้อง match กับ UserRecord.UserId เพื่อสร้าง FK ได้
	Amount    int       `gorm:"column:Amount"`
	Status    string    `gorm:"column:Status;size:10;comment:BUY/TOPUP"`
	CreatedAt time.Time `gorm:"column:CreatedAt;autoCreateTime"`

	User UserRecord `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (UserBalanceRecord) TableName() string {
	return "GOST_UserBalances"
}
