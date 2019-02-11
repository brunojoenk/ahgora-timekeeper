package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/rogerfernandes/ahgora-timekeeper/service"
)

//StartScheduler - Starts scheduling point punchers
func StartScheduler(service *service.Service) {
	gocron.Every(1).Day().At("08:20").Do(service.PunchPoint)
	gocron.Every(1).Day().At("11:50").Do(service.PunchPoint)
	gocron.Every(1).Day().At("13:00").Do(service.PunchPoint)
	gocron.Every(1).Day().At("18:13").Do(service.PunchPoint)
	gocron.Start()
}
