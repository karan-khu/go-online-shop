package purchase

import (
	"time"
)

type PurchaseHistoryEntity struct {
	PurchaseId      int
	BuyerId         string
	ItemId          int
	ItemName        string
	ItemDescription string
	ItemPrice       int
	Quantity        int
	Type            string
	ActiveStatus    string
	CreatedAt       time.Time
}
