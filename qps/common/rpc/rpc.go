package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"msqp/config"
	"msqp/discovery"
	"msqp/logs"
	"user/pb"
)

var (
	UserClient pb.UserServiceClient
)

func Init() {
	r := discovery.NewResolver(config.Conf.Etcd)
	resolver.Register(r)
	userDomain := config.Conf.Domain["user"]
	initClient(userDomain.Name, userDomain.LoadBalance, &UserClient)
}

func initClient(name string, loadBalance bool, client interface{}) {
	// 找服务的地址
	addr := fmt.Sprintf("etcd:///%s", name)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials())}
	if loadBalance {
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{	"LoadBalancingPolicy":"%s"}`, "round_robin")))
	}
	conn, err := grpc.DialContext(context.TODO(), addr, opts...)
	if err != nil {
		logs.Fatal("rpc connect etcd err:%v", err)
	}
	switch c := client.(type) {
	case *pb.UserServiceClient:
		*c = pb.NewUserServiceClient(conn)
	default:
		logs.Fatal("unsupported client type")
	}
}
