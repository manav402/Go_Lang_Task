package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"manav402/server/DBhandler"
	"manav402/server/models"
)

// check if the user is redirected or landed directly and send data accordingly
type dataSend struct {
	IsRedirected bool
	Arr          []models.Profile
}

func ThrowErrorToClient(res http.ResponseWriter, req *http.Request, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Header().Set("content-type", "text/json")
	res.Write([]byte(fmt.Sprintf("{code:500,message:server error :- %s}", err)))
	http.Error(res,"err",http.StatusInternalServerError)
}

// a handler function for route /register which store profile data in database
// @params :- a response writer to write data back to client and req which give us header file from client
func HandleRegsiter(res http.ResponseWriter, req *http.Request) {

	var err error
	// in case some unusual happens the server will send error data
	// calling parse form to decrypt the form data
	err = req.ParseForm()
	if err != nil {
		ThrowErrorToClient(res, req, err)
		return
	}
	// storing created data to user profile structure
	var userData = models.Profile{
		Fname:  req.FormValue("first_name"),
		Lname:  sql.NullString{String: req.FormValue("last_name")},
		Dob:    req.FormValue("dob"),
		Email:  req.FormValue("email"),
		Number: req.FormValue("mo_number"),
	}
	
	// calling insert function to store data in database
	err = DBhandler.Insert(userData)
	if err != nil {
		ThrowErrorToClient(res, req, err)
		return
	}

	// redirecting user to all result route to show each available users
	http.Redirect(res, req, "/allResult?valid=true", http.StatusSeeOther)
}

// handler function to retrive the all result data from database and send result to server
// @params :- a response writer to write data back to client and req which give us header file from client
func HandleALlResult(res http.ResponseWriter, req *http.Request) {
	// var url *url.URL;

	var err error
	// a data send for checking if a query present and rendering html with the flag
	var data dataSend

	// parsing the url query
	url, err := req.URL.Parse(req.RequestURI)
	if err != nil {
		ThrowErrorToClient(res, req, err)
	}
	// getting a map for all query value
	a := url.Query()

	// if valid query present than the user is redirected
	if _, ok := a["valid"]; ok {
		data = dataSend{IsRedirected: true}
	}

	// if req.URL
	// creating array to store each data from database
	dataArr := make([]models.Profile, 0)
	// calling get all user method to retrive the all user datas
	dataArr, err = DBhandler.GetAllUser()
	if err != nil {
		ThrowErrorToClient(res, req, err)
		return
	}

	// parsign the template html file which identify the variable needed in place
	temp, err := template.ParseFiles("./static/response.html")
	if err != nil {
		ThrowErrorToClient(res, req, err)
		return
	}

	// pushing data to the client
	data.Arr = dataArr
	// filling the blanks with data from database
	temp.Execute(res, data)
}
