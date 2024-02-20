package models

// import "gorm.io/gorm"

// a general structure to be used around the module
type Book struct {
	// gorm.Model
	Title     string `json:"title" gorm:"column:title"`
	Author    string `json:"author" gorm:"column:author"`
	ISBN      string `json:"isbn" gorm:"column:isbn;primaryKey;autoIncrement:false"`
	Publisher string `json:"publisher" gorm:"column:publisher;default:unknown"`
	Year      int    `json:"year" gorm:"column:year"`
	Genre     string `json:"genre" gorm:"column:genre;default:unknown"`
	Quantity  int    `json:"quantity" gorm:"column:quantity;default:1"`
	User      []User `gorm:"many2many:user_books"`
}

type User struct {
	Name     string `json:"name" gorm:"column:name"`
	Age      int    `json:"age" gorm:"coulmn:age"`
	UserId   int    `json:"userid" gorm:"column:userid;primaryKey"`
	Email    string `json:"email" gorm:"column:email;unique"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role;default:user"`
	Book     []Book `gorm:"many2many:user_books"`
}
