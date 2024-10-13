package services

import (
	"client/repositories"
	"errors"
	"time"
)

// KYCRequestService defines service methods for KYC requests
type KYCRequestService interface {
	GetKYCRequestsByRangeAndStatus(status uint8, dateRange string) (int64, error)
}

type kycRequestService struct {
	kycRequestRepo repositories.KYCRequestRepository
}

// NewKYCRequestService initializes a new KYCRequestService
func NewKYCRequestService(kycRequestRepo repositories.KYCRequestRepository) KYCRequestService {
	return &kycRequestService{kycRequestRepo: kycRequestRepo}
}

// GetKYCRequestsByRangeAndStatus returns the number of KYC requests with a given status for the specified date range (day, week, month)
func (s *kycRequestService) GetKYCRequestsByRangeAndStatus(status uint8, dateRange string) (int64, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch dateRange {
	case "day":
		startDate = now.Truncate(24 * time.Hour)
		endDate = startDate.Add(24 * time.Hour)
	case "week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	default:
		return 0, errors.New("invalid date range")
	}

	// Use the repository method to get the count of KYC requests with the given status
	count, err := s.kycRequestRepo.GetKYCRequestsByStatusAndDateRange(status, startDate, endDate)
	if err != nil {
		return 0, err
	}

	return count, nil
}
