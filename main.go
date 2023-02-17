package main

import (
	api "jump-backend-interview/src/api"
	db "jump-backend-interview/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	println("Starting server...");
	err := db.GetDb().Ping();
	if err != nil {
		panic(err)
	}
	router := gin.New()
	router.SetTrustedProxies(nil)
	api.InitApi(router)
	router.Run("0.0.0.0:8080")
}