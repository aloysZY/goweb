package login

import (
	"net/http"

	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/internal/settings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 登录请求处理
func LoginHandler(c *gin.Context) {
	// 验证登录请求参数
	p := new(model.ParamLoginUser)
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("login with invalid parameters", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			// 如果类型断言，发现错误不是可以反应的错误，直接返回
			c.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"msg":  "请求参数有误(不翻译)",
			})
			return
		} else {
			// 否则进行翻译
			c.JSON(http.StatusOK, gin.H{
				// "code": errs.Translate(settings.Trans),
				// 进一步删除结构体信息，根据需求来修改
				"code": settings.RemoveTopStruct(errs.Translate(settings.Trans)),
				"msg":  "请求参数有误(翻译)",
			})
			return
		}
	}

	// 验证登录密码
	// 业务处理,到这里，绑定成功，参数校验也基本满足要求,执行业务逻辑处理
	if err := logic.Login(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": err.Error(),
			"msg":  "登录失败",
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "login success",
	})
}
