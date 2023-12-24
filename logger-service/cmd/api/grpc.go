package main

import (
	"context"

	"github.com/Morpa/microservices-go/logger-service/data"
	"github.com/Morpa/microservices-go/logger-service/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Model data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetEntry()

	// write the WriteLog
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Model.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}
