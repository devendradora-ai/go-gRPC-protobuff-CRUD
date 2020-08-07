// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/grpc.proto

package project

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)




const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InventoryProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode     string  `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Store       string  `protobuf:"bytes,2,opt,name=store,proto3" json:"store,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Category    string  `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Image       string  `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Restricted  bool    `protobuf:"varint,7,opt,name=restricted,proto3" json:"restricted,omitempty"`
	Inventory   int32   `protobuf:"varint,8,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Msrp        float32 `protobuf:"fixed32,9,opt,name=msrp,proto3" json:"msrp,omitempty"`
	Price       float32 `protobuf:"fixed32,10,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *InventoryProduct) Reset() {
	*x = InventoryProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryProduct) ProtoMessage() {}

func (x *InventoryProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryProduct.ProtoReflect.Descriptor instead.
func (*InventoryProduct) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryProduct) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *InventoryProduct) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *InventoryProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InventoryProduct) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *InventoryProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InventoryProduct) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *InventoryProduct) GetRestricted() bool {
	if x != nil {
		return x.Restricted
	}
	return false
}

func (x *InventoryProduct) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *InventoryProduct) GetMsrp() float32 {
	if x != nil {
		return x.Msrp
	}
	return 0
}

func (x *InventoryProduct) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_grpc_proto protoreflect.FileDescriptor

var file_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x92, 0x02, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x72, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x73, 0x72, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe2, 0x01, 0x0a, 0x04, 0x43, 0x52, 0x55, 0x44, 0x12,
	0x39, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_proto_rawDescOnce sync.Once
	file_proto_grpc_proto_rawDescData = file_proto_grpc_proto_rawDesc
)

func file_proto_grpc_proto_rawDescGZIP() []byte {
	file_proto_grpc_proto_rawDescOnce.Do(func() {
		file_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_proto_rawDescData)
	})
	return file_proto_grpc_proto_rawDescData
}

var file_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_grpc_proto_goTypes = []interface{}{
	(*InventoryProduct)(nil), // 0: project.InventoryProduct
	(*ID)(nil),               // 1: project.ID
}
var file_proto_grpc_proto_depIdxs = []int32{
	0, // 0: project.CRUD.CreateProduct:input_type -> project.InventoryProduct
	1, // 1: project.CRUD.ReadProduct:input_type -> project.ID
	0, // 2: project.CRUD.UpdateProduct:input_type -> project.InventoryProduct
	1, // 3: project.CRUD.DeleteProduct:input_type -> project.ID
	1, // 4: project.CRUD.CreateProduct:output_type -> project.ID
	0, // 5: project.CRUD.ReadProduct:output_type -> project.InventoryProduct
	1, // 6: project.CRUD.UpdateProduct:output_type -> project.ID
	1, // 7: project.CRUD.DeleteProduct:output_type -> project.ID
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_grpc_proto_init() }
func file_proto_grpc_proto_init() {
	if File_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryProduct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_proto_goTypes,
		DependencyIndexes: file_proto_grpc_proto_depIdxs,
		MessageInfos:      file_proto_grpc_proto_msgTypes,
	}.Build()
	File_proto_grpc_proto = out.File
	file_proto_grpc_proto_rawDesc = nil
	file_proto_grpc_proto_goTypes = nil
	file_proto_grpc_proto_depIdxs = nil
}


// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	CreateProduct(ctx context.Context, in *InventoryProduct, opts ...grpc.CallOption) (*ID, error)
	ReadProduct(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InventoryProduct, error)
	UpdateProduct(ctx context.Context, in *InventoryProduct, opts ...grpc.CallOption) (*ID, error)
	DeleteProduct(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) CreateProduct(ctx context.Context, in *InventoryProduct, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/project.CRUD/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) ReadProduct(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InventoryProduct, error) {
	out := new(InventoryProduct)
	err := c.cc.Invoke(ctx, "/project.CRUD/ReadProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateProduct(ctx context.Context, in *InventoryProduct, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/project.CRUD/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteProduct(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/project.CRUD/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	CreateProduct(context.Context, *InventoryProduct) (*ID, error)
	ReadProduct(context.Context, *ID) (*InventoryProduct, error)
	UpdateProduct(context.Context, *InventoryProduct) (*ID, error)
	DeleteProduct(context.Context, *ID) (*ID, error)
	//mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (*UnimplementedCRUDServer) CreateProduct(context.Context, *InventoryProduct) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedCRUDServer) ReadProduct(context.Context, *ID) (*InventoryProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProduct not implemented")
}
func (*UnimplementedCRUDServer) UpdateProduct(context.Context, *InventoryProduct) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (*UnimplementedCRUDServer) DeleteProduct(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (*UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

func RegisterCRUDServer(s *grpc.Server, srv CRUDServer) {
	s.RegisterService(&_CRUD_serviceDesc, srv)
}

func _CRUD_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.CRUD/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).CreateProduct(ctx, req.(*InventoryProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_ReadProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).ReadProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.CRUD/ReadProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).ReadProduct(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.CRUD/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateProduct(ctx, req.(*InventoryProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.CRUD/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteProduct(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRUD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "project.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _CRUD_CreateProduct_Handler,
		},
		{
			MethodName: "ReadProduct",
			Handler:    _CRUD_ReadProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _CRUD_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _CRUD_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc.proto",
}

