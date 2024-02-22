package router

import (
	"context"
	"manav402/FiberMongo/services"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(ctx context.Context, app *fiber.App) error {
	services, err := services.InitServices(ctx)
	if err != nil {
		return err
	}

	// router for Students
	studentRouter := app.Group("/student")
	studentRouter.Post("/create", services.StudentServer.CreateAStudent)
	studentRouter.Get("/getAll", services.StudentServer.GetAllStudent)
	studentRouter.Get("/getOne/{studentid}", services.StudentServer.GetAStudent)
	studentRouter.Put("/editOne/{studentid}", services.StudentServer.EditStudent)
	studentRouter.Delete("/delete/{studentid}", services.StudentServer.DeleteStudent)

	// router for Department
	departmentRouter := app.Group("/department")
	departmentRouter.Post("/create", services.DepartmentServer.CreateADepartment)
	departmentRouter.Get("/getAll", services.DepartmentServer.GetAllDepartment)
	departmentRouter.Get("/getOne/{departmentid}", services.DepartmentServer.GetADepartment)
	departmentRouter.Put("/editOne/{departmentid}", services.DepartmentServer.EditDepartment)
	departmentRouter.Delete("/delete/{departmentid}", services.DepartmentServer.DeleteDepartment)

	// router for Subject
	subjectRouter := app.Group("/subject")
	subjectRouter.Post("/create", services.SubjectServer.CreateASubject)
	subjectRouter.Get("/getAll", services.SubjectServer.GetAllSubject)
	subjectRouter.Get("/getOne/{subjectid}", services.SubjectServer.GetASubject)
	subjectRouter.Put("/editOne/{subjectid}", services.SubjectServer.EditSubject)
	subjectRouter.Delete("/delete/{subjectid}", services.SubjectServer.DeleteSubject)

	return nil
}
