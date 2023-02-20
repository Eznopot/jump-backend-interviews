package function

import (
	"jump-backend-interview/src/database"
	"jump-backend-interview/src/model"

	"github.com/gin-gonic/gin"
)

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

func PostInvoice(c *gin.Context) {
	invoice := model.Invoice{}
	err := c.Bind(&invoice)
	if err != nil {
		println(err.Error())
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

// 404 si pas trouvé, 400 si pas le bon motant, 422 si deja payé 500 si erreur de la base de données, 204 si ok
func PostTransaction(c *gin.Context) {
	transaction := model.Transaction{}
	err := c.Bind(&transaction)
	if err != nil {
		println(err.Error())
		c.JSON(501, gin.H{
			"message": "bad arguments",
		})
		return
	}
	code := database.PostTransaction(transaction)
	if code != 204 {
		c.JSON(code, gin.H{
			"message": code,
		})
		return
	}
	c.Data(code, "application/json", nil)
}
