package routes

import (
	"gouth/controllers"
	"net/http"
)

func RegisterRoutes() {
	// Root health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("All GOOD!"))
	})

	// Signup endpoint
	http.HandleFunc("/signup", controllers.SignUp)
}
