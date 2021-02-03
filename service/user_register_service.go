package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mimigo/database"
	"mimigo/database/collections"
	"mimigo/serializer"
	"time"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	WorkName string `form:"workName" json:"workName" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// valid 验证
func (userService *UserRegisterService) valid() error {
	filter := bson.D{{"mobile", userService.Mobile}}

	if exist, err := database.SupportUser.CountDocuments(context.TODO(), filter); err != nil {
		return err
	} else {
		if exist > 0 {
			return fmt.Errorf("%d 该手机号已经注册", serializer.MobileExist)
		}
	}

	return nil

}

// Register 用户注册
func (userService *UserRegisterService) Register() error {
	account := collections.User{
		Id:        primitive.NewObjectID(),
		WorkName:  userService.WorkName,
		Mobile:    userService.Mobile,
		Password:  userService.Password,
		Authority: "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 表单验证
	if err := userService.valid(); err != nil {
		return err
	}

	// 加密密码
	if err := account.SetPassword(userService.Password); err != nil {
		return err
	}

	go func() {
		_, _ = database.SupportUser.InsertOne(context.TODO(), account)
	}()

	return nil
}
