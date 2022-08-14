package router

import (
	"github.com/jagadwp/gin-crud-web/controller"

	"github.com/gin-gonic/gin"
)

//Make : create endpoints
func RegisterBeratPath(r *gin.Engine, beratController *controller.BeratController) {
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", beratController.GetAllBerat)
	r.GET("/show", beratController.GetBeratById)
	r.GET("/new", beratController.NewBerat)
	r.POST("/insert", beratController.CreateBerat)
	r.GET("/edit", beratController.EditBerat)
	r.POST("/update", beratController.UpdateBerat)
	r.GET("/delete", beratController.DeleteBerat)
}
