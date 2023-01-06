package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin = errors.New("user not logged in")
	ContextUserIDKey  = "userID"
)

// GetCurrentUserID 调用这个就可以获取到用户登录后的 userID
func GetCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
