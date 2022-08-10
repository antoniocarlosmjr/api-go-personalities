package services

import (
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/repositories"
)

type PersonalityService interface {
	GetAllPersonalities() []models.Personality
	GetPersonalityById(id string) (models.Personality, error)
	DeletePersonality(id string) bool
	CreatePersonality(newPersonality models.Personality) models.Personality
	UpdatePersonality(personality models.Personality) models.Personality
}

type personalityService struct {
	repository repositories.PersonalityRepository
}

func NewPersonalityService(repository repositories.PersonalityRepository) PersonalityService {
	return &personalityService{repository: repository}
}

func (p *personalityService) GetAllPersonalities() []models.Personality {
	return p.repository.GetAllPersonalities()
}

func (p *personalityService) GetPersonalityById(id string) (models.Personality, error) {
	return p.repository.GetPersonalityById(id)
}

func (p *personalityService) DeletePersonality(id string) bool {
	return p.repository.DeletePersonality(id)
}

func (p *personalityService) CreatePersonality(newPersonality models.Personality) models.Personality {
	return p.repository.CreatePersonality(newPersonality)
}

func (p *personalityService) UpdatePersonality(personality models.Personality) models.Personality {
	return p.repository.UpdatePersonality(personality)
}
