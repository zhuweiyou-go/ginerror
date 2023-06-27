package ginerror

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Processor func(err *gin.Error) any

func Handler(processor ...Processor) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.Last()
		if lastError == nil {
			return
		}

		if len(processor) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": lastError.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, processor[0](lastError))
	}
}
