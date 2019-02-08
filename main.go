package main

import (
	"encoding/json"

	"github.com/apex/log"
	"github.com/rogerfernandes/ahgora-timekeeper/ahgora"
	"github.com/rogerfernandes/ahgora-timekeeper/config"
)

func main() {
	cfg := config.MustGet()

	ahgoraCfg := ahgora.Config{
		Account:  cfg.Account,
		Identity: cfg.Identity,
		Password: cfg.Password,
	}

	ahgoraClient, err := ahgora.New(ahgoraCfg)
	if err != nil {
		return
	}

	response, err := ahgoraClient.PunchPoint()
	if err != nil {
		log.WithError(err).Error("Erro ao Bater Ponto")
	}
	if !response.Result {
		log.Error("Erro ao bater ponto, motivo: " + response.Reason)
	} else {
		log.Info(printResponse(response))
	}
}

func printResponse(r *ahgora.PunchResponse) string {
	out, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(out)
}
