package server

import (
	"net/http"
)

type HealthStatus struct {
	Alive string `json:"alive"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	alive := HealthStatus{Alive: "ok"}

	respondWithJSON(w, http.StatusOK, alive)
}
