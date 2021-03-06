// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Values struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *Values) Reset()                    { *m = Values{} }
func (m *Values) String() string            { return proto1.CompactTextString(m) }
func (*Values) ProtoMessage()               {}
func (*Values) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Values) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type Response struct {
	StatusCode int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Headers    map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       []byte             `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetRequest struct {
	Uri     string             `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Headers map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *GetRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *GetRequest) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

type PutRequest struct {
	Uri     string             `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Headers map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body    []byte             `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto1.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *PutRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *PutRequest) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *PutRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type PostRequest struct {
	Uri     string             `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Headers map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body    []byte             `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto1.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *PostRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *PostRequest) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *PostRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type PatchRequest struct {
	Uri     string             `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Headers map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body    []byte             `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *PatchRequest) Reset()                    { *m = PatchRequest{} }
func (m *PatchRequest) String() string            { return proto1.CompactTextString(m) }
func (*PatchRequest) ProtoMessage()               {}
func (*PatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PatchRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *PatchRequest) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *PatchRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type DeleteRequest struct {
	Uri     string             `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Headers map[string]*Values `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *DeleteRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *DeleteRequest) GetHeaders() map[string]*Values {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto1.RegisterType((*Values)(nil), "proto.Values")
	proto1.RegisterType((*Response)(nil), "proto.Response")
	proto1.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto1.RegisterType((*PutRequest)(nil), "proto.PutRequest")
	proto1.RegisterType((*PostRequest)(nil), "proto.PostRequest")
	proto1.RegisterType((*PatchRequest)(nil), "proto.PatchRequest")
	proto1.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Response, error)
	Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Api/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Api/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Api/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Api/Patch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Api/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	Get(context.Context, *GetRequest) (*Response, error)
	Put(context.Context, *PutRequest) (*Response, error)
	Post(context.Context, *PostRequest) (*Response, error)
	Patch(context.Context, *PatchRequest) (*Response, error)
	Delete(context.Context, *DeleteRequest) (*Response, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Api/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Api/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Api/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Api/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Patch(ctx, req.(*PatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Api/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Api_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Api_Put_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Api_Post_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _Api_Patch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Api_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}

func init() { proto1.RegisterFile("platform.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0xb3, 0x94, 0xa2, 0x4c, 0xc1, 0x8f, 0x91, 0x03, 0x21, 0x06, 0x2a, 0x5e, 0x6a, 0x8c,
	0x98, 0x60, 0x62, 0x10, 0x4f, 0x46, 0x0d, 0x7a, 0x6b, 0x7a, 0xf0, 0x6a, 0x0a, 0x5d, 0x03, 0x11,
	0xd9, 0xda, 0xdd, 0x35, 0xe1, 0x79, 0x3c, 0xf9, 0x04, 0xc6, 0x78, 0xf3, 0x69, 0x7c, 0x0c, 0xd3,
	0xaf, 0x14, 0x5b, 0xf0, 0x60, 0x34, 0x9c, 0x3a, 0xdb, 0xfe, 0xff, 0x33, 0xbf, 0x9d, 0x9d, 0x2e,
	0xac, 0xb9, 0x63, 0x5b, 0xdc, 0x31, 0xef, 0xa1, 0xe5, 0x7a, 0x4c, 0x30, 0x54, 0x83, 0x47, 0xb3,
	0x0e, 0x85, 0x1b, 0x7b, 0x2c, 0x29, 0xc7, 0x0a, 0xa8, 0x4f, 0x7e, 0x54, 0x25, 0xba, 0x62, 0x14,
	0xad, 0x70, 0xd1, 0xfc, 0x20, 0xb0, 0x6a, 0x51, 0xee, 0xb2, 0x09, 0xa7, 0xd8, 0x00, 0x8d, 0x0b,
	0x5b, 0x48, 0x7e, 0x3b, 0x60, 0x8e, 0x2f, 0x24, 0x86, 0x6a, 0x41, 0xf8, 0xea, 0x9c, 0x39, 0x14,
	0x8f, 0x61, 0x65, 0x48, 0x6d, 0x87, 0x7a, 0xbc, 0x9a, 0xd3, 0x15, 0x43, 0x6b, 0x6f, 0x87, 0xd5,
	0x5a, 0x71, 0x8a, 0xd6, 0x55, 0xf8, 0xf9, 0x72, 0x22, 0xbc, 0xa9, 0x15, 0x8b, 0x11, 0x21, 0xdf,
	0x67, 0xce, 0xb4, 0xaa, 0xe8, 0xc4, 0x28, 0x59, 0x41, 0x5c, 0xbb, 0x86, 0xd2, 0xac, 0x18, 0x37,
	0x40, 0xb9, 0xa7, 0xd3, 0xa0, 0x68, 0xd1, 0xf2, 0x43, 0xdc, 0x8d, 0x89, 0x73, 0x3a, 0x31, 0xb4,
	0x76, 0x39, 0xaa, 0x15, 0xee, 0x27, 0xda, 0x40, 0x37, 0xd7, 0x21, 0xcd, 0x67, 0x02, 0xd0, 0xa3,
	0xc2, 0xa2, 0x8f, 0x92, 0x72, 0xe1, 0x67, 0x92, 0xde, 0x28, 0xce, 0x24, 0xbd, 0x11, 0x76, 0xd2,
	0xdc, 0xf5, 0x28, 0x57, 0xe2, 0x9a, 0x4f, 0xfe, 0x97, 0x94, 0xaf, 0x04, 0xc0, 0x94, 0xbf, 0xa1,
	0x4c, 0x5c, 0xcb, 0xe9, 0xef, 0x1b, 0x01, 0xcd, 0x64, 0xfc, 0x07, 0xf4, 0x93, 0x34, 0x7a, 0x23,
	0x46, 0x4f, 0x6c, 0xcb, 0x61, 0x7f, 0x27, 0x50, 0x32, 0x6d, 0x31, 0x18, 0x2e, 0x86, 0xef, 0xa6,
	0xe1, 0xf5, 0x18, 0x7e, 0xc6, 0xb7, 0x1c, 0xfa, 0x17, 0x02, 0xe5, 0x0b, 0x3a, 0xa6, 0x82, 0x2e,
	0xc6, 0x3f, 0x4d, 0xe3, 0xef, 0x44, 0xe9, 0xbe, 0x19, 0xff, 0x7d, 0xbe, 0xdb, 0x9f, 0x04, 0x94,
	0x33, 0x77, 0x84, 0x7b, 0xa0, 0xf4, 0xa8, 0xc0, 0xcd, 0xcc, 0x2f, 0x56, 0x5b, 0x4f, 0xdd, 0x16,
	0xbe, 0xd4, 0x94, 0x89, 0x34, 0x99, 0xf3, 0xac, 0x74, 0x1f, 0xf2, 0xfe, 0x2c, 0x21, 0x66, 0x07,
	0x2b, 0x2b, 0x3e, 0x00, 0x35, 0x38, 0x3b, 0xdc, 0x9a, 0x73, 0x92, 0x59, 0xf9, 0x21, 0x14, 0xc2,
	0x5e, 0x61, 0x65, 0x5e, 0xeb, 0x32, 0x86, 0x7e, 0x21, 0x58, 0x1f, 0x7d, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0x18, 0xf4, 0x89, 0x75, 0x05, 0x00, 0x00,
}
