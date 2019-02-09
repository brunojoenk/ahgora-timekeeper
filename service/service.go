package service

import (
	"time"

	"github.com/apex/log"
	"github.com/rogerfernandes/ahgora-timekeeper/ahgora"
)

const (
	saturday = 6
	sunday   = 7
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
func (s *Service) PunchPoint() {
	if shouldPunchPoint() {
		resp, err := s.ahgoraClient.PunchPoint()
		if err != nil {
			log.WithError(err).Error("Service Error ")
		}

		if !resp.Result {
			log.Error("ResponseReason: " + resp.Reason)
		}
	}
}

func shouldPunchPoint() bool {
	weekday := time.Now().Weekday()
	return weekday != saturday && weekday != sunday
}
