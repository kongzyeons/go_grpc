package services_test

import (
	"context"
	"errors"
	"my-package/models"
	"my-package/repository"
	"my-package/services"
	"my-package/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		nameTest    string
		req         *services.RegisterRequest
		errGetQuery error
		users       []models.User
		errCreate   error
		expect      bool
	}{
		{nameTest: "test Register success", req: &services.RegisterRequest{Username: "A", Password: "B"}, errGetQuery: nil, users: []models.User{}, errCreate: nil, expect: false},
		{nameTest: "test error repo Create", req: &services.RegisterRequest{Username: "A", Password: "B"}, errGetQuery: nil, users: []models.User{}, errCreate: errors.New(""), expect: true},
		{nameTest: "test error repo GetQuery", req: &services.RegisterRequest{Username: "A", Password: "B"}, errGetQuery: errors.New(""), users: []models.User{}, errCreate: nil, expect: true},
		{nameTest: "test error alredy exits", req: &services.RegisterRequest{Username: "A", Password: "B"}, errGetQuery: nil, users: []models.User{{}}, errCreate: nil, expect: true},
		{nameTest: "test requie request", req: &services.RegisterRequest{Username: "", Password: ""}, errGetQuery: nil, users: []models.User{}, errCreate: nil, expect: true},
		{nameTest: "test requie request is nill", req: nil, errGetQuery: nil, users: []models.User{}, errCreate: nil, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			userRepo := repository.NewuserRepoMock()
			userRepo.On("GetQuery", mock.AnythingOfType("models.User")).Return(
				tests[i].users,
				tests[i].errGetQuery,
			)
			userRepo.On("Create", mock.AnythingOfType("models.User")).Return(
				tests[i].errCreate,
			)
			userSrv := services.NewUserGrpcServer(userRepo)
			res, _ := userSrv.Register(context.Background(), tests[i].req)
			assert.Equal(t, tests[i].expect, res.Error)
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		nameTest    string
		req         *services.LoginRequest
		errGetQuery error
		users       []models.User
		expect      bool
	}{
		{nameTest: "test Login success", req: &services.LoginRequest{Username: "A", Password: "A"}, errGetQuery: nil, users: []models.User{{Password: utils.HashPassword("A")}}, expect: false},
		{nameTest: "test error repo GetQuery", req: &services.LoginRequest{Username: "A", Password: "A"}, errGetQuery: errors.New(""), users: []models.User{{Password: utils.HashPassword("A")}}, expect: true},
		{nameTest: "test not found username", req: &services.LoginRequest{Username: "A", Password: "A"}, errGetQuery: nil, users: []models.User{}, expect: true},
		{nameTest: "test invalid password", req: &services.LoginRequest{Username: "A", Password: "B"}, errGetQuery: nil, users: []models.User{{Password: utils.HashPassword("A")}}, expect: true},
		{nameTest: "test requie request", req: &services.LoginRequest{}, errGetQuery: nil, users: []models.User{}, expect: true},
		{nameTest: "test requie request is nill", req: nil, errGetQuery: nil, users: []models.User{}, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			userRepo := repository.NewuserRepoMock()
			userRepo.On("GetQuery", mock.AnythingOfType("models.User")).Return(
				tests[i].users,
				tests[i].errGetQuery,
			)
			userSrv := services.NewUserGrpcServer(userRepo)
			res, _ := userSrv.Login(context.Background(), tests[i].req)
			assert.Equal(t, tests[i].expect, res.Error)
		})
	}
}

func TestGetAllUser(t *testing.T) {
	tests := []struct {
		nameTest    string
		errGetQuery error
		users       []models.User
		expectUsers []*services.User
		expect      bool
	}{
		{nameTest: "test GetAllUser success", errGetQuery: nil, users: []models.User{{}}, expect: false},
		{nameTest: "test error repo GetQuery", errGetQuery: errors.New(""), users: []models.User{{}}, expect: true},
		{nameTest: "test users not found", errGetQuery: nil, users: []models.User{}, expect: true},
		{nameTest: "test check response users", errGetQuery: nil, users: []models.User{{Username: "A"}}, expectUsers: []*services.User{{Username: "A"}}, expect: false},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			userRepo := repository.NewuserRepoMock()
			userRepo.On("GetQuery", mock.AnythingOfType("models.User")).Return(
				tests[i].users,
				tests[i].errGetQuery,
			)

			userSrv := services.NewUserGrpcServer(userRepo)

			res, _ := userSrv.GetAllUser(context.Background(), &services.GetAllUserRequest{})

			if tests[i].nameTest == "test check response users" {
				assert.Equal(t, tests[i].expectUsers, res.Users)
				return
			}

			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}

func TestGetByID(t *testing.T) {
	tests := []struct {
		nameTest    string
		req         *services.GetByIDRequest
		errGetQuery error
		users       []models.User
		expectUser  *services.User
		expect      bool
	}{
		{nameTest: "test GetByID success", req: &services.GetByIDRequest{Id: 1}, errGetQuery: nil, users: []models.User{{}}, expect: false},
		{nameTest: "test error repo GetQuery", req: &services.GetByIDRequest{Id: 1}, errGetQuery: errors.New(""), users: []models.User{{}}, expect: true},
		{nameTest: "test user not found", req: &services.GetByIDRequest{Id: 1}, errGetQuery: nil, users: []models.User{}, expect: true},
		{nameTest: "test check response user", req: &services.GetByIDRequest{Id: 1}, errGetQuery: nil, users: []models.User{{ID: 1}}, expectUser: &services.User{Id: 1}, expect: false},
		{nameTest: "test error requie request", req: &services.GetByIDRequest{}, errGetQuery: nil, users: []models.User{{}}, expect: true},
		{nameTest: "test error requie request is mill", req: nil, errGetQuery: nil, users: []models.User{{}}, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			userRepo := repository.NewuserRepoMock()
			userRepo.On("GetQuery", mock.AnythingOfType("models.User")).Return(
				tests[i].users,
				tests[i].errGetQuery,
			)

			userSrv := services.NewUserGrpcServer(userRepo)

			res, _ := userSrv.GetByID(context.Background(), tests[i].req)

			if tests[i].nameTest == "test check response user" {
				assert.Equal(t, tests[i].expectUser, res.User)
				return
			}

			assert.Equal(t, tests[i].expect, res.Error)

		})
	}

}
