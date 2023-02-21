package function

import (
	"jump-backend-interview/src/database"
	"jump-backend-interview/src/model"
	"log"

	"github.com/gin-gonic/gin"
)

// It gets all the users from the database and returns them in JSON format
func GetUsers(c *gin.Context) {
	users, code := database.GetUsers()
	if code != 200 {
		c.JSON(code, gin.H{
			"message": "can't get users from database",
		})
		return
	}
	c.JSON(code, users)
}

// It takes a JSON object from the request body, and if it's valid, it adds it to the database
func PostInvoice(c *gin.Context) {
	invoice := model.Invoice{}
	err := c.Bind(&invoice)
	if err != nil {
		log.Default().Println(err.Error())
		c.JSON(400, gin.H{
			"message": "bad arguments",
		})
		return
	}
	code := database.PostInvoice(invoice)
	if code != 204 {
		c.JSON(code, gin.H{
			"message": "can't post invoice to database",
		})
		return
	}
	c.Data(code, "application/json", nil)
}

// It takes a transaction from the request body, and then adds it to the database
func PostTransaction(c *gin.Context) {
	transaction := model.Transaction{}
	err := c.Bind(&transaction)
	if err != nil {
		log.Default().Println(err.Error())
		c.JSON(501, gin.H{
			"message": "bad arguments",
		})
		return
	}
	code := database.PostTransaction(transaction)
	if code != 204 {
		c.JSON(code, gin.H{
			"message": "can't post invoice to database",
		})
		return
	}
	c.Data(code, "application/json", nil)
}
