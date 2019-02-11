package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/apex/log"
	"github.com/rogerfernandes/ahgora-timekeeper/ahgora"
	"github.com/rogerfernandes/ahgora-timekeeper/config"
	"github.com/rogerfernandes/ahgora-timekeeper/scheduler"
	"github.com/rogerfernandes/ahgora-timekeeper/service"
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
		Addr:         ":8080",
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
	ahgoraClient, err := ahgora.New(ahgoraCfg)
	if err != nil {
		panic(err)
	}

	service := service.New(ahgoraClient)

	scheduler.StartScheduler(service)
}

func status(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK at " + time.Now().Format(time.Stamp)))
	if err != nil {
		return
	}
}
