package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	wbePort  = "80"
	rpcPort  = "5001"
	grpcPort = "50001"
	mongoURL = "mongodb://mongo:27017"
)

var mongoClient *mongo.Client

type Config struct {
	Models data.Models
}

func main() {

	// create a mongo client
	mongo, err := connectMongo()
	if err != nil {
		log.Panic(err)
	}
	mongoClient = mongo

	// create a context to disconnect form mongo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(mongoClient),
	}

}

func (app *Config) Serve() {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", wbePort),
	}
}

func connectMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting to mongo:", err)
		return nil, err
	}

	return conn, nil

}
