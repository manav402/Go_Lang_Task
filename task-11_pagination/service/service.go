package service

import (
	"context"
	"errors"
	"manav/pagination/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// service structure that contains a mongo collection and collection methods
type Service struct{
	trainCollection 		*mongo.Collection
}

// initializing service
func NewService(ctx context.Context)(*Service,error){
	// first connecting to database
	var db,err = database.ConnectDB(ctx)
	if err != nil {
		return nil,err
	}

	if db != nil{
		return &Service{db.Collection("TrainCollection")},nil
	}
	
	return nil,errors.New("db pointer is null please initialize the database first")
}