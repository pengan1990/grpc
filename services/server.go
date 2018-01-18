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




