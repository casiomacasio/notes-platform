package server

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context, startFunc func(ctx context.Context) error) error {
	logrus.Info("Notification Server is starting...")
	return startFunc(ctx)
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Notification Server is shutting down...")
	return nil
}
