package main

import (
	"context"
	"flag"
	"fmt"
	"gate/app"
	"log"
	"msqp/config"
	"msqp/metrics"
	"os"
)

var configFile = flag.String("config", "application.yml", "config file")

func main() {
	//1.加载配置
	flag.Parse()
	config.InitConfig(*configFile)
	//2.启动监控
	go func() {
		err := metrics.Serve(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort))
		if err != nil {
			panic(err)
		}
	}()
	err := app.Run(context.Background())
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
