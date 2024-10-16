package services

import (
	"client/repositories"
	"time"
)

// SuccessfulKYCRatioServiceInterface defines the service interface
type SuccessfulKYCRatioServiceInterface interface {
	GetSuccessfulKYCRatioByDay(date time.Time) (map[string]repositories.KYCRequestRatioData, error)
	GetSuccessfulKYCRatioByWeek() (map[string]repositories.KYCRequestRatioData, error)
	GetSuccessfulKYCRatioByMonth() (map[string]repositories.KYCRequestRatioData, error)
}

// SuccessfulKYCRatioService defines the service structure
type SuccessfulKYCRatioService struct {
	repository repositories.SuccessfulKYCRatioRepositoryInterface
}

// NewSuccessfulKYCRatioService creates a new instance of the service
func NewSuccessfulKYCRatioService(repo repositories.SuccessfulKYCRatioRepositoryInterface) SuccessfulKYCRatioServiceInterface {
	return &SuccessfulKYCRatioService{repository: repo}
}

// GetSuccessfulKYCRatioByDay aggregates the KYC ratio data for a single day
func (s *SuccessfulKYCRatioService) GetSuccessfulKYCRatioByDay(date time.Time) (map[string]repositories.KYCRequestRatioData, error) {
	results, err := s.repository.GetSuccessfulKYCRatioByDay(date)
	if err != nil {
		return nil, err
	}

	aggregatedData := aggregateKYCData(results)
	return aggregatedData, nil
}

// GetSuccessfulKYCRatioByWeek aggregates the KYC ratio data for the last 7 days
func (s *SuccessfulKYCRatioService) GetSuccessfulKYCRatioByWeek() (map[string]repositories.KYCRequestRatioData, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -6) // Last 7 days
	results, err := s.repository.GetSuccessfulKYCRatioByRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	aggregatedData := aggregateKYCData(results)
	return aggregatedData, nil
}

// GetSuccessfulKYCRatioByMonth aggregates the KYC ratio data for the last 30 days
func (s *SuccessfulKYCRatioService) GetSuccessfulKYCRatioByMonth() (map[string]repositories.KYCRequestRatioData, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30) // Last 30 days
	results, err := s.repository.GetSuccessfulKYCRatioByRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	aggregatedData := aggregateKYCData(results)
	return aggregatedData, nil
}

// aggregateKYCData aggregates data by date
func aggregateKYCData(results map[string][]repositories.KYCRequestRatioData) map[string]repositories.KYCRequestRatioData {
	aggregatedData := make(map[string]repositories.KYCRequestRatioData)

	for date, data := range results {
		var totalRequests, status5Requests, status15Requests int64

		for _, entry := range data {
			totalRequests += entry.TotalRequests
			status5Requests += entry.Status5Requests
			status15Requests += entry.Status15Requests
		}

		aggregatedData[date] = repositories.KYCRequestRatioData{
			TotalRequests:    totalRequests,
			Status5Requests:  status5Requests,
			Status15Requests: status15Requests,
		}
	}

	return aggregatedData
}
