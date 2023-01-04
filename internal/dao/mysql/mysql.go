package mysql

import (
	"fmt"

	"github.com/aloysZy/goweb/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// 这个 db 是小写的，因为只有MySQL 文件里面才能用上这个
var db *sqlx.DB

// Mysql Init 初始化MySQL连接
func Mysql(cfg *config.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	zap.L().Info("init mysql success")
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
