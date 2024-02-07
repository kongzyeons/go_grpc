package routes

import (
	"my-package/config"
	"my-package/controller"
	"my-package/grpcClient"
	"my-package/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouterApp(app *gin.Engine, cfg config.Config) {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGrpc := grpcClient.NewUserGrpcClient(cfg.UserGrpc)
	productGrpc := grpcClient.NewProductGrpcClient(cfg.ProductGrpc)
	orderGrpc := grpcClient.NewOrderGrpcClient(cfg.OrderGrpc)

	useSrv := services.NewUserSrv(userGrpc)
	productSrv := services.NewProductSrv(productGrpc)
	orderSrv := services.NewOrderSrv(orderGrpc, productGrpc)

	userRest := controller.NewUserRest(useSrv)
	app.POST("api/v1/user/register", userRest.Register)
	app.POST("api/v1/user/login", userRest.Login)
	app.GET("api/v1/users", userRest.GetAllUser)
	app.GET("api/v1/user/:id", userRest.GetByID)

	productRest := controller.NewProductRest(productSrv)
	app.POST("api/v1/product/create", productRest.CreateProduct)
	app.GET("api/v1/products", productRest.GetAllProduct)
	app.GET("api/v1/product/:id", productRest.GetProductID)

	orderRest := controller.NewOrderRest(orderSrv)
	app.POST("api/v1/order/create/:user_id", orderRest.CreateOrder)
	app.GET("api/v1/order/:user_id", orderRest.GetOrderByUser)
	app.PUT("api/v1/order/:order_id", orderRest.AddProduct)

}
