package main

import (
	"github.com/gin-gonic/gin"
	"github.com/htetoonaing/go-gin-gorm-postgress-docker-rest-api/config"
	"github.com/htetoonaing/go-gin-gorm-postgress-docker-rest-api/routes"
)

func main() {
	router := gin.New()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Hello World")
	// })

	config.Connect()

	routes.UserRoute(router)

	router.Run(":8080")
}
