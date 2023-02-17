package api

import (
	"jump-backend-interview/src/function"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	c.GET("/users", function.GetUsers)

	c.POST("/invoice", function.PostInvoice)

	c.POST("/transaction", function.PostTransaction)
}
