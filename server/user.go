package main

import (
	"context"
	"douban/middleware"
	"douban/model"
	"douban/proto"
	"douban/service"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"net"
)

func InitServer() {
	fmt.Println("服务端已启动")
	// 监听端口
	lis, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDoubanServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println(1)
}

type server struct {
	proto.UnimplementedDoubanServer
}

func (ser *server) Login(ctx context.Context, in *proto.User) (*proto.SuccessfulResp, error) {
	var identity = "用户"
	var resp = &proto.SuccessfulResp{}
	resp.OK = false

	var iUser = model.User{
		0,
		(*in).UserName,
		(*in).PassWord,
		"",
	}

	// 检查是否为管理员账号
	flag := service.CheckAdministratorUsername(iUser.Username)
	if flag {
		identity = "管理员"
	}

	// 检查账号是否正确
	err, res := service.CheckPassword(iUser.Username, iUser.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.OK = true
			resp.Token = "无此账号"
			return resp, nil
		}
		err := status.New(codes.Internal, err.Error())
		return resp, err.Err()
	}
	if res {
		token, flag := middleware.SetToken(iUser.Username, identity)
		if !flag {
			err := status.New(codes.OK, err.Error())
			return resp, err.Err()
		}
		resp.OK = true
		resp.Token = token
		err := status.New(codes.OK, "成功")
		return resp, err.Err()
	} else {
		resp.OK = true
		resp.Token = "密码错误"
		return resp, nil
	}
}

func (ser *server) Register(ctx context.Context, user *proto.User) (*proto.SuccessfulResp, error) {
	var res = service.CheckSensitiveWords(user.UserName)
	var resp = &proto.SuccessfulResp{OK: true}
	var err error
	if !res {
		resp.Token = "用户名含有敏感词汇"
		return resp, nil
	}
	res = service.CheckSensitiveWords(user.Nickname)
	if !res {
		resp.Token = "昵称含有敏感词汇"
		return resp, nil
	}

	var iUser = model.User{
		0,
		user.UserName,
		user.PassWord,
		user.Nickname,
	}
	err, flag := service.CheckUsername(iUser)
	if err != nil {
		resp.OK = false
		return resp, status.Errorf(codes.Internal, "check username failed, error: "+err.Error())
	}
	if flag == false {
		resp.Token = " 用户名已存在!"
		return resp, nil
	}

	res = service.CheckLength(user.PassWord)
	if !res {
		resp.Token = "密码长度不合法"
		return resp, nil
	}

	err, iUser.Password = service.Encryption(user.PassWord)
	if err != nil {
		err = errors.New("internet error")
		return resp, status.Errorf(codes.Internal, "encryption failed , err :"+err.Error())
	}

	err = service.WriteIn(iUser)
	if err != nil {
		resp.OK = false
		resp.Token = "internet error"
		return resp, status.Errorf(codes.Internal, "register failed,err:"+err.Error())
	}

	resp.Token = "成功"
	return resp, err
}
