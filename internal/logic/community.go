package logic

import (
	"database/sql"

	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/model"
	"go.uber.org/zap"
)

// 调用数据库，查询社区列表
func GetCommunityList() (data []*model.Community, err error) {
	data, err = mysql.GetCommunityList()
	if err != nil {
		if err == sql.ErrNoRows {
			// 这里应该记录一个错误码，先不写了
			zap.L().Warn("查询数据库为空", zap.Error(err))
			return
		}
		zap.L().Error("查询数据库错误", zap.Error(err))
		return
	}
	return
}

func GetCommunityDetail(id string) (detail *model.CommunityDetail, err error) {
	detail, err = mysql.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		return
	}
	return
}
