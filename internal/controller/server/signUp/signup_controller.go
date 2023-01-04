package signUp

import (
	"net/http"

	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/controller/response"
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/aloysZy/goweb/internal/settings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 业务路由，参数效验，请求转发

// 返回信息有一个大的原则，就是放回前端的时候，都是返回一个大概的信息，或者提示信息，后端自己记录日志，不能把错误信息都返回给前端

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 	获取参数参数效验，参数处理
	// 获取参数，简单的参数类型效验，根据结构体定义的类型，传入的类型不匹配会出错
	p := new(model.ParamSignUpUser)
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("signUp with invalid parameters", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 如果类型断言，发现错误不是可以反应的错误，直接返回
			response.Error(c, controller.CodeInvalidParams)
			return
		}
		// 否则进行翻译
		response.ErrorWithMsg(c, controller.CodeInvalidParams, settings.RemoveTopStruct(errs.Translate(settings.Trans)))
		// c.JSON(http.StatusOK, gin.H{
		// 	// "code": errs.Translate(settings.Trans),
		// 	// 进一步删除结构体信息，根据需求来修改
		// 	"code": settings.RemoveTopStruct(errs.Translate(settings.Trans)),
		// 	"msg":  "请求参数有误(翻译)",
		// })
		return
	}
	// 业务处理,到这里，绑定成功，参数校验也基本满足要求,执行业务逻辑处理
	if err := logic.SignUp(p); err != nil {
		// zap.L().Error("用户注册失败", zap.String("username", p.Username), zap.Error(err))
		response.Error(c, controller.CodeRegistrationFailed)
		// c.JSON(http.StatusOK, gin.H{
		// 	"code": err.Error(),
		// 	"msg":  "注册失败",
		// })
		return
	}
	// 整体流程不存问题返回成功
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}
