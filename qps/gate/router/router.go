package router

import (
	"gate/api"
	"gate/auth"
	"github.com/gin-gonic/gin"
	"msqp/config"
	"msqp/rpc"
)

// RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	if config.Conf.Log.Level == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化grpc的client gate是作为grpc的客户端，去调用user grpc服务

	rpc.Init()
	r := gin.Default()
	r.Use(auth.Cors())
	userHandler := api.NewUserHandler()
	r.POST("/register", userHandler.Register)
	return r
}
