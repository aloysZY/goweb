package mysql

// 用户登录，查询，注册相关的数据库操作
import (
	"github.com/aloysZy/goweb/internal/model"
)

// 每一步的数据库操作都进行封装，登录 logic的业务需求调用

// CheckUserExist 查看用户名是都存在
func CheckUserExist(u string) (count int64, err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	err = db.Get(&count, sqlStr, u)
	return
}

// InsertUser 注册用户,插入数据库
func InsertUser(user *model.SignUpUser) (err error) {
	sqlStr := `insert into user(user_id,username, password,email) values(?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password, user.Email)
	return
}

// GetPassword 查询用户密码
func GetPassword(user *model.LoginUser) (err error) {
	// 这要把 userid 查询出来赋值
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	return
}

// func GetPassword(user *model.LoginUser) error {
// 	sqlStr := "select username, password from user where username = ?"
// 	// err == sql.ErrNoRows感觉没必要，和之前查询用户存在不存在重复了，修改一下代码，直接在这里查询吧
// 	if err := db.Get(user, sqlStr, user.UserName); err == sql.ErrNoRows {
// 		zap.L().Info("用户不存在", zap.Error(err))
// 		// 返回的信息不返回sql: no rows in result set 就返回用户不存在
// 		return errors.New("用户不存在")
// 	} else if err != nil {
// 		zap.L().Error("查询数据库中密码失败", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }
