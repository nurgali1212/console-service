package service

import (
	"console-service/model"
	"fmt"
	"log"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ObjectService(data model.Object) model.Object {

	log.Print(fmt.Sprintf(data.Message))
	return data
}
