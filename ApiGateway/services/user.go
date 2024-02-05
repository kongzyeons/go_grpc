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
	GetAllUser() (res models.Response)
	GetByID(id int) (res models.Response)
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

	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
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

	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
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

func (obj userSrv) GetAllUser() (res models.Response) {
	result, err := obj.userGrpc.GetAllUser(context.Background(), &grpcClient.GetAllUserRequest{})

	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
		return res
	}

	var users []models.GetAllUserRes
	for i := range result.Users {
		users = append(users, models.GetAllUserRes{
			ID:       int(result.Users[i].Id),
			Username: result.Users[i].Username,
		})
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
		Data:    users,
	}
	return res
}

func (obj userSrv) GetByID(id int) (res models.Response) {
	result, err := obj.userGrpc.GetByID(context.Background(), &grpcClient.GetByIDRequest{
		Id: int64(id),
	})

	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
		Data: models.GetByIDRes{
			ID:       int(result.User.Id),
			Username: result.User.Username,
		},
	}
	return res
}
