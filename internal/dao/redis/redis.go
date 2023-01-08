package redis

import (
	"fmt"

	"github.com/aloysZy/goweb/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	rdb *redis.Client
	Nil = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// Redis Init 初始化连接
func Redis(cfg *config.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if _, err = rdb.Ping().Result(); err != nil {
		return err
	}
	zap.L().Info("init redis success")
	return
}

func Close() {
	_ = rdb.Close()
}
