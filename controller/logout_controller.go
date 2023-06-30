package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutController struct{}

func NewLogoutController(r *gin.RouterGroup) *LogoutController {
	controller := LogoutController{}
	r.GET("/logout", controller.LogoutHandler)
	return &controller
}

func (l *LogoutController) LogoutHandler(ctx *gin.Context) {

	ctx.SetCookie("token", "", -1, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"msg": "logout successful"})
}
