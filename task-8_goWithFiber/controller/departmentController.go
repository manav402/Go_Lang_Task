package controller

import (
	"context"
	"errors"
	"manav402/FiberMongo/models"

	"go.mongodb.org/mongo-driver/bson"
)

//TODO: create buisness logics for department crud operations

func (department *DepartmentController) GetAllDepartment(ctx context.Context) ([]models.Department, error) {
	var data []models.Department
	cursor, err := department.MongoClient.Database("university").Collection("Department").Find(ctx, map[string]interface{}{})
	if err != nil {
		return []models.Department{}, err
	}

	cursor.All(ctx, &data)

	return data, nil
}

func (department *DepartmentController) GetOneDepartment(ctx context.Context, departmentId string) (models.Department, error) {
	var data models.Department
	cursor := department.MongoClient.Database("university").Collection("Department").FindOne(ctx, bson.D{{Key: "departmentid",Value: departmentId}})

	err := cursor.Decode(&data)
	if err != nil {
		return models.Department{},nil
	}
	return data, nil
}

func (department *DepartmentController) GetAllStudentFromDepartment(ctx context.Context, departmentId string) (map[string]interface{}, error) {
	return nil,nil
}

func (department *DepartmentController) GetAllSubjects(ctx context.Context, departmentId string) ([]models.Department, error) {
	return []models.Department{}, nil
}

func (department *DepartmentController) CretaeDepartment(ctx context.Context, departmentData models.Department) error {
	_,err := department.MongoClient.Database("university").Collection("Department").InsertOne(ctx,departmentData)
	if err != nil {
		return err
	}

	return nil

}

func (department *DepartmentController) EditDepartment(ctx context.Context, departmentData models.Department) error {
	result := department.MongoClient.Database("university").Collection("Department").FindOneAndReplace(ctx,bson.D{{Key:"departmentid",Value: departmentData.DepartmentId}},departmentData)
	
	return result.Err()
}

func (department *DepartmentController) DeleteDepartment(ctx context.Context, departmentid string) error {
	result,err := department.MongoClient.Database("university").Collection("Department").DeleteOne(ctx,bson.D{{Key:"departmentid",Value:departmentid}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no id found to delete the data maybe data alredy deleted")
	}
	return nil
}
