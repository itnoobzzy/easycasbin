package router

import "akcasbin/server/router/system"

type routerGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(routerGroup)
