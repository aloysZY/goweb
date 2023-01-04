package settings

import (
	"flag"
	"fmt"

	"github.com/aloysZy/goweb/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 初始化配置文件
var Conf = new(config.ServerConfig)

// Viper 配置文件读取初始化
func Viper() error {
	var myConfig string
	flag.StringVar(&myConfig, "c", "./config.yaml", "choose config file.")
	flag.Parse()
	viper.SetConfigFile(myConfig)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		viper.Unmarshal(&Conf)

	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	fmt.Printf("init viper success\n")
	return err
}
