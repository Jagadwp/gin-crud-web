package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jagadwp/gin-crud-web/helper"
	"github.com/jagadwp/gin-crud-web/service"

	"github.com/gin-gonic/gin"
)

type BeratController struct {
	beratService *service.BeratService
}

func NewBeratController(beratService *service.BeratService) *BeratController {
	return &BeratController{beratService: beratService}
}

//GetBerat : get a berat which find by ID
func (ctr *BeratController) GetBeratById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		helper.RespondJSON(c, 400, "Parameter id is not valid", nil, nil)
	}

	res, _ := ctr.beratService.GetBeratById(id)

	if (*res).ID == 0 {
		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
	}

	c.HTML(http.StatusOK, "Show", gin.H{
		"data": *res,
	})
}

//GetBerats : get all berat
func (ctr *BeratController) GetAllBerat(c *gin.Context) {
	res, err := ctr.beratService.GetAllBerat()

	if err != nil {
		helper.RespondJSON(c, 500, "Failed to process request", nil, nil)
	}

	// fmt.Println("\nList berat\nData: ", *res)
	c.HTML(http.StatusOK, "Index", gin.H{
		"data": *res,
	})

}

func (ctr *BeratController) NewBerat(c *gin.Context) {
	c.HTML(http.StatusOK, "New", nil)
}

// //CreateBerat : create a berat
func (ctr *BeratController) CreateBerat(c *gin.Context) {
	req := helper.BeratRequest{}

	if err := c.Bind(&req); err != nil {
		helper.RespondJSON(c, 400, "Required fields are empty or not valid", nil, nil)
	}

	res, err := ctr.beratService.CreateBerat(&req)

	if err != nil {
		helper.RespondJSON(c, 500, "Failed to process request", nil, nil)
	}

	fmt.Println("\nNew berat created\nData: ", res)
	c.Redirect(http.StatusFound, "http://localhost:8080/")
}

func (ctr *BeratController) EditBerat(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		helper.RespondJSON(c, 400, "Parameter id is not valid", nil, nil)
	}

	res, _ := ctr.beratService.GetBeratById(id)

	if (*res).ID == 0 {
		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
	}

	c.HTML(http.StatusOK, "Edit", gin.H{
		"data": *res,
	})
}

// //UpdateBerat : edit a berat which find by ID
func (ctr *BeratController) UpdateBerat(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		helper.RespondJSON(c, 400, "Parameter id is not valid", nil, nil)
	}

	isExistRes, _ := ctr.beratService.GetBeratById(id)

	if (*isExistRes).ID == 0 {
		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
	}

	req := helper.BeratRequest{}
	if err := c.Bind(&req); err != nil {
		helper.RespondJSON(c, 400, "Required fields are empty or not valid", nil, nil)
	}

	res, err := ctr.beratService.UpdateBerat(id, &req)

	if err != nil {
		helper.RespondJSON(c, 500, err.Error(), nil, nil)
	}

	fmt.Println("\nBerat updated\nData: ", res)
	c.Redirect(http.StatusFound, "http://localhost:8080/")
}

//DeleteBerat : delete a berat which find by ID

func (ctr *BeratController) DeleteBerat(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		helper.RespondJSON(c, 400, "Parameter id is not valid", nil, nil)
	}

	isExistRes, _ := ctr.beratService.GetBeratById(id)

	if (*isExistRes).ID == 0 {
		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
	}

	res, err := ctr.beratService.DeleteBerat(id)

	if err != nil {
		helper.RespondJSON(c, 500, "Failed to process request", nil, nil)
	}

	fmt.Println("\nBerat deleted\nData: ", res)
	c.Redirect(http.StatusFound, "http://localhost:8080/")
}

// //OptionsBerat : supporting options for CORS
// func OptionsBerat(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	c.Next()
// }
