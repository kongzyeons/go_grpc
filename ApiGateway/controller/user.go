package controller

import (
	"log"
	"my-package/models"
	"my-package/services"
	"net/http"

	_ "my-package/docs"

	"github.com/gin-gonic/gin"
)

type UserRest interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userRest struct {
	userSrv services.UserSrv
}

func NewUserRest(userSrv services.UserSrv) UserRest {
	return userRest{userSrv}
}

// Register godoc
// @summary Register
// @description Register
// @tags User
// @id Register
// @security ApiKeyAuth
// @accept json
// @produce json
// @param User body models.RegisterReq true "User data to be created"
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/user/register [post]
func (obj userRest) Register(c *gin.Context) {
	var req models.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error Register:", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: err.Error(),
		})
		return
	}
	res := obj.userSrv.Register(req)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// Login godoc
// @summary Login
// @description Login
// @tags User
// @id Login
// @security ApiKeyAuth
// @accept json
// @produce json
// @param User body models.LoginReq true "User data to be created"
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/user/login [post]
func (obj userRest) Login(c *gin.Context) {
	var req models.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error Login :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: err.Error(),
		})
		return
	}
	res := obj.userSrv.Login(req)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}
