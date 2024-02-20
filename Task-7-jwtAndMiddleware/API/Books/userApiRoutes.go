package Books

import (
	"encoding/json"
	"errors"
	"manav402/crudBooks/middleware"
	"manav402/crudBooks/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// return list of all books issued by the user
func (user *Api) GetAllUsersBook(res http.ResponseWriter, req *http.Request) {
	var data []models.User
	var id string

	m := mux.Vars(req)
	id = m["userid"]

	data, err := user.userService.GetAllUsersBook(id)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write([]byte("here is your data"))
	json.NewEncoder(res).Encode(data)
}

// create a user if not found in database
func (user *Api) CreateUser(res http.ResponseWriter, req *http.Request) {

	var userData models.User = models.User{}

	err := json.NewDecoder(req.Body).Decode(&userData)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	err = user.userService.CreateUser(&userData)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("user created succesfull"))
}

// login functionality with simple password checking and cookie setting
func (user *Api) Login(res http.ResponseWriter, req *http.Request) {

	var userData models.User
	json.NewDecoder(req.Body).Decode(&userData)
	if userData.Email == "" || userData.Password == "" {
		errorToClient(res, req, errors.New("broken body please give all required details"))
		return
	}
	
	err := user.userService.Login(&userData)
	if err != nil {
		errorToClient(res, req, errors.New("maybe user email alredy exist or "+err.Error()))
		return
	}
	
	// creating cookie if the user is valid and password is matched
	jwtToken, err := middleware.GenerateJwt(userData.Role)
	if err != nil {
		errorToClient(res, req, errors.New("error generating token ERR:-"+err.Error()))
	}
	var cookie = http.Cookie{
		Name:   "cookie",
		Value:  jwtToken,
		MaxAge: int(time.Now().Add(5 * time.Minute).Unix()),
		Path:   "/",
	}

	http.SetCookie(res, &cookie)
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("user login success"))
}

// TODO :-  issue a book
func (user *Api) IssueABook(res http.ResponseWriter, req *http.Request) {

}

func (user *Api) ReturnABook(res http.ResponseWriter, req *http.Request) {

}

//TODO :- Return a book
