package model

// 解析前端传入参数

// ParamSignUpUser 注册，根据数据库字段，和前端传入，以及必要字段来设定
// binding:"required"` 进行参数效验，这是因为gin 使用了validator
type ParamSignUpUser struct {
	// binding:"required 表示不能为空
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// eqfield=Password 表示这个字段的内容和另一个字段相同
	// binding:"required,eqfield=Password"` 中间要有逗号，不能有空格
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	// binding:"email"` 这里设置了必须是一个有效的邮箱，所以也能为空了
	// 可以设置为为空和有效的邮箱一起设置，那样，只有一个用户的邮箱可以为空，其他就不能了
	Email string `json:"email" binding:"email"`
}

// ParamLoginUser 登录，根据数据库字段，和前端传入，以及必要字段来设定
// 登录验证用户名密码就行了
type ParamLoginUser struct {
	// UserID uint64 `json:"user_id" db:"user_id"`
	// binding:"required 表示不能为空
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
