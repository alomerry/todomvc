// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: todo/todo.proto

package todo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_UNSET Status = 0
	Status_UNDO  Status = 1
	Status_DONE  Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNSET",
		1: "UNDO",
		2: "DONE",
	}
	Status_value = map[string]int32{
		"UNSET": 0,
		"UNDO":  1,
		"DONE":  2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_todo_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_todo_todo_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{0}
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DoneAt    int64  `protobuf:"varint,4,opt,name=doneAt,proto3" json:"doneAt,omitempty"`
	Color     string `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Status    Status `protobuf:"varint,6,opt,name=status,proto3,enum=proto.todo.Status" json:"status,omitempty"`
	Content   string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Todo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Todo) GetDoneAt() int64 {
	if x != nil {
		return x.DoneAt
	}
	return 0
}

func (x *Todo) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Todo) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNSET
}

func (x *Todo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy    string `protobuf:"bytes,1,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	IsAscend  bool   `protobuf:"varint,2,opt,name=isAscend,proto3" json:"isAscend,omitempty"`
	Page      int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	UserId    string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	StartedAt int64  `protobuf:"varint,6,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt   int64  `protobuf:"varint,7,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	Keyword   string `protobuf:"bytes,8,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Status    Status `protobuf:"varint,9,opt,name=status,proto3,enum=proto.todo.Status" json:"status,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{1}
}

func (x *GetTodoRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetTodoRequest) GetIsAscend() bool {
	if x != nil {
		return x.IsAscend
	}
	return false
}

func (x *GetTodoRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetTodoRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetTodoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetTodoRequest) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *GetTodoRequest) GetEndedAt() int64 {
	if x != nil {
		return x.EndedAt
	}
	return 0
}

func (x *GetTodoRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetTodoRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNSET
}

type GetTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetTodoResponse) Reset() {
	*x = GetTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoResponse) ProtoMessage() {}

func (x *GetTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoResponse.ProtoReflect.Descriptor instead.
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{2}
}

func (x *GetTodoResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

func (x *GetTodoResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AddTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color   string `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddTodoRequest) Reset() {
	*x = AddTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoRequest) ProtoMessage() {}

func (x *AddTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoRequest.ProtoReflect.Descriptor instead.
func (*AddTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{3}
}

func (x *AddTodoRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *AddTodoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddTodoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RemoveTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveTodoRequest) Reset() {
	*x = RemoveTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTodoRequest) ProtoMessage() {}

func (x *RemoveTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTodoRequest.ProtoReflect.Descriptor instead.
func (*RemoveTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoRequest) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{6}
}

var File_todo_todo_proto protoreflect.FileDescriptor

var file_todo_todo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0xc0, 0x01,
	0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x6f,
	0x6e, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x8a, 0x02, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4f, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x58,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa1, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x27, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x4e, 0x44, 0x4f, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_todo_proto_rawDescOnce sync.Once
	file_todo_todo_proto_rawDescData = file_todo_todo_proto_rawDesc
)

func file_todo_todo_proto_rawDescGZIP() []byte {
	file_todo_todo_proto_rawDescOnce.Do(func() {
		file_todo_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_todo_proto_rawDescData)
	})
	return file_todo_todo_proto_rawDescData
}

var file_todo_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_todo_todo_proto_goTypes = []interface{}{
	(Status)(0),               // 0: proto.todo.Status
	(*Todo)(nil),              // 1: proto.todo.Todo
	(*GetTodoRequest)(nil),    // 2: proto.todo.GetTodoRequest
	(*GetTodoResponse)(nil),   // 3: proto.todo.GetTodoResponse
	(*AddTodoRequest)(nil),    // 4: proto.todo.AddTodoRequest
	(*RemoveTodoRequest)(nil), // 5: proto.todo.RemoveTodoRequest
	(*UpdateTodoRequest)(nil), // 6: proto.todo.UpdateTodoRequest
	(*EmptyResponse)(nil),     // 7: proto.todo.EmptyResponse
	nil,                       // 8: proto.todo.UpdateTodoRequest.FieldsEntry
}
var file_todo_todo_proto_depIdxs = []int32{
	0, // 0: proto.todo.Todo.status:type_name -> proto.todo.Status
	0, // 1: proto.todo.GetTodoRequest.status:type_name -> proto.todo.Status
	1, // 2: proto.todo.GetTodoResponse.todos:type_name -> proto.todo.Todo
	8, // 3: proto.todo.UpdateTodoRequest.fields:type_name -> proto.todo.UpdateTodoRequest.FieldsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_todo_todo_proto_init() }
func file_todo_todo_proto_init() {
	if File_todo_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todo_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
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
		file_todo_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoResponse); i {
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
		file_todo_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoRequest); i {
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
		file_todo_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTodoRequest); i {
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
		file_todo_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoRequest); i {
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
		file_todo_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_todo_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_todo_todo_proto_goTypes,
		DependencyIndexes: file_todo_todo_proto_depIdxs,
		EnumInfos:         file_todo_todo_proto_enumTypes,
		MessageInfos:      file_todo_todo_proto_msgTypes,
	}.Build()
	File_todo_todo_proto = out.File
	file_todo_todo_proto_rawDesc = nil
	file_todo_todo_proto_goTypes = nil
	file_todo_todo_proto_depIdxs = nil
}
