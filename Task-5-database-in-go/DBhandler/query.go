package DBhandler

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"manav402/server/models"
)

// exporting db variable to user around the module
type DB struct {
	Dbptr *sql.DB
}
var db = DB{}
// var DB *sql.DB
var err error
var envMap = map[string]string{}

// setter type function for env map
// a read file function uses godotenv package to read env file and store the data on envMap
func ReadEnvFile()error{
	envMap,err = godotenv.Read(".env")
	if err != nil {
		return err
	}
	return nil
}

// getter function for env data
// return data from envMap if map is not yet initialized first the read function is called and the value from map is returned
// @params :- a string corrosponding of key in env map
// @returns :- the value store in map from env file
func GetEnvData(key string)string{
	if envMap == nil {
		ReadEnvFile()
	}
	return envMap[key]
}
// connect db funtion to initialize DB variable and connect to postgres database
func ConnectDB() (*DB,error) {
	m, err := godotenv.Read(".env")
	if err != nil {
		return nil,err
	}
	connStr := fmt.Sprintf("host = %s password = %s user = %s dbname = %s sslmode = disable", m["HOST"], m["PASSWORD"], m["UNAME"], m["DBNAME"])
	dbvar,err := sql.Open("postgres", connStr)
	db.Dbptr = dbvar
	if err != nil {
		return nil,err
	}

	return &DB{dbvar},nil
}

// insert function insert the given struct back to database
// @params :- a struct populated with user profile
func Insert(data models.Profile) error {
	// using prepare statement for query
	query, err := db.Dbptr.Prepare(`INSERT INTO profile (fname,lname,dob,email,number) VALUES ($1,$2,$3,$4,$5)`)
	defer query.Close()
	if err != nil {
		return err
	}
	// executing prepare statement with data from user profile
	_, err = query.Exec(data.Fname, data.Lname.String, data.Dob, data.Email, data.Number)
	if err != nil {
		return err
	}
	return nil
}

// get all user data from profile data base and return back to server
// @return :- a array of profile struct populated with whole data from database of profile
func GetAllUser() ([]models.Profile, error) {

	query := `SELECT * FROM profile`

	rows, err := db.Dbptr.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	// temparory data variable to store in array
	var dummy models.Profile
	var ansArr []models.Profile

	for rows.Next() {
		// scan method to store data at given address
		err := rows.Scan(&dummy.Id, &dummy.Fname, &dummy.Lname, &dummy.Dob, &dummy.Email, &dummy.Number)
		if err != nil {
			return nil, err
		}
		// trim the unnecessary info from date of birth
		dummy.Dob = dummy.Dob[:10]
		ansArr = append(ansArr, dummy)
	}

	return ansArr, nil
}
