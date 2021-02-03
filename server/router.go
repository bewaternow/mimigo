package server

import (
	"github.com/gin-gonic/gin"
	"mimigo/api"
	"mimigo/middleware"
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
		auth := v1.Group("")
		auth.Use(middleware.TokenAuth())
		{
			auth.GET("/user/profile", api.GetUserInfo)
			auth.GET("/user/logout", api.Logout)
			auth.POST("/aliyun/oss/token", api.AliyunOssSTSToken)
		}

	}

	return r
}
