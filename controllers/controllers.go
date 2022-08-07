package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/antoniocarlosmjr/api-go-personalities/database"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

// GetPersonalities Todo: remove acces to database this and put in a repository, for example
func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

// GetPersonalityById todo: remove acces to database this and put in a repository, for example
func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
