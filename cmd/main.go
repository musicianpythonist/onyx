package main

import (
	"client/config"
	"client/controllers"
	"client/repositories"
	"client/routes"
	"client/services"

	_ "client/docs" // swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Client and KYCRequest API
// @version 1.0
// @description This is the API for fetching client and KYC request data
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Gin router
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define login route for obtaining JWT token (POST /api/login)
	r.POST("/api/login", controllers.Login)

	// Connect to the database
	db := config.ConnectClientDatabase()

	// Initialize the repository, service, and controller for clients
	clientRepo := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo)
	clientController := controllers.NewClientController(clientService)

	// Register client routes
	routes.RegisterClientRoutes(r, clientController)

	// Initialize the repository, service, and controller for KYC requests
	kycRequestRepo := repositories.NewKYCRequestRepository(db)
	kycRequestService := services.NewKYCRequestService(kycRequestRepo)
	kycRequestController := controllers.NewKYCRequestController(kycRequestService)

	// Register KYC request routes
	routes.RegisterKYCRequestRoutes(r, kycRequestController)

	// Start the server
	r.Run(":8080")
}
