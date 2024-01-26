package services

import (
	context "context"
	"log"
	"my-package/models"
	"my-package/repository"
	"my-package/utils"
	"net/http"
)

type userGrpcServer struct {
	userRepo repository.UserRepo
}

func NewUserGrpcServer(userRepo repository.UserRepo) UserGrpcServer {
	return userGrpcServer{userRepo}
}

func (obj userGrpcServer) mustEmbedUnimplementedUserGrpcServer() {}

func (obj userGrpcServer) Register(ctx context.Context, req *RegisterRequest) (res *RegisterResponse, err error) {
	if req == nil {
		res = &RegisterResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie username and password",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.Username == "" || req.Password == "" {
		res = &RegisterResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie username and password",
		}
		log.Println(res.Message)
		return res, nil
	}

	users, err := obj.userRepo.GetQuery(models.User{
		Username: req.Username,
	})
	if err != nil {
		res = &RegisterResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(users) > 0 {
		res = &RegisterResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: "user already exites",
		}
		log.Println(res.Message)
		return res, nil
	}

	err = obj.userRepo.Create(models.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password),
	})
	if err != nil {
		res = &RegisterResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	res = &RegisterResponse{
		Status:  http.StatusCreated,
		Message: "Register success",
	}
	log.Println(res.Message)
	return res, nil
}

func (obj userGrpcServer) Login(ctx context.Context, req *LoginRequest) (res *LoginResponse, err error) {
	if req == nil {
		res = &LoginResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie username and password",
		}
		log.Println(res.Message)
		return res, nil
	}
	if req.Username == "" || req.Password == "" {
		res = &LoginResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie username and password",
		}
		log.Println(res.Message)
		return res, nil
	}

	users, err := obj.userRepo.GetQuery(models.User{
		Username: req.Username,
	})
	if err != nil {
		res = &LoginResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}
	if len(users) == 0 {
		res = &LoginResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "not found username",
		}
		log.Println(res.Message)
		return res, nil
	}

	if utils.HashPassword(req.Password) != users[0].Password {
		res = &LoginResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: "invalid password",
		}
		log.Println(res.Message)
		return res, nil
	}

	res = &LoginResponse{
		Status:   http.StatusOK,
		Message:  "Login success",
		Id:       int64(users[0].ID),
		Username: users[0].Username,
	}
	log.Println(res.Message)
	return res, nil
}

func (obj userGrpcServer) GetAllUser(ctx context.Context, req *GetAllUserRequest) (res *GetAllUserResponse, err error) {

	users, err := obj.userRepo.GetQuery(models.User{})
	if err != nil {
		res = &GetAllUserResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(users) == 0 {
		res = &GetAllUserResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "users not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res = &GetAllUserResponse{
		Status:  http.StatusOK,
		Message: "GetAllUser success",
		Users:   nil,
	}

	res.Users = make([]*User, 0, len(users))
	for i := range users {
		res.Users = append(res.Users, &User{
			Id:       int64(users[i].ID),
			Username: users[i].Username,
			Password: users[i].Password,
		})
	}

	log.Println(res.Message)
	return res, nil
}

func (obj userGrpcServer) GetByID(ctx context.Context, req *GetByIDRequest) (res *GetByIDResponse, err error) {
	if req == nil {
		res = &GetByIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie id user",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.Id == 0 {
		res = &GetByIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie id user",
		}
		log.Println(res.Message)
		return res, nil
	}

	users, err := obj.userRepo.GetQuery(models.User{
		ID: uint(req.Id),
	})
	if err != nil {
		res = &GetByIDResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(users) == 0 {
		res = &GetByIDResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res = &GetByIDResponse{
		Status:  http.StatusOK,
		Message: "GetByID success",
		User: &User{
			Id:       int64(users[0].ID),
			Username: users[0].Username,
			Password: users[0].Password,
		},
	}
	log.Println(res.Message)
	return res, nil
}

func (obj userGrpcServer) UpdatePassword(ctx context.Context, req *UpdatePasswordRequest) (res *UpdatePasswordResponse, err error) {
	// Your implementation here
	return res, nil
}

func (userGrpcServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (res *DeleteUserResponses, err error) {
	// Your implementation here
	return res, nil
}
