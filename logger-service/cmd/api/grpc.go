package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "failed by grpc",
		}

		return res, err
	}

	// return response
	res := &logs.LogResponse{
		Result: "logged by grpc",
	}

	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Panicf("failed to listen for grpc: %v\n", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Printf("grpc server started on port %s\n", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Panicf("failed to serve for grpc: %v\n", err)
	}
}