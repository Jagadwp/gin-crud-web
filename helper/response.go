package helper

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BeratResponse struct {
	ID   uint      `json:"id" form:"id"`
	Max  int       `json:"max" form:"max"`
	Min  int       `json:"min" form:"min"`
	Diff int       `json:"diff" form:"diff"`
	Date time.Time `json:"date" form:"date"`
}

type IndexResponse struct {
	Berat []BeratResponse `json:"berat" form:"berat"`
	Max  float32 `json:"max" form:"max"`
	Min  float32 `json:"min" form:"min"`
	Diff float32 `json:"diff" form:"diff"`
}

type AvgResponse struct {
	Max  float32 `json:"max" form:"max"`
	Min  float32 `json:"min" form:"min"`
	Diff float32 `json:"diff" form:"diff"`
}

//ResponseData : response format
type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

//ErrorField : 422 error validation format
type ErrorField struct {
	ID      string `json:"id"`
	Value   string `json:"value"`
	Caused  string `json:"caused"`
	Message string `json:"message"`
}

//RespondJSON : send a proper response json to the client
func RespondJSON(w *gin.Context, status int, message string, payload interface{}, errors interface{}) {
	var res ResponseData

	if status >= 200 && status < 300 {
		res.Success = true
	}

	if len(message) != 0 {
		res.Message = message
	}

	if payload != nil {
		res.Data = payload
	}

	if errors != nil {
		res.Error = errors
	}

	w.JSON(status, res)
}
