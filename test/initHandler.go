package test

import (
	"github.com/gin-gonic/gin"
	"github.com/jagadwp/gin-crud-web/config"
	"github.com/jagadwp/gin-crud-web/controller"
	"github.com/jagadwp/gin-crud-web/database"
	"github.com/jagadwp/gin-crud-web/repository"
	"github.com/jagadwp/gin-crud-web/router"
	"github.com/jagadwp/gin-crud-web/service"
)

func InitHandler() (*gin.Engine, *controller.BeratController) {
	config.Load()

	database.DatabaseInit()
	DB := database.DB()
	defer database.DB().Close()

	beratRepo := repository.NewBeratRepository(DB)
	beratService := service.NewBeratService(beratRepo)
	beratController := controller.NewBeratController(beratService)

	r := gin.Default()

	router.RegisterBeratPath(r, beratController)

	return r, beratController
}
