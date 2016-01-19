package routers

import (
	"goBill/controllers"

	"github.com/gorilla/mux"
)

// InitHealth - initialise the routes for the Health Controller
func InitHealth(router *mux.Router) *mux.Router {
	router.HandleFunc("/health/{service}", controllers.GetHealth).Methods("GET")
	return router
}
