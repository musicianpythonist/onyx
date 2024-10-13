package repositories

import (
	"client/models"
	"time"

	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClientCountByDateRange(startDate, endDate time.Time) (int64, error)
	GetKYCClientCountByDateRange(startDate, endDate time.Time) (int64, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

// GetClientCountByDateRange returns the number of clients created between startDate and endDate
func (r *clientRepository) GetClientCountByDateRange(startDate, endDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&models.Client{}).Where("CreateDate BETWEEN ? AND ?", startDate, endDate).Count(&count).Error
	return count, err
}

// GetClientCountByKYCAndDateRange counts the number of clients who have completed KYC within the given date range
func (r *clientRepository) GetKYCClientCountByDateRange(startDate, endDate time.Time) (int64, error) {
	var count int64
	// Query the database to count clients where NationalCode is not null and within the date range
	if err := r.db.Model(&models.Client{}).
		Where("NationalCode IS NOT NULL AND CreateDate BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
