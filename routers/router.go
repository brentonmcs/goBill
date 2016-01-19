package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes does stuff
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = InitHealth(router)
	router = InitDiscovery(router)
	return router
}
