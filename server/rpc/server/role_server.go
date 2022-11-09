package main

import (
	"akcasbin/forms"
	pb "akcasbin/rpc/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoleServer struct {
}

// AddDomainRole 添加域角色
func (s *RoleServer) AddDomainRole(ctx context.Context, request *pb.AddDomainRoleReq) (*pb.AddDomainRoleRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	form := forms.DomainRole{
		RoleName:   request.RoleName,
		DomainName: request.DomainName,
	}

	role, err := roleService.AddRole(&form)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return &pb.AddDomainRoleRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
		Data: &pb.AddDomainRoleRpl_Data{
			Id:        int64(role.ID),
			CreatTime: timestamppb.New(role.CreatedAt),
			Name:      role.Name,
			Domain:    role.Domain,
		},
	}, nil
}

// UpdateRoleInfo 更新域角色信息
func (s *RoleServer) UpdateRoleInfo(ctx context.Context, request *pb.UpdateDomainRoleReq) (*pb.UpdateDomainRoleRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	form := forms.UpdateDomainRole{
		RoleName:    request.RoleName,
		DomainName:  request.DomainName,
		NewRoleName: request.NewRoleName,
	}

	role, err := roleService.UpdateRoleInfo(&form)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return &pb.UpdateDomainRoleRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
		Data: &pb.UpdateDomainRoleRpl_Data{
			Id:         int64(role.ID),
			UpdateTime: timestamppb.New(role.UpdatedAt),
			Name:       role.Name,
			Domain:     role.Domain,
		},
	}, nil
}

// DeleteRole 删除对应域的角色
func (s *RoleServer) DeleteRole(ctx context.Context, request *pb.DeleteDomainRoleReq) (*pb.DeleteDomainRoleRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	form := forms.DeleteDomainRole{
		RoleName:   request.RoleName,
		DomainName: request.DomainName,
	}

	err := roleService.DeleteRole(&form)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return &pb.DeleteDomainRoleRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
		Data: &pb.DeleteDomainRoleRpl_Data{
			DeleteTime: timestamppb.Now(),
		},
	}, nil
}
