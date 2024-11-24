package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	// --- Dapetin cookie
	tokenString, err := ctx.Cookie("accessToken")

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}

	// --- Decode
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_INVALID_TOKEN_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Cek `exp`
		exp, ok := claims["exp"].(float64)
		if !ok {
			res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid token expiration time")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if time.Now().Unix() > int64(exp) {
			res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Token expired")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		// Cek `sub`
		userId, ok := claims["sub"].(string)
		if !ok {
			res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid user ID in token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		ctx.Set("user", userId)
		ctx.Next()
	} else {
		res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid token")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}

	// --- Check exp

	// --- FindUser by user

	// Attach to req

	// --- Continue

	ctx.Next()
}
