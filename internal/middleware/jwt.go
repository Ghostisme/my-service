package middleware

import (
	"my-service/pkg/app"
	"my-service/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			errCode = errcode.Success
		)
		//token放header内
		token = c.GetHeader("token")

		if token == "" {
			errCode = errcode.UnauthorizedTokenError
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					errCode = errcode.UnauthorizedTokenTimeout
				default:
					errCode = errcode.UnauthorizedTokenError
				}
			}
		}

		//直接回复错误码
		if errCode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(errCode)
			c.Abort()
			return
		}
		//继续请求
		c.Next()
	}
}
