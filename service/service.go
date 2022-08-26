package service

import (
	"console-service/model"
	"fmt"
	"log"
	"strconv"
)

const layout = "2006-01-02"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ObjectService(data model.Object) model.Object {

	log.Print(fmt.Sprintf("%s %s %s", data.Message, strconv.Itoa(data.Code), data.Date))
	return data
}
