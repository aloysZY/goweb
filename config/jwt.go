package config

import "time"

type Jwt struct {
	SigningKey  string        `mapstructure:"signing_key"`  // jwt签名
	ExpiresTime time.Duration `mapstructure:"expires_time"` // 过期时间后面设置为分钟
	BufferTime  int64         `mapstructure:"buffer_time"`  // 缓冲时间
	Issuer      string        `mapstructure:"issuer"`       // 签发者
}
