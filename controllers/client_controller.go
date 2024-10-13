package controllers

import (
	"client/services" // Corrected import path
	"net/http"

	"github.com/gin-gonic/gin"
)

// ClientController struct
type ClientController struct {
	clientService services.ClientService
}

// NewClientController initializes a new ClientController
func NewClientController(clientService services.ClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

// GetNewClientsByRange handles the request to get new clients based on the date range
// @Summary Get new clients
// @Description Get new clients by day, week, or month
// @Tags Clients
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Invalid or expired token"
// @Router /api/clients/new [get]
func (ctrl *ClientController) GetNewClientsByRange(c *gin.Context) {
	dateRange := c.Query("range")

	if dateRange == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'range' is required"})
		return
	}

	result, err := ctrl.clientService.GetNewClientsByRange(dateRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetKYCClientsByRange handles the request to get KYC-completed clients based on the date range
// @Summary Get KYC-completed clients
// @Description Get KYC-completed clients by day, week, or month
// @Tags Clients
// @Security ApiKeyAuth
// @Param range query string true "Range (day, week, month)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 401 {object} map[string]interface{} "Invalid or expired token"
// @Router /api/clients/kyc [get]
func (ctrl *ClientController) GetKYCClientsByRange(c *gin.Context) {
	dateRange := c.Query("range")

	if dateRange == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'range' is required"})
		return
	}

	result, err := ctrl.clientService.GetKYCClientsByRange(dateRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
