package main

import (
	"flag"
	"fmt"
	"net"
	"github.com/binlake/grpc_test/services" // 实现了服务接口的包service
	"github.com/binlake/grpc_test/protos"   // 此为自定义的protos包，存放的是.proto文件和对应的.pb.go文件

	"google.golang.org/grpc"
)

var (
	// 命令行参数-host，默认服务监听端口在9000
	addr = flag.String("host", "127.0.0.1:9000", "")
)

func main() {
	// 开启服务监听
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Println("listen error!")
		return
	}
	// 创建一个grpc服务
	grpcServer := grpc.NewServer()
	// 重点：向grpc服务中注册一个api服务，这里是UserService，处理相关请求
	protos.RegisterIUserServiceServer(grpcServer, services.NewUserService())

	// 可以添加多个api
	// TODO...

	// 启动grpc服务
	grpcServer.Serve(lis)
}
