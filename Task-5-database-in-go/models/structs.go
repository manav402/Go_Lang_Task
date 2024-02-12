package models

import (
	"database/sql"
)
// exported profile struct to use across module
type Profile struct {
	Id 	   int				`sql:"id"`
	Fname  string         	`sql:"fname"`
	Lname  sql.NullString   `sql:"lname"`
	Dob    string 			`sql:"dob"`
	Email  string 			`sql:"email"`
	Number string           `sql:"number"`
}
