package routes

import (
	"context"
	"manav/pagination/controller"

	"github.com/gofiber/fiber/v2"
)

// initialize get and post routes of app
// @params : context,fiber app and controller with mongo client pointer
func InitRoute(ctx context.Context, app *fiber.App,controller controller.Controller) error {
	app.Get("/getAll", controller.GetAll)
	app.Get("/get/:index",controller.GetNextPage)
	app.Post("/search",controller.Search)
	return nil
}
