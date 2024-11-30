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

func RoleAllow(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// --- Dapetin cookie
		tokenString, err := ctx.Cookie("accessToken")

		if err != nil {
			res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
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
				res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, dto.ERR_TOKEN_EXP_TIME)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			if time.Now().Unix() > int64(exp) {
				res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, dto.ERR_TOKEN_EXP)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			// Cek `sub`
			userType, ok := claims["role"].(string)
			if !ok {
				res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, dto.ERR_TOKEN_USER_ID)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			ctx.Set("role", userType)

			roleMatch := false
			for _, allowedRole := range allowedRoles {
				if userType == allowedRole {
					roleMatch = true
					break
				}
			}

			if !roleMatch {
				res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, "user type unauthorized")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			ctx.Next()
		} else {
			res := utils.ResponseFailed(dto.MSG_AUTH_FAILED, dto.ERR_INVALID_TOKEN)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
}
