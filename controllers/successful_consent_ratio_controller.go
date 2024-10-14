package controllers

import (
	"client/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SuccessfulConsentRatioController defines the controller structure
type SuccessfulConsentRatioController struct {
	service services.SuccessfulConsentRatioServiceInterface
}

// NewSuccessfulConsentRatioController creates a new instance of the controller
func NewSuccessfulConsentRatioController(service services.SuccessfulConsentRatioServiceInterface) *SuccessfulConsentRatioController {
	return &SuccessfulConsentRatioController{service: service}
}

// GetSuccessfulConsentRatio handles the API request to get consent ratios
// @Summary Get successful consent ratios
// @Description Get successful consent ratios by day, week, or month
// @Tags ServiceProvider
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Invalid or expired token"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/ServiceProvider/SuccessfulConsentRatioController [get]
func (c *SuccessfulConsentRatioController) GetSuccessfulConsentRatio(ctx *gin.Context) {
	// Retrieve the 'range' query parameter
	dateRange := ctx.Query("range")

	if dateRange == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query param 'range' is required"})
		return
	}

	switch dateRange {
	case "day":
		// Get today's data
		data, err := c.service.GetSuccessfulConsentRatioByDay(time.Now())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, data)

	case "week":
		// Get last 7 days data
		data, err := c.service.GetSuccessfulConsentRatioByWeek()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, data)

	case "month":
		// Get last 30 days data
		data, err := c.service.GetSuccessfulConsentRatioByMonth()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, data)

	default:
		// Invalid range provided
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range value. Use 'day', 'week', or 'month'."})
	}
}
