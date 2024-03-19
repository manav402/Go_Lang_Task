package router

import (
	"context"
	"manav/pagination/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(ctx context.Context, app *fiber.App,controller controller.Controller) error {
	// var router = app.
	app.Get("/getAll", controller.GetAll)
	app.Get("/get/:index",controller.GetNextPage)
	app.Post("/search",controller.Search)
	return nil
}
