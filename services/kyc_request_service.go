package services

import (
	"client/dto"
	"client/repositories"
	"errors"
	"time"
)

// Declare KYCRequestService as an interface
type KYCRequestService interface {
	GetKYCRequestsByRange(statusId int, dateRange string) (dto.KYCRequestsRangeResponseDTO, error)
}

// Create a struct that implements the interface
type kycRequestService struct {
	kycRequestRepo repositories.KYCRequestRepository
}

// NewKYCRequestService returns an instance of the service implementing the KYCRequestService interface
func NewKYCRequestService(repo repositories.KYCRequestRepository) KYCRequestService {
	return &kycRequestService{kycRequestRepo: repo}
}

// Implement the GetKYCRequestsByRange method for the kycRequestService struct
func (s *kycRequestService) GetKYCRequestsByRange(statusId int, dateRange string) (dto.KYCRequestsRangeResponseDTO, error) {
	now := time.Now()
	var result []dto.KYCRequestsDayResponseDTO

	switch dateRange {
	case "day":
		start := now.Truncate(24 * time.Hour)
		end := start.Add(24 * time.Hour)
		count, err := s.kycRequestRepo.GetKYCRequestCountByDateRange(statusId, start, end)
		if err != nil {
			return dto.KYCRequestsRangeResponseDTO{}, err
		}
		result = append(result, dto.KYCRequestsDayResponseDTO{
			Date:  start.Format("2006-01-02"),
			Count: count,
		})

	case "week":
		for i := 6; i >= 0; i-- {
			start := now.AddDate(0, 0, -i).Truncate(24 * time.Hour)
			end := start.Add(24 * time.Hour)
			count, err := s.kycRequestRepo.GetKYCRequestCountByDateRange(statusId, start, end)
			if err != nil {
				return dto.KYCRequestsRangeResponseDTO{}, err
			}
			result = append(result, dto.KYCRequestsDayResponseDTO{
				Date:  start.Format("2006-01-02"),
				Count: count,
			})
		}

	case "month":
		for i := 29; i >= 0; i-- {
			start := now.AddDate(0, 0, -i).Truncate(24 * time.Hour)
			end := start.Add(24 * time.Hour)
			count, err := s.kycRequestRepo.GetKYCRequestCountByDateRange(statusId, start, end)
			if err != nil {
				return dto.KYCRequestsRangeResponseDTO{}, err
			}
			result = append(result, dto.KYCRequestsDayResponseDTO{
				Date:  start.Format("2006-01-02"),
				Count: count,
			})
		}

	default:
		return dto.KYCRequestsRangeResponseDTO{}, errors.New("invalid date range")
	}

	return dto.KYCRequestsRangeResponseDTO{
		Range:       dateRange,
		KYCRequests: result,
	}, nil
}
