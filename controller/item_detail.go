package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goberman15/sandbox-l4/model"
	"github.com/goberman15/sandbox-l4/repository"
)

type ItemDetailController struct {
	r repository.ItemDetailRepo
	ir repository.ItemRepo
}

func NewItemDetailController(r repository.ItemDetailRepo, ir repository.ItemRepo) *ItemDetailController {
	return &ItemDetailController{r: r, ir: ir}
}

func (c *ItemDetailController) CreateItemDetail(w http.ResponseWriter, r *http.Request) {
	var itemDetail model.ItemDetail
	err := json.NewDecoder(r.Body).Decode(&itemDetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.CreateItemDetail(itemDetail.Name, itemDetail.ItemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success Create Item Detail"))
}

func (c *ItemDetailController) ListItemDetailByItemId(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "itemId")

	_, err := c.ir.GetItem(itemId)
	if err != nil {
		if err.Error() == "item not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	itemDetails, err := c.r.ListItemDetailById(itemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(itemDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *ItemDetailController) UpdateItemDetail(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "id")

	var itemDetail model.ItemDetail
	err := json.NewDecoder(r.Body).Decode(&itemDetail)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.UpdateItemDetail(itemId, itemDetail.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Update Item Detail"))
}

func (c *ItemDetailController) DeleteItemDetail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	itemDetail, err := c.r.GetItemDetail(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.DeleteItemDetail(id, itemDetail.ItemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Delete Item Detail"))
}
