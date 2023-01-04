package main

import (
	"fmt"

	"github.com/aloysZy/goweb/global/conf"
	"github.com/aloysZy/goweb/internal/dao/mysql"
	"github.com/aloysZy/goweb/internal/dao/redis"
	"github.com/aloysZy/goweb/internal/logger"
	"github.com/aloysZy/goweb/internal/router"
	"github.com/aloysZy/goweb/pkg/viper"
)

func main() {
	if err := viper.Viper(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Zap(conf.Config.LogConfig, conf.Config.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Mysql(conf.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Redis(conf.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 初始化并启动 gin
	router.SetupRouter()

}
