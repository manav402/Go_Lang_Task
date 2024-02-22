package services

import (
	// "manav402/FiberMongo/models"

	"manav402/FiberMongo/controller"
	"manav402/FiberMongo/models"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/mongo"
)

type DepartmentServer interface {
	CreateADepartment(*fiber.Ctx) error
	GetAllDepartment(*fiber.Ctx) error
	GetADepartment(*fiber.Ctx) error
	EditDepartment(*fiber.Ctx) error
	DeleteDepartment(*fiber.Ctx) error
}

type DepartmentService struct {
	DepartmentController controller.DepartmentController
}

// create deparment will create a department from body parsed
func (subjctService *DepartmentService) CreateADepartment(ctx *fiber.Ctx) error {
	var departmentData models.Department
	err := ctx.BodyParser(&departmentData)
	if err != nil {
		return err
	}
	err = subjctService.DepartmentController.CretaeDepartment(ctx.Context(),departmentData)
	if err != nil {
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data inserted successfully",nil})

	return nil
}

// get all department will fetch all department data from collections
func (subjctService *DepartmentService) GetAllDepartment(ctx *fiber.Ctx) error {
	var departmentData []models.Department
	
	departmentData,err := subjctService.DepartmentController.GetAllDepartment(ctx.Context())
	if err != nil{
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data retrived succesfully",nil})
	ctx.JSON(departmentData)
	return nil
}

// get a department will fetch data from collection matching a document id
func (subjctService *DepartmentService) GetADepartment(ctx *fiber.Ctx) error {
	var departmentData models.Department
	departmentId := ctx.Params("departmentid","1")
	departmentData,err := subjctService.DepartmentController.GetOneDepartment(ctx.Context(),departmentId)
	if err != nil{
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data retrived succesfully",nil})
	ctx.JSON(departmentData)
	return nil
}

// edit department will replace old documetn with new one if id found in collection
func (subjctService *DepartmentService) EditDepartment(ctx *fiber.Ctx) error {
	var departmentData models.Department
	err := ctx.BodyParser(&departmentData)
	if err != nil {
		return err
	}
	err = subjctService.DepartmentController.EditDepartment(ctx.Context(),departmentData)
	if err != nil{
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data edited succesfully",nil})
	return nil
}

// delete department can delete the department documentation if id found in collections
func (subjctService *DepartmentService) DeleteDepartment(ctx *fiber.Ctx) error {
	departmentId := ctx.Params("departmentid","1")
	err := subjctService.DepartmentController.DeleteDepartment(ctx.Context(),departmentId)
	if err != nil{
		return err
	}

	ctx.JSON(OutputFormat{fiber.StatusOK,"data deleted succesfully",nil})
	return nil
}
