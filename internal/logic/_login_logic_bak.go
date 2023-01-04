package logic

// import (
// 	"errors"
//
// 	"github.com/aloysZy/goweb/internal/dao/mysql"
// 	"github.com/aloysZy/goweb/internal/model"
// 	"github.com/aloysZy/goweb/internal/settings"
// 	"go.uber.org/zap"
// )
//
// // 逻辑层，负责处理业务逻辑，存在业务逻辑代码
//
// // Login 登录
// func Login(p *model.ParamLoginUser) (err error) {
// 	// 判断用户是否存在
// 	if c, err := mysql.CheckUserExist(p.Username); err != nil {
// 		// 数据库查询出错
// 		zap.L().Error("查询用户错误", zap.Error(err))
// 		// 这里任务用户还没注册，专挑到注册api，这是后端做的还是前端做的？
// 		return err
// 	} else if c > 0 {
// 		// 构建用户的信息
// 		user := &model.LoginUser{
// 			UserName: p.Username,
// 		}
// 		// 查询用户密码
// 		if err = mysql.GetPassword(user); err != nil {
// 			return err
// 		}
// 		// 验证用户输入的密码是否正确 user.Password 数据库密码，p.Password 用户输入密码
// 		if b := settings.ComparePasswords(user.Password, p.Password); b != true {
// 			zap.L().Error("密码错误", zap.String("Username", p.Username), zap.String("Password", p.Password), zap.Error(err))
// 			return errors.New("密码错误")
// 		}
// 		zap.L().Info("用户登录成功", zap.String("Username", p.Username), zap.String("Password", p.Password))
// 		return nil
// 	}
// 	zap.L().Info("用户不存在", zap.String("Username", p.Username))
// 	return errors.New("用户不存在")
// }