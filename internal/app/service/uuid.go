package service

import "github.com/google/uuid"

type Service struct{}

func (s Service) GenerateID() string {
	return uuid.New().String()
}
