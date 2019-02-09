package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/rogerfernandes/ahgora-timekeeper/service"
)

//StartScheduler - Starts scheduling point punchers
func StartScheduler(service *service.Service) {
	gocron.Every(1).Day().At("08:00").Do(service.PunchPoint)
	gocron.Every(1).Day().At("12:00").Do(service.PunchPoint)
	gocron.Every(1).Day().At("13:30").Do(service.PunchPoint)
	gocron.Every(1).Day().At("18:00").Do(service.PunchPoint)

	//Test
	gocron.Every(2).Seconds().Do(service.PunchPoint)

	gocron.Start()
}
