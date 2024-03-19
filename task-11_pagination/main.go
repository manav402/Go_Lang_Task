package main

import (
	"context"
	"flag"
	"log"
	"manav/pagination/controller"
	"manav/pagination/csv"
	"manav/pagination/database"
	router "manav/pagination/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var fileName string
	var parseCsv bool
	flag.StringVar(&fileName, "file", "All_Indian_Trains.csv", "read csv file from specified file name")
	flag.BoolVar(&parseCsv, "readCsv", false, "this flag specify to read the csv data")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var app = fiber.New()

	controller, err := controller.NewController(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(parseCsv,fileName)
	if parseCsv {
		log.Println("parsing the csv file")
		err = csv.ReadCSV(ctx, controller.Service, fileName)
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = router.InitRoute(ctx, app, *controller)
	if err != nil {
		log.Println(err)
		return
	}

	port, err := database.Getenv("PORT")
	if err != nil {
		port = "8000"
	}
	err = app.Listen(":" + port)
	log.Println(err)
}
