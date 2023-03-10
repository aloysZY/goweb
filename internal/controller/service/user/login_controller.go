package user

import (
	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/dao/redis"
	"github.com/aloysZy/goweb/internal/logic/user"
	"github.com/aloysZy/goweb/internal/model"
	"github.com/gin-gonic/gin"
)

// 这层返回前端错误，都是封装过的

// LoginHandler 登录请求处理
func LoginHandler(c *gin.Context) {
	// 验证登录请求参数
	p := new(model.LoginUser)
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	// 将用户输入的用户名和密码绑定到p
	if err := c.ShouldBind(p); err != nil {
		controller.ErrorWithMsg(c, controller.CodeInvalidParams, err.Error())
		return
		// if errs, ok := err.(validator.ValidationErrors); !ok {
		// 	// 如果类型断言，发现错误不是可以反应的错误，直接返回
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"code": err.Error(),
		// 		"msg":  "请求参数有误(不翻译)",
		// 	})
		// 	return
		// } else {
		// 	// 否则进行翻译
		// 	c.JSON(http.StatusOK, gin.H{
		// 		// "code": errs.Translate(settings.Trans),
		// 		// 进一步删除结构体信息，根据需求来修改
		// 		"code": settings.RemoveTopStruct(errs.Translate(settings.Trans)),
		// 		"msg":  "请求参数有误(翻译)",
		// 	})
		// 	return
		// }
	}
	// 验证登录密码
	// 业务处理,到这里，绑定成功，参数校验也基本满足要求,执行业务逻辑处理
	if err := user.Login(p); err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
		// c.JSON(http.StatusOK, gin.H{
		// 	"code": err.Error(),
		// 	"msg":  "登录失败",s
		// })
	}

	// fmt.Printf("p%v\n", p)
	// 还是感觉在登录里面直接做 token 吧,还是在这里调用吧
	aToken, rToken, err := user.GetLoginToken(p.UserId)
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
	}

	// 设置账户同一时间只能一个用户登录，存储userId 和对于的 atoken到 redis缓存中
	// jwt 认证后，从 redis 中匹配传入的 token，和存入是不是一致的
	// 存入 token 和 userid
	err = redis.SetToken(p.UserId, aToken)
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
	}
	// 返回响应
	controller.Success(c, aToken, rToken)
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":    200,
	// 	"message": "login success",
	// })
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	aToken, rToken, err := user.RefreshLoginToken(rt, authHeader)
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
	}
	controller.Success(c, aToken, rToken)
}
