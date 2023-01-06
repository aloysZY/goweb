package user

import (
	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/logic/user"
	"github.com/gin-gonic/gin"
)

// 这层应该是调用的 函数来进行判断

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	aToken, rToken, err := user.RefreshToken(rt, authHeader)
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
	}
	controller.Success(c, aToken, rToken)
}
