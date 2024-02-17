package controller

import (
	"net/http"

	"github.com/goberman15/sandbox-l4/customErr"
	"github.com/goberman15/sandbox-l4/model"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
)

type ItemController struct {
	r repository.ItemRepo
}

func NewItemController(r repository.ItemRepo) *ItemController {
	return &ItemController{r}
}

func (c *ItemController) ListItems(ctx *fiber.Ctx) error {
	items, err := c.r.ListItems()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(items)
}

func (c *ItemController) CreateItem(ctx *fiber.Ctx) error {
	var item model.Item

	if err := ctx.BodyParser(&item); err != nil {
		return err
	}

	if err := c.r.CreateItem(item.Name); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Success Create Item",
	})
}

func (c *ItemController) GetItemById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return customErr.NewBadRequestError("invalid id")
	}
	item, err := c.r.GetItem(id)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(item)
}

func (c *ItemController) UpdateItemStatus(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return customErr.NewBadRequestError("invalid id")
	}

	var item model.Item

	if err := ctx.BodyParser(&item); err != nil {
		return err
	}

	if err := c.r.UpdateItemField(id, "status", item.Status); err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Update Item Status",
	})
}

func (c *ItemController) UpdateItemAmount(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return customErr.NewBadRequestError("invalid id")
	}

	var item model.Item

	if err := ctx.BodyParser(&item); err != nil {
		return err
	}

	if err := c.r.UpdateItemField(id, "amount", item.Amount); err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Update Item Amount",
	})
}

func (c *ItemController) DeleteItem(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return customErr.NewBadRequestError("invalid id")
	}

	if err := c.r.DeleteItem(id); err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Item",
	})
}
