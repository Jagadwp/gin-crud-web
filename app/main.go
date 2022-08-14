package main

import (
	"github.com/bonggar/gorestapi/config"
	"github.com/bonggar/gorestapi/controller"
	"github.com/bonggar/gorestapi/database"
	"github.com/bonggar/gorestapi/repository"
	"github.com/bonggar/gorestapi/router"
	"github.com/bonggar/gorestapi/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.Load()

	// Open connection to database
	database.DatabaseInit()
	DB := database.DB()
	defer database.DB().Close()

	// Init Repo, Service, Controller
	beratRepo := repository.NewBeratRepository(DB)
	beratService := service.NewBeratService(beratRepo)
	beratController := controller.NewBeratController(beratService)

	// Init Gin
	r := gin.Default()
	
	// Register routes
	router.RegisterBeratPath(r, beratController)

	// Run Server
	r.Run("localhost:" + config.HTTPPort)
}
