package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl   string
	ImageSavePath    string
	ImagePreviewPath string
	ImageMaxSize     int
	ImageAllowExts   []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	ExportSavePath string
}

type Server struct {
	RunMode      string
	HttpHost     string
	HttpPort     string
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
	Cfg.SetConfigType("yaml")
	Cfg.SetConfigName("app") // 名称

	err := Cfg.ReadInConfig() // 读配置
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	err = Cfg.Unmarshal(AppSetting)
	if err != nil {
		log.Fatal("AppSetting转换Json失败")
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Unmarshal(ServerSetting)
	if err != nil {
		log.Fatal("ServerSetting转换Json失败")
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Unmarshal(DatabaseSetting)
	if err != nil {
		log.Fatal("DatabaseSetting转换Json失败")
	}
}
