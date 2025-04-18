package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gouth/routes" // Update this if your module name is different
)

func respond() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "All GOOD!")
		if err != nil {
			return
		}
	})
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found (skipping)")
	}

	// Set default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	// Initialize router
	router := mux.NewRouter()

	// Root health check route
	router.Handle("/", respond()).Methods("GET")

	// Register /signup route
	routes.RegisterAuthRoutes(router)

	// Start server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
