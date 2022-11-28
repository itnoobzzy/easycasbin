package router

import "easycasbin/router/system"

type routerGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(routerGroup)
