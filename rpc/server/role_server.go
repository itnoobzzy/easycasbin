package main

import (
	"easycasbin/forms"
	pb "easycasbin/rpc/proto"
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

// GetDomainRoles 获取指定域下所有角色
func (s *RoleServer) GetDomainRoles(ctx context.Context, request *pb.GetAllRolesReq) (*pb.GetAllRolesRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}
	form := forms.GetAllRoles{
		Domain: request.DomainName,
	}
	roles := roleService.GetDomainRoles(&form)

	resData := make([]*pb.GetAllRolesRpl_Data, 0)
	for _, role := range roles {
		resData = append(resData, &pb.GetAllRolesRpl_Data{
			Id:         int64(role.ID),
			CreateTime: timestamppb.New(role.CreatedAt),
			UpdateTime: timestamppb.New(role.UpdatedAt),
			Name:       role.Name,
			Domain:     role.Domain,
		})
	}
	return &pb.GetAllRolesRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
		Data: resData,
	}, nil
}

// GetDomainSubsForRole 获取指定域角色下所有鉴权主体（用户或者角色）
func (s *RoleServer) GetDomainSubsForRole(ctx context.Context, request *pb.GetDomainSubsForRoleReq) (*pb.GetDomainSubsForRoleRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}
	form := forms.DomainRole{
		DomainName: request.DomainName,
		RoleName:   request.RoleName,
	}

	subs := roleService.GetDomainSubsForRole(&form)
	resData := make([]*pb.GetDomainSubsForRoleRpl_Data, 0)
	for _, sub := range subs {
		resData = append(resData, &pb.GetDomainSubsForRoleRpl_Data{Sub: sub})
	}

	return &pb.GetDomainSubsForRoleRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
		Data: resData,
	}, nil
}

// AddRoleForSubInDomain 为用户添加域角色或者为角色继承另一个角色权限
func (s *RoleServer) AddRoleForSubInDomain(ctx context.Context, request *pb.AddRoleForSubInDomainReq) (*pb.AddRoleForSubInDomainRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}
	form := forms.RoleForSubInDomain{
		Sub:    request.Sub,
		Domain: request.Domain,
		Role:   request.Role,
	}
	err := roleService.AddRoleForSubInDomain(&form)
	if err != nil {
		st := status.New(codes.NotFound, err.Error())
		return nil, st.Err()
	}
	return &pb.AddRoleForSubInDomainRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
	}, nil
}

// DeleteRoleForSubInDomain  删除角色下subject（鉴权主体）
func (s *RoleServer) DeleteRoleForSubInDomain(ctx context.Context, request *pb.DeleteRoleForSubInDomainReq) (*pb.DeleteRoleForSubInDomainRpl, error) {
	if err := request.Validate(); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}
	form := forms.RoleForSubInDomain{
		Sub:    request.Sub,
		Domain: request.Domain,
		Role:   request.Role,
	}
	err := roleService.DeleteRoleForSubInDomain(&form)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return &pb.DeleteRoleForSubInDomainRpl{
		Code: uint32(codes.OK),
		Msg:  "ok",
	}, nil
}
