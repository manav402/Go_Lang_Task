package database

import (
	"context"
	"errors"
	"log"
	// "manav402/FiberMongo/models"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envMap = make(map[string]string)

func ReadEnv() error {
	var err error
	envMap, err = godotenv.Read(".env")
	return err
}

func GetEnv(key string) (string, error) {
	if len(envMap) == 0 {
		err := ReadEnv()
		if err != nil {
			return "", err
		}
	}
	if v, ok := envMap[key]; ok {
		return v, nil
	}
	return "", errors.New("the key provided is not found in env file")
}

func ConnectDB(ctx context.Context) (*mongo.Client, error) {

	mongoUrl, err := GetEnv("MONGOURL")
	if errors.Is(err, errors.New(("the key provided is not found in env file"))) {
		mongoUrl = "mongodb://localhost:27017"
	} else if err != nil {
		return nil, err
	} else {
		password, err := GetEnv("PASSWORD")
		if err != nil {
			return nil, err
		}
		mongoUrl = strings.Replace(mongoUrl, "<password>", password, -1)
	}

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}

	// create collections
	// // departmentOpt := options.CreateCollection().SetValidator(models.DepartmentSchema)
	// err = mongoClient.Database("university").CreateCollection(ctx, "Department")
	// if err != nil {
	// 	log.Println(err)
	// 	return nil,err
	// }

	// studentOpt := options.CreateCollection().SetValidator(models.StudentSchema)
	// err = mongoClient.Database("university").CreateCollection(ctx, "Student")
	// if err != nil {
	// 	log.Println(err)
	// 	// return nil,err
	// }

	// subjectOpt := options.CreateCollection().SetValidator(models.SubjectSchema)
	// err = mongoClient.Database("university").CreateCollection(ctx, "Subject")
	// if err != nil {
	// 	log.Println(err)
	// 	// return nil,err
	// }	

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Println(err)

		// return nil, err
	}
	log.Println("yaay database connected!!!")

	return mongoClient, nil
}
