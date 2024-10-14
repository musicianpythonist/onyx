package controllers

import (
	"client/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// KYCRequestController handles KYC request-related actions
type KYCRequestController struct {
	kycRequestService services.KYCRequestService
}

func NewKYCRequestController(service services.KYCRequestService) *KYCRequestController {
	return &KYCRequestController{kycRequestService: service}
}

// GetKYCRequestsByRange returns KYC requests filtered by statusId and date range
// @Summary Get KYC requests
// @Description Returns the number of KYC requests with a specific status (e.g., submitted) for the given date range (day, week, month)
// @Tags KYC Requests
// @Produce json
// @Param status_id query int true "Status ID of the KYC request"
// @Param date_range query string true "Date Range (day, week, month)"
// @Success 200 {object} dto.KYCRequestsRangeResponseDTO
// @Failure 400 {object} map[string]string "Invalid status ID or date range"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/kyc/request [get] // Updated route path in the Swagger annotation
func (c *KYCRequestController) GetKYCRequestsByRange(ctx *gin.Context) {
	statusIdParam := ctx.Query("status_id")
	statusId, err := strconv.Atoi(statusIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	dateRange := ctx.Query("date_range")
	result, err := c.kycRequestService.GetKYCRequestsByRange(statusId, dateRange)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
