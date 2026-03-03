package item

type ItemModel struct {
	ItemId      int    `json:"item_id"`
	AdminId     int    `json:"admin_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       int    `json:"price"`
}

type RequestItemFilter struct {
	SearchText string `query:"search_text"`
}
