package middleware

import (
	"my-service/pkg/app"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Cipher() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取加密指针地址
		cip := app.NewCaesar((15))
		fmt.Println("Encrypt Key(15) abcd =>", cip.Encryption("123456"))
		//直接回复错误码
		// if errCode != errcode.Success {
		// 	response := app.NewResponse(c)
		// 	response.ToErrorResponse(errCode)
		// 	c.Abort()
		// 	return
		// }
		//继续请求
		c.Next()
	}
}
