package controller

import (
	"fmt"
	"log"
	"my-package/models"
	"my-package/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderRest interface {
	CreateOrder(c *gin.Context)
	GetOrderByUser(c *gin.Context)
	AddProduct(c *gin.Context)
}

type orderRest struct {
	orderSrv services.OrderSrv
}

func NewOrderRest(orderSrv services.OrderSrv) OrderRest {
	return orderRest{orderSrv}
}

// CreateOrder godoc
// @summary CreateOrder
// @description CreateOrder
// @tags Order
// @id CreateOrder
// @security ApiKeyAuth
// @accept json
// @produce json
// @param user_id path string true "user_id of user to be get"
// @param Order body models.CreateOrderReq true "Order data to be created"
// @response 200 {object} models.Response "OK"
// @Router /api/v1/order/create/{user_id} [post]
func (obj orderRest) CreateOrder(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println("error GetByID :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: "error invalid ID",
		})
		return
	}
	var req models.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error CreateOrder", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: err.Error(),
		})
		return
	}
	res := obj.orderSrv.CreateOrder(user_id, req)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// GetOrderByUser godoc
// @summary GetOrderByUser
// @description GetOrderByUser
// @tags Order
// @id GetOrderByUser
// @security ApiKeyAuth
// @accept json
// @produce json
// @param user_id path string true "user_id of user to be get"
// @response 200 {object} models.Response "OK"
// @Router /api/v1/order/{user_id} [get]
func (obj orderRest) GetOrderByUser(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println("error GetByID :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: "error invalid ID",
		})
		return
	}
	res := obj.orderSrv.GetOrderByUser(user_id)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// AddProduct godoc
// @summary AddProduct
// @description AddProduct
// @tags Order
// @id AddProduct
// @security ApiKeyAuth
// @accept json
// @produce json
// @param order_id path string true "order_id"
// @param Order body models.AddProductReq true "AddProductReq"
// @response 200 {object} models.Response "OK"
// @Router /api/v1/order/{order_id} [put]
func (obj orderRest) AddProduct(c *gin.Context) {
	order_id := c.Param("order_id")
	var req models.AddProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error AddProduct :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: "error invalid ID",
		})
		return
	}
	fmt.Println(order_id, "-----")

	res := obj.orderSrv.AddProduct(order_id, req)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}
