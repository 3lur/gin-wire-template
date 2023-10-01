package service

import "log"

type Service struct {
	logger *log.Logger
}

// NewService creates a new Service
func NewService(logger *log.Logger) *Service {
	return &Service{logger}
}
