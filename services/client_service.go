package services

import (
	"client/dto"
	"client/repositories"
	"errors"
	"fmt"
	"time"
)

// ClientService defines the service methods available
type ClientService interface {
	GetNewClientsByRange(dateRange string) (dto.NewClientsRangeResponseDTO, error)
	GetKYCClientsByRange(dateRange string) (map[string]interface{}, error)
}

// clientService struct implements ClientService and holds the repository
type clientService struct {
	clientRepo repositories.ClientRepository
}

// NewClientService initializes a new ClientService
func NewClientService(clientRepo repositories.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

// GetNewClientsByRange returns the number of new clients for the given date range (day, week, month)
func (s *clientService) GetNewClientsByRange(dateRange string) (dto.NewClientsRangeResponseDTO, error) {
	now := time.Now()
	var result []dto.NewClientsDayResponseDTO

	// Logic for day, week, month as described in your original code
	switch dateRange {
	case "day":
		start := now.Truncate(24 * time.Hour)
		end := start.Add(24 * time.Hour)
		count, err := s.clientRepo.GetClientCountByDateRange(start, end)
		if err != nil {
			return dto.NewClientsRangeResponseDTO{}, err
		}
		result = append(result, dto.NewClientsDayResponseDTO{
			Date:  start.Format("2006-01-02"),
			Count: count,
		})

	case "week":
		for i := 6; i >= 0; i-- {
			start := now.AddDate(0, 0, -i).Truncate(24 * time.Hour)
			end := start.Add(24 * time.Hour)
			count, err := s.clientRepo.GetClientCountByDateRange(start, end)
			if err != nil {
				return dto.NewClientsRangeResponseDTO{}, err
			}
			result = append(result, dto.NewClientsDayResponseDTO{
				Date:  start.Format("2006-01-02"),
				Count: count,
			})
		}

	case "month":
		for i := 29; i >= 0; i-- {
			start := now.AddDate(0, 0, -i).Truncate(24 * time.Hour)
			end := start.Add(24 * time.Hour)
			count, err := s.clientRepo.GetClientCountByDateRange(start, end)
			if err != nil {
				return dto.NewClientsRangeResponseDTO{}, err
			}
			result = append(result, dto.NewClientsDayResponseDTO{
				Date:  start.Format("2006-01-02"),
				Count: count,
			})
		}

	default:
		return dto.NewClientsRangeResponseDTO{}, errors.New("invalid date range")
	}

	// Return NewClientsRangeResponseDTO
	return dto.NewClientsRangeResponseDTO{
		Range:      dateRange,
		NewClients: result,
	}, nil
}

// GetKYCClientsByRange returns the number of clients who have completed KYC for the given date range
func (s *clientService) GetKYCClientsByRange(dateRange string) (map[string]interface{}, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch dateRange {
	case "day":
		startDate = now.Truncate(24 * time.Hour) // Start of the day
		endDate = startDate.Add(24 * time.Hour)  // End of the day
	case "week":
		startDate = now.AddDate(0, 0, -7) // Start of the week (7 days ago)
		endDate = now                     // End of the current day
	case "month":
		startDate = now.AddDate(0, -1, 0) // Start of the month (1 month ago)
		endDate = now                     // End of the current day
	default:
		return nil, fmt.Errorf("invalid date range")
	}

	// Use the method in the repository to get the KYC client count
	count, err := s.clientRepo.GetKYCClientCountByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Return the count in a map
	return map[string]interface{}{
		"range":      dateRange,
		"kycClients": count,
	}, nil
}
