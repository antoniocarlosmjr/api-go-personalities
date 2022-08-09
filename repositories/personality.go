package repositories

import (
	"github.com/antoniocarlosmjr/api-go-personalities/database"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
)

func GetAllPersonalities() []models.Personality {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	return personalities
}

func GetPersonalityById(id string) models.Personality {
	var personality models.Personality
	database.DB.First(&personality, id)
	return personality
}

func CreatePersonality(newPersonality models.Personality) models.Personality {
	database.DB.Create(&newPersonality)
	return newPersonality
}

func UpdatePersonality(personality models.Personality) models.Personality {
	database.DB.Save(&personality)
	return personality
}

func DeletePersonality(id string) bool {
	var personality models.Personality
	database.DB.Delete(&personality, id)
	return true
}
