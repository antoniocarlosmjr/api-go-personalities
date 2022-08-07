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

	log.Fatal(http.ListenAndServe(":8000", route))
}
