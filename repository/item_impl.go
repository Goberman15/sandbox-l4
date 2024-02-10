package repository

import (
	"errors"

	"github.com/goberman15/sandbox-l4/model"
	"github.com/jmoiron/sqlx"
)

type itemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *itemRepo {
	return &itemRepo{db: db}
}

func (r *itemRepo) CreateItem(itemName string) error {
	_, err := r.db.Exec("INSERT INTO item (name) VALUES ($1)", itemName)
	return err
}

func (r *itemRepo) ListItems() ([]*model.Item, error) {
	var items []*model.Item
	err := r.db.Select(&items, "SELECT * FROM item")
	return items, err
}

func (r *itemRepo) GetItem(id string) (*model.Item, error) {
	var item model.Item
	err := r.db.Get(&item, "SELECT * FROM item WHERE id = $1", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("item not found")
		}

		return nil, err
	}
	return &item, nil
}

func (r *itemRepo) UpdateItemField(id string, field string, value any) error {
	_, err := r.db.Exec("UPDATE item SET "+field+" = $1 WHERE id = $2", value, id)
	return err
}

func (r *itemRepo) DeleteItem(id string) error {
	_, err := r.db.Exec("DELETE FROM item WHERE id = $1", id)
	return err
}
