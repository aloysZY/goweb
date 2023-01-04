package logic

import (
	"errors"

	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/internal/settings"
	"go.uber.org/zap"
)

// 逻辑层，负责处理业务逻辑，存在业务逻辑代码

// SignUp 注册
func SignUp(p *model.ParamSignUpUser) error {
	// 判断用户是否存在
	if c, err := mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		zap.L().Error(controller.CodeQueryDatabase.Msg(), zap.String("username", p.Username), zap.Error(err))
		return err
	} else if c > 0 {
		zap.L().Error(controller.CodeUserExist.Msg(), zap.String("username", p.Username), zap.Error(err))
		return errors.New(controller.CodeRegistrationFailed.Msg())
	}
	// 生成 UID
	userID, err := settings.GetID()
	if err != nil {
		zap.L().Error("user_ID错误", zap.Error(err))
		return err
	}
	// 密码存入之前进行加密处理
	newPassword, err := settings.HashAndSalt(p.Password)
	if err != nil {
		zap.L().Error("密码注册加密错误", zap.String("username", p.Username), zap.Error(err))
		return err
	}
	// 存入数据到数据库，查询数据库
	user := &model.SignUpUser{
		UserID:   userID,
		UserName: p.Username,
		Password: newPassword,
		Email:    p.Email,
	}
	if err := mysql.InsertUser(user); err != nil {
		zap.L().Error("创建用户错误", zap.Error(err))
		return err
	}
	zap.L().Info("用户创建成功", zap.String("username", p.Username))
	return nil
}
