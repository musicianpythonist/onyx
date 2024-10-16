package controllers

import (
	"client/repositories"
	"client/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SuccessfulKYCRatio defines the controller structure
type SuccessfulKYCRatio struct {
	service services.SuccessfulKYCRatioServiceInterface
}

// NewSuccessfulKYCRatio creates a new instance of the controller
func NewSuccessfulKYCRatio(service services.SuccessfulKYCRatioServiceInterface) *SuccessfulKYCRatio {
	return &SuccessfulKYCRatio{service: service}
}

// / KYCRatioResponse represents the format for the final API response
type KYCRatioResponse struct {
	Date  string                           `json:"date"`
	Stats repositories.KYCRequestRatioData `json:"stats"` // Aggregated data for the date
}

// GetSuccessfulKYCRatio handles the API request to get aggregated KYC ratios
// @Summary Get successful KYC ratios
// @Description Get aggregated KYC request ratios by day, week, or month
// @Tags KYCRequest
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {array} KYCRatioResponse
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Invalid or expired token"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/KYC/SuccessfulKYCRatio [get]
func (c *SuccessfulKYCRatio) GetSuccessfulKYCRatio(ctx *gin.Context) {
	// Get the date range
	dateRange := ctx.Query("range")

	// Fetch aggregated data from the service
	var result map[string]repositories.KYCRequestRatioData
	var err error

	switch dateRange {
	case "day":
		result, err = c.service.GetSuccessfulKYCRatioByDay(time.Now())
	case "week":
		result, err = c.service.GetSuccessfulKYCRatioByWeek()
	case "month":
		result, err = c.service.GetSuccessfulKYCRatioByMonth()
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range value"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Build the response
	var response []KYCRatioResponse
	for date, stats := range result {
		response = append(response, KYCRatioResponse{
			Date:  date,
			Stats: stats,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
