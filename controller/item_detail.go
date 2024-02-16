package controller

import (
	"net/http"

	"github.com/goberman15/sandbox-l4/model"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
)

type ItemDetailController struct {
	r  repository.ItemDetailRepo
	ir repository.ItemRepo
}

func NewItemDetailController(r repository.ItemDetailRepo, ir repository.ItemRepo) *ItemDetailController {
	return &ItemDetailController{r: r, ir: ir}
}

func (c *ItemDetailController) CreateItemDetail(ctx *fiber.Ctx) error {
	var itemDetail model.ItemDetail

	if err := ctx.BodyParser(&itemDetail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.r.CreateItemDetail(itemDetail.Name, itemDetail.ItemId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Success Create Item Detail",
	})
}

func (c *ItemDetailController) ListItemDetailByItemId(ctx *fiber.Ctx) error {
	itemId := ctx.Params("itemId")

	if _, err := c.ir.GetItem(itemId); err != nil {
		if err.Error() == "item not found" {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return err
	}

	itemDetails, err := c.r.ListItemDetailById(itemId)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(itemDetails)
}

func (c *ItemDetailController) UpdateItemDetail(ctx *fiber.Ctx) error {
	itemId := ctx.Params("id")

	var itemDetail model.ItemDetail

	if err := ctx.BodyParser(&itemDetail); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.r.UpdateItemDetail(itemId, itemDetail.Name); err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Update Item Detail",
	})
}

func (c *ItemDetailController) DeleteItemDetail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	itemDetail, err := c.r.GetItemDetail(id)
	if err != nil {
		return err
	}

	if err := c.r.DeleteItemDetail(id, itemDetail.ItemId); err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Item Detail",
	})
}
