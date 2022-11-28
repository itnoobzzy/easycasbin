package main

import (
	"easycasbin/rpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"os"
)

func HandleGrpcErrorToHttp(err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				fmt.Printf("参数错误: %v\n", e.Message())
			case codes.Unavailable:
				fmt.Println("连接服务失败")
			case codes.Internal:
				fmt.Println("服务错误")
			case codes.NotFound:
				fmt.Println(e.Message())
			default:
				fmt.Println(e.Message())
			}
			os.Exit(1)
		}
	}
}

func main() {
	dial, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	c := proto.NewDomainRoleClient(dial)

	//reply, err := c.AddDomainRole(context.Background(),
	//	&proto.AddDomainRoleReq{RoleName: "role1", DomainName: "domain12"},
	//)

	//reply, err := c.UpdateRoleInfo(context.Background(),
	//	&proto.UpdateDomainRoleReq{RoleName: "role1", DomainName: "domain12", NewRoleName: "role2"},
	//)

	//reply, err := c.DeleteRole(context.Background(),
	//	&proto.DeleteDomainRoleReq{RoleName: "role2", DomainName: "domain12"},
	//)

	//reply, err := c.GetDomainRoles(context.Background(),
	//	&proto.GetAllRolesReq{DomainName: "domain1"},
	//)

	//reply, err := c.GetDomainSubsForRole(context.Background(),
	//	&proto.GetDomainSubsForRoleReq{DomainName: "domain1", RoleName: "admin"},
	//)

	//reply, err := c.AddRoleForSubInDomain(context.Background(),
	//	&proto.AddRoleForSubInDomainReq{
	//		Sub:    "user:zhouzy1",
	//		Domain: "domain:domain1",
	//		Role:   "role:admin3",
	//	},
	//)

	reply, err := c.DeleteRoleForSubInDomain(context.Background(),
		&proto.DeleteRoleForSubInDomainReq{
			Sub:    "user:zhouzy1",
			Domain: "domain:domain1",
			Role:   "role:admin3",
		},
	)

	HandleGrpcErrorToHttp(err)
	fmt.Println(reply.Code)
}
