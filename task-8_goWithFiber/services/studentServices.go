package services

import (
	// "manav402/FiberMongo/models"

	"manav402/FiberMongo/controller"
	"manav402/FiberMongo/models"

	"github.com/gofiber/fiber/v2"
)

type StudentServer interface {
	CreateAStudent(*fiber.Ctx) error
	GetAllStudent(*fiber.Ctx) error
	GetAStudent(*fiber.Ctx) error
	EditStudent(*fiber.Ctx) error
	DeleteStudent(*fiber.Ctx) error
}

// student service struct which will implement the above interface
type StudentService struct {
	StudentController controller.StudentController
}

// create a student function gather body data from client and pass to controller function
func (subjctService *StudentService) CreateAStudent(ctx *fiber.Ctx) error {
	var Data models.Student
	// parsing the body data from body
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	// sending data to controller functions
	err = subjctService.StudentController.CreateAstudent(ctx.Context(),Data)
	if err != nil {
		return err
	}

	// if every thing is ok than sending back the response to client
	ctx.JSON(OutputFormat{fiber.StatusOK,"data inserted successfully",nil})
	return nil
}

// getallstudent method can return all student found in collections
func (subjctService *StudentService) GetAllStudent(ctx *fiber.Ctx) error {
	var Data []models.Student
	// finding all students from the collections
	Data,err := subjctService.StudentController.GetAllStudents(ctx.Context())
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}

// GETone student function can return a single student based on the id from url params
func (subjctService *StudentService) GetAStudent(ctx *fiber.Ctx) error {
	var Data models.Student
	var id = ctx.Params("studentid")
	// if collection has data with the params id we will return it
	Data,err := subjctService.StudentController.GetAStudent(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}

// editastudent  method can reinsert the data with same id if the id found in collections
func (subjctService *StudentService) EditStudent(ctx *fiber.Ctx) error {
	var Data models.Student

	// parsing the data from body
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	// reinserting data while updating the value in edit student method
	err = subjctService.StudentController.EditStudent(ctx.Context(), Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Edited successfully", nil})

	return nil
}

// delete student method delets the student data if the id is found in collections
func (subjctService *StudentService) DeleteStudent(ctx *fiber.Ctx) error {

	id := ctx.Params("studentId","1")
	err := subjctService.StudentController.DeleteStudent(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Deleted successfully", nil})

	return nil
}
