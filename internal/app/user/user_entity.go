package user

import "time"

type UserEntity struct {
	UserId       string    `gorm:"column:UserId;primaryKey"`
	FullName     string    `gorm:"column:FullName;size:255"`
	Email        string    `gorm:"column:Email;size:255"`
	Picture      string    `gorm:"column:Picture;size:255"`
	Role         string    `gorm:"column:Role;size:10;comment:ADMIN/USER;default:USER"`
	ActiveStatus string    `gorm:"column:ActiveStatus;size:20;default:AVAILABLE"`
	CreatedAt    time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`
}

func (UserEntity) TableName() string {
	return "GOST_Users"
}
