package database

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// sync.map can be used for thread safe searching on large application
var m map[string]string

// getter method to read env data
func Getenv(key string) (string, error) {
	if m == nil {
		err := ParseEnv()
		if err != nil {
			return "", err
		}
	}

	if _, ok := m[key]; !ok {
		return "", errors.New("the key not present in env file yet do your forgot to read the env file ?")
	}

	return m[key], nil
}

// run time setter method for env variable
func Setenv(key string, value string) error {
	if len(m) == 0 {
		return errors.New("env file is not yet parsed")
	}

	m[key] = value
	return nil
}

// parse env function parse the env file one time on server startup
func ParseEnv() error {
	var err error
	m, err = godotenv.Read(".env")
	return err
}

// connectdb based on env value connect to either global or local mongo
func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	// get hosted mongo url
	mongoURL, err := Getenv("MONGOURL")
	if err == nil {
		// if exist replace username and password
		user, err := Getenv("USERNAME")
		if err != nil {
			return nil, err
		}
		password, err := Getenv("PASSWORD")
		if err != nil {
			return nil, err
		}
		mongoURL = strings.Replace(mongoURL, "<username>", user, 1)
		mongoURL = strings.Replace(mongoURL, "<password>", password, 1)
	} else {
		// if not found connect to local mongodb
		mongoURL = "mongodb://localhost:27017/"
		log.Println("connecting to local mongodb")
	}

	// connecting with mongo client
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if mongoClient == nil {
		return nil,errors.New("why mongoclient is nill")
	}
	if err != nil {
		return nil,err
	}

	// if cant ping return error
	err = mongoClient.Ping(ctx, nil)
	if  err != nil {
		return nil,err
	}
	log.Println("connected with database")

	return mongoClient.Database("TrainDatabase"), err
}
