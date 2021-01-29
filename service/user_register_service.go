package service

import (
	"Flamingo/config"
	"Flamingo/database"
	"Flamingo/database/collections"
	"Flamingo/serializer"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	WorkName string `form:"workName" json:"workName" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// valid 验证
func (userService *UserRegisterService) valid() serializer.Response {
	collection := database.MongoDB.Database(config.MongoDefaultDB).Collection(database.User)
	filter := bson.D{{"mobile", userService.Mobile}}

	if exist, err := collection.CountDocuments(context.TODO(), filter); err != nil {
		CodeDBError := serializer.Response{
			ErrCode: serializer.DbQueryError,
			Message: err.Error(),
		}
		return CodeDBError
	} else {
		if exist > 0 {
			mobileExist := serializer.Response{
				ErrCode: serializer.MobileExist,
				Message: "该手机号已经注册",
			}
			return mobileExist
		}
	}

	return serializer.Response{
		ErrCode: 0,
	}

}

// Register 用户注册
func (userService *UserRegisterService) Register() serializer.Response {
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
	if err := userService.valid(); err.ErrCode != 0 {
		return err
	}

	// 加密密码
	if err := account.SetPassword(userService.Password); err != nil {
		cryptRes := serializer.Response{
			ErrCode: serializer.CodeEncryptError,
			Message: err.Error(),
		}
		return cryptRes
	}

	collection := database.MongoDB.Database(config.MongoDefaultDB).Collection(database.User)
	go func() {
		_, _ = collection.InsertOne(context.TODO(), account)
	}()

	return serializer.Response{
		ErrCode: 0,
		Message: "success",
	}
}
