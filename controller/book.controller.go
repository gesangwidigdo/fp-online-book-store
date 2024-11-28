package controller

import (
	"net/http"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/service"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bs service.BookService) BookController {
	return &bookController{
		bookService: bs,
	}
}

func (bc *bookController) Create(ctx *gin.Context) {
	var bookReq dto.BookReq

	if err := ctx.ShouldBind(&bookReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := bc.bookService.Create(bookReq)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_CREATE_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

func (bc *bookController) GetAll(ctx *gin.Context) {
	books, err := bc.bookService.GetAll()
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_GET_ALL_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_GET_ALL_SUCCESS, books)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	book, err := bc.bookService.GetBySlug(slug)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_GET_BY_SLUG_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_GET_BY_SLUG_SUCCESS, book)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	convertedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_DELETE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var bookReq dto.BookReq

	if err := ctx.ShouldBind(&bookReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := bc.bookService.Update(convertedID, bookReq)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_UPDATE_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}

func (bc *bookController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	convertedID, err := uuid.Parse(id)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_DELETE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = bc.bookService.Delete(convertedID)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_BOOK_DELETE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_BOOK_DELETE_SUCCESS, nil)
	ctx.JSON(http.StatusOK, res)
}
