package main

import (
	"fmt"
	"jump-backend-interview/src/api"
	"jump-backend-interview/src/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Starting server...")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Default().Println(err.Error())
		os.Exit(-1)
	}
	router := gin.New()
	router.SetTrustedProxies(nil)
	api.InitApi(router)
	router.Run(fmt.Sprintf("%s:%d", cfg.Server.Ip, cfg.Server.Port))
}
