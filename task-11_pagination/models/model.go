package model

type TrainModel struct {
	No       int `json:"no." csv:"No."`
	Train_No int `json:"train no." csv:"Train No."`
	Train_Name string `json:"train name" csv:"Train Name"`
	Starts 		string `json:"starts" csv:"Starts"`
	Ends 		string 	`json:"ends" csv:"Ends"`
}
