package services

import (
	"context"
	"manav402/FiberMongo/controller"
	"manav402/FiberMongo/database"
)

type Services struct {
	DepartmentServer
	StudentServer
	SubjectServer
}

type OutputFormat struct{
	Code    int 		`json:"code"`
	Message string		`json:"message"`
	Error   error		`json:"error"`
}

func InitServices(ctx context.Context) (*Services, error) {
	// to refactor it
	// init database here please
	mongoClient,err := database.ConnectDB(ctx)
	if err != nil {
		return nil,err
	}
	api,err := controller.InitApi(mongoClient) 
	if err != nil {
		return nil, err
	}
	return &Services{DepartmentServer:&DepartmentService{api.DepartmentController}, StudentServer: &StudentService{api.StudentController}, SubjectServer: &SubjectService{api.SubjectController}}, nil
}
