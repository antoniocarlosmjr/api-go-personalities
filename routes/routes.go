package routes

import (
	"github.com/antoniocarlosmjr/api-go-personalities/controllers"
	"github.com/antoniocarlosmjr/api-go-personalities/middleware"
	"github.com/antoniocarlosmjr/api-go-personalities/repositories"
	"github.com/antoniocarlosmjr/api-go-personalities/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	route := mux.NewRouter()
	route.Use(middleware.ContentTypeMiddleware)

	// Repositories
	repositoryPersonality := repositories.NewPersonalityRepository()

	// Services
	servicePersonality := services.NewPersonalityService(repositoryPersonality)

	// Controllers
	personalityController := controllers.NewPersonalityController(servicePersonality)

	route.HandleFunc("/", personalityController.Home)
	route.HandleFunc("/api/personalities", personalityController.GetPersonalities).Methods("GET")
	route.HandleFunc("/api/personalities/{id}", personalityController.GetPersonalityById).Methods("GET")
	route.HandleFunc("/api/personalities", personalityController.CreatePersonality).Methods("POST")
	route.HandleFunc("/api/personalities/{id}", personalityController.DeletePersonality).Methods("DELETE")
	route.HandleFunc("/api/personalities/{id}", personalityController.UpdatePersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(route)))
}
