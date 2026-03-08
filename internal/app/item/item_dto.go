package item

import (
	"net/url"
	"strings"
	"time"
)

type ItemModel struct {
	ItemId      int    `json:"item_id"`
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

type RequestItemCreate struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required,min=1"`
	Picture     string `json:"picture" validate:"required,url"`
}

func (r *RequestItemCreate) ToEntity() *ItemEntity {
	if src, err := url.Parse(r.Picture); err == nil && src.Path != "" {
		r.Picture = src.Path
	}
	return &ItemEntity{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Picture:     strings.TrimPrefix(r.Picture, "/"),
	}
}

type RequestItemEdit struct {
	ItemId int `json:"item_id" validate:"required"`
	RequestItemCreate
}

func (r *RequestItemEdit) ToEntity() *ItemEntity {
	if src, err := url.Parse(r.Picture); err == nil && src.Path != "" {
		r.Picture = src.Path
	}
	return &ItemEntity{
		ItemId:      r.ItemId,
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Picture:     strings.TrimPrefix(r.Picture, "/"),
	}
}

type ResponseItemDetail struct {
	AdminId      int       `json:"admin_id"`
	ActiveStatus string    `json:"active_status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ItemModel
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
