package heroku

import (
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
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
	resp, err := http.Get(herokuAppURL)
	if err != nil {
		log.WithError(err).Error("WakeUp HerokuApp! Error")
	}
	log.Debug("WakeUp HerokuApp! Response: " + getBodyContent(resp))
}

func getBodyContent(resp *http.Response) string {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(content)
}
