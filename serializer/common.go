package serializer

import (
	"time"
)

// Response 基础序列化器
type Response struct {
	ErrCode uint        `json:"errCode"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
	Error   interface{} `json:"error"`
	ISODate time.Time   `json:"ISODate"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"trackId"`
}

func (response Response) TimeMarked() Response {
	response.ISODate = time.Now()
	return response
}
