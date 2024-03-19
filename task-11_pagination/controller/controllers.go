package controller

import (
	model "manav/pagination/models"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	var data []model.TrainModel
	// TODO:
	data, err := c.Service.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	err = ctx.SendStatus(fiber.StatusOK)
	if err != nil {
		return err
	}

	err = ctx.JSON(data)
	return err
}
