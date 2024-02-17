package repository

import (
	"github.com/goberman15/sandbox-l4/model"
	"github.com/jmoiron/sqlx"
)

type itemDetailRepo struct {
	db *sqlx.DB
}

func NewItemDetailRepo(db *sqlx.DB) *itemDetailRepo {
	return &itemDetailRepo{db: db}
}

func (r *itemDetailRepo) CreateItemDetail(name string, itemId int) (err error) {
	tx := r.db.MustBegin()

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	_, err = tx.Exec("INSERT INTO item_detail (name, item_id) VALUES ($1, $2)", name, itemId)
	if err != nil {
		return
	}

	_, err = tx.Exec("UPDATE item SET amount = amount + 1, status = $1 WHERE id = $2", "Available", itemId)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return err
}

func (r *itemDetailRepo) ListItemDetailById(itemId int) ([]*model.ItemDetail, error) {
	var itemDetail []*model.ItemDetail
	err := r.db.Select(&itemDetail, "SELECT * FROM item_detail WHERE item_id = $1", itemId)
	return itemDetail, err
}

func (r *itemDetailRepo) GetItemDetail(id int) (*model.ItemDetail, error) {
	var itemDetail model.ItemDetail
	err := r.db.Get(&itemDetail, "SELECT * FROM item_detail WHERE id = $1", id)
	return &itemDetail, err
}

func (r *itemDetailRepo) UpdateItemDetail(id int, value string) error {
	_, err := r.db.Exec("UPDATE item_detail SET name = $1 WHERE id = $2", value, id)
	return err
}

func (r *itemDetailRepo) DeleteItemDetail(id int, itemId int) (err error) {
	tx := r.db.MustBegin()

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	_, err = tx.Exec("DELETE FROM item_detail WHERE id = $1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE item 
		SET amount = amount - 1, 
		status = (CASE WHEN amount = 1 THEN $1 ELSE $2 END) 
		WHERE id = $3`,
		"Restock", "Available", itemId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return
}
