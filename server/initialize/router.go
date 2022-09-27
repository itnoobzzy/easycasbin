package initialize

import (
	"akcasbin/global"
	"akcasbin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitInitRouter(PublicGroup)
		systemRouter.InitUserRouter(PublicGroup)
		systemRouter.InitRoleRouter(PublicGroup)
		systemRouter.InitCasbinRouter(PublicGroup)
	}

	global.CASBIN_LOG.Info("router register success")
	return Router
}
