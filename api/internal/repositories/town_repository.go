package repositories

import (
	"SportsClubs/internal/models"

	"gorm.io/gorm"
)

type TownRepository struct {
	DB *gorm.DB
}

func NewTownRepository(db *gorm.DB) *TownRepository {
	return &TownRepository{DB: db}
}

func (r *TownRepository) GetAll(search string, page int, limit int) ([]models.Town, int64, error) {
	var towns []models.Town
	var total int64

	query := r.DB.Model(&models.Town{})
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit).Order("id ASC")
	}

	result := query.Find(&towns)
	return towns, total, result.Error
}

func (r *TownRepository) GetClubsByTownName(search string, page int, limit int) ([]models.Club, int64, error) {
	if search == "" {
		return nil, 0, gorm.ErrInvalidData
	}
	var clubs []models.Club
	var total int64

	query := r.DB.Model(&models.Club{}).
		Preload("Town").
		Joins("JOIN towns ON towns.id = clubs.town_id").
		Where("towns.name ILIKE ?", "%"+search+"%")

	query.Count(&total)

	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit).Order("id ASC")
	}

	result := query.Find(&clubs)
	return clubs, total, result.Error
}

func (r *TownRepository) CreateTown(town *models.Town) error {
	return r.DB.Create(town).Error
}

func (r *TownRepository) UpdateTown(town *models.Town) error {
	return r.DB.Save(town).Error
}

func (r *TownRepository) DeleteTown(id uint) error {
	return r.DB.Delete(&models.Town{}, id).Error
}

func (r *TownRepository) Upsert(town *models.Town) error {
	return r.DB.Save(town).Error
}
