package service

import (
	"github.com/apex/log"
	"github.com/rogerfernandes/ahgora-timekeeper/ahgora"
)

//Service - application service
type Service struct {
	ahgoraClient *ahgora.Client
}

//New - creates a new service
func New(client *ahgora.Client) *Service {
	return &Service{
		ahgoraClient: client,
	}
}

//PunchPoint - punches a point in ahgora app
func (s *Service) PunchPoint() error {
	_, err := s.ahgoraClient.PunchPoint()
	if err != nil {
		log.WithError(err).Error("Service Error ")
	}
	return err
}
