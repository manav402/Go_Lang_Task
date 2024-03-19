package service

import (
	"context"
	"errors"
	"manav/pagination/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct{
	trainCollection 		*mongo.Collection
}

func NewService(ctx context.Context)(*Service,error){
	var db,err = database.ConnectDB(ctx)
	if err != nil {
		return nil,err
	}
	if db != nil{
		return &Service{db.Collection("TrainCollection")},nil
	}
	return nil,errors.New("db pointer is null please initialize the database first")
}