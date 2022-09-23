package initialize

import (
	"akcasbin/global"
	"akcasbin/utils"
	"go.uber.org/zap"
)

func init() {
	err := utils.RegisterRule("PageVerify", utils.Rules{
		"Page":     {utils.NotEmpty()},
		"PageSize": {utils.NotEmpty()},
	})

	err = utils.RegisterRule("IdVerify",
		utils.Rules{
			"Id": {utils.NotEmpty()},
		},
	)
	if err != nil {
		global.CASBIN_LOG.Error("初始化自定义规则失败：", zap.Error(err))
	}
}
