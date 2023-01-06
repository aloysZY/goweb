package user

import (
	"errors"
	"strings"

	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/pkg/jwt"
	"go.uber.org/zap"
)

var (
	errorAuthToken = errors.New("请求头缺少Auth Token")

	errorInvalidAuthFormat = errors.New("token格式不对")
)

func RefreshToken(rt, authHeader string) (aToken, rToken string, err error) {
	// rt := c.Query("refresh_token")
	// // 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// // 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// // 这里的具体实现方式要依据你的实际业务情况决定
	// authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		// controller.Error(c, controller.CodeInvalidAuthFormat, "请求头缺少Auth Token")
		zap.L().Error(logic.ErrorAuthToken)
		// c.Abort()
		err = errorAuthToken
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		// ResponseErrorWithMsg(c, controller.CodeInvalidAuthFormat, "Token格式不对")
		zap.L().Error(logic.ErrorInvalidAuthFormat)
		// c.Abort()
		err = errorInvalidAuthFormat
		return
	}
	aToken, rToken, err = jwt.RefreshToken(parts[1], rt)
	// fmt.Println(err)
	// c.JSON(http.StatusOK, gin.H{
	// 	"access_token":  aToken,
	// 	"refresh_token": rToken,
	// })
	return
}
