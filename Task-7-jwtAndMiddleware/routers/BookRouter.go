package routers

import (
	"manav402/crudBooks/API/Books"
	"manav402/crudBooks/middleware"
	"manav402/crudBooks/utils"

	"github.com/gorilla/mux"
)

// InitRouter initialize mux routers for rest api purpose
// Params :- a struct with different db pointer and handler functions and a mux router pointer
// Returns :- error if any
func InitRouter(db *utils.DBstruct, router *mux.Router) error {
	// bookApi structer has a db pointer and also has http handler methods
	Api, err := Books.InitBookApi(db)
	if err != nil {
		return err
	}
	// creating http handler rotes filter with state of api
	userRoute := router.PathPrefix("/User").Subrouter()
	userRoute.HandleFunc("/login",Api.Login).Methods("POST")
	userRoute.HandleFunc("/register", Api.CreateUser).Methods("POST")
	userRoute.Handle("/{userid}/Books",middleware.VerifyJwt(Api.GetAllUsersBook,"user")).Methods("GET")
	userRoute.Handle("/IssueBook/{isbn}", middleware.VerifyJwt(Api.IssueABook,"user")).Methods("GET")
	userRoute.Handle("/Returnbook/{isbn}", middleware.VerifyJwt(Api.ReturnABook,"user")).Methods("GET")

	bookRoute := router.PathPrefix("/Book").Subrouter()
	bookRoute.Handle("/getAll", middleware.VerifyJwt(Api.GetAllBook,"admin")).Methods("GET")
	bookRoute.Handle("/{id}", middleware.VerifyJwt(Api.GetOneBook,"admin")).Methods("GET")
	bookRoute.Handle("/create", middleware.VerifyJwt(Api.CreateBook,"admin")).Methods("POST")
	bookRoute.Handle("/Book/{id}", middleware.VerifyJwt(Api.EditBook,"admin")).Methods("PUT")
	return nil
}
