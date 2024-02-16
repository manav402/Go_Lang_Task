package routers
import (
	"github.com/gorilla/mux"
	"manav402/crudBooks/API/Books"
	"manav402/crudBooks/utils"
)
// InitRouter initialize mux routers for rest api purpose
// Params :- a struct with different db pointer and handler functions and a mux router pointer
// Returns :- error if any
func InitRouter(db *utils.DBstruct,router *mux.Router)(error){
	// bookApi structer has a db pointer and also has http handler methods
	bookApi,err := Books.InitBookApi(db)
	if err != nil{
		return err
	}
	// creating http handler rotes filter with state of api
	router.HandleFunc("/Books",bookApi.GetAllBook).Methods("GET")
	router.HandleFunc("/Book/{id}",bookApi.GetOneBook).Methods("GET")
	router.HandleFunc("/Book",bookApi.CreateBook).Methods("POST")
	router.HandleFunc("/Book/{id}",bookApi.EditBook).Methods("PUT")
	return nil
}


