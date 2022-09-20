package router

import "akcasbin/router/system"

type routerGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(routerGroup)
