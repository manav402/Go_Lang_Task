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
// envMap is a private variable which will be used by getter and setter method
var envMap = make(map[string]string)

// setter method for envmap which will populate it with .env files data
func ReadEnv() error {
	var err error
	envMap, err = godotenv.Read(".env")
	return err
}

// getter for envmap which will return value stored in envMap
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


// database connection code will be invoked from here and will return a mongo cliend
func ConnectDB(ctx context.Context) (*mongo.Client, error) {
	// mongourl not found than localhost will be used
	mongoUrl, err := GetEnv("MONGOURL")
	if errors.Is(err, errors.New(("the key provided is not found in env file"))) {
		mongoUrl = "mongodb://localhost:27017"
	} else if err != nil {
		return nil, err
	} else {
		// replacing password with env file password
		password, err := GetEnv("PASSWORD")
		if err != nil {
			return nil, err
		}
		mongoUrl = strings.Replace(mongoUrl, "<password>", password, -1)
	}

	// connecting to mongodb
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
