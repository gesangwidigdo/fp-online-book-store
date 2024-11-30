package controller

import (
	"net/http"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/service"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Create(ctx *gin.Context)
	GetTransactionStatus(ctx *gin.Context)
	GetAllTransactionByUserLogin(ctx *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(ts service.TransactionService) TransactionController {
	return &transactionController{
		transactionService: ts,
	}
}

func (tc *transactionController) Create(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var transactionReq dto.TransactionCreateRequest

	if err := ctx.ShouldBind(&transactionReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_TRANSACTION_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := tc.transactionService.CreateTransaction(userId.(string))

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_TRANSACTION_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_TRANSACTION_CREATE_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

func (tc *transactionController) GetTransactionStatus(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	response, err := tc.transactionService.GetTransactionStatus(userId.(string))

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_TRANSACTION_STATUS_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_TRANSACTION_STATUS_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transactionController) GetAllTransactionByUserLogin(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	response, err := tc.transactionService.GetAllTransactionByUserLogin(userId.(string))

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_TRANSACTION_STATUS_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_TRANSACTION_STATUS_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}