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

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	config.DB.Create(&user)
	ctx.JSON(200, &user)
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", ctx.Param("id")).First(&user)
	ctx.BindJSON(&user)
	config.DB.Save(&user)
	ctx.JSON(200, &user)
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", ctx.Param("id")).Delete(&user)
	ctx.JSON(200, &user)
}
