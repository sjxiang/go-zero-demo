package main

import (
	"flag"
	"fmt"

	"github.com/sjxiang/go-zero-demo/user-api/internal/config"
	"github.com/sjxiang/go-zero-demo/user-api/internal/handler"
	"github.com/sjxiang/go-zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 参数配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 依赖
	ctx := svc.NewServiceContext(c)
	
	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
