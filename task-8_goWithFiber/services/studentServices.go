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

type StudentService struct {
	StudentController controller.StudentController
}


func (subjctService *StudentService) CreateAStudent(ctx *fiber.Ctx) error {
	var Data models.Student
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	err = subjctService.StudentController.CreateAstudent(ctx.Context(),Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data inserted successfully",nil})

	return nil
}
func (subjctService *StudentService) GetAllStudent(ctx *fiber.Ctx) error {
	var Data []models.Student
	// err := ctx.BodyParser(&Data)
	// if err != nil {
	// 	return err
	// }
	Data,err := subjctService.StudentController.GetAllStudents(ctx.Context())
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}
func (subjctService *StudentService) GetAStudent(ctx *fiber.Ctx) error {
	var Data models.Student
var id = ctx.Params("studentid","1")
	// err := ctx.BodyParser(&Data)
	// if err != nil {
	// 	return err
	// }
	Data,err := subjctService.StudentController.GetAStudent(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}
func (subjctService *StudentService) EditStudent(ctx *fiber.Ctx) error {
	var Data models.Student
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	err = subjctService.StudentController.EditStudent(ctx.Context(), Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Edited successfully", nil})

	return nil
}
func (subjctService *StudentService) DeleteStudent(ctx *fiber.Ctx) error {
		// var Data models.Subjects
	// err := ctx.BodyParser(&Data)
	// if err != nil {
		// return err
	// }
	id := ctx.Params("studentId","1")
	err := subjctService.StudentController.DeleteStudent(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Deleted successfully", nil})

	return nil
}
