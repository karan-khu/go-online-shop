package inventory

import "time"

type InventoryEntity struct {
	InventoryId     int
	UserId          string
	ItemId          int
	ActiveStatus    string
	CreatedAt       time.Time
	ItemName        string
	ItemDescription string
	ItemPicture     string
}
