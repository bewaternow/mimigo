package service

import (
	"Flamingo/config"
	"Flamingo/database"
	"Flamingo/database/collections"
	"Flamingo/middleware"
	"context"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UserLoginService struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

type UserLogoutService struct{}

// UserCheck 用户记录验证
func (userService UserLoginService) UserExist() (collections.User, error) {
	collection := database.MongoDB.Database(config.MongoDefaultDB).Collection(database.User)
	filter := bson.D{{"mobile", userService.Mobile}}

	if exist, err := collection.CountDocuments(context.TODO(), filter); err != nil {
		return collections.User{}, err
	} else {
		if exist > 0 {
			var account collections.User
			if err := collection.FindOne(context.TODO(), filter).Decode(&account); err != nil {
				return collections.User{}, err
			}
			return account, nil
		} else {
			return collections.User{}, fmt.Errorf("账户不存在")
		}

	}

}

type AuthInfo struct {
	AccessToken string `json:"access_token"`
	ExpiredAt  int64  `json:"expired_at"`
	Authority   string `json:"authority" bson:"authority"`
}

// 生成令牌
func (userService UserLoginService) GenerateJwtToken() (AuthInfo, error) {

	j := &middleware.JWT{
		SigningKey: []byte(config.SigningKey),
	}

	user, err := userService.UserExist()
	if err != nil {
		return AuthInfo{}, err
	}

	if !user.CheckPassword(userService.Password) {
		return AuthInfo{}, fmt.Errorf("账户或密码错误")
	}

	claims := middleware.CustomClaims{
		Id:        user.Id,
		WorkName:  user.WorkName,
		Mobile:    user.Mobile,
		Authority: user.Authority,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),                         // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*int64(config.TokenLife)), // 过期时间设置在配置文件中
			Issuer:    config.SigningKey,                                       //签名的发行者
		},
	}

	accessToken, err := j.CreateToken(claims)

	if err != nil {
		return AuthInfo{}, err
	}
	//	插入到数据库中
	collection := database.MongoDB.Database(config.MongoDefaultDB).Collection(database.PersonalAccessToken)

	tokenRecord := collections.PersonalAccessToken{
		UserId:    user.Id,
		Token:     accessToken,
	}.FormatToken()

	_, insertErr := collection.InsertOne(context.Background(), tokenRecord)

	if insertErr != nil {
		return AuthInfo{}, insertErr
	}

	return AuthInfo{
		accessToken,
		claims.ExpiresAt,
		user.Authority,
	}, nil
}

func(logoutService UserLogoutService) Logout(loginUser collections.LoginUser) error {
	//	把token从数据库中删除
	tokenCollection := database.MongoDB.Database(config.MongoDefaultDB).Collection(database.PersonalAccessToken)
	if result,err := tokenCollection.DeleteOne(context.Background(), bson.M{"_id": loginUser.TokenId}); err != nil {
		return err
	}else{
		fmt.Println(result)
	}

	return nil
}
