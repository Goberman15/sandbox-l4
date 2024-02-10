package repository

import (
	"github.com/goberman15/sandbox-l4/model"
)

type ItemRepo interface {
	CreateItem(string) error
	ListItems() ([]*model.Item, error)
	GetItem(string) (*model.Item, error)
	UpdateItemField(string, string, any) error
	DeleteItem(string) error
}
