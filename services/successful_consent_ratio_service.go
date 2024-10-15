package services

import (
	"client/repositories"
	"time"
)

// SuccessfulConsentRatioServiceInterface defines the service interface
type SuccessfulConsentRatioServiceInterface interface {
	GetSuccessfulConsentRatioByDay(date time.Time) (map[string][]repositories.ConsentRatioData, error)
	GetSuccessfulConsentRatioByWeek() (map[string][]repositories.ConsentRatioData, error)
	GetSuccessfulConsentRatioByMonth() (map[string][]repositories.ConsentRatioData, error)
}

// SuccessfulConsentRatioService defines the service structure
type SuccessfulConsentRatioService struct {
	repository repositories.SuccessfulConsentRatioRepositoryInterface
}

// NewSuccessfulConsentRatioService creates a new instance of the service
func NewSuccessfulConsentRatioService(repo repositories.SuccessfulConsentRatioRepositoryInterface) SuccessfulConsentRatioServiceInterface {
	return &SuccessfulConsentRatioService{repository: repo}
}

// GetSuccessfulConsentRatioByDay retrieves the consent ratio for a single day
func (s *SuccessfulConsentRatioService) GetSuccessfulConsentRatioByDay(date time.Time) (map[string][]repositories.ConsentRatioData, error) {
	return s.repository.GetSuccessfulConsentRatioByDay(date)
}

// GetSuccessfulConsentRatioByWeek retrieves the consent ratio for the last 7 days
func (s *SuccessfulConsentRatioService) GetSuccessfulConsentRatioByWeek() (map[string][]repositories.ConsentRatioData, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -6)
	return s.repository.GetSuccessfulConsentRatioByRange(startDate, endDate)
}

// GetSuccessfulConsentRatioByMonth retrieves the consent ratio for the last 30 days
func (s *SuccessfulConsentRatioService) GetSuccessfulConsentRatioByMonth() (map[string][]repositories.ConsentRatioData, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)
	return s.repository.GetSuccessfulConsentRatioByRange(startDate, endDate)
}
