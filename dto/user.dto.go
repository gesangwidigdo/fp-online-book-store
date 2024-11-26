package dto

import "errors"

type UserRegisterReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

type UserRegisterRes struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserLoginReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginRes struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserMeRes struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

const (
	MSG_USER_REGISTER_SUCCESS = "user registered successfully"
	MSG_USER_LOGIN_SUCCESS    = "user login successfully"
	MSG_AUTH_SUCCESS          = "authorized"

	MSG_USER_REGISTER_FAILED = "user registration failed"
	MSG_USER_LOGIN_FAILED    = "user login failed"
	MSG_AUTH_FAILED          = "unauthorized"
	MSG_INVALID_TOKEN_FAILED = "invalid token"
	MSG_METHOD_TOKEN_FAILED  = "unexpected signing method"

	MSG_USER_NOT_FOUND = "user not found"
)

var (
	ERR_USERNAME_ALREADY_EXISTS = errors.New("username already exists")
	ERR_EMAIL_ALREADY_EXISTS    = errors.New("email already exists")
	ERR_LOGIN                   = errors.New("invalid email or password")
	ERR_USER_NOT_FOUND          = errors.New("user not found")
	ERR_TOKEN_EXP               = errors.New("token expired")
	ERR_TOKEN_EXP_TIME          = errors.New("Invalid token expiration time")
	ERR_TOKEN_USER_ID           = errors.New("Invalid user ID in token")
	ERR_INVALID_TOKEN           = errors.New("Invalid token")
)
