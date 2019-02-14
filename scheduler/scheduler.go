package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/rogerfernandes/ahgora-timekeeper/scheduler/randTime"
	"github.com/rogerfernandes/ahgora-timekeeper/service"
)

//Scheduler - Struct
type Scheduler struct {
	Scheduler *gocron.Scheduler
	Service   *service.Service
	CronTimes []string
}

//New - Creates a new Scheduler
func New(service *service.Service, cronTimes []string) *Scheduler {
	return &Scheduler{
		Scheduler: gocron.NewScheduler(),
		Service:   service,
		CronTimes: cronTimes,
	}
}

//StartScheduler - Starts scheduling point punchers
func (s *Scheduler) StartScheduler() {
	s.schedule()
	s.Scheduler.Start()
	s.rescheduler()
}

func (s *Scheduler) schedule() {
	for i := range s.CronTimes {
		s.Scheduler.Every(1).Day().At(randTime.RandomizeTime(s.CronTimes[i])).Do(s.Service.PunchPoint)
	}
}

func (s *Scheduler) reschedule() {
	s.Scheduler.Clear()
	s.schedule()
}

func (s *Scheduler) rescheduler() {
	gocron.Every(1).Day().At("00:00").Do(s.reschedule)
	gocron.Start()
}
