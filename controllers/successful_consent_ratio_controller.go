package controllers

import (
	"client/repositories"
	"client/services"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

// SuccessfulConsentRatio defines the controller structure
type SuccessfulConsentRatio struct {
	service services.SuccessfulConsentRatioServiceInterface
}

// NewSuccessfulConsentRatio creates a new instance of the controller
func NewSuccessfulConsentRatio(service services.SuccessfulConsentRatioServiceInterface) *SuccessfulConsentRatio {
	return &SuccessfulConsentRatio{service: service}
}

// ConsentRatioResponse represents the format for the final API response
type ConsentRatioResponse struct {
	Date      string                          `json:"date"`
	Merchants []repositories.ConsentRatioData `json:"merchants"` // Updated to reference repositories package
}

// GetSuccessfulConsentRatio handles the API request to get consent ratios
// @Summary Get successful consent ratios
// @Description Get successful consent ratios by day, week, or month
// @Tags ServiceProvider
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {array} ConsentRatioResponse
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Invalid or expired token"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/ServiceProvider/SuccessfulConsentRatio [get]
func (c *SuccessfulConsentRatio) GetSuccessfulConsentRatio(ctx *gin.Context) {
	// Get the date range
	dateRange := ctx.Query("range")

	// Fetch data from the service
	var result map[string][]repositories.ConsentRatioData
	var err error

	switch dateRange {
	case "day":
		result, err = c.service.GetSuccessfulConsentRatioByDay(time.Now())
	case "week":
		result, err = c.service.GetSuccessfulConsentRatioByWeek()
	case "month":
		result, err = c.service.GetSuccessfulConsentRatioByMonth()
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range value"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Sort the dates before responding
	var dates []string
	for date := range result {
		dates = append(dates, date)
	}
	sort.Strings(dates) // Ensure the dates are in ascending order

	// Build the response
	var response []ConsentRatioResponse
	for _, date := range dates {
		response = append(response, ConsentRatioResponse{
			Date:      date,
			Merchants: result[date],
		})
	}

	ctx.JSON(http.StatusOK, response)
}
