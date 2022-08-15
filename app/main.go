package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jagadwp/gin-crud-web/config"
	"github.com/jagadwp/gin-crud-web/controller"
	"github.com/jagadwp/gin-crud-web/database"
	"github.com/jagadwp/gin-crud-web/repository"
	"github.com/jagadwp/gin-crud-web/router"
	"github.com/jagadwp/gin-crud-web/service"
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

	// Using interface for testing for ease of testing
	var IBeratRepo repository.IBeratRepository = beratRepo
	beratService := service.NewBeratService(IBeratRepo)
	beratController := controller.NewBeratController(beratService)

	// Init Gin
	r := gin.Default()
	
	// Register routes
	router.RegisterBeratPath(r, beratController)

	// Run Server
	r.Run("localhost:" + config.HTTPPort)
}
