package controller

import (
	"net/http"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/service"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	Create(ctx *gin.Context)
	CreateStandard(ctx *gin.Context)
}

type paymentController struct {
	paymentService service.PaymentService
	midtransService service.MidtransService
}

func NewPaymentController (ps service.PaymentService, m service.MidtransService) PaymentController {
	return &paymentController{
		paymentService: ps,
		midtransService: m,
	}
}

func (pc *paymentController) Create(ctx *gin.Context) {
	var paymentReq dto.CreatePaymentRequest

	if err := ctx.ShouldBind(&paymentReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_CREATE_PAYMENT_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	// fmt.Println(paymentReq)

	response, err := pc.midtransService.GenerateSnapUrl(paymentReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_CREATE_PAYMENT_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_CREATE_PAYMENT_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

func (pc *paymentController) CreateStandard(ctx *gin.Context){
	var paymentReq dto.CreatePaymentRequest
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if err := ctx.ShouldBind(&paymentReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_CREATE_PAYMENT_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := pc.paymentService.CreateStandard(paymentReq, userId.(string))

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_CREATE_PAYMENT_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_CREATE_PAYMENT_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}