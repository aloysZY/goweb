package router

import (
	"fmt"
	"net/http"

	"github.com/aloysZy/goweb/global/conf"
	"github.com/aloysZy/goweb/internal/controller/middleare"
	"github.com/aloysZy/goweb/internal/controller/service/user"
	"github.com/aloysZy/goweb/internal/logger"
	"github.com/aloysZy/goweb/internal/settings"
	"github.com/aloysZy/goweb/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter() {
	// 设置model
	settings.SetModel()
	// 初始化 validator,这个是 gin记录翻译的包
	if err := validator.InitTrans(conf.Config.Locale); err != nil {
		zap.L().Error("init Trans failed", zap.Error(err))
	}
	// 初始化 gin
	r := gin.Default()
	// 加载设置的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册
	r.POST("/signUp", user.SignUpHandler)
	// 登录
	r.POST("/login", user.LoginHandler)
	// 在这里启动 gin

	// 判断当前登录的用户是否是登录用户，判断请求头是否有有效的 token
	r.POST("/xxx", middleare.JWTAuthMiddleware(), func(c *gin.Context) {
		userId := c.MustGet("userId").(uint64)
		fmt.Printf("userid=%v\n", userId)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"user_id": userId},
		})
	})

	settings.Start(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

}
