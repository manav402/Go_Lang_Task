package controller

import (
	"context"
	"errors"
	"manav402/FiberMongo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (student *StudentController) CreateAstudent(ctx context.Context, studentData models.Student) error {
	_, err := student.MongoClient.Database("university").Collection("Student").InsertOne(ctx, studentData)
	if err != nil {
		return err
	}

	return nil
}

func (student *StudentController) GetAllStudents(ctx context.Context) ([]models.Student, error) {
	var data []models.Student
	cursor, err := student.MongoClient.Database("university").Collection("Student").Find(ctx, map[string]interface{}{})
	if err != nil {
		return []models.Student{}, err
	}

	cursor.All(ctx, &data)

	return data, nil
}

func (student *StudentController) GetAStudent(ctx context.Context, studentId string) (models.Student, error) {
	var data models.Student
	cursor := student.MongoClient.Database("university").Collection("Student").FindOne(ctx, bson.D{{Key: "enrollment", Value: studentId}})

	err := cursor.Decode(&data)
	if err != nil {
		return models.Student{}, nil
	}
	return data, nil
}

func (student *StudentController) EditStudent(ctx context.Context, studentData models.Student) error {
	result := student.MongoClient.Database("university").Collection("Student").FindOneAndReplace(ctx, bson.D{{Key: "enrollment", Value: studentData.Enrollment}}, studentData)

	return result.Err()
}
func (student *StudentController) DeleteStudent(ctx context.Context, enrollment string) error {
	result, err := student.MongoClient.Database("university").Collection("Student").DeleteOne(ctx, bson.D{{Key: "enrollment", Value: enrollment}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no id found to delete the data maybe data alredy deleted")
	}
	return nil
}
