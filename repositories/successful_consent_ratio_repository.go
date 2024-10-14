package repositories

import (
	"time"

	"gorm.io/gorm"
)

// ConsentRatioData represents the aggregated data for consent ratios (without date)
type ConsentRatioData struct {
	MerchantId         int   `json:"merchantId"`
	TotalClients       int64 `json:"totalClients"`
	SuccessfulConsents int64 `json:"successfulConsents"`
}

// ConsentRatioDataWithDate is used internally for organizing data by date (includes date for internal use)
type ConsentRatioDataWithDate struct {
	MerchantId         int       `json:"merchantId"`
	TotalClients       int64     `json:"totalClients"`
	SuccessfulConsents int64     `json:"successfulConsents"`
	Date               time.Time `json:"-"` // Date is used internally but not exposed in the final JSON response
}

// SuccessfulConsentRatioRepositoryInterface defines the interface for the repository
type SuccessfulConsentRatioRepositoryInterface interface {
	GetSuccessfulConsentRatioByDay(date time.Time) (map[string][]ConsentRatioData, error)
	GetSuccessfulConsentRatioByRange(startDate time.Time, endDate time.Time) (map[string][]ConsentRatioData, error)
}

// SuccessfulConsentRatioRepository defines the repository structure
type SuccessfulConsentRatioRepository struct {
	db *gorm.DB
}

// NewSuccessfulConsentRatioRepository creates a new instance of the repository
func NewSuccessfulConsentRatioRepository(db *gorm.DB) *SuccessfulConsentRatioRepository {
	return &SuccessfulConsentRatioRepository{db: db}
}

// GetSuccessfulConsentRatioByDay retrieves consent ratios for a single day
func (r *SuccessfulConsentRatioRepository) GetSuccessfulConsentRatioByDay(date time.Time) (map[string][]ConsentRatioData, error) {
	var results []ConsentRatioDataWithDate
	err := r.db.Table("ConsentToServiceProvider").
		Select("MerchantId, CONVERT(date, CreateDate) as Date, COUNT(DISTINCT ClientId) as TotalClients, COUNT(DISTINCT CASE WHEN ConsentToServiceProviderStatusId = 2 THEN ClientId END) as SuccessfulConsents").
		Where("CONVERT(date, CreateDate) = ?", date.Format("2006-01-02")).
		Group("MerchantId, CONVERT(date, CreateDate)").
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	// Organize results by date but exclude the date field from the objects
	dataByDate := make(map[string][]ConsentRatioData)
	for _, result := range results {
		dateStr := result.Date.Format("2006-01-02T00:00:00Z")
		dataByDate[dateStr] = append(dataByDate[dateStr], ConsentRatioData{
			MerchantId:         result.MerchantId,
			TotalClients:       result.TotalClients,
			SuccessfulConsents: result.SuccessfulConsents,
		})
	}
	return dataByDate, nil
}

// GetSuccessfulConsentRatioByRange retrieves consent ratios for a range of dates (week or month)
func (r *SuccessfulConsentRatioRepository) GetSuccessfulConsentRatioByRange(startDate time.Time, endDate time.Time) (map[string][]ConsentRatioData, error) {
	var results []ConsentRatioDataWithDate
	err := r.db.Table("ConsentToServiceProvider").
		Select("MerchantId, CONVERT(date, CreateDate) as Date, COUNT(DISTINCT ClientId) as TotalClients, COUNT(DISTINCT CASE WHEN ConsentToServiceProviderStatusId = 2 THEN ClientId END) as SuccessfulConsents").
		Where("CONVERT(date, CreateDate) BETWEEN ? AND ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Group("MerchantId, CONVERT(date, CreateDate)").
		Order("CONVERT(date, CreateDate) ASC").
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	// Organize results by date but exclude the date field from the objects
	dataByDate := make(map[string][]ConsentRatioData)
	for _, result := range results {
		dateStr := result.Date.Format("2006-01-02T00:00:00Z")
		dataByDate[dateStr] = append(dataByDate[dateStr], ConsentRatioData{
			MerchantId:         result.MerchantId,
			TotalClients:       result.TotalClients,
			SuccessfulConsents: result.SuccessfulConsents,
		})
	}
	return dataByDate, nil
}
