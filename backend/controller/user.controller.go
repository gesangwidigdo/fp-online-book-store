package controller

import (
	"net/http"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/service"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (uc *userController) GetAll(ctx *gin.Context) {
	// --- SEBAGAI CONTOH
	ctx.JSON(200, gin.H{
		"message": "Get all users",
	})
}

func (uc *userController) Register(ctx *gin.Context) {
	var userReq dto.UserRegisterReq

	if err := ctx.ShouldBind(&userReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := uc.userService.Register(userReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_REGISTER_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

func (uc *userController) Login(ctx *gin.Context) {
	var userReq dto.UserLoginReq

	if err := ctx.ShouldBind(&userReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_LOGIN_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := uc.userService.Login(userReq)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_LOGIN_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("accessToken", response.Token, 3600*24*1, "", "", false, true)

	res := utils.ResponseSuccess(dto.MSG_USER_LOGIN_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) Me(ctx *gin.Context) {
	userId, exists := ctx.Get("user")
	if !exists {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	user, err := uc.userService.Me(userId.(string))
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_AUTH_SUCCESS, user)
	ctx.JSON(http.StatusOK, res)
}
