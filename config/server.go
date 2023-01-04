package config

type ServerConfig struct {
	Name         string `mapstructure:"name"`
	Addr         string `mapstructure:"addr"`
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	MachineID    uint16 `mapstructure:"machine_id"`
	Locale       string `mapstructure:"locale"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
