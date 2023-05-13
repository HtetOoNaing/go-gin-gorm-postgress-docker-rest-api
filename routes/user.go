package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/htetoonaing/go-gin-gorm-postgress-docker-rest-api/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetAllUsers)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
