package repository

import (
	"github.com/goberman15/sandbox-l4/model"
)

type ItemDetailRepo interface {
	CreateItemDetail(string, int) error
	ListItemDetailById(int) ([]*model.ItemDetail, error)
	GetItemDetail(int) (*model.ItemDetail, error)
	UpdateItemDetail(int, string) error
	DeleteItemDetail(int, int) error
}
