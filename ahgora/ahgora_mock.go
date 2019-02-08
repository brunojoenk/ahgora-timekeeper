package ahgora

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

//StartMockServer - Is a mock server to local tests
func StartMockServer() {
	router := mux.NewRouter()
	ahgoraRouter := router.PathPrefix("/batidaonline").Subrouter()
	ahgoraRouter.Path("/status").Methods(http.MethodGet).HandlerFunc(status)
	ahgoraRouter.Path("/verifyIdentification").Methods(http.MethodPost).HandlerFunc(punch)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.WithField("address", server.Addr).Info("Ahgora Mock Server - Started")

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("Ahgora Mock Server - starting failed")
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func punch(w http.ResponseWriter, r *http.Request) {
	log.Info("Ahgora Mock Server - Punching point!")
	resp := &PunchResponse{
		Day:     "2019-01-02",
		Name:    "CRONOS",
		Punches: []string{"0800", "1200", "1330"},
		Reason:  "",
		Result:  true,
		Time:    "1330",
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
