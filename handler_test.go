package ginerror

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHandler(t *testing.T) {
	r := gin.Default()
	r.Use(Handler())
}

func TestHandlerCustom(t *testing.T) {
	r := gin.Default()
	r.Use(Handler(func(err *gin.Error) any {
		return gin.H{"error": err.Error()}
	}))
}
