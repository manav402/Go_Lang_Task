package initialize

import (
	"fmt"
	"encoding/json"
	"os"
	"manav402/jsonParse/models"
)

// initialized data variables for saving the json file data
var fUser,fTech,fContact []byte
var UserData []models.User;
var TechData []models.Tech;
var ContactData []models.Contact;

// function which can read file the by defined file name
// @params :- fileName is the name of just file not path to read
// @returns :- a slice of byte with the data of whole json file
func ReadFile(fileName string)([]byte){
	// opening the user.json file with pointer to file
	var name string= fmt.Sprintf("./json-data/%s.json",fileName)
	fi,err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return fi
}

// init function to initialize the variables with the data from json file
func init() {
	fUser = ReadFile("user")
	fTech = ReadFile("tech")
	fContact = ReadFile("contact")
}

// exported function to initialize struct with proper json data with json file
func InitializeStructs(){
	var err error

	// getting struct for the user type
	err = json.Unmarshal(fUser,&UserData)
	if err != nil {
		panic(err)
	}

	// getting struct for tech type
	err = json.Unmarshal(fTech,&TechData)
	if err != nil{
		panic(err)
	}
	
	// getting struct for Contact type
	err = json.Unmarshal(fContact,&ContactData)
	if err!= nil {
		panic(err)
	}

}