package logic

import (
	"errors"

	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/internal/settings"
	"go.uber.org/zap"
)

// 逻辑层，负责处理业务逻辑，存在业务逻辑代码

var (
	ErrorUserExit = errors.New("用户名重复")
)

// SignUp 注册
func SignUp(p *model.ParamSignUpUser) (err error) {
	// 判断用户是否存在
	if count, err := mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error(errorQueryFailed, zap.Error(err))
		return err
	} else if count > 0 {
		zap.L().Error(errorUserExist, zap.String(username, p.Username), zap.Error(err))
		return ErrorUserExit
	}
	// 生成 UID
	userID, err := settings.GetID()
	if err != nil {
		zap.L().Error(errorCreateUserId, zap.Error(err))
		return
	}
	// 密码存入之前进行加密处理
	newPassword, err := settings.HashAndSalt(p.Password)
	if err != nil {
		zap.L().Error(errorEncryptionPassword, zap.String(username, p.Username), zap.Error(err))
		return
	}
	// 存入数据到数据库，查询数据库
	// 初始化
	user := &model.SignUpUser{
		UserID:   userID,
		UserName: p.Username,
		Password: newPassword,
		Email:    p.Email,
	}
	if err = mysql.InsertUser(user); err != nil {
		zap.L().Error(errorInsertFailed, zap.Error(err))
		return
	}
	zap.L().Info(success, zap.String(username, p.Username))
	return
}
