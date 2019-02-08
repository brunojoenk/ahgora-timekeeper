package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/apex/log"
	"github.com/rogerfernandes/ahgora-timekeeper/ahgora"
	"github.com/rogerfernandes/ahgora-timekeeper/config"
)

func main() {
	cfg := config.MustGet()

	if cfg.AhgoraMockServerEnable {
		go ahgora.StartMockServer()
		cfg.AhgoraURL = "http://localhost:8081"
	}

	ahgoraCfg := ahgora.Config{
		AhgoraURL: cfg.AhgoraURL,
		Account:   cfg.Account,
		Identity:  cfg.Identity,
		Password:  cfg.Password,
	}

	ahgoraClient, err := ahgora.New(ahgoraCfg)
	if err != nil {
		return
	}

	punchPoint(ahgoraClient)

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

func status(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func punchPoint(ahgoraClient *ahgora.Client) {
	response, err := ahgoraClient.PunchPoint()
	if err != nil {
		log.WithError(err).Error("Erro ao Bater Ponto")
	}
	if !response.Result {
		log.Error("Erro ao bater ponto, motivo: " + response.Reason)
	} else {
		log.Info("Timekeeper Server - response: " + printResponse(response))
	}
}

func printResponse(r *ahgora.PunchResponse) string {
	out, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(out)
}
