package config

import (
	"Flamingo/database"
	"Flamingo/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var (
	SigningKey     = "PersistUntilSucceed"
	MongoDefaultDB = "Support"
	TokenLife      = time.Hour * 24 * 7
)

func Load() {
	//	从本地读取环境变量
	_ = godotenv.Load()

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//	设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}
	//	连接数据库
	database.Mongo()

}
