package server

import (
	"Flamingo/api"
	"Flamingo/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		// 获得token
		v1.POST("/register", api.RegisterUser)
		v1.POST("/login", api.Login)
		// 使用中间件验证.
		jwt := v1.Group("")
		jwt.Use(middleware.TokenAuth())
		{
			jwt.GET("/account/info", api.GetUserInfo)
			jwt.GET("/account/logout", api.Logout)
			jwt.POST("/aliyun/oss/token", api.AliyunOssSTSToken)
		}

	}

	return r
}
