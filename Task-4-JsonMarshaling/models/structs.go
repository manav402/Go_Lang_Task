package models

// structure defination from tech.json file with appropriete tag and exported fields
type TechDets struct {
	Tech 	string 			`json:"tech"`
	Exp 	float64			`json:"exp"`
}

type Tech struct {
	Id       int        `json:"id"`
	TechDets []TechDets `json:"techDets"`
}

// user.json schema or structure defination with proper tag and exported field
type Address struct {
	Area    string 		`json:"area"`
	Country string 		`json:"country"`
}

type User struct {
	Id 		int 		`json:"id"`
	Name 	string		`json:"name"`
	Address				`json:"address"`
}

// contact.json field and corrosponding json tag
type ContactDets struct{
	Email 	string		`json:"email"`
	Phone	string		`json:"phone"`
}	

type Contact struct{
	Id 		int 		`json:"id"`
	ContactDets			`json:"contactDets"`
}

// final out put struct file
type NewTech struct{
	Techdata string 	`json:"techdata"`
	Exp 	 float64	`json:"exp"`
}

type Output struct{
	Userid	int
	Name	string
	Address
	TechDetails []NewTech
	Email 	string
	Phone 	string
}
 