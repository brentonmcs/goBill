package controllers

import (
	"encoding/json"
	"goBill/models"
	"goBill/util/redis"
	"net/http"
	"time"
)

//DiscoveryRegister - registers a new service
func DiscoveryRegister(w http.ResponseWriter, r *http.Request) {

	requestService := new(models.RegisterService)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestService)

	if len(requestService.BaseURI) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Base URI is missing"))
		return
	}
	if len(requestService.ServiceName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Service Name is missing"))
		return
	}

	marshelledService, _ := json.Marshal(requestService)

	client := redis.NewRedis()

	client.Set("Bill-Service-LastUpdated", requestService.ServiceName, time.Minute)
	client.HSet("Bill-Service", requestService.ServiceName, string(marshelledService))
	w.WriteHeader(http.StatusOK)
}
