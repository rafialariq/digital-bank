package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/service"
)

type LoginController struct {
	loginService service.LoginService
}

func NewLoginController(r *gin.RouterGroup, ls service.LoginService) *LoginController {
	controller := LoginController{
		loginService: ls,
	}
	r.GET("/login", controller.LoginHandler)
	return &controller
}

func (l *LoginController) LoginHandler(ctx *gin.Context) {
	var user dto.LoginDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := l.loginService.FindUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": res,
	})
}
