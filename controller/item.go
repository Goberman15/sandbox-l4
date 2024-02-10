package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goberman15/sandbox-l4/model"
	"github.com/goberman15/sandbox-l4/repository"
)

type ItemController struct {
	r repository.ItemRepo
}

func NewItemController(r repository.ItemRepo) *ItemController {
	return &ItemController{r}
}

func (c *ItemController) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := c.r.ListItems()
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
}

func (c *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.CreateItem(item.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success Create Item"))
}

func (c *ItemController) GetItemById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := c.r.GetItem(id)
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

	json.NewEncoder(w).Encode(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *ItemController) UpdateItemStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.UpdateItemField(id, "status", item.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Update Item Status"))
}

func (c *ItemController) UpdateItemAmount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.r.UpdateItemField(id, "amount", item.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Update Item Amount"))
}

func (c *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := c.r.DeleteItem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Delete Item"))
}
