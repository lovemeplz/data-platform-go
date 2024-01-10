package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	ExportSavePath string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var (
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	DatabaseSetting = &Database{}
	Cfg             = viper.New()
)

func Setup() {
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

	err = Cfg.Unmarshal("app")
	fmt.Println("setting:::", Cfg.AllSettings())
}
