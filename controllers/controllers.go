package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/services"
	"github.com/gorilla/mux"
	"net/http"
)

type PersonalityController interface {
	Home(w http.ResponseWriter, r *http.Request)
	GetPersonalities(w http.ResponseWriter, r *http.Request)
	GetPersonalityById(w http.ResponseWriter, r *http.Request)
	CreatePersonality(w http.ResponseWriter, r *http.Request)
	DeletePersonality(w http.ResponseWriter, r *http.Request)
	UpdatePersonality(w http.ResponseWriter, r *http.Request)
}

type personalityController struct {
	service services.PersonalityService
}

func NewPersonalityController(service services.PersonalityService) PersonalityController {
	return &personalityController{service: service}
}

func (p *personalityController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func (p *personalityController) GetPersonalities(w http.ResponseWriter, r *http.Request) {
	personalities := p.service.GetAllPersonalities()
	json.NewEncoder(w).Encode(personalities)
}

func (p *personalityController) GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	personality, err := p.service.GetPersonalityById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func (p *personalityController) CreatePersonality(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	newPersonality = p.service.CreatePersonality(newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func (p *personalityController) DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p.service.DeletePersonality(id)
	json.NewEncoder(w).Encode("Personality deleted")
}

func (p *personalityController) UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	personality, _ = p.service.GetPersonalityById(id)
	json.NewDecoder(r.Body).Decode(&personality)

	personality = p.service.UpdatePersonality(personality)
	json.NewEncoder(w).Encode(&personality)
}
