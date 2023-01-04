package response

import (
	"net/http"

	"github.com/aloysZy/goweb/internal/controller"
	"github.com/gin-gonic/gin"
)

// ResponseData 返回响应的固定格式
type ResponseData struct {
	Code    controller.ResCode `json:"code"`
	Message any                `json:"message"`
	Data    any                `json:"data"`
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

// 返回错误
func Error(c *gin.Context, code controller.ResCode) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

// 返回正常
func Successr(c *gin.Context, data any) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    controller.CodeSuccess,
		Message: controller.CodeSuccess.Msg(),
		Data:    data,
	})
}

// 返回自定义错误
func ErrorWithMsg(c *gin.Context, code controller.ResCode, msg any) {
	c.JSONP(http.StatusOK, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
