package model

type ItemDetail struct {
	Id     int
	ItemId int `db:"item_id"`
	Name   string
}
