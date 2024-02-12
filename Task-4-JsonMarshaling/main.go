package main

import (
	"encoding/json"
	"os"

	. "manav402/jsonParse/initialize"
	"manav402/jsonParse/models"
)
// global map so i can preppend the country code in mobile number
var m map[string]string = map[string]string{
	"IND": "+91",
	"UK":  "+41",
}

func main() {
	// calling my initialize method so i can read data from the files
	InitializeStructs()

	// variable for the data manipulation use for json key changes
	var lastDataIndex = len(UserData)
	var outputStruct []models.Output
	var newTech []models.NewTech
	var userDataIndex, contactDataIndex, techdataIndex int

	outputStruct = make([]models.Output, lastDataIndex)

	// looping over user data and creating a defined new output structure
	for userDataIndex = 0; userDataIndex < lastDataIndex; userDataIndex++ {
		// getting index for contact data
		for contactDataIndex = 0; contactDataIndex < len(ContactData); contactDataIndex++ {
			if ContactData[contactDataIndex].Id == UserData[userDataIndex].Id {
				break
			}
		}
		// getting index for tech related json data
		for techdataIndex = 0; techdataIndex < len(TechData); techdataIndex++ {
			if TechData[techdataIndex].Id == UserData[userDataIndex].Id {
				break
			}
		}
		// transforming the form of tech data to my new formate of newTech data
		newTech = make([]models.NewTech, 0)
		for j := 0; j < len(TechData[techdataIndex].TechDets); j++ {
			var subTech models.NewTech
			subTech.Techdata = TechData[techdataIndex].TechDets[j].Tech
			subTech.Exp = TechData[techdataIndex].TechDets[j].Exp
			newTech = append(newTech, subTech)
		}

		// creating the final output data of predefined formate
		var newStruct models.Output = models.Output{
			Userid: UserData[userDataIndex].Id,
			Name:   UserData[userDataIndex].Name,
			Address: models.Address{
				Area:    UserData[userDataIndex].Area,
				Country: UserData[userDataIndex].Country,
			},
			TechDetails: newTech,
			Email:       ContactData[contactDataIndex].Email,
			Phone:       m[UserData[userDataIndex].Country] + "-" + ContactData[contactDataIndex].Phone,
		}

		// inserting the data as an array on ouput variable
		outputStruct[userDataIndex] = newStruct

	}

	// converting the output struct to a json string
	outputByte, err := json.MarshalIndent(outputStruct,"","   ")
	if err != nil {
		panic(err)
	}

	// wrting the data back to ouput.json file with read write permission
	err = os.WriteFile("output.json", outputByte, 0666)
}
