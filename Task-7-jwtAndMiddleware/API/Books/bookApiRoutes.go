package Books

import (
	"encoding/json"
	"manav402/crudBooks/models"
	"manav402/crudBooks/services"
	"manav402/crudBooks/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// Api consist of DB pointer,a service Struct which implement all the services and http handler methods
type Api struct {
	DB          *utils.DBstruct
	bookService services.Service
	userService services.Service
}

// InitBookApi initialize the BookApi struct and return pointer to it for caller
func InitBookApi(db *utils.DBstruct) (*Api, error) {
	// InitServie return a struct which has service implemented which is an interface
	bookServiece, err := services.InitService(db)
	if err != nil {
		return nil, err
	}
	return &Api{DB: db, bookService: bookServiece,userService: bookServiece}, nil
}

// Return Message is an output formate for client in json
type ReturnMessage struct {
	Code    int
	Message string
	Data    []models.Book
}

// errorToClient usefull for throwing error to client in json
func errorToClient(res http.ResponseWriter, req *http.Request, err error) {
	errorMessage := ReturnMessage{Code: http.StatusInternalServerError, Message: err.Error()}

	res.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(res).Encode(errorMessage)

	// http.Error(res, err.Error(), http.StatusInternalServerError)
}

// GetALlBook api returns to client in json of all book in table
// a Http handler method
func (bookApi *Api) GetAllBook(res http.ResponseWriter, req *http.Request) {

	// call GetAllBook api service which deals with tables
	data, err := bookApi.bookService.GetAllBook()

	if err != nil {
		errorToClient(res, req, err)
		return
	}
	// if no data found return error to client
	if len(data) == 0 {
		res.WriteHeader(http.StatusOK)
		outputMessage := ReturnMessage{Code: http.StatusNotFound, Message: "Success, but no data available", Data: nil}
		json.NewEncoder(res).Encode(outputMessage)
		return
	}

	res.WriteHeader(http.StatusOK)
	outputMessage := ReturnMessage{Code: http.StatusOK, Message: "Success", Data: data}
	json.NewEncoder(res).Encode(outputMessage)

}

// GetOneBook method returns the book asked from url in json basis
// A http handler method
func (bookApi *Api) GetOneBook(res http.ResponseWriter, req *http.Request) {
	// mux.vars to parse url and gather info about the book they asked
	val := mux.Vars(req)
	data, err := bookApi.bookService.GetOneBook(val["id"])

	if err != nil {
		errorToClient(res, req, err)
		return
	}

	// if no data found return friendly error to client
	if data.ISBN == "" {
		res.WriteHeader(http.StatusOK)
		outputMessage := ReturnMessage{Code: http.StatusNotFound, Message: "Success, but no data available", Data: nil}
		json.NewEncoder(res).Encode(outputMessage)
		return
	}

	res.WriteHeader(http.StatusOK)
	outputMessage := ReturnMessage{Code: http.StatusOK, Message: "Success", Data: []models.Book{data}}
	json.NewEncoder(res).Encode(outputMessage)

}

// CreteBook method create a book in table and return success message in json
func (bookApi *Api) CreateBook(res http.ResponseWriter, req *http.Request) {
	var book models.Book
	// json new decoder decode req.body and populate teh book struct
	json.NewDecoder(req.Body).Decode(&book)

	// CreateBook service create book in table if not exists
	err := bookApi.bookService.CreateBook(book)
	if err != nil {
		errorToClient(res, req, err)
		return
	}

	res.WriteHeader(http.StatusOK)
	outputMessage := ReturnMessage{Code: http.StatusOK, Message: "Data Inserted"}
	json.NewEncoder(res).Encode(outputMessage)

}

// EditBook method can change table data from provided book id
func (bookApi *Api) EditBook(res http.ResponseWriter, req *http.Request) {
	var book models.Book

	json.NewDecoder(req.Body).Decode(&book)

	// EditBook service play with table and can edit table data
	err := bookApi.bookService.EditBook(book)
	if err != nil {
		errorToClient(res, req, err)
		return
	}

	res.WriteHeader(http.StatusOK)
	outputMessage := ReturnMessage{Code: http.StatusOK, Message: "Data Updated"}
	json.NewEncoder(res).Encode(outputMessage)

}
