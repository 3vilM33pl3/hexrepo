// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/hexagon.proto

package hexcloud

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HexagonServiceClient is the client API for HexagonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HexagonServiceClient interface {
	RepoAddHexagonInfo(ctx context.Context, in *HexInfoList, opts ...grpc.CallOption) (*Result, error)
	RepoGetHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*HexInfoList, error)
	RepoGetHexagonInfoData(ctx context.Context, in *HexIDData, opts ...grpc.CallOption) (*HexIDData, error)
	RepoGetAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HexInfoList, error)
	RepoDelHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*Result, error)
	RepoDelHexagonInfoData(ctx context.Context, in *HexIDData, opts ...grpc.CallOption) (*Result, error)
	MapAdd(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error)
	MapAddData(ctx context.Context, in *HexLocDataList, opts ...grpc.CallOption) (*Result, error)
	MapGet(ctx context.Context, in *HexagonGetRequest, opts ...grpc.CallOption) (*HexLocationList, error)
	MapUpdate(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error)
	MapUpdateData(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error)
	MapRemove(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error)
	MapRemoveData(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error)
	GetStatusServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	GetStatusStorage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	GetStatusClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	RepoDelAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
	MapRemoveAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
}

type hexagonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHexagonServiceClient(cc grpc.ClientConnInterface) HexagonServiceClient {
	return &hexagonServiceClient{cc}
}

func (c *hexagonServiceClient) RepoAddHexagonInfo(ctx context.Context, in *HexInfoList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoAddHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoGetHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*HexInfoList, error) {
	out := new(HexInfoList)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoGetHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoGetHexagonInfoData(ctx context.Context, in *HexIDData, opts ...grpc.CallOption) (*HexIDData, error) {
	out := new(HexIDData)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoGetHexagonInfoData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoGetAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HexInfoList, error) {
	out := new(HexInfoList)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoGetAllHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoDelHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoDelHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoDelHexagonInfoData(ctx context.Context, in *HexIDData, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoDelHexagonInfoData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapAdd(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapAddData(ctx context.Context, in *HexLocDataList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapAddData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapGet(ctx context.Context, in *HexagonGetRequest, opts ...grpc.CallOption) (*HexLocationList, error) {
	out := new(HexLocationList)
	err := c.cc.Invoke(ctx, "/HexagonService/MapGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapUpdate(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapUpdateData(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapUpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapRemove(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapRemoveData(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapRemoveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/HexagonService/GetStatusServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusStorage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/HexagonService/GetStatusStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/HexagonService/GetStatusClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoDelAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/RepoDelAllHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapRemoveAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/HexagonService/MapRemoveAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HexagonServiceServer is the server API for HexagonService service.
// All implementations must embed UnimplementedHexagonServiceServer
// for forward compatibility
type HexagonServiceServer interface {
	RepoAddHexagonInfo(context.Context, *HexInfoList) (*Result, error)
	RepoGetHexagonInfo(context.Context, *HexIDList) (*HexInfoList, error)
	RepoGetHexagonInfoData(context.Context, *HexIDData) (*HexIDData, error)
	RepoGetAllHexagonInfo(context.Context, *Empty) (*HexInfoList, error)
	RepoDelHexagonInfo(context.Context, *HexIDList) (*Result, error)
	RepoDelHexagonInfoData(context.Context, *HexIDData) (*Result, error)
	MapAdd(context.Context, *HexLocationList) (*Result, error)
	MapAddData(context.Context, *HexLocDataList) (*Result, error)
	MapGet(context.Context, *HexagonGetRequest) (*HexLocationList, error)
	MapUpdate(context.Context, *HexLocation) (*Result, error)
	MapUpdateData(context.Context, *HexLocation) (*Result, error)
	MapRemove(context.Context, *HexLocationList) (*Result, error)
	MapRemoveData(context.Context, *HexLocation) (*Result, error)
	GetStatusServer(context.Context, *Empty) (*Status, error)
	GetStatusStorage(context.Context, *Empty) (*Status, error)
	GetStatusClients(context.Context, *Empty) (*Status, error)
	RepoDelAllHexagonInfo(context.Context, *Empty) (*Result, error)
	MapRemoveAll(context.Context, *Empty) (*Result, error)
	mustEmbedUnimplementedHexagonServiceServer()
}

// UnimplementedHexagonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHexagonServiceServer struct {
}

func (UnimplementedHexagonServiceServer) RepoAddHexagonInfo(context.Context, *HexInfoList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoAddHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoGetHexagonInfo(context.Context, *HexIDList) (*HexInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoGetHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoGetHexagonInfoData(context.Context, *HexIDData) (*HexIDData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoGetHexagonInfoData not implemented")
}
func (UnimplementedHexagonServiceServer) RepoGetAllHexagonInfo(context.Context, *Empty) (*HexInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoGetAllHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoDelHexagonInfo(context.Context, *HexIDList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoDelHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoDelHexagonInfoData(context.Context, *HexIDData) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoDelHexagonInfoData not implemented")
}
func (UnimplementedHexagonServiceServer) MapAdd(context.Context, *HexLocationList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapAdd not implemented")
}
func (UnimplementedHexagonServiceServer) MapAddData(context.Context, *HexLocDataList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapAddData not implemented")
}
func (UnimplementedHexagonServiceServer) MapGet(context.Context, *HexagonGetRequest) (*HexLocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapGet not implemented")
}
func (UnimplementedHexagonServiceServer) MapUpdate(context.Context, *HexLocation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapUpdate not implemented")
}
func (UnimplementedHexagonServiceServer) MapUpdateData(context.Context, *HexLocation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapUpdateData not implemented")
}
func (UnimplementedHexagonServiceServer) MapRemove(context.Context, *HexLocationList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapRemove not implemented")
}
func (UnimplementedHexagonServiceServer) MapRemoveData(context.Context, *HexLocation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapRemoveData not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusServer(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusServer not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusStorage(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusStorage not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusClients(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusClients not implemented")
}
func (UnimplementedHexagonServiceServer) RepoDelAllHexagonInfo(context.Context, *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoDelAllHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) MapRemoveAll(context.Context, *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapRemoveAll not implemented")
}
func (UnimplementedHexagonServiceServer) mustEmbedUnimplementedHexagonServiceServer() {}

// UnsafeHexagonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HexagonServiceServer will
// result in compilation errors.
type UnsafeHexagonServiceServer interface {
	mustEmbedUnimplementedHexagonServiceServer()
}

func RegisterHexagonServiceServer(s grpc.ServiceRegistrar, srv HexagonServiceServer) {
	s.RegisterService(&HexagonService_ServiceDesc, srv)
}

func _HexagonService_RepoAddHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexInfoList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoAddHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoAddHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoAddHexagonInfo(ctx, req.(*HexInfoList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoGetHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoGetHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoGetHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoGetHexagonInfo(ctx, req.(*HexIDList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoGetHexagonInfoData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoGetHexagonInfoData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoGetHexagonInfoData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoGetHexagonInfoData(ctx, req.(*HexIDData))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoGetAllHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoGetAllHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoGetAllHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoGetAllHexagonInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoDelHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoDelHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoDelHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoDelHexagonInfo(ctx, req.(*HexIDList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoDelHexagonInfoData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoDelHexagonInfoData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoDelHexagonInfoData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoDelHexagonInfoData(ctx, req.(*HexIDData))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocationList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapAdd(ctx, req.(*HexLocationList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapAddData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocDataList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapAddData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapAddData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapAddData(ctx, req.(*HexLocDataList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexagonGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapGet(ctx, req.(*HexagonGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapUpdate(ctx, req.(*HexLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapUpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapUpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapUpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapUpdateData(ctx, req.(*HexLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocationList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapRemove(ctx, req.(*HexLocationList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapRemoveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapRemoveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapRemoveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapRemoveData(ctx, req.(*HexLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/GetStatusServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/GetStatusStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusStorage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/GetStatusClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusClients(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoDelAllHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoDelAllHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/RepoDelAllHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoDelAllHexagonInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapRemoveAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapRemoveAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HexagonService/MapRemoveAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapRemoveAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HexagonService_ServiceDesc is the grpc.ServiceDesc for HexagonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HexagonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HexagonService",
	HandlerType: (*HexagonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RepoAddHexagonInfo",
			Handler:    _HexagonService_RepoAddHexagonInfo_Handler,
		},
		{
			MethodName: "RepoGetHexagonInfo",
			Handler:    _HexagonService_RepoGetHexagonInfo_Handler,
		},
		{
			MethodName: "RepoGetHexagonInfoData",
			Handler:    _HexagonService_RepoGetHexagonInfoData_Handler,
		},
		{
			MethodName: "RepoGetAllHexagonInfo",
			Handler:    _HexagonService_RepoGetAllHexagonInfo_Handler,
		},
		{
			MethodName: "RepoDelHexagonInfo",
			Handler:    _HexagonService_RepoDelHexagonInfo_Handler,
		},
		{
			MethodName: "RepoDelHexagonInfoData",
			Handler:    _HexagonService_RepoDelHexagonInfoData_Handler,
		},
		{
			MethodName: "MapAdd",
			Handler:    _HexagonService_MapAdd_Handler,
		},
		{
			MethodName: "MapAddData",
			Handler:    _HexagonService_MapAddData_Handler,
		},
		{
			MethodName: "MapGet",
			Handler:    _HexagonService_MapGet_Handler,
		},
		{
			MethodName: "MapUpdate",
			Handler:    _HexagonService_MapUpdate_Handler,
		},
		{
			MethodName: "MapUpdateData",
			Handler:    _HexagonService_MapUpdateData_Handler,
		},
		{
			MethodName: "MapRemove",
			Handler:    _HexagonService_MapRemove_Handler,
		},
		{
			MethodName: "MapRemoveData",
			Handler:    _HexagonService_MapRemoveData_Handler,
		},
		{
			MethodName: "GetStatusServer",
			Handler:    _HexagonService_GetStatusServer_Handler,
		},
		{
			MethodName: "GetStatusStorage",
			Handler:    _HexagonService_GetStatusStorage_Handler,
		},
		{
			MethodName: "GetStatusClients",
			Handler:    _HexagonService_GetStatusClients_Handler,
		},
		{
			MethodName: "RepoDelAllHexagonInfo",
			Handler:    _HexagonService_RepoDelAllHexagonInfo_Handler,
		},
		{
			MethodName: "MapRemoveAll",
			Handler:    _HexagonService_MapRemoveAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hexagon.proto",
}
