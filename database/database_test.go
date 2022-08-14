package database

import (
	"testing"

	"github.com/bonggar/gorestapi/model"
	"github.com/gin-gonic/gin"
)

func TestDBInitialization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	DatabaseInit()
	if database != nil {
		t.Logf("DB Connection Success")
		if database.HasTable(model.User{}) {
			t.Logf("DB Migration Success")
		} else {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
