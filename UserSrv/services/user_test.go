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
