package repositories

import (
	"client/models"
	"time"

	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClientCountByDateRange(startDate, endDate time.Time) (int64, error)
}

type clientRepository struct {
	db *gorm.DB
}

// NewClientRepository creates a new instance of ClientRepository
func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

// GetClientCountByDateRange counts clients whose status is 5 or 10 within the given date range
func (r *clientRepository) GetClientCountByDateRange(startDate, endDate time.Time) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Client{}).
		Where("ClientTypeId IN (?, ?)", 5, 10).
		Where("CreateDate BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
