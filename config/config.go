package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"problems-everyday-bot/utils"
)

var Conf = new(config)

type config struct {
	Lark     *Lark     `mapstructure:"lark" json:"lark" yaml:"lark"`
	LarkBots *LarkBots `mapstructure:"larkBots" json:"larkBots" yaml:"larkBots"`
}

// InitConfig 初始化配置
func InitConfig(configPath string) {
	// 检查入参是否为空
	if configPath == "" {
		fmt.Println("请指定配置文件")
		os.Exit(1)
	}
	// 检查配置文件是否存在
	if ok := utils.FileExists(configPath); !ok {
		fmt.Println("配置文件不存在")
		os.Exit(1)
	}
	// 初始化配置
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 监控配置文件变化并热加载程序
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改，重新加载")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
	})
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}
}

// Lark 飞书AK/SK配置
type Lark struct {
	AccessKey string `mapstructure:"accessKey" json:"accessKey"`
	SecretKey string `mapstructure:"secretKey" json:"secretKey"`
}

// LarkBots 飞书机器人配置
type LarkBots []string
