package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/service"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(r *gin.RouterGroup, ps service.PaymentService) *PaymentController {
	controller := PaymentController{
		paymentService: ps,
	}
	r.PUT("/payment", controller.PaymentHandler)
	return &controller
}

func (p *PaymentController) PaymentHandler(ctx *gin.Context) {

	var payment dto.PaymentDTO
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := p.paymentService.MakePayment(payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"msg": "Transaction Successfull",
	})

}
