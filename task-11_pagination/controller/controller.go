package controller

import (
	"context"
	"errors"
	"manav/pagination/service"
)

type Controller struct{
	Service	*service.Service
}

// inticontroller call service creator and return a conntroller struct
func NewController(ctx context.Context)(*Controller,error){
	service,err := service.NewService(ctx)
	if err != nil {
		return nil,err
	}
	
	if service == nil {
		return nil,errors.New("please initialize the services first")
	}

	return &Controller{Service: service},nil
}