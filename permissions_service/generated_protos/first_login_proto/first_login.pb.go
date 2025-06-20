// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: protos/first_login.proto

package first_login_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FirstLogin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *string                `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	CreatedBy     uint32                 `protobuf:"varint,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy     uint32                 `protobuf:"varint,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DeletedBy     *uint32                `protobuf:"varint,7,opt,name=deleted_by,json=deletedBy,proto3,oneof" json:"deleted_by,omitempty"`
	UserId        uint32                 `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FirstLogin) Reset() {
	*x = FirstLogin{}
	mi := &file_protos_first_login_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FirstLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirstLogin) ProtoMessage() {}

func (x *FirstLogin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirstLogin.ProtoReflect.Descriptor instead.
func (*FirstLogin) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{0}
}

func (x *FirstLogin) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FirstLogin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FirstLogin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FirstLogin) GetDeletedAt() string {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return ""
}

func (x *FirstLogin) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *FirstLogin) GetUpdatedBy() uint32 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *FirstLogin) GetDeletedBy() uint32 {
	if x != nil && x.DeletedBy != nil {
		return *x.DeletedBy
	}
	return 0
}

func (x *FirstLogin) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_protos_first_login_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_protos_first_login_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateResponse) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_protos_first_login_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_protos_first_login_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{4}
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	Limit         uint32                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        uint32                 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_protos_first_login_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

func (x *GetRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_protos_first_login_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetResponse) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type ListRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Limit          *uint32                `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset         *uint32                `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	IncludeDeleted *bool                  `protobuf:"varint,3,opt,name=include_deleted,json=includeDeleted,proto3,oneof" json:"include_deleted,omitempty"`
	OrderBy        *OrderBy               `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
	UserId         *uint32                `protobuf:"varint,5,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_protos_first_login_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{7}
}

func (x *ListRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() uint32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ListRequest) GetIncludeDeleted() bool {
	if x != nil && x.IncludeDeleted != nil {
		return *x.IncludeDeleted
	}
	return false
}

func (x *ListRequest) GetOrderBy() *OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *ListRequest) GetUserId() uint32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

type OrderBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt     *string                `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	mi := &file_protos_first_login_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBy.ProtoReflect.Descriptor instead.
func (*OrderBy) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{8}
}

func (x *OrderBy) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *OrderBy) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rows          []*FirstLogin          `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	Count         uint32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_protos_first_login_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{9}
}

func (x *ListResponse) GetRows() []*FirstLogin {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *ListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        *uint32                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,3,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_protos_first_login_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetUserId() uint32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *UpdateRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_protos_first_login_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_first_login_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_first_login_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateResponse) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

var File_protos_first_login_proto protoreflect.FileDescriptor

const file_protos_first_login_proto_rawDesc = "" +
	"\n" +
	"\x18protos/first_login.proto\x12\x11first_login_proto\"\x97\x02\n" +
	"\n" +
	"FirstLogin\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\tR\tupdatedAt\x12\"\n" +
	"\n" +
	"deleted_at\x18\x04 \x01(\tH\x00R\tdeletedAt\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"created_by\x18\x05 \x01(\rR\tcreatedBy\x12\x1d\n" +
	"\n" +
	"updated_by\x18\x06 \x01(\rR\tupdatedBy\x12\"\n" +
	"\n" +
	"deleted_by\x18\a \x01(\rH\x01R\tdeletedBy\x88\x01\x01\x12\x17\n" +
	"\auser_id\x18\b \x01(\rR\x06userIdB\r\n" +
	"\v_deleted_atB\r\n" +
	"\v_deleted_by\"K\n" +
	"\rCreateRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId\"L\n" +
	"\x0eCreateResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId\"B\n" +
	"\rDeleteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId\"\x10\n" +
	"\x0eDeleteResponse\"m\n" +
	"\n" +
	"GetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x16\n" +
	"\x06offset\x18\x04 \x01(\rR\x06offset\"I\n" +
	"\vGetResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId\"\x8f\x02\n" +
	"\vListRequest\x12\x19\n" +
	"\x05limit\x18\x01 \x01(\rH\x00R\x05limit\x88\x01\x01\x12\x1b\n" +
	"\x06offset\x18\x02 \x01(\rH\x01R\x06offset\x88\x01\x01\x12,\n" +
	"\x0finclude_deleted\x18\x03 \x01(\bH\x02R\x0eincludeDeleted\x88\x01\x01\x12:\n" +
	"\border_by\x18\x04 \x01(\v2\x1a.first_login_proto.OrderByH\x03R\aorderBy\x88\x01\x01\x12\x1c\n" +
	"\auser_id\x18\x05 \x01(\rH\x04R\x06userId\x88\x01\x01B\b\n" +
	"\x06_limitB\t\n" +
	"\a_offsetB\x12\n" +
	"\x10_include_deletedB\v\n" +
	"\t_order_byB\n" +
	"\n" +
	"\b_user_id\"X\n" +
	"\aOrderBy\x12\x13\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x88\x01\x01\x12\"\n" +
	"\n" +
	"created_at\x18\x02 \x01(\tH\x01R\tcreatedAt\x88\x01\x01B\x05\n" +
	"\x03_idB\r\n" +
	"\v_created_at\"W\n" +
	"\fListResponse\x121\n" +
	"\x04rows\x18\x01 \x03(\v2\x1d.first_login_proto.FirstLoginR\x04rows\x12\x14\n" +
	"\x05count\x18\x02 \x01(\rR\x05count\"l\n" +
	"\rUpdateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1c\n" +
	"\auser_id\x18\x02 \x01(\rH\x00R\x06userId\x88\x01\x01\x12!\n" +
	"\frequester_id\x18\x03 \x01(\rR\vrequesterIdB\n" +
	"\n" +
	"\b_user_id\"L\n" +
	"\x0eUpdateResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12!\n" +
	"\frequester_id\x18\x02 \x01(\rR\vrequesterId2\x8f\x03\n" +
	"\x11FirstLoginService\x12M\n" +
	"\x06Create\x12 .first_login_proto.CreateRequest\x1a!.first_login_proto.CreateResponse\x12M\n" +
	"\x06Delete\x12 .first_login_proto.DeleteRequest\x1a!.first_login_proto.DeleteResponse\x12G\n" +
	"\x04List\x12\x1e.first_login_proto.ListRequest\x1a\x1f.first_login_proto.ListResponse\x12M\n" +
	"\x06Update\x12 .first_login_proto.UpdateRequest\x1a!.first_login_proto.UpdateResponse\x12D\n" +
	"\x03Get\x12\x1d.first_login_proto.GetRequest\x1a\x1e.first_login_proto.GetResponseB&Z$./generated_protos/first_login_protob\x06proto3"

var (
	file_protos_first_login_proto_rawDescOnce sync.Once
	file_protos_first_login_proto_rawDescData []byte
)

func file_protos_first_login_proto_rawDescGZIP() []byte {
	file_protos_first_login_proto_rawDescOnce.Do(func() {
		file_protos_first_login_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_first_login_proto_rawDesc), len(file_protos_first_login_proto_rawDesc)))
	})
	return file_protos_first_login_proto_rawDescData
}

var file_protos_first_login_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protos_first_login_proto_goTypes = []any{
	(*FirstLogin)(nil),     // 0: first_login_proto.FirstLogin
	(*CreateRequest)(nil),  // 1: first_login_proto.CreateRequest
	(*CreateResponse)(nil), // 2: first_login_proto.CreateResponse
	(*DeleteRequest)(nil),  // 3: first_login_proto.DeleteRequest
	(*DeleteResponse)(nil), // 4: first_login_proto.DeleteResponse
	(*GetRequest)(nil),     // 5: first_login_proto.GetRequest
	(*GetResponse)(nil),    // 6: first_login_proto.GetResponse
	(*ListRequest)(nil),    // 7: first_login_proto.ListRequest
	(*OrderBy)(nil),        // 8: first_login_proto.OrderBy
	(*ListResponse)(nil),   // 9: first_login_proto.ListResponse
	(*UpdateRequest)(nil),  // 10: first_login_proto.UpdateRequest
	(*UpdateResponse)(nil), // 11: first_login_proto.UpdateResponse
}
var file_protos_first_login_proto_depIdxs = []int32{
	8,  // 0: first_login_proto.ListRequest.order_by:type_name -> first_login_proto.OrderBy
	0,  // 1: first_login_proto.ListResponse.rows:type_name -> first_login_proto.FirstLogin
	1,  // 2: first_login_proto.FirstLoginService.Create:input_type -> first_login_proto.CreateRequest
	3,  // 3: first_login_proto.FirstLoginService.Delete:input_type -> first_login_proto.DeleteRequest
	7,  // 4: first_login_proto.FirstLoginService.List:input_type -> first_login_proto.ListRequest
	10, // 5: first_login_proto.FirstLoginService.Update:input_type -> first_login_proto.UpdateRequest
	5,  // 6: first_login_proto.FirstLoginService.Get:input_type -> first_login_proto.GetRequest
	2,  // 7: first_login_proto.FirstLoginService.Create:output_type -> first_login_proto.CreateResponse
	4,  // 8: first_login_proto.FirstLoginService.Delete:output_type -> first_login_proto.DeleteResponse
	9,  // 9: first_login_proto.FirstLoginService.List:output_type -> first_login_proto.ListResponse
	11, // 10: first_login_proto.FirstLoginService.Update:output_type -> first_login_proto.UpdateResponse
	6,  // 11: first_login_proto.FirstLoginService.Get:output_type -> first_login_proto.GetResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_first_login_proto_init() }
func file_protos_first_login_proto_init() {
	if File_protos_first_login_proto != nil {
		return
	}
	file_protos_first_login_proto_msgTypes[0].OneofWrappers = []any{}
	file_protos_first_login_proto_msgTypes[7].OneofWrappers = []any{}
	file_protos_first_login_proto_msgTypes[8].OneofWrappers = []any{}
	file_protos_first_login_proto_msgTypes[10].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_first_login_proto_rawDesc), len(file_protos_first_login_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_first_login_proto_goTypes,
		DependencyIndexes: file_protos_first_login_proto_depIdxs,
		MessageInfos:      file_protos_first_login_proto_msgTypes,
	}.Build()
	File_protos_first_login_proto = out.File
	file_protos_first_login_proto_goTypes = nil
	file_protos_first_login_proto_depIdxs = nil
}
