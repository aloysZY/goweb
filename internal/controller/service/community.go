package service

import (
	"github.com/aloysZy/goweb/internal/controller"
	"github.com/aloysZy/goweb/internal/logic"
	"github.com/gin-gonic/gin"
)

// CommunityHandler 社区相关函数
func CommunityHandler(c *gin.Context) {
	// 访问这个接口，返回现有社区的community_id,community_name,列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
	}
	controller.Success(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 	根据请求参数，获取id对于的帖子详细信息
	id := c.Param("id")
	detail, err := logic.GetCommunityDetail(id)
	if err != nil {
		controller.Error(c, controller.CodeServerBusy)
		return
	}
	controller.Success(c, detail)
}
