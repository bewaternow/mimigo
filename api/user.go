package api

import (
	"Flamingo/serializer"
	"Flamingo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerService service.UserRegisterService
	if bindErr := c.ShouldBindJSON(&registerService); bindErr != nil {
		c.JSON(http.StatusOK, ErrorResponse(bindErr).TimeMarked())
		return
	}

	response := registerService.Register()
	c.JSON(http.StatusOK, response)

}

// Login 登录
func Login(c *gin.Context) {
	var loginService service.UserLoginService
	if parseErr := c.ShouldBindJSON(&loginService); parseErr != nil {
		c.JSON(http.StatusOK, ErrorResponse(parseErr).TimeMarked())
		return
	}

	accessToken, err := loginService.GenerateJwtToken(c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.CreateTokenError,
			Message: err.Error(),
			Content: nil,
			Error:   err.Error(),
			ISODate: time.Now(),
		})
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: 0,
			Content: accessToken,
		})
	}

	return
}

func Logout(c *gin.Context)  {
	var logoutService service.UserLogoutService
	userInfo := AuthUser(c)
	if err := logoutService.Logout(userInfo);err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.DbDeleteError,
			Message: "注销登录状态失败",
			Content: nil,
			Error:   err.Error(),
			ISODate: time.Now(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		ErrCode: 0,
		Message: "注销登录状态成功",
	})

}

func GetUserInfo(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, serializer.Response{
				ErrCode: 0,
				Error:   r,
			})
		}
	}()

	user := AuthUser(c)
	c.JSON(http.StatusOK, serializer.Response{
		ErrCode: 0,
		Content: user,
	}.TimeMarked())
	return
}
