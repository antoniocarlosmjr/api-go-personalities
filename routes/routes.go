package routes

import (
	"github.com/antoniocarlosmjr/api-go-personalities/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
