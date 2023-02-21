package api

import (
	"jump-backend-interview/src/function"

	"github.com/gin-gonic/gin"
)

// It creates a route for the GET and POST requests.
func Routes(c *gin.Engine) {
	c.GET("/users", function.GetUsers)
	c.POST("/invoice", function.PostInvoice)
	c.POST("/transaction", function.PostTransaction)
}
