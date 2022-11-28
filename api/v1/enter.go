package v1

import "easycasbin/api/v1/system"

type apiGroup struct {
	SystemApiGroup system.SysApiGroup
}

var ApiGroupApp = new(apiGroup)
