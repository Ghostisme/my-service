package routers

import (
	_ "my-service/docs"
	"my-service/global"
	"my-service/internal/middleware"
	v1 "my-service/internal/routers/api/v1"

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
		apiDirect.POST("/login", v1.Login)
		// 获取验证码
		apiDirect.POST("/code", v1.CreateCode)
		// 注册
		apiDirect.POST("/register", v1.Register)
	}
	// global.Logger.Infof("查看login接口数据: %v", v1.GetUser)
	apiV1 := r.Group("api/v1")
	user := v1.NewUser()
	role := v1.NewRole()
	apiV1.Use(middleware.JWT()) // 增加token有效性验证
	{
		// 登出
		apiV1.GET("/logout", v1.Logout)
		// 注册
		// apiV1.POST("/register", )
		// 获取用户列表
		apiV1.POST("/user", user.List)
		// 获取角色列表
		apiV1.POST("/role", role.List)
		// 创建角色
		apiV1.POST("/role/save", role.Create)
		// 编辑角色
		apiV1.PUT("/role", role.Update)
		// 删除角色
		apiV1.DELETE("/role", role.Del)
		// 创建用户
		apiV1.POST("/user/save", user.Create)
		// 编辑用户
		apiV1.PUT("/user", user.Update)
		// 删除用户
		apiV1.DELETE("/user", user.Del)
	}

	return r
}
