package logic

const (
	Success                 = "Success"
	CodeInvalidParams       = "请求参数错误"
	ErrorUserExist          = "用户名重复"
	ErrorUserNotExist       = "用户不存在"
	ErrorInvalidPassword    = "用户名或密码错误"
	ErrorGetToken           = "获取 token 错误"
	codeServerBusy          = "服务繁忙"
	ErrorQueryFailed        = "查询数据库错误"
	ErrorCreateUserId       = "userId创建错误"
	ErrorEncryptionPassword = "密码加密错误"
	ErrorInsertFailed       = "插入数据库错误"
	ErrorInvalidToken       = "无效的Token"
	ErrorInvalidAuthFormat  = "认证格式有误"
	codeNotLogin            = "未登录"
	ErrorAuthToken          = "请求头缺少Auth Token"

	Username = "Username"
	Password = "Password"
)
