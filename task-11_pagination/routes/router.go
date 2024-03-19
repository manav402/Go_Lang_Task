package router

import (
	"context"
	"manav/pagination/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(ctx context.Context, app *fiber.App,controller controller.Controller) error {
	app.Get("/getAll", controller.GetAll)
	return nil
}
