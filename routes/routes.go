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
		kycRoutes.Use(middleware.JWTMiddleware())
		kycRoutes.GET("/request", kycRequestController.GetKYCRequestsByRange)
	}
}

// RegisterServiceProviderRoutes registers the routes for the Service Provider API
func RegisterServiceProviderRoutes(router *gin.Engine, consentController *controllers.SuccessfulConsentRatio) {
	serviceProviderRoutes := router.Group("/api/ServiceProvider")
	{
		serviceProviderRoutes.Use(middleware.JWTMiddleware()) // Protect with JWT
		serviceProviderRoutes.GET("/SuccessfulConsentRatio", consentController.GetSuccessfulConsentRatio)
	}
}
