package models

// import "gorm.io/gorm"

// a general structure to be used around the module
type Book struct {
	// gorm.Model
	Title     string `json:"title" gorm:"column:title"`
	Author    string `json:"author" gorm:"column:author"`
	ISBN      string `json:"isbn" gorm:"column:isbn;primaryKey"`
	Publisher string `json:"publisher" gorm:"column:publisher"`
	Year      int    `json:"year" gorm:"column:year"`
	Genre     string `json:"genre" gorm:"column:genre"`
}
