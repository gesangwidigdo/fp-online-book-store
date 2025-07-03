package controller

import (
	"net/http"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/service"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookTransactionController interface {
	Create(ctx *gin.Context)
	GetByTransactionID(ctx *gin.Context)
	UpdateQuantity(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookTransactionController struct {
	bookTransactionService service.BookTransactionService
}

func NewBookTransactionController(bts service.BookTransactionService) BookTransactionController {
	return &bookTransactionController{
		bookTransactionService: bts,
	}
}

func (btc *bookTransactionController) Create(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var bookTransactionReq dto.BookTransactionReq

	if err := ctx.ShouldBind(&bookTransactionReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := btc.bookTransactionService.Create(userId.(string), bookTransactionReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_TRANSACTION_CREATE_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

func (btc *bookTransactionController) GetByTransactionID(ctx *gin.Context) {
	transactionID := ctx.Param("transaction_id")
	uuidTransactionID, err := uuid.Parse(transactionID)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_GET_ALL_FAILED, "Invalid transaction ID format")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	bookTransactions, err := btc.bookTransactionService.GetByTransactionID(uuidTransactionID)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_GET_ALL_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_TRANSACTION_GET_ALL_SUCCESS, bookTransactions)
	ctx.JSON(http.StatusOK, res)
}

func (btc *bookTransactionController) UpdateQuantity(ctx *gin.Context) {
	transaction_id := ctx.Param("transaction_id")

	uuid_transaction_id, err := uuid.Parse(transaction_id)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_UPDATE_FAILED, "Invalid transaction ID format")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var bookTransactionReq dto.BookTransactionReq

	if err := ctx.ShouldBind(&bookTransactionReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := btc.bookTransactionService.UpdateQuantity(uuid_transaction_id, bookTransactionReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_TRANSACTION_UPDATE_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}

func (btc *bookTransactionController) Delete(ctx *gin.Context) {
	transaction_id := ctx.Param("transaction_id")
	
	uuid_transaction_id, err := uuid.Parse(transaction_id)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_UPDATE_FAILED, "Invalid transaction ID format")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	
	var bookTransactionReq dto.BookTransactionDeleteReq

	if err := ctx.ShouldBind(&bookTransactionReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_DELETE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deleteErr := btc.bookTransactionService.Delete(uuid_transaction_id, bookTransactionReq)
	if deleteErr != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_TRANSACTION_DELETE_FAILED, deleteErr.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_TRANSACTION_DELETE_SUCCESS, nil)
	ctx.JSON(http.StatusOK, res)
}
