package inventory

import "time"

type InventoryListing struct {
	ItemId          int       `json:"item_id"`
	ItemName        string    `json:"item_name"`
	ItemDescription string    `json:"item_description"`
	ItemPicture     string    `json:"item_picture"`
	CreatedAt       time.Time `json:"created_at"`
	Quantity        int       `json:"quantity"`
}
