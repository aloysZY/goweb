package user

import (
	"database/sql"

	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/pkg/scrypt"
	"go.uber.org/zap"
)

// 逻辑层，负责处理业务逻辑，存放业务逻辑代码

// 想让处理的错误，都放在这层进行处理,调用的时候处理错误,在这层记录错误日志

// Login 登录
func Login(p *model.LoginUser) (err error) {
	// 根据用户输入信息构建结构体，存储一下密码就行，后面进行密码匹配验证
	user := &model.LoginUser{
		Password: p.Password,
	}
	// 根据
	if err = mysql.GetPassword(p); err == sql.ErrNoRows {
		// 基本就是登录失败，直接返回
		zap.L().Warn(logic.ErrorUserNotExist, zap.String(logic.Username, p.UserName), zap.Error(err))
		return
	} else if err != nil {
		zap.L().Error(logic.ErrorQueryFailed, zap.String(logic.Username, p.UserName), zap.Error(err))
		return
	}
	// 验证用户输入的密码是否正确 user.Password 数据库密码，p.Password 用户输入密码
	if err = scrypt.ComparePasswords(p.Password, user.Password); err != nil {
		zap.L().Warn(logic.ErrorInvalidPassword, zap.String(logic.Username, p.UserName), zap.String(logic.Password, user.Password), zap.Error(err))
		return
	}
	zap.L().Info(logic.Success, zap.String(logic.Username, p.UserName), zap.String(logic.Password, user.Password))
	return
}
