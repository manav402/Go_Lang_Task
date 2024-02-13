package main

// importing modularatiy packages
import (
	"fmt"
	"log"
	handle "manav402/server/controller"
	"net/http"
	"manav402/server/DBhandler"
)

func main() {
	var err error
	// reading env file for single time which parse the env file data
	err = DBhandler.ReadEnvFile()
	if err != nil {
		log.Println(err)
	}
	// now using getenvdata function to parse the env data from env files
	var port = DBhandler.GetEnvData("PORT")

	// a file handler of http.Handler type that will be used to serve static file to client on / route
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)

	// different handler defined in controller file associated with each route
	http.HandleFunc("/register", handle.HandleRegsiter)
	http.HandleFunc("/allResult", handle.HandleALlResult)

	// calling connect db method to manipulate databases
	db,err := DBhandler.ConnectDB()
	defer func(){
		if err := db.Dbptr.Close(); err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Println(err)
	}

	log.Println("database connected")
	log.Printf("server is started at http://localhost:%s/\n", port)

	// connecting go application and binding with port defined above
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println(err)
	}

}
