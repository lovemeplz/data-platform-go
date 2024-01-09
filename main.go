package main

import (
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	"github.com/lovemeplz/data-platform-go/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           "127.0.0.1:9000", // TODO 改为从配置文件读取
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
