package repositories

import (
	"errors"
	"github.com/antoniocarlosmjr/api-go-personalities/database"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
)

type PersonalityRepository interface {
	GetAllPersonalities() []models.Personality
	GetPersonalityById(id string) (models.Personality, error)
	CreatePersonality(newPersonality models.Personality) models.Personality
	UpdatePersonality(personality models.Personality) models.Personality
	DeletePersonality(id string) bool
}

type personalityRepository struct{}

func NewPersonalityRepository() PersonalityRepository {
	return &personalityRepository{}
}

func (p *personalityRepository) GetAllPersonalities() []models.Personality {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	return personalities
}

func (p *personalityRepository) GetPersonalityById(id string) (models.Personality, error) {
	var personality models.Personality
	err := database.DB.First(&personality, id).Error
	if err != nil {
		return models.Personality{}, errors.New("personality not found")
	}
	return personality, nil
}

func (p *personalityRepository) CreatePersonality(newPersonality models.Personality) models.Personality {
	database.DB.Create(&newPersonality)
	return newPersonality
}

func (p *personalityRepository) UpdatePersonality(personality models.Personality) models.Personality {
	database.DB.Save(&personality)
	return personality
}

func (p *personalityRepository) DeletePersonality(id string) bool {
	var personality models.Personality
	database.DB.Delete(&personality, id)
	return true
}
