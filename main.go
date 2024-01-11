package main

import (
	"fmt"
	"github.com/lovemeplz/data-platform-go/models"
	"github.com/lovemeplz/data-platform-go/pkg/logging"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	"github.com/lovemeplz/data-platform-go/routers"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
