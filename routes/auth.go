package routes

import (
	"muse-dashboard-api/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRoutes := r.PathPrefix("/auth").Subrouter()

	authRoutes.HandleFunc("/login", controllers.Login).Methods("POST")
	authRoutes.HandleFunc("/refresh-token", controllers.RefreshToken).Methods("POST")
}