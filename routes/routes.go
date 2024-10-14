package routes

import (
	"client/controllers"
	"client/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes registers the routes for the client API
func RegisterClientRoutes(router *gin.Engine, clientController *controllers.ClientController) {
	clientRoutes := router.Group("/api/clients")
	{
		clientRoutes.Use(middleware.JWTMiddleware()) // Protect with JWT
		clientRoutes.GET("/new", clientController.GetNewClientsByRange)
	}
}

// RegisterKYCRequestRoutes registers the routes for KYC requests
func RegisterKYCRequestRoutes(router *gin.Engine, kycRequestController *controllers.KYCRequestController) {
	kycRoutes := router.Group("/api/kyc")
	{
		kycRoutes.Use(middleware.JWTMiddleware())                             // Apply JWT middleware to protect the route
		kycRoutes.GET("/request", kycRequestController.GetKYCRequestsByRange) // Update the route to /api/kyc/request
	}
}
