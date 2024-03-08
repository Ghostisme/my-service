package routers

import (
	_ "my-project-admin-service/docs"
	"my-project-admin-service/global"
	"my-project-admin-service/internal/middleware"
	v1 "my-project-admin-service/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.ServerSettings.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Recovery())
	r.MaxMultipartMemory = 8 * 1024
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiDirect := r.Group("api/v1")
	{
		//登录验证
		apiDirect.POST("/login", v1.GetUser)
	}
	global.Logger.Infof("查看login接口数据: %v", v1.GetUser)
	return r
}