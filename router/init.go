package apirouter

import (
	"atlantis/conf"
	"atlantis/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Router  新建路由，增加中间件，分组
func Router(app conf.App) *gin.Engine {
	router := gin.New()

	if app.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.Logger(), gin.Recovery())
	api := router.Group("")
	api.Use(middleware.Cors())
	InitUserRouter(api)
	InitURouter(api)
	return router
}
