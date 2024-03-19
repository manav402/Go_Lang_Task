package service

import (
	"context"
	model "manav/pagination/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) InsertOne(ctx context.Context, data model.TrainModel) error {
	_, err := s.trainCollection.InsertOne(ctx, data)
	return err
}

func (s *Service) GetAll(ctx context.Context) ([]model.TrainModel, error) {
	var data = []model.TrainModel{}
	var dummy = model.TrainModel{}
	cur, err := s.trainCollection.Find(ctx, bson.D{})
	if err != nil {
		return []model.TrainModel{}, err
	}

	for cur.TryNext(ctx) {
		err = cur.Decode(&dummy)
		if err != nil {
			return []model.TrainModel{}, err
		}
		data = append(data, dummy)
	}
	return data, nil
}
