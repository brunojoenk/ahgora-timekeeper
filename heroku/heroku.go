package heroku

import (
	"net/http"

	"github.com/jasonlvhit/gocron"
)

//CronHeroku - calls HerokuApp
func CronHeroku() {
	hcron := gocron.NewScheduler()
	hcron.Every(28).Minutes().Do(pingHerokuApp)
	hcron.Start()
}

func pingHerokuApp() {
	http.Get("ahgora-timekeeper.herokuapp.com/status")
}
