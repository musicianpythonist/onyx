package repositories

import (
	"client/models"
	"time"

	"gorm.io/gorm"
)

// KYCRequestRepository defines methods for interacting with the KYCRequest table
type KYCRequestRepository interface {
	GetKYCRequestsByStatusAndDateRange(status uint8, startDate, endDate time.Time) (int64, error)
}

type kycRequestRepository struct {
	db *gorm.DB
}

// NewKYCRequestRepository creates a new instance of kycRequestRepository
func NewKYCRequestRepository(db *gorm.DB) KYCRequestRepository {
	return &kycRequestRepository{db: db}
}

// GetKYCRequestsByStatusAndDateRange returns the number of KYC requests with a specific status between startDate and endDate
func (r *kycRequestRepository) GetKYCRequestsByStatusAndDateRange(status uint8, startDate, endDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&models.KYCRequest{}).
		Where("KYCRequestStatusId = ? AND CreateDate BETWEEN ? AND ?", status, startDate, endDate).
		Count(&count).Error
	return count, err
}
