package heroku

import (
	"net/http"

	"github.com/jasonlvhit/gocron"
)

var herokuAppURL string

//CronHeroku - calls HerokuApp
func CronHeroku(url string) {
	herokuAppURL = url
	hcron := gocron.NewScheduler()
	hcron.Every(28).Minutes().Do(pingHerokuApp)
	hcron.Start()
}

func pingHerokuApp() {
	http.Get(herokuAppURL)
}
