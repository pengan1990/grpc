# 如题： go-grpc

## QA
接触vitess必定会接触到grpc，那么如何使用grpc，并且有哪些特性值得vitess使用，并且考虑到网络端服务的并发和吞吐，grpc又是如何做到的？  


## 代码
个人代码[地址](https://pan.baidu.com/s/1nwS94UH)  

## 个人实践
下面根据网页的参考 以及grpc客户端hellowworld的说明，写如下代码  


### 工程目录

* 当前目录
pengan@pengan:binlake$ pwd
/export/go/work/src/github.com/binlake/grpc_test

* 文件结构如下  
```txt
pengan@pengan:grpc_test$ tree -L 4
.
├── client
│   └── main.go
├── protos
│   ├── test.pb.go
│   └── test.proto
├── server
│   └── main.go
└── services
└── server.go

4 directories, 5 files
```

* 使用介绍  

* 创建proto文件    
```protobuf

syntax = "proto3";

// compile shell
// /export/tools/protobuf-3.1.0/bin/protoc -I protos/ protos/test.proto --go_out=plugins=grpc:protos

package protos;

message User {
	int64 id = 1;
	string name = 2;
}

message UserRequest {
	int64 id = 1;
}

service IUserService {
	rpc Get(UserRequest) returns (User);
}
```
注意点: 
编译命令  
```
/export/tools/protobuf-3.1.0/bin/protoc -I protos/ protos/test.proto --go_out=plugins=grpc:protos
```  
指定包名： protos;

* 编译完成
protos目录下生成 test.pb.go 文件

* 创建接口实现对象{service/server.go}    
创建UserService对象 并且对象实现相应的protobuf的接口
```
package services;


import (
		"github.com/binlake/grpc_test/protos"
		"golang.org/x/net/context"
	   )

//对外提供的工厂函数
func NewUserService() *UserService {
	return &UserService{}
}

// 接口实现对象，属性成员根据而业务自定义
type UserService struct {
}

// Get接口方法实现
func (this *UserService) Get(ctx context.Context, req *protos.UserRequest) (*protos.User, error) {
	return &protos.User{Id: 1, Name: "shuai"}, nil
}

```
注意点: 仅仅是定义UserService对象并且实现protos当中定义的接口

* 注册接口实现对象启动server端服务{server/main.go}  
```
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
```
注意点: 重点注册UserService 直接将对象注册进去

* 客户端创建链接{client/main.go}  
```
package main

import (
		"log"

		"golang.org/x/net/context"
		"google.golang.org/grpc"
		pb "github.com/binlake/grpc_test/protos"
	   )

var (
		address     = "localhost:9000"
		defaultName = "world"
	)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
	defer conn.Close()
		c := pb.NewIUserServiceClient(conn)

		// Contact the server and print out its response.
		r, err := c.Get(context.Background(), &pb.UserRequest{Id:1})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	log.Printf("Greeting: %s %s", r.Id, r.Name)
}
```
客户请求服务端，代码需要注意两个地方，dial tcp链接， 创建链接， 请求rpc 输出  


到此，则完成grpc的go版本的初级使用。
