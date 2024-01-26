package services

import (
	"context"
	"my-package/grpcClient"
	"my-package/models"
	"my-package/utils"
)

type UserSrv interface {
	Register(req models.RegisterReq) (res models.Response)
	Login(req models.LoginReq) (res models.Response)
}

type userSrv struct {
	userGrpc grpcClient.UserGrpcClient
}

func NewUserSrv(userGrpc grpcClient.UserGrpcClient) UserSrv {
	return userSrv{userGrpc}
}

func (obj userSrv) Register(req models.RegisterReq) (res models.Response) {
	result, err := obj.userGrpc.Register(context.Background(), &grpcClient.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	res = utils.HandlerErrGrpcCleint(result, err)
	if res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
	}
	return res
}

func (obj userSrv) Login(req models.LoginReq) (res models.Response) {
	result, err := obj.userGrpc.Login(context.Background(), &grpcClient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	res = utils.HandlerErrGrpcCleint(result, err)
	if res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
		Data: models.LoginRes{
			ID:       int(result.Id),
			Username: result.Username,
			AccToken: "accToken",
			RefToken: "refToken",
		},
	}
	return res
}
