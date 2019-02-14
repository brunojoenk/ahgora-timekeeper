package randTime

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/apex/log"
)

//RandomizeTime - Randomize string time (08:00) with 10 minutes range
func RandomizeTime(time string) string {
	hour, minute := splitTime(time)
	minute = randomize(minute)
	if minute > 60 {
		minute -= 60
		hour++
	}
	return stringify(hour) + ":" + stringify(minute)
}

func splitTime(time string) (int, int) {
	sTime := strings.Split(time, ":")
	hour, err := strconv.Atoi(sTime[0])
	if err != nil {
		log.WithError(err).Error("Error when try to convert hour string to int")
		hour = 0
	}
	minute, err := strconv.Atoi(sTime[1])
	if err != nil {
		log.WithError(err).Error("Error when try to convert minute string to int")
		minute = 0
	}
	return hour, minute
}

func randomize(minute int) int {
	rand.Seed(time.Now().UnixNano())
	delta := (minute + 10) - minute
	return rand.Intn(delta) + minute
}

func stringify(value int) string {
	sValue := strconv.Itoa(value)
	if len(sValue) == 1 {
		return "0" + sValue
	}
	return sValue
}
