// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: opsAgent/model.proto

package pbgo

import (
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

type HostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HostnameRequest) Reset() {
	*x = HostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostnameRequest) ProtoMessage() {}

func (x *HostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostnameRequest.ProtoReflect.Descriptor instead.
func (*HostnameRequest) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{0}
}

type HostnameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *HostnameReply) Reset() {
	*x = HostnameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostnameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostnameReply) ProtoMessage() {}

func (x *HostnameReply) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostnameReply.ProtoReflect.Descriptor instead.
func (*HostnameReply) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{1}
}

func (x *HostnameReply) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type ExecTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cmd  string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *ExecTaskRequest) Reset() {
	*x = ExecTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecTaskRequest) ProtoMessage() {}

func (x *ExecTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecTaskRequest.ProtoReflect.Descriptor instead.
func (*ExecTaskRequest) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{2}
}

func (x *ExecTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecTaskRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type ExecTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cmd    string `protobuf:"bytes,3,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Result string `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ExecTaskReply) Reset() {
	*x = ExecTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecTaskReply) ProtoMessage() {}

func (x *ExecTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecTaskReply.ProtoReflect.Descriptor instead.
func (*ExecTaskReply) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{3}
}

func (x *ExecTaskReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecTaskReply) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *ExecTaskReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ExecTaskReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type WriteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Filepath string `protobuf:"bytes,2,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WriteFileRequest) Reset() {
	*x = WriteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileRequest) ProtoMessage() {}

func (x *WriteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileRequest.ProtoReflect.Descriptor instead.
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{4}
}

func (x *WriteFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *WriteFileRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *WriteFileRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type WriteFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*WriteFileRequest `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *WriteFilesRequest) Reset() {
	*x = WriteFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFilesRequest) ProtoMessage() {}

func (x *WriteFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFilesRequest.ProtoReflect.Descriptor instead.
func (*WriteFilesRequest) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{5}
}

func (x *WriteFilesRequest) GetFiles() []*WriteFileRequest {
	if x != nil {
		return x.Files
	}
	return nil
}

type WriteFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Result   string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	Code     int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *WriteFileReply) Reset() {
	*x = WriteFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileReply) ProtoMessage() {}

func (x *WriteFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileReply.ProtoReflect.Descriptor instead.
func (*WriteFileReply) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{6}
}

func (x *WriteFileReply) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *WriteFileReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *WriteFileReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *WriteFileReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type WriteFilesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WfRes []*WriteFileReply `protobuf:"bytes,1,rep,name=wfRes,proto3" json:"wfRes,omitempty"`
}

func (x *WriteFilesReply) Reset() {
	*x = WriteFilesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opsAgent_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFilesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFilesReply) ProtoMessage() {}

func (x *WriteFilesReply) ProtoReflect() protoreflect.Message {
	mi := &file_opsAgent_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFilesReply.ProtoReflect.Descriptor instead.
func (*WriteFilesReply) Descriptor() ([]byte, []int) {
	return file_opsAgent_model_proto_rawDescGZIP(), []int{7}
}

func (x *WriteFilesReply) GetWfRes() []*WriteFileReply {
	if x != nil {
		return x.WfRes
	}
	return nil
}

var File_opsAgent_model_proto protoreflect.FileDescriptor

var file_opsAgent_model_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x11, 0x0a, 0x0f, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x0d, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x45,
	0x78, 0x65, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x6d, 0x64, 0x22, 0x63, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x64, 0x0a, 0x10, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x4f, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x22, 0x6e, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x4b, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x77, 0x66, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x77, 0x66, 0x52, 0x65, 0x73, 0x42, 0x10, 0x5a,
	0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opsAgent_model_proto_rawDescOnce sync.Once
	file_opsAgent_model_proto_rawDescData = file_opsAgent_model_proto_rawDesc
)

func file_opsAgent_model_proto_rawDescGZIP() []byte {
	file_opsAgent_model_proto_rawDescOnce.Do(func() {
		file_opsAgent_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_opsAgent_model_proto_rawDescData)
	})
	return file_opsAgent_model_proto_rawDescData
}

var file_opsAgent_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_opsAgent_model_proto_goTypes = []interface{}{
	(*HostnameRequest)(nil),   // 0: opsAgent.model.v1.HostnameRequest
	(*HostnameReply)(nil),     // 1: opsAgent.model.v1.HostnameReply
	(*ExecTaskRequest)(nil),   // 2: opsAgent.model.v1.ExecTaskRequest
	(*ExecTaskReply)(nil),     // 3: opsAgent.model.v1.ExecTaskReply
	(*WriteFileRequest)(nil),  // 4: opsAgent.model.v1.WriteFileRequest
	(*WriteFilesRequest)(nil), // 5: opsAgent.model.v1.WriteFilesRequest
	(*WriteFileReply)(nil),    // 6: opsAgent.model.v1.WriteFileReply
	(*WriteFilesReply)(nil),   // 7: opsAgent.model.v1.WriteFilesReply
}
var file_opsAgent_model_proto_depIdxs = []int32{
	4, // 0: opsAgent.model.v1.WriteFilesRequest.files:type_name -> opsAgent.model.v1.WriteFileRequest
	6, // 1: opsAgent.model.v1.WriteFilesReply.wfRes:type_name -> opsAgent.model.v1.WriteFileReply
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_opsAgent_model_proto_init() }
func file_opsAgent_model_proto_init() {
	if File_opsAgent_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opsAgent_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostnameRequest); i {
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
		file_opsAgent_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostnameReply); i {
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
		file_opsAgent_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecTaskRequest); i {
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
		file_opsAgent_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecTaskReply); i {
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
		file_opsAgent_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileRequest); i {
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
		file_opsAgent_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFilesRequest); i {
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
		file_opsAgent_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileReply); i {
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
		file_opsAgent_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFilesReply); i {
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
			RawDescriptor: file_opsAgent_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opsAgent_model_proto_goTypes,
		DependencyIndexes: file_opsAgent_model_proto_depIdxs,
		MessageInfos:      file_opsAgent_model_proto_msgTypes,
	}.Build()
	File_opsAgent_model_proto = out.File
	file_opsAgent_model_proto_rawDesc = nil
	file_opsAgent_model_proto_goTypes = nil
	file_opsAgent_model_proto_depIdxs = nil
}