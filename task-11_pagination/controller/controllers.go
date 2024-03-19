package controller

import (
	model "manav/pagination/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	var data []model.TrainModel
	data, err := c.Service.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	err = ctx.SendStatus(fiber.StatusOK)
	if err != nil {
		return err
	}

	err = ctx.Render("index", fiber.Map{
		"No":   0,
		"Data": data,
	})

	return err
}

func (c *Controller) GetNextPage(ctx *fiber.Ctx) error {
	var data []model.TrainModel
	var page int
	var err error
	params := ctx.AllParams()
	if v, ok := params["index"]; !ok {
		page = 0
	} else {
		page, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	if page < 0 {
		page = 0
	}
	data, err = c.Service.GetNextPage(ctx.Context(), page)
	if err != nil {
		return err
	}

	err = ctx.SendStatus(fiber.StatusOK)
	if err != nil {
		return err
	}

	return ctx.Render("index", fiber.Map{
		"No":   page,
		"Data": data,
	})
}

func (c *Controller) Search(ctx *fiber.Ctx)error{
	var query = ctx.Body()
	data,err := c.Service.Search(ctx.Context(),string(query))
	if err != nil {
		return err
	}
	ctx.JSON(data)
	return nil
}
