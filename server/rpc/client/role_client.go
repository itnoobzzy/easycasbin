package main

import (
	"akcasbin/rpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	c := proto.NewDomainRoleClient(dial)
	reply, err := c.AddDomainRole(context.Background(), &proto.DomainRole{RoleName: "role1", DomainName: "domain1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.Id)

}
