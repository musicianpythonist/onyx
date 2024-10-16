package repositories

import (
	"time"

	"gorm.io/gorm"
)

// KYCRequestRatioData represents the aggregated data for KYC request ratios (without date)
type KYCRequestRatioData struct {
	ClientId         int   `json:"clientId"`
	TotalRequests    int64 `json:"totalRequests"`
	Status5Requests  int64 `json:"status5Requests"`
	Status15Requests int64 `json:"status15Requests"`
}

// KYCRequestRatioDataWithDate is used internally for organizing data by date (includes date for internal use)
type KYCRequestRatioDataWithDate struct {
	ClientId         int       `json:"clientId"`
	TotalRequests    int64     `json:"totalRequests"`
	Status5Requests  int64     `json:"status5Requests"`
	Status15Requests int64     `json:"status15Requests"`
	Date             time.Time `json:"-"` // Date is used internally but not exposed in the final JSON response
}

// SuccessfulKYCRatioRepositoryInterface defines the interface for the repository
type SuccessfulKYCRatioRepositoryInterface interface {
	GetSuccessfulKYCRatioByDay(date time.Time) (map[string][]KYCRequestRatioData, error)
	GetSuccessfulKYCRatioByRange(startDate time.Time, endDate time.Time) (map[string][]KYCRequestRatioData, error)
}

// SuccessfulKYCRatioRepository defines the repository structure
type SuccessfulKYCRatioRepository struct {
	db *gorm.DB
}

// NewSuccessfulKYCRatioRepository creates a new instance of the repository
func NewSuccessfulKYCRatioRepository(db *gorm.DB) *SuccessfulKYCRatioRepository {
	return &SuccessfulKYCRatioRepository{db: db}
}

// GetSuccessfulKYCRatioByDay retrieves KYC request ratios for a single day
func (r *SuccessfulKYCRatioRepository) GetSuccessfulKYCRatioByDay(date time.Time) (map[string][]KYCRequestRatioData, error) {
	var results []KYCRequestRatioDataWithDate
	err := r.db.Table("KYCRequest").
		Select("ClientId, CONVERT(date, CreateDate) as Date, COUNT(Id) as TotalRequests, COUNT(CASE WHEN KYCRequestStatusId = 5 THEN 1 END) as Status5Requests, COUNT(CASE WHEN KYCRequestStatusId = 15 THEN 1 END) as Status15Requests").
		Where("CONVERT(date, CreateDate) = ?", date.Format("2006-01-02")).
		Group("ClientId, CONVERT(date, CreateDate)").
		Order("CONVERT(date, CreateDate) ASC"). // Sort by date (even though it's a single day)
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	// Organize results by date but exclude the date field from the objects
	dataByDate := make(map[string][]KYCRequestRatioData)
	for _, result := range results {
		dateStr := result.Date.Format("2006-01-02T00:00:00Z")
		dataByDate[dateStr] = append(dataByDate[dateStr], KYCRequestRatioData{
			ClientId:         result.ClientId,
			TotalRequests:    result.TotalRequests,
			Status5Requests:  result.Status5Requests,
			Status15Requests: result.Status15Requests,
		})
	}
	return dataByDate, nil
}

// GetSuccessfulKYCRatioByRange retrieves KYC request ratios for a range of dates (week or month)
func (r *SuccessfulKYCRatioRepository) GetSuccessfulKYCRatioByRange(startDate time.Time, endDate time.Time) (map[string][]KYCRequestRatioData, error) {
	var results []KYCRequestRatioDataWithDate
	err := r.db.Table("KYCRequest").
		Select("ClientId, CONVERT(date, CreateDate) as Date, COUNT(Id) as TotalRequests, COUNT(CASE WHEN KYCRequestStatusId = 5 THEN 1 END) as Status5Requests, COUNT(CASE WHEN KYCRequestStatusId = 15 THEN 1 END) as Status15Requests").
		Where("CONVERT(date, CreateDate) BETWEEN ? AND ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Group("ClientId, CONVERT(date, CreateDate)").
		Order("Date ASC"). // Ensure ordering by date
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	// Organize results by date but exclude the date field from the objects
	dataByDate := make(map[string][]KYCRequestRatioData)
	for _, result := range results {
		dateStr := result.Date.Format("2006-01-02T00:00:00Z")
		dataByDate[dateStr] = append(dataByDate[dateStr], KYCRequestRatioData{
			ClientId:         result.ClientId,
			TotalRequests:    result.TotalRequests,
			Status5Requests:  result.Status5Requests,
			Status15Requests: result.Status15Requests,
		})
	}
	return dataByDate, nil
}
