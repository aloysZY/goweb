package logic

import (
	"database/sql"
	"errors"

	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/internal/settings"
	"go.uber.org/zap"
)

// 逻辑层，负责处理业务逻辑，存在业务逻辑代码

// 想让处理的错误，都放在这层进行处理,调用的时候处理错误

// Login 登录
func Login(p *model.ParamLoginUser) (err error) {
	// 根据用户输入信息构建结构体
	user := &model.LoginUser{
		UserName: p.Username,
	}
	// 查询用户密码
	if err = mysql.GetPassword(user); err == sql.ErrNoRows {
		// 基本就是登录失败，直接返回
		zap.L().Info("用户不存在", zap.String("username", p.Username))
		return errors.New("用户不存在")
	} else if err != nil {
		zap.L().Info("数据库查询用户错误", zap.String("username", p.Username))
		return err
	}
	// 验证用户输入的密码是否正确 user.Password 数据库密码，p.Password 用户输入密码
	if b := settings.ComparePasswords(user.Password, p.Password); b != true {
		zap.L().Error("密码错误", zap.String("username", p.Username), zap.String("password", p.Password), zap.Error(err))
		return errors.New("密码错误")
	}
	zap.L().Info("用户登录成功", zap.String("username", p.Username), zap.String("password", p.Password))
	return nil
}
