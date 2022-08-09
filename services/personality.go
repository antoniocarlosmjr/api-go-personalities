package services

import (
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/repositories"
)

func GetAllPersonalities() []models.Personality {
	return repositories.GetAllPersonalities()
}

func GetPersonalityById(id string) models.Personality {
	return repositories.GetPersonalityById(id)
}

func DeletePersonality(id string) bool {
	return repositories.DeletePersonality(id)
}

func CreatePersonality(newPersonality models.Personality) models.Personality {
	return repositories.CreatePersonality(newPersonality)
}

func UpdatePersonality(personality models.Personality) models.Personality {
	return repositories.UpdatePersonality(personality)
}
