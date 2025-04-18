package routes

import (
	"github.com/gorilla/mux"
	"gouth/controllers"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
}
