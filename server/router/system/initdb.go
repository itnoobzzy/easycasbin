package system

import (
	v1 "akcasbin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)
		initRouter.POST("initdb", dbApi.CheckDB)
	}
}
