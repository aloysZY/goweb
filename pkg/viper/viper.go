package viper

import (
	"flag"
	"fmt"

	"github.com/aloysZy/goweb/global/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 配置文件读取初始化
func Viper() error {
	var myConfig string
	flag.StringVar(&myConfig, "c", "./config.yaml", "choose config file.")
	flag.Parse()
	viper.SetConfigFile(myConfig)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		viper.Unmarshal(&conf.Config)

	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&conf.Config); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	fmt.Printf("init viper success\n")
	return err
}
