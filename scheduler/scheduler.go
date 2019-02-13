package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/rogerfernandes/ahgora-timekeeper/service"
)

//StartScheduler - Starts scheduling point punchers
func StartScheduler(service *service.Service) {
	gocron.Every(1).Day().At("08:00").Do(service.PunchPoint)
	gocron.Every(1).Day().At("11:53").Do(service.PunchPoint)
	gocron.Every(1).Day().At("13:04").Do(service.PunchPoint)
	gocron.Every(1).Day().At("18:23").Do(service.PunchPoint)
	gocron.Start()
}
