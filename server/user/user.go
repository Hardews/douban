package main

import (
	"context"
	"douban/middleware"
	"douban/model"
	"douban/proto"
	"douban/service"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (ser *server) Login(ctx context.Context, in *proto.LoginReq) (*proto.SuccessfulResp, error) {
	var identity = "用户"
	var resp = &proto.SuccessfulResp{}
	resp.OK = false

	var iUser = model.User{
		0,
		(*in).Username,
		(*in).Password,
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
			resp.Msg = "无此账号"
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
		resp.Msg = token
		err := status.New(codes.OK, "成功")
		return resp, err.Err()
	} else {
		resp.OK = true
		resp.Msg = "密码错误"
		return resp, nil
	}
}

func (ser *server) Register(ctx context.Context, user *proto.RegisterReq) (*proto.SuccessfulResp, error) {
	var res = service.CheckSensitiveWords(user.Username)
	var resp = &proto.SuccessfulResp{OK: true}
	var err error
	if !res {
		resp.Msg = "用户名含有敏感词汇"
		return resp, nil
	}
	res = service.CheckSensitiveWords(user.Nickname)
	if !res {
		resp.Msg = "昵称含有敏感词汇"
		return resp, nil
	}

	var iUser = model.User{
		0,
		user.Username,
		user.Password,
		user.Nickname,
	}
	err, flag := service.CheckUsername(iUser)
	if err != nil {
		resp.OK = false
		return resp, status.Errorf(codes.Internal, "check username failed, error: "+err.Error())
	}
	if flag == false {
		resp.Msg = " 用户名已存在!"
		return resp, nil
	}

	res = service.CheckLength(user.Password)
	if !res {
		resp.Msg = "密码长度不合法"
		return resp, nil
	}

	err, iUser.Password = service.Encryption(user.Password)
	if err != nil {
		err = errors.New("internet error")
		return resp, status.Errorf(codes.Internal, "encryption failed , err :"+err.Error())
	}

	err = service.WriteIn(iUser)
	if err != nil {
		resp.OK = false
		resp.Msg = "internet error"
		return resp, status.Errorf(codes.Internal, "register failed,err:"+err.Error())
	}

	resp.Msg = "成功"
	return resp, err
}
