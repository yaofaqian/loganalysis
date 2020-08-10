package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type ConfigData struct {
	DirPath         string `json:"DirPath"`
	HowManyInTheTop int    `json:"HowManyInTheTop"`
	DBName          string `json:"DBName"`
	DBUser          string `json:"DBUser"`
	DBPwd           string `json:"DBPwd"`
	DBHost          string `json:"DBHost"`
	DBPort          string `json:"DBPort"`
	DBCharset       string `json:"DBCharset"`
	LogPath         string `json:"LogPath"`
	LogPkgTime      string `json:"LogPkgTime"`
}

//
var Config ConfigData

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	viper.WatchConfig()
	viper.Unmarshal(&Config)
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&Config)
	})
}

////读取配置
//func (this *Config) InitConfig() error {
//	if this.Name != "" {
//		viper.SetConfigFile(this.Name)
//	} else {
//		viper.AddConfigPath("./config")
//		viper.SetConfigName("config")
//	}
//	viper.SetConfigType("json")
//	//从环境变量总读取
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix("web")
//	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))
//	return viper.ReadInConfig()
//}
//// 监控配置改动
//func (this *Config) WatchConfig(change chan int) {
//	viper.WatchConfig()
//	viper.OnConfigChange(func(e fsnotify.Event) {
//		log.Printf("配置已经被改变: %s", e.Name)
//		change <- 1
//	})
//}
