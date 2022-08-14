package model

import "time"

//Berat : Berat struct
type Berat struct {
	ID        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Max       int       `json:"max" form:"max" gorm:"not null"`
	Min       int       `json:"min" form:"min" gorm:"not null"`
	Diff      int       `json:"diff" form:"diff" gorm:"default:null"`
	Date      time.Time `json:"date" form:"date" gorm:"type:date;unique;not null"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
