package database

import (
	"fmt"
	"log"
	"manav402/crudBooks/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var m map[string]string = make(map[string]string)

// read env function read env file and parse the contant as a map
func ReadEnv() error {

	var err error

	m, err = godotenv.Read(".env")
	if err != nil {
		return err
	}

	return nil
}

// getEnv function return the environment value for the key provided
// Params :- key to which the env value is required
// Returns :- the value and err if error occured in between
func GetEnv(key string) (string, error) {
	if m == nil {
		err := ReadEnv()
		if err != nil {
			return "", err
		}
	}
	if v, ok := m[key]; ok {
		return v, nil
	}
	return "", nil

}

// init db start database connect with table as user and auto migrate the structure to table format
// Returns :- a pointer to db and an error if can't connect to db
func InitDb() (*gorm.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", m["USER"], m["PASSWORD"], m["HOST"], m["DBNAME"])

	// connecting to db at provided connection string
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// checking if we succesfully connected to db or not
	sqlDb,_ := db.DB()
	if err := sqlDb.Ping(); err != nil {
		return nil,err
	}

	log.Println("connected to database")
	// creating table migration from our book model
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil,err
	}
	return db, nil
}
