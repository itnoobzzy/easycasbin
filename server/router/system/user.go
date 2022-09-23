package system

import (
	v1 "akcasbin/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("register", userApi.Register)
	}
	return userRouter
}
