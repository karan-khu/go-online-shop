package purchase

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type PurchaseHttpHandler interface {
	Listing(c *echo.Context) error
}

type PurchaseUsecase interface {
	Listing(userId string) ([]*PurchaseHistoryEntity, error)
}

type PurchaseRepository interface {
	Listing(userId string) ([]*PurchaseHistoryEntity, error)
	Create(tx *gorm.DB, purchase *PurchaseHistoryEntity) (*PurchaseHistoryEntity, error)
}
