package services

import (
	"SportsClubs/internal/models"
	"SportsClubs/internal/repositories"
	"errors"

	"gorm.io/gorm"
)

type TownService struct {
	repo *repositories.TownRepository
}

func NewTownService(repo *repositories.TownRepository) *TownService {
	return &TownService{repo: repo}
}

func (s *TownService) GetTowns(search string, page int, limit int) ([]map[string]interface{}, int64, error) {
	towns, total, err := s.repo.GetAll(search, page, limit)
	if err != nil {
		return nil, 0, err
	}
	var response []map[string]interface{}
	for _, town := range towns {
		response = append(response, map[string]interface{}{
			"id":   town.ID,
			"name": town.Name,
		})
	}
	return response, total, nil
}

func (s *TownService) GetClubsByTownName(search string, page int, limit int) ([]map[string]interface{}, int64, error) {
	if search == "" {
		return nil, 0, gorm.ErrInvalidData
	}
	clubs, total, err := s.repo.GetClubsByTownName(search, page, limit)
	if err != nil {
		return nil, 0, err
	}
	var response []map[string]interface{}
	for _, club := range clubs {
		response = append(response, map[string]interface{}{
			"id":                  club.ID,
			"name":                club.Name,
			"quantity_tituls":     club.QuantityTituls,
			"average_age_players": club.AverageAgePlayers,
			"town": map[string]interface{}{
				"name": club.Town.Name,
			},
		})
	}
	return response, total, nil
}

func (s *TownService) CreateTown(town *models.Town) error {
	err := s.repo.CreateTown(town)
	if err != nil {
		return err
	}
	return nil
}

func (s *TownService) UpdateTown(id uint, updated *models.Town) (*models.Town, error) {
	town := &models.Town{
		ID:   id,
		Name: updated.Name,
	}
	err := s.repo.UpdateTown(town)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return town, nil
}

func (s *TownService) DeleteTown(id uint) error {
	err := s.repo.DeleteTown(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TownService) GetByID(id uint) (*models.Town, error) {
	var town models.Town
	if err := s.repo.DB.First(&town, id).Error; err != nil {
		return nil, err
	}
	return &town, nil
}
