package helper

import "time"

type BeratRequest struct {
	Date string `json:"date" form:"date"`
	Max  int    `json:"max" form:"max"`
	Min  int    `json:"min" form:"min"`
}

type BeratResponse struct {
	ID   uint      `json:"id" form:"id"`
	Max  int       `json:"max" form:"max"`
	Min  int       `json:"min" form:"min"`
	Diff int       `json:"diff" form:"diff"`
	Date time.Time `json:"date" form:"date"`
}
