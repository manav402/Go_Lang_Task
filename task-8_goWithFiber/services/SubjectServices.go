package services

import (
	// "manav402/FiberMongo/models"

	"manav402/FiberMongo/controller"
	"manav402/FiberMongo/models"

	"github.com/gofiber/fiber/v2"
)

type SubjectServer interface {
	CreateASubject(*fiber.Ctx) error
	GetAllSubject(*fiber.Ctx) error
	GetASubject(*fiber.Ctx) error
	EditSubject(*fiber.Ctx) error
	DeleteSubject(*fiber.Ctx) error
}

type SubjectService struct {
	SubjectController controller.SubjectController
}

// func InitSubjectService()(SubjectServer,error){
// 	return &SubjectService{},nil
// }

func (subjctService *SubjectService) CreateASubject(ctx *fiber.Ctx) error {
	var Data models.Subjects
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	err = subjctService.SubjectController.CreateAsubject(ctx.Context(), Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data inserted successfully", nil})

	return nil
}

func (subjctService *SubjectService) GetAllSubject(ctx *fiber.Ctx) error {
	var Data []models.Subjects
	// err := ctx.BodyParser(&Data)
	// if err != nil {
	// 	return err
	// }
	Data,err := subjctService.SubjectController.GetAllSubject(ctx.Context())
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}
func (subjctService *SubjectService) GetASubject(ctx *fiber.Ctx) error {
var Data models.Subjects
var id = ctx.Params("subjectId","1")
	// err := ctx.BodyParser(&Data)
	// if err != nil {
	// 	return err
	// }
	Data,err := subjctService.SubjectController.GetASubject(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}
func (subjctService *SubjectService) EditSubject(ctx *fiber.Ctx) error {
	var Data models.Subjects
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}
	err = subjctService.SubjectController.EditSubject(ctx.Context(), Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Edited successfully", nil})

	return nil
}
func (subjctService *SubjectService) DeleteSubject(ctx *fiber.Ctx) error {
	// var Data models.Subjects
	// err := ctx.BodyParser(&Data)
	// if err != nil {
		// return err
	// }
	id := ctx.Params("subjectId","1")
	err := subjctService.SubjectController.DeleteSubject(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Deleted successfully", nil})

	return nil
}
