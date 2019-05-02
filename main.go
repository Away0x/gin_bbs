package main

import (
	"fmt"
	"net/http"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"

	"gin_bbs/config"

	"gin_bbs/bootstrap"

	"github.com/spf13/pflag"
)

var (
	// 需要 mock data，注意该操作会覆盖数据库；只在非 release 时生效
	needMock = pflag.BoolP("mock", "m", false, "need mock data")
)

func main() {
	// 解析命令行参数
	pflag.Parse()
	// 初始化配置
	config.InitConfig()

	// gin setup
	g := gin.New()
	bootstrap.SetupGin(g)

	// db setup
	db, err := bootstrap.SetupDB(*needMock)
	if err != nil {
		return
	}
	defer db.Close()

	// router setup
	bootstrap.SetupRouter(g)

	// 启动
	fmt.Printf("\n\n-------------------------------------------------- Start to listening the incoming requests on http address: %s --------------------------------------------------\n\n", config.AppConfig.Addr)
	if err := http.ListenAndServe(config.AppConfig.Addr, g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}
