package aliyun

import (
	"os"
)

var (
	OssEndpoint = os.Getenv("OSS_ENDPOINT")
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	OssAccessKeyId     = os.Getenv("OSS_ACCESS_KEY_ID")
	OssAccessKeySecret = os.Getenv("OSS_ACCESS_KEY_SECRET")
	OssBucketName      = os.Getenv("OSS_BUCKET_NAME")
	OssSecurityToken	= os.Getenv("OSS_SECUITY_TOKEN")
)
