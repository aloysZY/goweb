package user

import (
	"database/sql"
	"fmt"

	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/pkg/jwt"
	"github.com/aloysZy/goweb/pkg/scrypt"
	"go.uber.org/zap"
)

// 逻辑层，负责处理业务逻辑，存在业务逻辑代码

// 想让处理的错误，都放在这层进行处理,调用的时候处理错误,在这层记录错误日志

// Login 登录
func Login(p *model.ParamLoginUser) (token string, err error) {
	// 根据用户输入信息构建结构体
	user := &model.LoginUser{
		UserName: p.Username,
		Password: p.Password,
	}
	// 查询用户密码
	if err = mysql.GetPassword(user); err == sql.ErrNoRows {
		// 基本就是登录失败，直接返回
		zap.L().Warn(logic.ErrorUserNotExist, zap.String(logic.Username, p.Username), zap.Error(err))
		return "", err
	} else if err != nil {
		zap.L().Error(logic.ErrorQueryFailed, zap.String(logic.Username, p.Username), zap.Error(err))
		return "", err
	}
	// 验证用户输入的密码是否正确 user.Password 数据库密码，p.Password 用户输入密码
	if err = scrypt.ComparePasswords(user.Password, p.Password); err != nil {
		zap.L().Warn(logic.ErrorInvalidPassword, zap.String(logic.Username, p.Username), zap.String(logic.Password, p.Password), zap.Error(err))
		return "", err
	}
	// 这样就找到了userid
	fmt.Printf("loginc userId= %d,username=%s\n", user.UserId, user.UserName)
	// 生成 token
	token, err = jwt.GenToken(user.UserId, user.UserName)
	if err != nil {
		zap.L().Error(logic.ErrorGetToken, zap.String(logic.Username, p.Username), zap.String(logic.Password, p.Password), zap.Error(err))
		return "", err
	}
	zap.L().Info(logic.Success, zap.String(logic.Username, p.Username), zap.String(logic.Password, p.Password))
	return
}
