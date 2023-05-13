package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/htetoonaing/go-gin-gorm-postgress-docker-rest-api/config"
	"github.com/htetoonaing/go-gin-gorm-postgress-docker-rest-api/models"
)

func GetAllUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(200, &users)
}
