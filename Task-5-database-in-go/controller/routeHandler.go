package controller

import (
	"database/sql"
	"html/template"
	"log"
	"manav402/server/DB"
	"manav402/server/models"
	"net/http"
)

// a handler function for route /register which store profile data in database
// @params :- a response writer to write data back to client and req which give us header file from client
func HandleRegsiter(res http.ResponseWriter, req *http.Request) {
	var err error
	// in case some unusual happens the server will send error data
	defer func() {
		if recover() != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Header().Set("content-type", "text/json")
			res.Write([]byte(`{code:500,message:"server error"}`))
		}
	}()
	
	// calling parse form to decrypt the form data
	err = req.ParseForm()
	if err != nil {
		log.Panic(err)
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
	err = DB.Insert(userData)
	if err != nil {
		log.Panic(err)
		return
	}

	// redirecting user to all result route to show each available users
	http.Redirect(res, req, "/allResult", http.StatusSeeOther)
}

// handler function to retrive the all result data from database and send result to server
// @params :- a response writer to write data back to client and req which give us header file from client
func HandleALlResult(res http.ResponseWriter, req *http.Request) {
	// creating array to store each data from database
	dataArr := make([]models.Profile, 0)
	var err error

	// calling get all user method to retrive the all user datas
	dataArr, err = DB.GetAllUser()
	if err != nil {
		panic(err)
	}

	// parsign the template html file which identify the variable needed in place
	temp, err := template.ParseFiles("./static/response.html")
	if err != nil {
		panic(err)
	}

	// filling the blanks with data from database
	temp.Execute(res, dataArr)
}
