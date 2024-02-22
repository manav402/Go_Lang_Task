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

// create subject method creats a collection of subject as well append the subject to department collection
func (subjctService *SubjectService) CreateASubject(ctx *fiber.Ctx) error {
	var Data models.Subjects
	err := ctx.BodyParser(&Data)
	if err != nil {
		return err
	}

	// create a new subject and append the same data to department section as well
	err = subjctService.SubjectController.CreateAsubject(ctx.Context(), Data)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data inserted successfully", nil})

	return nil
}

// get all subject returns all available data from collections
func (subjctService *SubjectService) GetAllSubject(ctx *fiber.Ctx) error {
	var Data []models.Subjects

	Data,err := subjctService.SubjectController.GetAllSubject(ctx.Context())
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}

// get a subject return a single subject data if the subject id from params found
func (subjctService *SubjectService) GetASubject(ctx *fiber.Ctx) error {
	var Data models.Subjects
	// getting subject id from params
	var id = ctx.Params("subjectid")
	Data,err := subjctService.SubjectController.GetASubject(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Retrived successfully", nil})
	ctx.JSON(Data)
	return nil
}

// edit method for subject will replace the old data with whole new data with same primary key
func (subjctService *SubjectService) EditSubject(ctx *fiber.Ctx) error {
	var Data models.Subjects
	//  finding the data with parsed from body
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

// delete method for subject can delete a subject if params id found in collection
func (subjctService *SubjectService) DeleteSubject(ctx *fiber.Ctx) error {
	// getting params from url usign params method
	id := ctx.Params("subjectid")
	err := subjctService.SubjectController.DeleteSubject(ctx.Context(), id)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK, "data Deleted successfully", nil})

	return nil
}
