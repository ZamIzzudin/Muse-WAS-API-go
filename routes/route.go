// routes/routes.go
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DefaultMessage struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	Contact string `json:"contact"`
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	welcome := DefaultMessage{
			Message: "Welcome to Muse Dashboard API.",
			Author:  "Zamizzudin",
			Contact: "azzamizzudinhasan@gmail.com",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(welcome)
}

// SetupRoutes initializes all the routes for the application
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", DefaultRoute).Methods("GET")

	// Prefix
	UserRoutes(r)
	AuthRoutes(r)

	return r
}