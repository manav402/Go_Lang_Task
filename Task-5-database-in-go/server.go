package main

// importing modularatiy packages
import (
	"fmt"
	"log"
	"net/http"

	"manav402/server/DB"
	handle "manav402/server/controller"
)

func main() {
	var err error
	var port = 8000

	// a file handler of http.Handler type that will be used to serve static file to client on / route
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)

	// different handler defined in controller file associated with each route
	http.HandleFunc("/register", handle.HandleRegsiter)
	http.HandleFunc("/allResult", handle.HandleALlResult)

	// calling connect db method to manipulate databases
	err = DB.ConnectDB()
	defer DB.DB.Close()
	if err != nil {
		log.Println(err)
	}

	log.Println("database connected")
	log.Printf("server is started at http://localhost:%d/\n", port)

	// connecting go application and binding with port defined above
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println(err)
	}

}
