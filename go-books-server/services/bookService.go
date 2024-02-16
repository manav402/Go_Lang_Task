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

// GetAllBook method is an api service which return all available book from table
// Returns :- a structure array of book and an error if any
func (s *service) GetAllBook() ([]models.Book, error) {
	var modelArr []models.Book

	// usign gorm.Find to find all data from table model
	err := s.DB.DB.Find(&modelArr).Error
	if err != nil {
		return nil,err
	}

	return modelArr, nil

}

// GetOneBook method is an api service which return one book based on id
// Params :- book id (ISBN) parsed from client
// Returns :- A book struct and an error if any
func (s *service) GetOneBook(id string) (models.Book, error) {
	var model models.Book

	// using gorm.First and id as condition parameter
	err := s.DB.DB.Model(&models.Book{}).First(&model,id).Error
	// err := s.DB.DB.First(&model).Error
	if err != nil{
		return models.Book{},err
	}

	return model, nil
}

// CreateBook service can create a book from user inputs
// Params :- a book struct parsed from body on user header
// Returns :- error if any
func (s *service) CreateBook(book models.Book) error {
	// using gorm.create method so we can create based on book model
	err := s.DB.DB.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

// Edit Book service can replace the book with new one or create one if not exist
// Params :- a new book stuct to replace parsed from user
// Returns :- an error if any
func (s *service) EditBook(book models.Book) error {
	// usign gorm.save if the matching id available the record is updated else new created
	err := s.DB.DB.Save(&book).Error
	if err != nil {
		return err
	}

	return nil
}
