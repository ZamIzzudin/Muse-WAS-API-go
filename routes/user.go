// routes/user_routes.go
package routes

import (
	"muse-dashboard-api/controllers"

	"github.com/gorilla/mux"
)

// UserRoutes defines routes for the user prefix
func UserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("", controllers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	userRouter.HandleFunc("", controllers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
}