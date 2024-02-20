package services

import (
	"manav402/crudBooks/models"
	"manav402/crudBooks/utils"
)

type Service interface {
	GetAllBook() ([]models.Book, error)
	GetOneBook(id string) (models.Book, error)
	CreateBook(models.Book) error
	EditBook(models.Book) error
	CreateUser(*models.User) error
	GetAllUsersBook(string) ([]models.User,error)
	IssueABook(string,string) error
	ReturnABook(string,string) error
	Login(*models.User) error
}

// This service structure will implement the Service interface
// And contains a pointer to db
type service struct {
	DB *utils.DBstruct
}

// InitSerivice only return service struct if the struct implements interface
func InitService(db *utils.DBstruct) (Service, error) {
	return &service{DB:db}, nil
}