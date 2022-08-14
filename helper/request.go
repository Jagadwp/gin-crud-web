package helper

type BeratRequest struct {
	Date string `json:"date" form:"date"`
	Max  int    `json:"max" form:"max"`
	Min  int    `json:"min" form:"min"`
}
