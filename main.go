// main.go
package main

import (
	"fmt"
	"log"
	"muse-dashboard-api/config"
	"muse-dashboard-api/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to the database
    config.ConnectDB()

    // Setup routes
    r := routes.SetupRoutes()

    // Get server port from environment variables
    port := os.Getenv("SERVER_PORT")

    // Start the server
    fmt.Printf("Server started at : %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}