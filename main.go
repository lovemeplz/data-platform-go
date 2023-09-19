package main

import (
	"fmt"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	"github.com/lovemeplz/data-platform-go/routers"
	"moul.io/banner"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(banner.Inline("data-platform-go"))
	s.ListenAndServe()

}
