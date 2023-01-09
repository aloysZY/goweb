package mysql

import (
	"github.com/aloysZy/goweb/internal/model"
)

// GetCommunityList 执行 sql 查询数据库
func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := "SELECT community_id,community_name from community"
	// err = db.Select(&communityList, sql)
	//
	// // for i, k := range communityList {
	// // 	fmt.Printf("GetCommunityList= \n%d\t%v\n", i, k)
	// // }
	//
	// if err != nil {
	// 	return
	// }
	// returns
	// dao层不想做判断
	// if err = db.Select(communityList,sqlStr);err != nil{
	// 	if err == sql.ErrNoRows{
	// 		return
	// 	}
	// 	return nil, err
	// }
	err = db.Select(&communityList, sqlStr)
	return
}

func GetCommunityDetail(id string) (detail *model.CommunityDetail, err error) {
	// idint, _ := strconv.Atoi(id)
	detail = new(model.CommunityDetail)
	// fmt.Printf("id = %T\n", id)
	sqlstr := `select community_id,community_name,introduction,create_time from community where community_id = ?`
	err = db.Get(detail, sqlstr, id)
	if err != nil {
		// fmt.Printf("GetCommunityDetail error: %v\n", err)
		return nil, err
	}
	// fmt.Println("GetCommunityDetail suess")
	// fmt.Printf("detail= %d\t%s\t%s\t%s\n", detail.Id, detail.Name, detail.Intr, detail.CreateTime)
	return
}
