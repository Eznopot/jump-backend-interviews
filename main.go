package main

import (
	api "jump-backend-interview/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	println("Starting server...");
	router := gin.New()
	router.SetTrustedProxies(nil)
	api.InitApi(router)
	router.Run("0.0.0.0:8080")
}