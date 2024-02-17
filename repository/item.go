package repository

import (
	"github.com/goberman15/sandbox-l4/model"
)

type ItemRepo interface {
	CreateItem(string) error
	ListItems() ([]*model.Item, error)
	GetItem(int) (*model.Item, error)
	UpdateItemField(int, string, any) error
	DeleteItem(int) error
}
