package repository

import (
	"github.com/goberman15/sandbox-l4/model"
)

type ItemDetailRepo interface {
	CreateItemDetail(string, int) error
	ListItemDetailById(string) ([]*model.ItemDetail, error)
	GetItemDetail(string) (*model.ItemDetail, error)
	UpdateItemDetail(string, string) error
	DeleteItemDetail(string, int) error
}
