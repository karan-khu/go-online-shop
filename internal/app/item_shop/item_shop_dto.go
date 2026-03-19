package itemshop

type RequestItemFilter struct {
	SearchText string `query:"search_text"`
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
}

type ResponseItemList struct {
	Items      []*ItemShopEntity `json:"items"`
	Pagination *Pagination       `json:"pagination"`
}

type Pagination struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type RequestItemBuying struct {
	UserId   string
	ItemId   int `json:"item_id" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type RequestItemSelling struct {
	UserId   string
	ItemId   int `json:"item_id" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}
