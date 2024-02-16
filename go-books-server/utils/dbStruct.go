package utils

import (
	"manav402/crudBooks/database"
	"manav402/crudBooks/models"

	"gorm.io/gorm"
)

// DBstruct store various database pointer and can be accessible around the module
type DBstruct struct {
	DB *gorm.DB
}

// InitDbStruct initialize the structure and return the pointer of that to caller which also initializez the database
// Returns :- pointer to dbStruct and error 
func InitDBstruct() (*DBstruct, error) {
	db, err := database.InitDb()
	if err != nil {
		return nil,err
	}

	return &DBstruct{DB:db}, nil
}

// Migrate method of DBstruct allow us to Perform Migration on run time
func (db *DBstruct) Migrate() error {
	err := db.DB.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}
	return nil
}

// Close method can be used to close database at run time
func (db *DBstruct) Close() error {
	sqldb,err := db.DB.DB()
	if err != nil{
		return err
	}
	return sqldb.Close()
}
