package api

import (
	"github.com/gin-gonic/gin"
	"mimigo/serializer"
	"mimigo/service"
	"net/http"
)

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerService service.UserRegisterService
	if err := c.ShouldBindJSON(&registerService); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err).TimeMarked())
		return
	}

	if err := registerService.Register(); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.DbCreateError,
			Message: err.Error(),
			Content: nil,
			Error:   err.Error(),
		}.TimeMarked())
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.Success,
			Message: "用户注册成功",
		}.TimeMarked())
	}

}

// Login 登录
func Login(c *gin.Context) {
	var loginService service.UserLoginService
	if parseErr := c.ShouldBindJSON(&loginService); parseErr != nil {
		c.JSON(http.StatusOK, ErrorResponse(parseErr).TimeMarked())
		c.Abort()
		return
	}

	accessToken, err := loginService.GenerateJwtToken(c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.CreateTokenError,
			Message: err.Error(),
			Content: nil,
			Error:   err.Error(),
		}.TimeMarked())
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.Success,
			Content: accessToken,
		}.TimeMarked())
	}

	return
}

func Logout(c *gin.Context) {
	var logoutService service.UserLogoutService
	userInfo := AuthUser(c)
	if err := logoutService.Logout(userInfo); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			ErrCode: serializer.DbDeleteError,
			Message: "注销登录状态失败",
			Content: nil,
			Error:   err.Error(),
		}.TimeMarked())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		ErrCode: serializer.Success,
		Message: "注销登录状态成功",
	})

}

func GetUserInfo(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, serializer.Response{
				ErrCode: serializer.Success,
				Error:   r,
			}.TimeMarked())
		}
	}()

	user := AuthUser(c)
	c.JSON(http.StatusOK, serializer.Response{
		ErrCode: serializer.Success,
		Content: user,
	}.TimeMarked())
}
