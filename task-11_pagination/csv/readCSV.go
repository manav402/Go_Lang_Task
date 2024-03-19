package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	model "manav/pagination/models"
	"manav/pagination/service"
	"os"
	"strconv"
)
// transfer function transfer data string to train model
func transfer(buffer []string) (model.TrainModel, error) {

	if len(buffer) != 5 {
		return model.TrainModel{}, errors.New("buffer has not read all the data")
	}

	no, err := strconv.Atoi(buffer[0])
	if err != nil {
		return model.TrainModel{}, err
	}

	trainNo, err := strconv.Atoi(buffer[1])
	if err != nil {
		return model.TrainModel{}, err
	}

	var data = model.TrainModel{
		No:         no,
		Train_No:   trainNo,
		Train_Name: buffer[2],
		Starts:     buffer[3],
		Ends:       buffer[4],
	}
	return data, nil
}

// read csv function read the data from line by line and push the data on mongo cluster
// @params : context, service with mongo client, file name
func ReadCSV(ctx context.Context, service *service.Service, fileName string) error {

	csvFile, err := os.OpenFile("./csv/"+fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	csvPtr := csv.NewReader(csvFile)
	csvPtr.ReuseRecord = true
	csvPtr.FieldsPerRecord = 5

	var buffer []string
	var model model.TrainModel
	var isFirst = true
	
	// read the cursor untill eof and push the data in collection
	for !errors.Is(err, io.EOF) {

		buffer, err = csvPtr.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		} else if !isFirst && !errors.Is(err, io.EOF) {

			model, err = transfer(buffer)
			if err != nil {
				return err
			}

			err = service.InsertOne(ctx, model)
			if err != nil {
				return err
			}

		}

		if isFirst {
			isFirst = false
		}
		
	}
	return nil
}
