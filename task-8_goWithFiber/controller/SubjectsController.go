package controller

import (
	"context"
	"errors"
	"log"
	"manav402/FiberMongo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (subject *SubjectController) CreateAsubject(ctx context.Context, subjectData models.Subjects) error {
	var output models.Department
	response := subject.MongoClient.Database("university").Collection("Department").FindOne(ctx, bson.D{{Key: "departmentid", Value: subjectData.Departmentid}})

	err := response.Decode(&output)
	if err != nil {
		return err
	}
	log.Println(output)
	output.Subjects = append(output.Subjects, subjectData.Subjectid)
	log.Println(output)

	result := subject.MongoClient.Database("university").Collection("Department").FindOneAndReplace(ctx, bson.D{{Key: "departmentid", Value: output.DepartmentId}}, output)
	if result.Err() != nil {
		return result.Err()
	}

	_, err = subject.MongoClient.Database("university").Collection("Subject").InsertOne(ctx, subjectData)
	if err != nil {
		return err
	}

	return nil
}

func (subject *SubjectController) GetAllSubject(ctx context.Context) ([]models.Subjects, error) {
	var data []models.Subjects
	cursor, err := subject.MongoClient.Database("university").Collection("Subject").Find(ctx, map[string]interface{}{})
	if err != nil {
		return []models.Subjects{}, err
	}

	cursor.All(ctx, &data)

	return data, nil
}

func (subject *SubjectController) GetASubject(ctx context.Context, subjectId string) (models.Subjects, error) {
	var data models.Subjects
	cursor := subject.MongoClient.Database("university").Collection("Subject").FindOne(ctx, bson.D{{Key: "subjectid", Value: subjectId}})

	err := cursor.Decode(&data)
	if err != nil {
		return models.Subjects{}, nil
	}
	return data, nil
}

func (subject *SubjectController) EditSubject(ctx context.Context, subjectData models.Subjects) error {
	result := subject.MongoClient.Database("university").Collection("Subject").FindOneAndReplace(ctx, bson.D{{Key: "subjectid", Value: subjectData.Subjectid}}, subjectData)

	return result.Err()
}
func (subject *SubjectController) DeleteSubject(ctx context.Context, subjectId string) error {
	result, err := subject.MongoClient.Database("university").Collection("Subject").DeleteOne(ctx, bson.D{{Key: "subjectid", Value: subjectId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no id found to delete the data maybe data alredy deleted")
	}
	return nil
}
