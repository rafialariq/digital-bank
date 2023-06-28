package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/service"
)

type RegisterController struct {
	registerService service.RegisterService
}

func NewRegisterController(r *gin.RouterGroup, rs service.RegisterService) *RegisterController {
	controller := RegisterController{
		registerService: rs,
	}
	r.POST("/signup", controller.RegisterHandler)
	return &controller
}

func (r *RegisterController) RegisterHandler(ctx *gin.Context) {
	var newUser dto.RegisterDTO

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := r.registerService.CreateUser(&newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "new user created successfully"})
}
