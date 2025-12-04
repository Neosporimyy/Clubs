package services

import (
	"SportsClubs/internal/models"
	"SportsClubs/internal/repositories"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ClubService struct {
	repo *repositories.ClubRepository
}

func NewClubService(repo *repositories.ClubRepository) *ClubService {
	return &ClubService{repo: repo}
}

func (s *ClubService) GetAll(page, limit int) ([]map[string]interface{}, int64, error) {
	clubs, total, err := s.repo.GetAll(page, limit)
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

func (s *ClubService) SearchByName(search string, page, limit int) ([]map[string]interface{}, int64, error) {
	clubs, total, err := s.repo.SearchByName(search, page, limit)
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

func (s *ClubService) CreateClub(club *models.Club) error {
	err := s.repo.CreateClub(club)
	if err != nil {
		return err
	}
	if err := s.repo.DB.Preload("Town").First(club, club.ID).Error; err != nil {
		log.Printf("Не удалось предварительно загрузить город для нового клуба %d: %v", club.ID, err)
	}
	return nil
}

func (s *ClubService) UpdateClub(id uint, updated *models.Club) (*models.Club, error) {
	club := &models.Club{
		ID:                id,
		Name:              updated.Name,
		QuantityTituls:    updated.QuantityTituls,
		AverageAgePlayers: updated.AverageAgePlayers,
		TownID:            updated.TownID,
	}
	err := s.repo.UpdateClub(club)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := s.repo.DB.Preload("Town").First(club, id).Error; err != nil {
		log.Printf("Не удалось предварительно загрузить город для обновленного клуба %d: %v", id, err)
	}
	return club, nil
}

func (s *ClubService) DeleteClub(id uint) error {
	err := s.repo.DeleteClub(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ClubService) GetByID(id uint) (*models.Club, error) {
	var club models.Club
	if err := s.repo.DB.Preload("Town").First(&club, id).Error; err != nil {
		return nil, err
	}
	return &club, nil
}
