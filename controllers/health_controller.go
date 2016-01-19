package controllers

import (
	"goBill/util/redis"
	"net/http"

	"github.com/gorilla/mux"
)

//GetHealth - returns the health for the selected service
func GetHealth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(getResultFromCache(params["service"]))
}

func getResultFromCache(service string) []byte {
	result, err := redis.NewRedis().HGet("health", service).Bytes()
	if err != nil {
		panic(err)
	}
	return result
}
