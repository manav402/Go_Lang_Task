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

var m map[string]string

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

func Setenv(key string, value string) error {
	if len(m) == 0 {
		return errors.New("env file is not yet parsed")
	}
	m[key] = value
	return nil
}

func ParseEnv() error {
	var err error
	m, err = godotenv.Read(".env")
	return err
}

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	mongoURL, err := Getenv("MONGOURL")
	if err == nil {
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
		mongoURL = "mongodb://localhost:27017/"
		log.Println("connecting to local mongodb")
	}
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if mongoClient == nil {
		return nil,errors.New("why mongoclient is nill")
	}
	err = mongoClient.Ping(ctx, nil)
	if  err != nil {
		return nil,err
	}
	log.Println("connected with database")

	return mongoClient.Database("TrainDatabase"), err
}
