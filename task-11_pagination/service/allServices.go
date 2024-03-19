package service

import (
	"context"
	model "manav/pagination/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *Service) GetNextPage(ctx context.Context, page int) ([]model.TrainModel, error) {
	var data []model.TrainModel
	var dummy model.TrainModel
	cur, err := s.trainCollection.Find(ctx, bson.D{}, options.Find().SetSkip(int64((page)*10)).SetLimit(10))
	if err != nil {
		return data, err
	}

	for cur.TryNext(ctx) {
		err = cur.Decode(&dummy)
		if err != nil {
			return data, err
		}
		data = append(data, dummy)
	}
	return data, nil
}

func (s *Service) Search(ctx context.Context, query string) ([]model.TrainModel, error) {
	var data []model.TrainModel
	var dummy model.TrainModel
	// log.Println(query)
	cur, err := s.trainCollection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"train_name": bson.M{"$regex": query, "$options": "i"}},
			{"starts": bson.M{"$regex": query, "$options": "i"}},
			{"ends": bson.M{"$regex": query, "$options": "i"}},
		},
	},options.Find().SetLimit(10))
	if err != nil {
		return data, err
	}
	for cur.TryNext(ctx) {
		err = cur.Decode(&dummy)
		if err != nil {
			return data, err
		}
		data = append(data, dummy)
	}
	return data, nil
}
