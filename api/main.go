package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mimigo/config"
	"mimigo/database/collections"
	"mimigo/serializer"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				ErrCode: serializer.UserInputError,
				Message: fmt.Sprintf("%s%s", field, tag),
				Error:   fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			ErrCode: serializer.UserInputError,
			Message: "JSON类型不匹配",
			Error:   fmt.Sprint(err),
		}
	}

	return serializer.Response{
		ErrCode: serializer.UserInputError,
		Message: "参数错误",
		Error:   fmt.Sprint(err),
	}
}

func AuthUser(c *gin.Context) collections.LoginUser {
	return c.MustGet("user").(collections.LoginUser)
}
