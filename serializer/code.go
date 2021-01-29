package serializer

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	UserInputError   = 40003
	DbCreateError    = 50001
	DbUpdateError    = 50002
	DbDeleteError    = 50003
	DbQueryError     = 50004
	DbRecordNotFound = 50005
	// CodeEncryptError 加密失败
	CodeEncryptError = 51000
)

const (
	//	需要登陆
	AccessDenied = 401
	// 	未授权访问
	CodeNoRightErr = 403
)

const (
	//	不论怎样，反正错了
	ErrorAnyway = 40001
	//	手机号已经存在
	MobileExist = 10002
	//	创建token失败
	CreateTokenError = 10003
	//	token过期了
	TokenExpired = 10006
)
