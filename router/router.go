package router

import (
	"github.com/bonggar/gorestapi/controller"

	"github.com/gin-gonic/gin"
)


//Make : create endpoints
func RegisterBeratPath(r *gin.Engine, beratController *controller.BeratController ) {
	r.LoadHTMLGlob("templates/*.html")
	
	r.GET("/", beratController.GetAllBerat)
	// r.GET("/", controller.GetAllBerat)
	// r.GET("/show", controller.GetAllBerat)
	// r.GET("/new", controller.GetAllBerat)
	// r.GET("/edit", controller.GetAllBerat)
	// r.GET("/insert", controller.GetAllBerat)
	// r.PUT("/update", controller.GetAllBerat)
	// r.DELETE("/delete", controller.GetAllBerat)

}
