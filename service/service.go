package service

import (
	"github.com/riny/go-rpc/logger"
)

type UserInfoService struct{}

type Response struct {
	msg  string
	user *UserInfo
}

var l = &logger.Logger{}

func (s *UserInfoService) Hello(request *UserInfo, reply *string) error {
	*reply = "Hello " + request.GetName()
	l.Info(request)

	return nil
}

func (s *UserInfoService) Login(request *UserInfo, reply *Response) error {
	l.Info(request)

	if request.Name != "test" && request.Age != 24 {
		reply.msg = "Login failed"
		return nil
	}
	reply.user = request.User

	return nil
}
