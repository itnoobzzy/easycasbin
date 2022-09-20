package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DBApi struct{}

func (i *DBApi) InitDB(c *gin.Context) {
	if global.CASBIN_DB != nil {
		global.CASBIN_LOG.Error("The db already exists!")
		response.FailWithMessage("The db already exists!", c)
		return
	}
	var dbInfo forms.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.CASBIN_LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过！", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.CASBIN_LOG.Error("自动创建数据库失败！", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}

func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)
	if global.CASBIN_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.CASBIN_LOG.Info(message)
	response.OkWithDetailed(gin.H{
		"needInit": needInit,
	}, message, c)
}
