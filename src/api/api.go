package api

import (
	"github.com/gin-gonic/gin"
)

// Initialize the API and middleware
func InitApi(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	//add more middleware here
	Routes(router)
}
