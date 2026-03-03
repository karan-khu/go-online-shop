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
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
}

type ResponseItemList struct {
	Items      []*ItemModel `json:"items"`
	Pagination *Pagination  `json:"pagination"`
}

type Pagination struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
