package model

// SignUpUser 解析数据库参数，往数据库写数据
type SignUpUser struct {
	UserID   uint64 `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

// LoginUser 解析数据库参数，往数据库写数据，根据用户
type LoginUser struct {
	UserName string `db:"username"`
	Password string `db:"password"`
}
