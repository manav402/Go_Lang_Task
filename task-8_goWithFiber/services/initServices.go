package services

import (
	"context"
	"manav402/FiberMongo/controller"
	"manav402/FiberMongo/database"
)

// services which has all the service interface implemented
type Services struct {
	DepartmentServer
	StudentServer
	SubjectServer
}

// json formate to send back the data to client
type OutputFormat struct{
	Code    int 		`json:"code"`
	Message string		`json:"message"`
	Error   error		`json:"error"`
}

// initializing service struct after initializing every interface struct and database
func InitServices(ctx context.Context) (*Services, error) {

	// init database here please
	mongoClient,err := database.ConnectDB(ctx)
	if err != nil {
		return nil,err
	}

	// after database is initialized the whole controller api will be initialized
	api,err := controller.InitApi(mongoClient) 
	if err != nil {
		return nil, err
	}

	// if everything is right the service structure will be returned
	return &Services{DepartmentServer:&DepartmentService{api.DepartmentController}, StudentServer: &StudentService{api.StudentController}, SubjectServer: &SubjectService{api.SubjectController}}, nil
}
