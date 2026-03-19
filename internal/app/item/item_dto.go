package item

import (
	"net/url"
	"strings"
)

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
