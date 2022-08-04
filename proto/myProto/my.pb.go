// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.3
// source: my.proto

package myproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request 请求结构
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"` //1 执行普通命令 2 执行cron命令
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Spec    string `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"` //cron 表达式
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Request) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

// Response 响应结构
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//文件上传请求
type UpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpRequest) Reset() {
	*x = UpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpRequest) ProtoMessage() {}

func (x *UpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpRequest.ProtoReflect.Descriptor instead.
func (*UpRequest) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{2}
}

func (x *UpRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *UpRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//文件上传响应
type UpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Code     int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"` //状态码
	Msg      string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`    //消息
}

func (x *UpResponse) Reset() {
	*x = UpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpResponse) ProtoMessage() {}

func (x *UpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpResponse.ProtoReflect.Descriptor instead.
func (*UpResponse) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{3}
}

func (x *UpResponse) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *UpResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//文件下载请求
type DlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *DlRequest) Reset() {
	*x = DlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DlRequest) ProtoMessage() {}

func (x *DlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DlRequest.ProtoReflect.Descriptor instead.
func (*DlRequest) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{4}
}

func (x *DlRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

//文件下载响应
type DlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"` //状态码
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`    //消息
}

func (x *DlResponse) Reset() {
	*x = DlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DlResponse) ProtoMessage() {}

func (x *DlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DlResponse.ProtoReflect.Descriptor instead.
func (*DlResponse) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{5}
}

func (x *DlResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DlResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DlResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_my_proto protoreflect.FileDescriptor

var file_my_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x79, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x0a, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x27, 0x0a, 0x09, 0x44, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x46, 0x0a, 0x0a, 0x44, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x3c, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12,
	0x35, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e,
	0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0x7e, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x6d,
	0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x79, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_proto_rawDescOnce sync.Once
	file_my_proto_rawDescData = file_my_proto_rawDesc
)

func file_my_proto_rawDescGZIP() []byte {
	file_my_proto_rawDescOnce.Do(func() {
		file_my_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_proto_rawDescData)
	})
	return file_my_proto_rawDescData
}

var file_my_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_my_proto_goTypes = []interface{}{
	(*Request)(nil),    // 0: myproto.Request
	(*Response)(nil),   // 1: myproto.Response
	(*UpRequest)(nil),  // 2: myproto.UpRequest
	(*UpResponse)(nil), // 3: myproto.UpResponse
	(*DlRequest)(nil),  // 4: myproto.DlRequest
	(*DlResponse)(nil), // 5: myproto.DlResponse
}
var file_my_proto_depIdxs = []int32{
	0, // 0: myproto.Cmd.ExecStream:input_type -> myproto.Request
	2, // 1: myproto.File.UploadFile:input_type -> myproto.UpRequest
	4, // 2: myproto.File.DownloadFile:input_type -> myproto.DlRequest
	1, // 3: myproto.Cmd.ExecStream:output_type -> myproto.Response
	3, // 4: myproto.File.UploadFile:output_type -> myproto.UpResponse
	5, // 5: myproto.File.DownloadFile:output_type -> myproto.DlResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_my_proto_init() }
func file_my_proto_init() {
	if File_my_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_my_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_my_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_my_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpRequest); i {
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
		file_my_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpResponse); i {
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
		file_my_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DlRequest); i {
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
		file_my_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DlResponse); i {
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
			RawDescriptor: file_my_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_my_proto_goTypes,
		DependencyIndexes: file_my_proto_depIdxs,
		MessageInfos:      file_my_proto_msgTypes,
	}.Build()
	File_my_proto = out.File
	file_my_proto_rawDesc = nil
	file_my_proto_goTypes = nil
	file_my_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmdClient is the client API for Cmd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdClient interface {
	//服务端推送流，响应都以数据流式处理
	ExecStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Cmd_ExecStreamClient, error)
}

type cmdClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdClient(cc grpc.ClientConnInterface) CmdClient {
	return &cmdClient{cc}
}

func (c *cmdClient) ExecStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Cmd_ExecStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cmd_serviceDesc.Streams[0], "/myproto.Cmd/ExecStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmdExecStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cmd_ExecStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type cmdExecStreamClient struct {
	grpc.ClientStream
}

func (x *cmdExecStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CmdServer is the server API for Cmd service.
type CmdServer interface {
	//服务端推送流，响应都以数据流式处理
	ExecStream(*Request, Cmd_ExecStreamServer) error
}

// UnimplementedCmdServer can be embedded to have forward compatible implementations.
type UnimplementedCmdServer struct {
}

func (*UnimplementedCmdServer) ExecStream(*Request, Cmd_ExecStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecStream not implemented")
}

func RegisterCmdServer(s *grpc.Server, srv CmdServer) {
	s.RegisterService(&_Cmd_serviceDesc, srv)
}

func _Cmd_ExecStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmdServer).ExecStream(m, &cmdExecStreamServer{stream})
}

type Cmd_ExecStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type cmdExecStreamServer struct {
	grpc.ServerStream
}

func (x *cmdExecStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Cmd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.Cmd",
	HandlerType: (*CmdServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecStream",
			Handler:       _Cmd_ExecStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "my.proto",
}

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileClient interface {
	//服务端推送流
	//请求的是流，响应的是普通响应
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (File_UploadFileClient, error)
	//请求的是普通请求，响应的是流式响应
	DownloadFile(ctx context.Context, in *DlRequest, opts ...grpc.CallOption) (File_DownloadFileClient, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (File_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[0], "/myproto.File/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadFileClient{stream}
	return x, nil
}

type File_UploadFileClient interface {
	Send(*UpRequest) error
	CloseAndRecv() (*UpResponse, error)
	grpc.ClientStream
}

type fileUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileUploadFileClient) Send(m *UpRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadFileClient) CloseAndRecv() (*UpResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileClient) DownloadFile(ctx context.Context, in *DlRequest, opts ...grpc.CallOption) (File_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[1], "/myproto.File/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type File_DownloadFileClient interface {
	Recv() (*DlResponse, error)
	grpc.ClientStream
}

type fileDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileDownloadFileClient) Recv() (*DlResponse, error) {
	m := new(DlResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	//服务端推送流
	//请求的是流，响应的是普通响应
	UploadFile(File_UploadFileServer) error
	//请求的是普通请求，响应的是流式响应
	DownloadFile(*DlRequest, File_DownloadFileServer) error
}

// UnimplementedFileServer can be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (*UnimplementedFileServer) UploadFile(File_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFileServer) DownloadFile(*DlRequest, File_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServer).UploadFile(&fileUploadFileServer{stream})
}

type File_UploadFileServer interface {
	SendAndClose(*UpResponse) error
	Recv() (*UpRequest, error)
	grpc.ServerStream
}

type fileUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileUploadFileServer) SendAndClose(m *UpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadFileServer) Recv() (*UpRequest, error) {
	m := new(UpRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _File_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServer).DownloadFile(m, &fileDownloadFileServer{stream})
}

type File_DownloadFileServer interface {
	Send(*DlResponse) error
	grpc.ServerStream
}

type fileDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileDownloadFileServer) Send(m *DlResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.File",
	HandlerType: (*FileServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _File_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _File_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "my.proto",
}
