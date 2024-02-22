package controller

import (
	"context"
	"errors"
	"log"
	"manav402/FiberMongo/models"

	"go.mongodb.org/mongo-driver/bson"
)

// create studene will populates the student collections with the req body data
func (student *StudentController) CreateAstudent(ctx context.Context, studentData models.Student) error {
	log.Println("got student data", studentData)
	// populating student data with department id in it
	departmentIdResult := student.MongoClient.Database("university").Collection("Department").FindOne(ctx, bson.D{{Key: "name", Value: studentData.DepartmentName}})
	var departmentId models.Department
	err := departmentIdResult.Decode(&departmentId)
	if err != nil {
		return err
	}

	studentData.DepartmentId = departmentId.DepartmentId
	log.Println("set department id to studentdat ",studentData)
	_, err = student.MongoClient.Database("university").Collection("Student").InsertOne(ctx, studentData)
	if err != nil {
		log.Println("err",err)
		return err
	}
	log.Println("data is insrted")
	return nil
}

// get all student will return all document from student collections
func (student *StudentController) GetAllStudents(ctx context.Context) ([]models.Student, error) {
	var data []models.Student
	cursor, err := student.MongoClient.Database("university").Collection("Student").Find(ctx, map[string]interface{}{})
	if err != nil {
		return []models.Student{}, err
	}

	cursor.All(ctx, &data)

	return data, nil
}

// get a student will return a document metching id from collections
func (student *StudentController) GetAStudent(ctx context.Context, studentId string) (models.Student, error) {
	var data models.Student
	cursor := student.MongoClient.Database("university").Collection("Student").FindOne(ctx, bson.D{{Key: "enrollment", Value: studentId}})

	err := cursor.Decode(&data)
	if err != nil {
		return models.Student{}, nil
	}
	return data, nil
}

// edit student record will replace the data matching id
func (student *StudentController) EditStudent(ctx context.Context, studentData models.Student) error {
	result := student.MongoClient.Database("university").Collection("Student").FindOneAndReplace(ctx, bson.D{{Key: "enrollment", Value: studentData.Enrollment}}, studentData)

	return result.Err()
}

// delete student record will delete documents matching ids
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
