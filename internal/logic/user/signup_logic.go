package user

import (
	"errors"
	"fmt"

	"github.com/aloysZy/goweb/global/conf"
	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/pkg/scrypt"
	"github.com/aloysZy/goweb/pkg/snowflake"
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
		zap.L().Error(logic.ErrorQueryFailed, zap.Error(err))
		return err
	} else if count > 0 {
		zap.L().Error(logic.ErrorUserExist, zap.String(logic.Username, p.Username), zap.Error(err))
		return ErrorUserExit
	}
	// 初始化
	if err := snowflake.Snowflake(conf.Config.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return err
	}
	// 生成 UID
	userID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error(logic.ErrorCreateUserId, zap.Error(err))
		return
	}
	// 密码存入之前进行加密处理
	newPassword, err := scrypt.HashAndSalt(p.Password)
	if err != nil {
		zap.L().Error(logic.ErrorEncryptionPassword, zap.String(logic.Username, p.Username), zap.Error(err))
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
		zap.L().Error(logic.ErrorInsertFailed, zap.Error(err))
		return
	}
	zap.L().Info(logic.Success, zap.String(logic.Username, p.Username))
	return
}
