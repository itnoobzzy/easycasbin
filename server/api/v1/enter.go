package v1

import "akcasbin/server/api/v1/system"

type apiGroup struct {
	SystemApiGroup system.SysApiGroup
}

var ApiGroupApp = new(apiGroup)
