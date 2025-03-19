package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	protos "logger/protos/gen"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	protos.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *protos.LogRequest) (*protos.LogResponse, error) {
	// get the inputs
	input := req.GetLogEntry()

	// wirte the log to mongo
	entry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(entry)
	if err != nil {
		res := &protos.LogResponse{Result: "failed to do logging"}
		return res, err
	}
	res := &protos.LogResponse{Result: "logged data"}
	return res, nil
}

func (app *Config) gRPCListen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterLogServiceServer(s, &LogServer{Models: app.Models})
	log.Printf("starting grpc server oon port: %s", grpcPort)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}
}
