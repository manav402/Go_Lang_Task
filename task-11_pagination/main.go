package main

import (
	"context"
	"flag"
	"log"

	"manav/pagination/controller"
	"manav/pagination/csv"
	"manav/pagination/database"
	"manav/pagination/routes"

 	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var fileName string
	var parseCsv bool

	// parsing flags to read the data from csv files
	flag.StringVar(&fileName, "file", "All_Indian_Trains.csv", "read csv file from specified file name")
	flag.BoolVar(&parseCsv, "readCsv", false, "this flag specify to read the csv data")
	flag.Parse()

	// creating a general context to perform gradually shutdown of server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// creating a render engine of html templates
	var engine = html.New("./templates",".html")
	var app = fiber.New(fiber.Config{
		Views:       engine,
	})

	// making public folder accessible so can grab css and js
	app.Static("/public","./public")

	// controller initialization to perform read csv operations
	controller, err := controller.NewController(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	// if flag is exist the data will be populated from csv
	if parseCsv {
		err = csv.ReadCSV(ctx, controller.Service, fileName)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// initialization of app routes
	err = routes.InitRoute(ctx, app, *controller)
	if err != nil {
		log.Println(err)
		return
	}

	// reading env variable if not possible assigning a pre defined port
	port, err := database.Getenv("PORT")
	if err != nil {
		port = "8000"
	}

	// starting server
	err = app.Listen(":" + port)
	log.Println(err)
}
