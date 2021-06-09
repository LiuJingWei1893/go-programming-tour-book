package main

import (
	"net/http"
	"time"

	"github.com/go-programing-tour-book/blog-service/global"
	"github.com/go-programing-tour-book/blog-service/internal/routers"
	"github.com/go-programing-tour-book/blog-service/pkg/setting"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
}
