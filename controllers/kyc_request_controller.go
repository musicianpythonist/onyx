package controllers

import (
	"client/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// KYCRequestController defines the KYC request controller
type KYCRequestController struct {
	kycRequestService services.KYCRequestService
}

// NewKYCRequestController initializes a new KYCRequestController
func NewKYCRequestController(kycRequestService services.KYCRequestService) *KYCRequestController {
	return &KYCRequestController{kycRequestService: kycRequestService}
}

// GetKYCRequestsByRange handles the request to get KYC requests with a given status and date range
// @Summary Get submitted KYC requests
// @Description Get submitted KYC requests (status 5) by day, week, or month
// @Tags KYCRequest
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/kycrequest/submitted [get]
func (ctrl *KYCRequestController) GetKYCRequestsByRange(c *gin.Context) {
	dateRange := c.Query("range")
	if dateRange == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'range' is required"})
		return
	}

	status := uint8(5) // Status 5 is for submitted KYC requests
	count, err := ctrl.kycRequestService.GetKYCRequestsByRangeAndStatus(status, dateRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "status": status, "range": dateRange})
}
