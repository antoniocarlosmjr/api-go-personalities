package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/services"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	personalities := services.GetAllPersonalities()
	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	personality, err := services.GetPersonalityById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	newPersonality = services.CreatePersonality(newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	services.DeletePersonality(id)
	json.NewEncoder(w).Encode("Personality deleted")
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	personality, _ = services.GetPersonalityById(id)
	json.NewDecoder(r.Body).Decode(&personality)

	personality = services.UpdatePersonality(personality)
	json.NewEncoder(w).Encode(&personality)
}
