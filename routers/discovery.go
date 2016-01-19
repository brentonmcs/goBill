package routers

import (
	"goBill/controllers"

	"github.com/gorilla/mux"
)

//InitDiscovery - initialises the Discovery routes
func InitDiscovery(router *mux.Router) *mux.Router {
	router.HandleFunc("/discovery", controllers.DiscoveryRegister).Methods("POST")
	return router
}
