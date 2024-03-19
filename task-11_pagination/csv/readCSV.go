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

func ReadCSV(ctx context.Context, service *service.Service, fileName string) error {
	csvFile, err := os.OpenFile("./csv/"+fileName, os.O_RDONLY, os.ModePerm)
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("file not exist")
	} else if errors.Is(err, os.ErrPermission) {
		return errors.New("file exist but not have permission to read the file")
	} else if err != nil {
		return err
	}

	csvPtr := csv.NewReader(csvFile)
	csvPtr.ReuseRecord = true
	csvPtr.FieldsPerRecord = 5
	var buffer []string
	var model model.TrainModel
	// var err error
	var isFirst = true
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
