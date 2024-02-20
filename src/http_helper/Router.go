package httphelper

import (
	routes "poj/src/routes"

	"github.com/gorilla/mux"
)

// Router returns a new mux.Router
// It contains each route of the application and its respective handler
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/health", routes.HealthCheck).Methods("GET")

	return router
}
