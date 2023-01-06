package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 返回响应的固定格式
type ResponseData struct {
	Code    ResCode `json:"code"`
	Message any     `json:"message"`
	Data    any     `json:"data"`
}

// // 这里使用方法感觉不是很好
// func (r *ResponseData) ResponseError(c *gin.Context) {
// 	rd := &ResponseData{
// 		Code:    r.Code,
// 		Message: getMsg(r.Code),
// 		Data:    nil,
// 	}
// 	c.JSONP(http.StatusOK, rd)
// }

// Error 返回错误
func Error(c *gin.Context, code ResCode) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

// Success 返回正常
func Success(c *gin.Context, data ...any) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}

// ErrorWithMsg 返回自定义错误信息
func ErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
