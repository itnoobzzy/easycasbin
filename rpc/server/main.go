package main

import (
	"easycasbin/core"
	"easycasbin/global"
	"easycasbin/initialize"
	pb "easycasbin/rpc/proto"
	"easycasbin/service"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	roleService = service.ServiceGroupApp.SystemServiceGroup.RoleService
)

func main() {
	global.CASBIN_VP = core.Viper("../server/config.yaml")
	global.CASBIN_LOG = core.Zap()
	zap.ReplaceGlobals(global.CASBIN_LOG)
	global.CASBIN_DB = initialize.Gorm()
	global.CASBIN_ENFORCER = core.Enforcer("../server/casbin_rbac_domain.conf")
	// 初始化表单校验翻译器
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	if global.CASBIN_DB != nil {
		initialize.RegisterTables(global.CASBIN_DB)
		db, _ := global.CASBIN_DB.DB()
		defer db.Close()
	}

	g := grpc.NewServer()
	pb.RegisterDomainRoleServer(g, &RoleServer{})
	addr := fmt.Sprintf("0.0.0.0:%v", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start gtpc:" + err.Error())
	}
}
