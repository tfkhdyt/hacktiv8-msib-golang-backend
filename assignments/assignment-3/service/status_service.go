package service

import (
	"assignment_3/dto"
	"assignment_3/entity"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"
	"time"
)

type StatusService interface {
	GenerateStatus(min int, max int)
	ReadStatus() (*dto.StatusResponse, error)
}

type statusService struct {
	r *rand.Rand
}

func NewStatusService(r *rand.Rand) StatusService {
	return &statusService{r: r}
}

func (s *statusService) GenerateStatus(min int, max int) {
	for {
		status := entity.Status{
			Water: s.r.Intn(max-min) + min,
			Wind:  s.r.Intn(max-min) + min,
		}

		b, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			log.Fatalln("Failed to marshal status to json.", err.Error())
		}

		if err := os.WriteFile("data/status.json", b, 0644); err != nil {
			log.Fatalln("Failed to write json to file.", err.Error())
		}

		time.Sleep(15 * time.Second)
	}
}

func (s *statusService) ReadStatus() (*dto.StatusResponse, error) {
	b, err := os.ReadFile("data/status.json")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var status entity.Status
	if err := json.Unmarshal(b, &status); err != nil {
		return nil, errors.New(err.Error())
	}

	response := dto.StatusResponse{}
	response.ParseStatusDescription(status)

	return &response, err
}
