package repositories

import (
	"client/models"
	"time"

	"gorm.io/gorm"
)

type KYCRequestRepository interface {
	GetKYCRequestCountByDateRange(statusId int, startDate, endDate time.Time) (int64, error)
}

type kycRequestRepository struct {
	db *gorm.DB
}

func NewKYCRequestRepository(db *gorm.DB) KYCRequestRepository {
	return &kycRequestRepository{db: db}
}

// GetKYCRequestCountByDateRange counts the KYC requests with the given statusId within the provided date range
func (r *kycRequestRepository) GetKYCRequestCountByDateRange(statusId int, startDate, endDate time.Time) (int64, error) {
	var count int64
	if err := r.db.Model(&models.KYCRequest{}).
		Where("KYCRequestStatusId = ?", statusId).
		Where("CreateDate BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
