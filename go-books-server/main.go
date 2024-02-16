package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"manav402/crudBooks/database"
	"manav402/crudBooks/routers"
	"manav402/crudBooks/utils"
)

func main(){
	// sparse .env file for environment variables
	err := database.ReadEnv()
	if err != nil {
		log.Println(err)
		return
	}

	// initialize first struct which pass the db pointer across module
	// this method internally initialize every required structure
	db,err := utils.InitDBstruct()
	defer func (){err := db.Close(); if err != nil {fmt.Println(err)}}()
	if err != nil {
		log.Println(err)
		return
	}


	// after initializing each struct the server will be started
	err = startServer(db)
	if err != nil {
		log.Println(err)
		return
	}

}

// after startup the server will be started to listen in port
func startServer(db *utils.DBstruct)error{
	router := mux.NewRouter()
	// starting each books routs
	err := routers.InitRouter(db,router)
	if err != nil {
		return err
	}

	// reading env variable for assigning a port value
	port,err := database.GetEnv("PORT")
	if err != nil {
		return err
	}

	log.Println("listening on port ",port)
	// starting the server at port
	err = http.ListenAndServe(":"+port,router)
	if err != nil {
		return err
	}
	return nil

}