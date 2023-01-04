package router

import (
	"github.com/aloysZy/goweb/internal/controller/server/login"
	"github.com/aloysZy/goweb/internal/controller/server/signUp"
	"github.com/aloysZy/goweb/internal/logger"
	"github.com/aloysZy/goweb/internal/settings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter() {
	// 设置model
	settings.SetModel()
	// 初始化 validator,这个是 gin记录翻译的包
	if err := settings.InitTrans(settings.Conf.Locale); err != nil {
		zap.L().Error("init Trans failed", zap.Error(err))
	}
	// 初始化 gin
	r := gin.Default()
	// 加载设置的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册
	r.POST("/signUp", signUp.SignUpHandler)
	// 登录
	r.POST("/login", login.LoginHandler)
	// 在这里启动 gin
	settings.Start(r)

}
