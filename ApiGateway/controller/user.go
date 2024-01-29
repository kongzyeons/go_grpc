package controller

import (
	"log"
	"my-package/models"
	"my-package/services"
	"net/http"
	"strconv"

	_ "my-package/docs"

	"github.com/gin-gonic/gin"
)

type UserRest interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetByID(c *gin.Context)
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

// GetAllUser godoc
// @summary GetAllUser
// @description GetAllUser
// @tags User
// @id GetAllUser
// @security ApiKeyAuth
// @accept json
// @produce json
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/users [get]
func (obj userRest) GetAllUser(c *gin.Context) {
	res := obj.userSrv.GetAllUser()
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// GetByID godoc
// @summary GetByID
// @description GetByID
// @tags User
// @id GetByID
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "id of user to be get"
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/user/{id} [get]
func (obj userRest) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error GetByID :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: "error invalid ID",
		})
		return
	}
	res := obj.userSrv.GetByID(id)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}
