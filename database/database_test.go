package database

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jagadwp/gin-crud-web/model"
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
