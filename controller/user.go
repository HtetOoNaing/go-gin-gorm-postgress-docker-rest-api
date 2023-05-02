package controller

import "github.com/gin-gonic/gin"

func GetAllUsers(ctx *gin.Context) {
	ctx.String(200, "Hello World From Controller")
}
