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
	// a simple context for passing around the fiber and router
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	// creating an object for fiber app
	app := fiber.New()

	// initializing the router with context and app which internally initialize the database and all structures
	err := router.InitRouter(ctx, app)
	if err != nil {
		log.Println("error at router ", err)
		log.Fatal(err)
	}


	// start server here and gathering port from env file
	port, err := database.GetEnv("PORT")
	if err != nil {
		port = "8000"
	}

	// listening on definded port
	err = app.Listen(":" + port)
	if err != nil {
		log.Println("error starting server ", err)
	}
}