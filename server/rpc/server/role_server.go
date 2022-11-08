package main

import (
	"akcasbin/core"
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/initialize"
	pb "akcasbin/rpc/proto"
	"akcasbin/service"
	"context"
	"flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	roleService = service.ServiceGroupApp.SystemServiceGroup.RoleService
)

type server struct {
}

func (s *server) AddDomainRole(ctx context.Context, request *pb.DomainRole) (*pb.AddDomainRoleRpl, error) {
	//st := status.New(codes.OK, "ok")
	if err := request.Validate(); err != nil {
		//msg := fmt.Printf("参数错误：%v", err.Error())
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	form := forms.DomainRole{
		RoleName:   request.RoleName,
		DomainName: request.DomainName,
	}

	role, err := roleService.AddRole(&form)
	if err != nil {
		return nil, err
	}
	return &pb.AddDomainRoleRpl{
		Domain:    role.Domain,
		Name:      role.Name,
		Id:        int64(role.ID),
		CreatTime: timestamppb.New(role.CreatedAt),
	}, nil
}

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
	pb.RegisterDomainRoleServer(g, &server{})
	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start gtpc:" + err.Error())
	}
}
