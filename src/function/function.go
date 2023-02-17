package function

import (
	db "jump-backend-interview/src/database"
	"github.com/gin-gonic/gin"
)

func GetUsers(c * gin.Context) {
	db.GetDb();
	c.JSON(204, gin.H{
		"message": "pong",
	});
}

func PostInvoice(c * gin.Context) {
	db.GetDb();
	c.JSON(204, gin.H{
		"message": "pong",
	});
}

func PostTransaction(c * gin.Context) {
	db.GetDb();
	c.JSON(204, gin.H{
		"message": "pong",
	});
}
