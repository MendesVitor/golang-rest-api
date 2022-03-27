package routes

import (
	"log"
	"net/http"
	"rest-api/controllers"
	"rest-api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.Use(middleware.SetContentTypeMiddleware)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
