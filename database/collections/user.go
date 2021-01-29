package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	WorkName  string             `json:"workName" bson:"workName"`
	Mobile    string             `json:"mobile" bson:"mobile"`
	Authority string             `json:"authority" bson:"authority"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type LoginUser struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	WorkName  string             `json:"workName" bson:"workName"`
	Mobile    string             `json:"mobile" bson:"mobile"`
	Authority string             `json:"authority" bson:"authority"`
	TokenId   primitive.ObjectID `json:"token_id" bson:"token_id"`
}

func (user User) FormatSimpleUser() interface{}{
	return struct{
		Id        primitive.ObjectID `json:"id" bson:"_id"`
		WorkName  string             `json:"workName" bson:"workName"`
		Mobile    string             `json:"mobile" bson:"mobile"`
		Authority string             `json:"authority" bson:"authority"`
	}{
		Id:        user.Id,
		WorkName:  user.WorkName,
		Mobile:    user.Mobile,
		Authority: user.Authority,
	}
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
