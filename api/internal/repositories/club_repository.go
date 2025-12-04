package repositories

import (
	"SportsClubs/internal/models"

	"gorm.io/gorm"
)

type ClubRepository struct {
	DB *gorm.DB
}

func NewClubRepository(db *gorm.DB) *ClubRepository {
	return &ClubRepository{DB: db}
}

func (r *ClubRepository) GetAll(page, limit int) ([]models.Club, int64, error) {
	var clubs []models.Club
	var total int64

	query := r.DB.Model(&models.Club{}).Preload("Town")
	query.Count(&total)

	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit).Order("id ASC")
	}

	result := query.Find(&clubs)
	return clubs, total, result.Error
}

func (r *ClubRepository) SearchByName(search string, page, limit int) ([]models.Club, int64, error) {
	var clubs []models.Club
	var total int64

	query := r.DB.Model(&models.Club{}).Preload("Town")
	query = query.Where("name ILIKE ?", "%"+search+"%")
	query.Count(&total)

	if page > 0 && limit > 0 {
		query = query.Offset((page - 1) * limit).Limit(limit).Order("id ASC")
	}

	result := query.Find(&clubs)
	return clubs, total, result.Error
}

func (r *ClubRepository) CreateClub(club *models.Club) error {
	return r.DB.Create(club).Error
}

func (r *ClubRepository) UpdateClub(club *models.Club) error {
	return r.DB.Save(club).Error
}

func (r *ClubRepository) DeleteClub(id uint) error {
	return r.DB.Delete(&models.Club{}, id).Error
}

func (r *ClubRepository) Upsert(club *models.Club) error {
	return r.DB.Save(club).Error
}
