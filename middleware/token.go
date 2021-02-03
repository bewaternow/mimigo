package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mimigo/config"
	"mimigo/database"
	"mimigo/database/collections"
	"mimigo/serializer"
	"net/http"
	"strings"
	"time"
)

// TokenAuth 中间件，检查 token 是否存在，以及是否有效
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		usefulAccessToken := strings.Split(authorization, " ")

		if usefulAccessToken[0] != "Bearer" || len(usefulAccessToken) < 2 {
			c.JSON(http.StatusOK, serializer.Response{
				ErrCode: serializer.AccessDenied,
				Message: "无访问权限",
			})
			c.Abort()
			return
		}

		var tokenRecord collections.PersonalAccessToken
		//	令牌有效的两个条件：1、记录存在；2、令牌没有过期
		if err := database.SupportPersonalAccessToken.FindOne(context.Background(), bson.D{{"ipAddress", c.ClientIP()}, {"userAgent", c.GetHeader("User-Agent")}, {"token", usefulAccessToken[1]}, {"expired_at", bson.D{{"$gt", time.Now()}}}}).Decode(&tokenRecord); err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				ErrCode: serializer.AccessDenied,
				Message: "无效的令牌",
			}.TimeMarked())
			c.Abort()
			return
		}

		var userRecord collections.User
		if err := database.SupportUser.FindOne(context.Background(), bson.M{"_id": tokenRecord.UserId}).Decode(&userRecord); err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				ErrCode: serializer.AccessDenied,
				Message: "无效的令牌",
			}.TimeMarked())
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		//c.Set("claims", claims)
		c.Set("user", collections.LoginUser{
			Id:        userRecord.Id,
			WorkName:  userRecord.WorkName,
			Mobile:    userRecord.Mobile,
			Authority: userRecord.Authority,
			TokenId:   tokenRecord.Id,
		})
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired ")
	TokenNotValidYet error  = errors.New("Token not active yet ")
	TokenMalformed   error  = errors.New("That's not even a token ")
	TokenInvalid     error  = errors.New("Couldn't handle this token: ")
	SignKey          string = config.SigningKey
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	WorkName  string             `json:"workName" bson:"workName"`
	Mobile    string             `json:"mobile" bson:"mobile"`
	Authority string             `json:"authority" bson:"authority"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
