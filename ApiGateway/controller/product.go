package controller

import (
	"log"
	"my-package/models"
	"my-package/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductRest interface {
	CreateProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
	GetProductID(c *gin.Context)
}

type productRest struct {
	productSrv services.ProductSrv
}

func NewProductRest(productSrv services.ProductSrv) ProductRest {
	return productRest{productSrv}
}

// CreateProduct godoc
// @summary CreateProduct
// @description CreateProduct
// @tags Product
// @id CreateProduct
// @security ApiKeyAuth
// @accept json
// @produce json
// @param User body models.CreateProductReq true "Product data to be created"
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/product/create [post]
func (obj productRest) CreateProduct(c *gin.Context) {
	var req models.CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error CreateProduct :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: err.Error(),
		})
		return
	}
	res := obj.productSrv.CreateProduct(req)
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// GetAllProduct godoc
// @summary GetAllProduct
// @description GetAllProduct
// @tags Product
// @id GetAllProduct
// @security ApiKeyAuth
// @accept json
// @produce json
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/products [get]
func (obj productRest) GetAllProduct(c *gin.Context) {
	res := obj.productSrv.GetAllProduct()
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}

// GetProductID godoc
// @summary GetProductID
// @description GetProductID
// @tags Product
// @id GetProductID
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "id of product to be get"
// @response 200 {object} models.Response "OK"
// @response 201 {object} models.Response "Create Ok"
// @response 400 {object} models.Response "Bad Request"
// @response 401 {object} models.Response "Unauthorized"
// @response 500 {object} models.Response "Internal Server Error"
// @Router /api/v1/product/{id} [get]
func (obj productRest) GetProductID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error GetProductID :", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   true,
			Status:  http.StatusBadRequest,
			Massage: "error invalid ID",
		})
		return
	}
	res := obj.productSrv.GetProductID(uint64(id))
	log.Println(res.Massage)
	c.JSON(int(res.Status), res)
}
