package user

import (
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/pkg/jwt"
	"go.uber.org/zap"
)

func GetToken(userId uint64) (aToken, rToken string, err error) {
	// 这样就找到了userid
	// fmt.Printf("loginc userId= %d,username=%s\n", user.UserId, user.UserName)
	// 生成 token
	aToken, rToken, err = jwt.GenToken(userId)
	if err != nil {
		zap.L().Error(logic.ErrorGetToken, zap.Uint64("userId", userId), zap.Error(err))
		return
	}
	return
}
