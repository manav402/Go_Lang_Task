package services

import (
	"errors"
	"fmt"
	"manav402/crudBooks/models"

	"gorm.io/gorm"
)

// check if email is alredy available or not in database
func (s *service) IsEmailAvailable(user *models.User) (bool, error) {
	err := s.DB.DB.First(&models.User{}, "email = ?", user.Email).Error
	if err != nil && errors.Is(err,gorm.ErrRecordNotFound) {
		return false, nil
	}else if err != nil{
		return false,err
	}else{
		return true,nil
	}
}

// create user creats a new user and save data to table
func (s *service) CreateUser(user *models.User) error {

	// check if email alredy exist or not in table
	ok, err := s.IsEmailAvailable(user)
	if err != nil {
		return err
	} else {
		if ok {
			return errors.New("user alredy exist please login")
		}
	}

	// if book not available than only create user
	if len(user.Book) == 0 {
		return s.DB.DB.Omit("Book").Create(user).Error
	}

	// if book is also available create data for junction table
	return s.DB.DB.Create(user).Error
}

func (s *service) Login(user *models.User) error {

	// check if the email is available or not
	ok, err := s.IsEmailAvailable(user)
	if err != nil {
		return err
	} else {
		if !ok {
			return errors.New("user not exist please register")
		}
	}

	// checking if user entered right password or not
	var password = user.Password
	err = s.DB.DB.First(user, "email = ?", user.Email).Error
	if err != nil {
		return err
	}
	
	if password == user.Password {
		return nil
	} else {
		return errors.New("wrong Password")
	}
}

// return all books the user currently have
func (s *service) GetAllUsersBook(userid string) ([]models.User, error) {

	var data []models.User
	err := s.DB.DB.Table("users").Preload("Book").Find(&data, userid).Error
	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}

	return data, nil
}

func (s *service) IssueABook(userId string, bookId string) error {
	// to implemetn the logic
	return nil
}

func (s *service) ReturnABook(userId string, bookId string) error {
	// to implement the logic
	return nil
}
