package services

import (
	"client/dto"
	"client/repositories"
	"errors"
	"time"
)

// ClientService defines the service methods available
type ClientService interface {
	GetNewClientsByRange(dateRange string) (dto.NewClientsRangeResponseDTO, error)
	// GetKYCClientsByRange(dateRange string) (map[string]interface{}, error)
}

// clientService struct implements ClientService and holds the repository
type clientService struct {
	clientRepo repositories.ClientRepository
}

// NewClientService initializes a new ClientService
func NewClientService(clientRepo repositories.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

// GetNewClientsByRange returns the number of new clients with status 5 or 10 for the given date range (day, week, month)
func (s *clientService) GetNewClientsByRange(dateRange string) (dto.NewClientsRangeResponseDTO, error) {
	now := time.Now()
	var result []dto.NewClientsDayResponseDTO

	// Logic for day, week, month
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

	// Return the result
	return dto.NewClientsRangeResponseDTO{
		Range:      dateRange,
		NewClients: result,
	}, nil
}
