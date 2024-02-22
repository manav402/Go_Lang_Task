package main

import (
	// "context"
	"context"
	"log"
	"manav402/FiberMongo/database"
	"manav402/FiberMongo/router"

	"github.com/gofiber/fiber/v2"
)

// TO-DO :- perform crud operations for department and students (many2one)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	app := fiber.New()

	err := router.InitRouter(ctx, app)
	if err != nil {
		log.Println("error at router ", err)
		log.Fatal(err)
	}

	// connect with database here

	// start server here
	port, err := database.GetEnv("PORT")
	if err != nil {
		port = "8000"
	}
	err = app.Listen(":" + port)
	if err != nil {
		log.Println("error starting server ", err)
	}
}
