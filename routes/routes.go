package routes

import (
	"github.com/antoniocarlosmjr/api-go-personalities/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Home)
	route.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	route.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	route.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	route.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	route.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", route))
}
