package api

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"mimigo/config/aliyun"
	"net/http"
	"os"
)

func AliyunOssSTSToken(c *gin.Context) {
	client, err := oss.New(aliyun.OssEndpoint, aliyun.OssAccessKeyId, aliyun.OssAccessKeySecret, oss.SecurityToken(aliyun.OssSecurityToken))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)

		// OSS操作。
	}
	c.JSON(http.StatusOK, client)
}
