package main

// importing modularatiy packages
import (
	"fmt"
	"log"
	handle "manav402/server/controller"
	"net/http"
	"manav402/server/DBhandler"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	m, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err)
	}
	var port = m["PORT"]

	// a file handler of http.Handler type that will be used to serve static file to client on / route
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)

	// different handler defined in controller file associated with each route
	http.HandleFunc("/register", handle.HandleRegsiter)
	http.HandleFunc("/allResult", handle.HandleALlResult)

	// calling connect db method to manipulate databases
	db,err := DBhandler.ConnectDB()
	defer db.Dbptr.Close()
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
