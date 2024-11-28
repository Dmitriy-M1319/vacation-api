// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vacation_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// VacationsServiceClient is the client API for VacationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacationsServiceClient interface {
	// Получить всех работников вместе с отпусками
	GetAllEmployeers(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*ManyEmployeersResponse, error)
	// Получить cотрудника по его номеру
	GetEmployeeById(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*EmployeeResponse, error)
	// Добавить нового сотрудника в базу данных
	InsertEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeResponse, error)
	// Обновить данные для сотрудника
	UpdateEmployee(ctx context.Context, in *UpdateEmpRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	// Удалить сотрудника из базы данных
	DeleteEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Получить список отпусков для определенного работника
	GetVacationsByEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*ManyVacationsResponse, error)
	// Получить отпуск по его уникальному идентификатору
	GetVacationById(ctx context.Context, in *VacationId, opts ...grpc.CallOption) (*VacationResponse, error)
	// Добавить новый отпуск для сотрудника в базу данных
	InsertVacation(ctx context.Context, in *Vacation, opts ...grpc.CallOption) (*VacationResponse, error)
	// Обновить данные для отпуска
	UpdateVacation(ctx context.Context, in *UpdateVacRequest, opts ...grpc.CallOption) (*VacationResponse, error)
	// Удалить отпуск сотрудника из базы данных
	DeleteVacation(ctx context.Context, in *VacationId, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Получить все нормы отпуска
	GetAllVacationNorms(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*ManyVacationNormsResponse, error)
	// Получить норму по его уникальному идентификатору
	GetVacationNormById(ctx context.Context, in *VacationNormId, opts ...grpc.CallOption) (*VacationNormResponse, error)
	// Добавить новый норму отпуска в базу данных
	InsertVacationNorm(ctx context.Context, in *VacationNorm, opts ...grpc.CallOption) (*VacationNormResponse, error)
	// Обновить данные для нормы отпуска
	UpdateVacationNorm(ctx context.Context, in *UpdateVacNormRequest, opts ...grpc.CallOption) (*VacationNormResponse, error)
	// Удалить норму отпуск из базы данных
	DeleteVacationNorm(ctx context.Context, in *VacationNormId, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type vacationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacationsServiceClient(cc grpc.ClientConnInterface) VacationsServiceClient {
	return &vacationsServiceClient{cc}
}

func (c *vacationsServiceClient) GetAllEmployeers(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*ManyEmployeersResponse, error) {
	out := new(ManyEmployeersResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetAllEmployeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) GetEmployeeById(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetEmployeeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) InsertEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/InsertEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) UpdateEmployee(ctx context.Context, in *UpdateEmpRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) DeleteEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) GetVacationsByEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*ManyVacationsResponse, error) {
	out := new(ManyVacationsResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetVacationsByEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) GetVacationById(ctx context.Context, in *VacationId, opts ...grpc.CallOption) (*VacationResponse, error) {
	out := new(VacationResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetVacationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) InsertVacation(ctx context.Context, in *Vacation, opts ...grpc.CallOption) (*VacationResponse, error) {
	out := new(VacationResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/InsertVacation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) UpdateVacation(ctx context.Context, in *UpdateVacRequest, opts ...grpc.CallOption) (*VacationResponse, error) {
	out := new(VacationResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/UpdateVacation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) DeleteVacation(ctx context.Context, in *VacationId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/DeleteVacation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) GetAllVacationNorms(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*ManyVacationNormsResponse, error) {
	out := new(ManyVacationNormsResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetAllVacationNorms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) GetVacationNormById(ctx context.Context, in *VacationNormId, opts ...grpc.CallOption) (*VacationNormResponse, error) {
	out := new(VacationNormResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/GetVacationNormById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) InsertVacationNorm(ctx context.Context, in *VacationNorm, opts ...grpc.CallOption) (*VacationNormResponse, error) {
	out := new(VacationNormResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/InsertVacationNorm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) UpdateVacationNorm(ctx context.Context, in *UpdateVacNormRequest, opts ...grpc.CallOption) (*VacationNormResponse, error) {
	out := new(VacationNormResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/UpdateVacationNorm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacationsServiceClient) DeleteVacationNorm(ctx context.Context, in *VacationNormId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/vacation_api.v1.VacationsService/DeleteVacationNorm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacationsServiceServer is the server API for VacationsService service.
// All implementations must embed UnimplementedVacationsServiceServer
// for forward compatibility
type VacationsServiceServer interface {
	// Получить всех работников вместе с отпусками
	GetAllEmployeers(context.Context, *EmptyResponse) (*ManyEmployeersResponse, error)
	// Получить cотрудника по его номеру
	GetEmployeeById(context.Context, *EmployeeId) (*EmployeeResponse, error)
	// Добавить нового сотрудника в базу данных
	InsertEmployee(context.Context, *Employee) (*EmployeeResponse, error)
	// Обновить данные для сотрудника
	UpdateEmployee(context.Context, *UpdateEmpRequest) (*EmployeeResponse, error)
	// Удалить сотрудника из базы данных
	DeleteEmployee(context.Context, *EmployeeId) (*EmptyResponse, error)
	// Получить список отпусков для определенного работника
	GetVacationsByEmployee(context.Context, *EmployeeId) (*ManyVacationsResponse, error)
	// Получить отпуск по его уникальному идентификатору
	GetVacationById(context.Context, *VacationId) (*VacationResponse, error)
	// Добавить новый отпуск для сотрудника в базу данных
	InsertVacation(context.Context, *Vacation) (*VacationResponse, error)
	// Обновить данные для отпуска
	UpdateVacation(context.Context, *UpdateVacRequest) (*VacationResponse, error)
	// Удалить отпуск сотрудника из базы данных
	DeleteVacation(context.Context, *VacationId) (*EmptyResponse, error)
	// Получить все нормы отпуска
	GetAllVacationNorms(context.Context, *EmptyResponse) (*ManyVacationNormsResponse, error)
	// Получить норму по его уникальному идентификатору
	GetVacationNormById(context.Context, *VacationNormId) (*VacationNormResponse, error)
	// Добавить новый норму отпуска в базу данных
	InsertVacationNorm(context.Context, *VacationNorm) (*VacationNormResponse, error)
	// Обновить данные для нормы отпуска
	UpdateVacationNorm(context.Context, *UpdateVacNormRequest) (*VacationNormResponse, error)
	// Удалить норму отпуск из базы данных
	DeleteVacationNorm(context.Context, *VacationNormId) (*EmptyResponse, error)
	mustEmbedUnimplementedVacationsServiceServer()
}

// UnimplementedVacationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVacationsServiceServer struct {
}

func (UnimplementedVacationsServiceServer) GetAllEmployeers(context.Context, *EmptyResponse) (*ManyEmployeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmployeers not implemented")
}
func (UnimplementedVacationsServiceServer) GetEmployeeById(context.Context, *EmployeeId) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeById not implemented")
}
func (UnimplementedVacationsServiceServer) InsertEmployee(context.Context, *Employee) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmployee not implemented")
}
func (UnimplementedVacationsServiceServer) UpdateEmployee(context.Context, *UpdateEmpRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedVacationsServiceServer) DeleteEmployee(context.Context, *EmployeeId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedVacationsServiceServer) GetVacationsByEmployee(context.Context, *EmployeeId) (*ManyVacationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacationsByEmployee not implemented")
}
func (UnimplementedVacationsServiceServer) GetVacationById(context.Context, *VacationId) (*VacationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacationById not implemented")
}
func (UnimplementedVacationsServiceServer) InsertVacation(context.Context, *Vacation) (*VacationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertVacation not implemented")
}
func (UnimplementedVacationsServiceServer) UpdateVacation(context.Context, *UpdateVacRequest) (*VacationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVacation not implemented")
}
func (UnimplementedVacationsServiceServer) DeleteVacation(context.Context, *VacationId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVacation not implemented")
}
func (UnimplementedVacationsServiceServer) GetAllVacationNorms(context.Context, *EmptyResponse) (*ManyVacationNormsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVacationNorms not implemented")
}
func (UnimplementedVacationsServiceServer) GetVacationNormById(context.Context, *VacationNormId) (*VacationNormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacationNormById not implemented")
}
func (UnimplementedVacationsServiceServer) InsertVacationNorm(context.Context, *VacationNorm) (*VacationNormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertVacationNorm not implemented")
}
func (UnimplementedVacationsServiceServer) UpdateVacationNorm(context.Context, *UpdateVacNormRequest) (*VacationNormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVacationNorm not implemented")
}
func (UnimplementedVacationsServiceServer) DeleteVacationNorm(context.Context, *VacationNormId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVacationNorm not implemented")
}
func (UnimplementedVacationsServiceServer) mustEmbedUnimplementedVacationsServiceServer() {}

// UnsafeVacationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacationsServiceServer will
// result in compilation errors.
type UnsafeVacationsServiceServer interface {
	mustEmbedUnimplementedVacationsServiceServer()
}

func RegisterVacationsServiceServer(s *grpc.Server, srv VacationsServiceServer) {
	s.RegisterService(&_VacationsService_serviceDesc, srv)
}

func _VacationsService_GetAllEmployeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetAllEmployeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetAllEmployeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetAllEmployeers(ctx, req.(*EmptyResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_GetEmployeeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetEmployeeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetEmployeeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetEmployeeById(ctx, req.(*EmployeeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_InsertEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).InsertEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/InsertEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).InsertEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).UpdateEmployee(ctx, req.(*UpdateEmpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).DeleteEmployee(ctx, req.(*EmployeeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_GetVacationsByEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetVacationsByEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetVacationsByEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetVacationsByEmployee(ctx, req.(*EmployeeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_GetVacationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetVacationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetVacationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetVacationById(ctx, req.(*VacationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_InsertVacation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).InsertVacation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/InsertVacation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).InsertVacation(ctx, req.(*Vacation))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_UpdateVacation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).UpdateVacation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/UpdateVacation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).UpdateVacation(ctx, req.(*UpdateVacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_DeleteVacation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).DeleteVacation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/DeleteVacation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).DeleteVacation(ctx, req.(*VacationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_GetAllVacationNorms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetAllVacationNorms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetAllVacationNorms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetAllVacationNorms(ctx, req.(*EmptyResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_GetVacationNormById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacationNormId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).GetVacationNormById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/GetVacationNormById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).GetVacationNormById(ctx, req.(*VacationNormId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_InsertVacationNorm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacationNorm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).InsertVacationNorm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/InsertVacationNorm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).InsertVacationNorm(ctx, req.(*VacationNorm))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_UpdateVacationNorm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVacNormRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).UpdateVacationNorm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/UpdateVacationNorm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).UpdateVacationNorm(ctx, req.(*UpdateVacNormRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacationsService_DeleteVacationNorm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacationNormId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacationsServiceServer).DeleteVacationNorm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacation_api.v1.VacationsService/DeleteVacationNorm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacationsServiceServer).DeleteVacationNorm(ctx, req.(*VacationNormId))
	}
	return interceptor(ctx, in, info, handler)
}

var _VacationsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vacation_api.v1.VacationsService",
	HandlerType: (*VacationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllEmployeers",
			Handler:    _VacationsService_GetAllEmployeers_Handler,
		},
		{
			MethodName: "GetEmployeeById",
			Handler:    _VacationsService_GetEmployeeById_Handler,
		},
		{
			MethodName: "InsertEmployee",
			Handler:    _VacationsService_InsertEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _VacationsService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _VacationsService_DeleteEmployee_Handler,
		},
		{
			MethodName: "GetVacationsByEmployee",
			Handler:    _VacationsService_GetVacationsByEmployee_Handler,
		},
		{
			MethodName: "GetVacationById",
			Handler:    _VacationsService_GetVacationById_Handler,
		},
		{
			MethodName: "InsertVacation",
			Handler:    _VacationsService_InsertVacation_Handler,
		},
		{
			MethodName: "UpdateVacation",
			Handler:    _VacationsService_UpdateVacation_Handler,
		},
		{
			MethodName: "DeleteVacation",
			Handler:    _VacationsService_DeleteVacation_Handler,
		},
		{
			MethodName: "GetAllVacationNorms",
			Handler:    _VacationsService_GetAllVacationNorms_Handler,
		},
		{
			MethodName: "GetVacationNormById",
			Handler:    _VacationsService_GetVacationNormById_Handler,
		},
		{
			MethodName: "InsertVacationNorm",
			Handler:    _VacationsService_InsertVacationNorm_Handler,
		},
		{
			MethodName: "UpdateVacationNorm",
			Handler:    _VacationsService_UpdateVacationNorm_Handler,
		},
		{
			MethodName: "DeleteVacationNorm",
			Handler:    _VacationsService_DeleteVacationNorm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacation_api/v1/vacation_api.proto",
}
