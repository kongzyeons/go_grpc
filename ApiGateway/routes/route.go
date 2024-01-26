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
	useSrv := services.NewUserSrv(userGrpc)
	userRest := controller.NewUserRest(useSrv)

	app.POST("api/v1/user/register", userRest.Register)
	app.POST("api/v1/user/login", userRest.Login)
}
