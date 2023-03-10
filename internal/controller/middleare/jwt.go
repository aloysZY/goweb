package middleare

import (
	"strings"

	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/logic/user"
	"github.com/aloysZy/goweb/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	errorAuthToken         = "请求头中auth为空"
	errorInvalidToken      = "无效的Token"
	errorInvalidAuthFormat = "请求头中auth格式有误"
	authInvalid            = "token认证成功"
)

// 这里记录日志好像只能在这里记录了，没有返回错误，并且是中间件也不需要返回错误

// JWTAuthMiddleware 认证 token
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			zap.L().Warn(errorInvalidAuthFormat)
			controller.Error(c, controller.CodeNotLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Warn(errorAuthToken)
			controller.Error(c, controller.CodeInvalidAuthFormat)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Warn(errorInvalidToken, zap.Error(err))
			controller.Error(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 刷新 token，前端触发，带着 rtoken 来访问，不在这里刷新
		// aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
		// if err != nil {
		// 	zap.L().Error(errorRefshTonek, zap.Error(err))
		// 	controller.Error(c, controller.CodeRefshToken)
		// 	c.Abort()
		// 	return
		// }

		// 这里获取当前token,查看是否和存入在 rdis 中的是一样的
		err = user.GetRedisToken(mc.UserID, parts[1])
		if err != nil {
			zap.L().Warn(errorInvalidToken, zap.Error(err))
			controller.Error(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		// fmt.Printf("middleware UserId = %v\n", mc.UserID)
		c.Set(controller.ContextUserIDKey, mc.UserID)
		zap.L().Debug(authInvalid)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
