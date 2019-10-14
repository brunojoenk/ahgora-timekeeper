package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/apex/log"
	"github.com/brunojoenk/ahgora-timekeeper/ahgora"
	"github.com/brunojoenk/ahgora-timekeeper/config"
	"github.com/brunojoenk/ahgora-timekeeper/scheduler"
	"github.com/brunojoenk/ahgora-timekeeper/service"
)

func main() {
	cfg := config.MustGet()

	log.SetLevel(log.MustParseLevel(strings.ToLower(cfg.LogLevel)))
	log.Info("initializing")

	if cfg.AhgoraMockServerEnabled {
		go ahgora.StartMockServer()
		cfg.AhgoraURL = "http://localhost:8081"
	}

	start(cfg)

	router := mux.NewRouter()
	router.Path("/status").Methods(http.MethodGet).HandlerFunc(status)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	log.WithField("address", server.Addr).Info("Timekeeper Server - Started")

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("Timekeeper Server - starting failed")
	}
}

func start(cfg config.Config) {
	ahgoraCfg := ahgora.Config{
		Account:   cfg.Account,
		AhgoraURL: cfg.AhgoraURL,
		Identity:  cfg.Identity,
		Password:  cfg.Password,
	}
	ahgoraClient := ahgora.New(ahgoraCfg)
	srvc := service.New(ahgoraClient)
	schdlr := scheduler.New(srvc, cfg.CronTimes)
	schdlr.Start()

	//heroku.CronHeroku(cfg.HerokuAppURL)
}

func status(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK at " + time.Now().Format(time.Stamp)))
	if err != nil {
		return
	}
}
