package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var (
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

var Cfg = viper.New()

func init() {
	Cfg.AddConfigPath("./conf") // 路径(当前路径下的conf文件夹)
	Cfg.SetConfigName("app")    // 名称

	err := Cfg.ReadInConfig() // 读配置
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	LoadBase()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	viper.SetDefault("HTTPPort", 9000)
}

func LoadServer() {
	HTTPPort = Cfg.GetInt("server.HTTP_PORT")
	ReadTimeout = time.Duration(Cfg.GetInt("server.READ_TIMEOUT")) * time.Second
	WriteTimeout = time.Duration(Cfg.GetInt("server.WRITE_TIMEOUT")) * time.Second
}

func LoadApp() {

}
