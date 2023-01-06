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

	// 这里面的路由都需要登录验证
	r.Use(middleare.JWTAuthMiddleware())
	{
		// 测试
		r.POST("/xxx", func(c *gin.Context) {
			userId := c.MustGet("userId").(uint64)
			fmt.Printf("route userid=%v\n", userId)
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": gin.H{"user_id": userId},
			})
		})
	}

	// 在这里启动 gin
	settings.Start(r)

	// 没有路由匹配
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

}
