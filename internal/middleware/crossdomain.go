package middleware

import (
	"fmt"
	"my-project-admin-service/global"
	"my-project-admin-service/pkg/app"
	"my-project-admin-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		fmt.Println(method, origin)
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型检验
		if method == "OPTIONS" {
			app.NewResponse(ctx).ToResponseSuccess()
		}

		defer func() {
			if err := recover(); err != nil {
				global.Logger.Infof("HttpError: ", err)
			}
		}()

		ctx.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				global.Logger.WithCallerFrames().Errorf("gin catch error: %v", r)
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
