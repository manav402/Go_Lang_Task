package controller

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type SubjectController struct {
	MongoClient   *mongo.Client
}

type StudentController struct{
	MongoClient  *mongo.Client
}


type DepartmentController struct{
	MongoClient		*mongo.Client
}

type API struct {
	SubjectController
	DepartmentController
	StudentController
}

// initapi will initialize the controller methods
func InitApi(mongoClient *mongo.Client)(*API,error){
	if mongoClient == nil {
		return nil , errors.New("mongo client shoudl not be nil")
	}
	return &API{SubjectController{mongoClient},DepartmentController{mongoClient},StudentController{mongoClient}},nil
}