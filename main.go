package main

import (
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

	apontamento := ahgora.Apontamento{
		//Account:  "51eec6356c615d3edf39d497c137d75b",
		//Identity: "454",
		//Password: "454",
		//Timestamp:    time.Now(),
		//TimestampLoc: time.Now(),
	}

	err = ahgoraClient.BaterPonto(apontamento)
	if err != nil {

	}
}
