package main

import (
	"context"
	"log"
	"logger/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(p RPCPayload, ag *string) error {
	collection := mongoClient.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      p.Name,
		Data:      p.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo")
		return err
	}

	*ag = "processed payload via RPC:" + p.Name
	return nil
}
