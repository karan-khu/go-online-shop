package itemshop

type ItemShopEntity struct {
	ItemId      int    `json:"item_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       int    `json:"price"`
}
