// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: protos/roles.proto

package roles_proto

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

type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *string                `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	CreatedBy     uint32                 `protobuf:"varint,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy     uint32                 `protobuf:"varint,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DeletedBy     *uint32                `protobuf:"varint,9,opt,name=deleted_by,json=deletedBy,proto3,oneof" json:"deleted_by,omitempty"`
	IsActive      bool                   `protobuf:"varint,10,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Permission    string                 `protobuf:"bytes,11,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_protos_roles_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Permission) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Permission) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Permission) GetDeletedAt() string {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return ""
}

func (x *Permission) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Permission) GetUpdatedBy() uint32 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *Permission) GetDeletedBy() uint32 {
	if x != nil && x.DeletedBy != nil {
		return *x.DeletedBy
	}
	return 0
}

func (x *Permission) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Permission) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *string                `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	CreatedBy     uint32                 `protobuf:"varint,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy     uint32                 `protobuf:"varint,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DeletedBy     *uint32                `protobuf:"varint,9,opt,name=deleted_by,json=deletedBy,proto3,oneof" json:"deleted_by,omitempty"`
	IsActive      bool                   `protobuf:"varint,10,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_protos_roles_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Role) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Role) GetDeletedAt() string {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return ""
}

func (x *Role) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Role) GetUpdatedBy() uint32 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *Role) GetDeletedBy() uint32 {
	if x != nil && x.DeletedBy != nil {
		return *x.DeletedBy
	}
	return 0
}

func (x *Role) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_protos_roles_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[2]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *Role                  `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_protos_roles_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[3]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_protos_roles_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[4]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *Role                  `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_protos_roles_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[5]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type ListRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Limit          uint32                 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset         uint32                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	IncludeDeleted *bool                  `protobuf:"varint,3,opt,name=include_deleted,json=includeDeleted,proto3,oneof" json:"include_deleted,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_protos_roles_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[6]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{6}
}

func (x *ListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest) GetIncludeDeleted() bool {
	if x != nil && x.IncludeDeleted != nil {
		return *x.IncludeDeleted
	}
	return false
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rows          []*Role                `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	Count         uint32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_protos_roles_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[7]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetRows() []*Role {
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
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,4,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_protos_roles_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[8]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *Role                  `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_protos_roles_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[9]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,2,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_protos_roles_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[10]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{10}
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
	mi := &file_protos_roles_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[11]
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
	return file_protos_roles_proto_rawDescGZIP(), []int{11}
}

type AddPermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleId        uint32                 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	PermissionId  uint32                 `protobuf:"varint,2,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,3,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPermissionRequest) Reset() {
	*x = AddPermissionRequest{}
	mi := &file_protos_roles_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionRequest) ProtoMessage() {}

func (x *AddPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionRequest) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{12}
}

func (x *AddPermissionRequest) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *AddPermissionRequest) GetPermissionId() uint32 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *AddPermissionRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type AddPermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *Role                  `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPermissionResponse) Reset() {
	*x = AddPermissionResponse{}
	mi := &file_protos_roles_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionResponse) ProtoMessage() {}

func (x *AddPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionResponse.ProtoReflect.Descriptor instead.
func (*AddPermissionResponse) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{13}
}

func (x *AddPermissionResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RemovePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleId        uint32                 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	PermissionId  uint32                 `protobuf:"varint,2,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	RequesterId   uint32                 `protobuf:"varint,3,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemovePermissionRequest) Reset() {
	*x = RemovePermissionRequest{}
	mi := &file_protos_roles_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionRequest) ProtoMessage() {}

func (x *RemovePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionRequest) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{14}
}

func (x *RemovePermissionRequest) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *RemovePermissionRequest) GetPermissionId() uint32 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *RemovePermissionRequest) GetRequesterId() uint32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

type RemovePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *Role                  `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemovePermissionResponse) Reset() {
	*x = RemovePermissionResponse{}
	mi := &file_protos_roles_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionResponse) ProtoMessage() {}

func (x *RemovePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionResponse.ProtoReflect.Descriptor instead.
func (*RemovePermissionResponse) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{15}
}

func (x *RemovePermissionResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type GetPermissionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleId        uint32                 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Limit         uint32                 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        uint32                 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPermissionsRequest) Reset() {
	*x = GetPermissionsRequest{}
	mi := &file_protos_roles_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsRequest) ProtoMessage() {}

func (x *GetPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{16}
}

func (x *GetPermissionsRequest) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *GetPermissionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPermissionsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetPermissionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rows          []*Permission          `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	Count         uint32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPermissionsResponse) Reset() {
	*x = GetPermissionsResponse{}
	mi := &file_protos_roles_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsResponse) ProtoMessage() {}

func (x *GetPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_roles_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_protos_roles_proto_rawDescGZIP(), []int{17}
}

func (x *GetPermissionsResponse) GetRows() []*Permission {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *GetPermissionsResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_roles_proto protoreflect.FileDescriptor

const file_protos_roles_proto_rawDesc = "" +
	"\n" +
	"\x12protos/roles.proto\x12\vroles_proto\"\xf1\x02\n" +
	"\n" +
	"Permission\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\tR\tupdatedAt\x12\"\n" +
	"\n" +
	"deleted_at\x18\x06 \x01(\tH\x00R\tdeletedAt\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"created_by\x18\a \x01(\rR\tcreatedBy\x12\x1d\n" +
	"\n" +
	"updated_by\x18\b \x01(\rR\tupdatedBy\x12\"\n" +
	"\n" +
	"deleted_by\x18\t \x01(\rH\x01R\tdeletedBy\x88\x01\x01\x12\x1b\n" +
	"\tis_active\x18\n" +
	" \x01(\bR\bisActive\x12\x1e\n" +
	"\n" +
	"permission\x18\v \x01(\tR\n" +
	"permissionB\r\n" +
	"\v_deleted_atB\r\n" +
	"\v_deleted_by\"\xcb\x02\n" +
	"\x04Role\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\tR\tupdatedAt\x12\"\n" +
	"\n" +
	"deleted_at\x18\x06 \x01(\tH\x00R\tdeletedAt\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"created_by\x18\a \x01(\rR\tcreatedBy\x12\x1d\n" +
	"\n" +
	"updated_by\x18\b \x01(\rR\tupdatedBy\x12\"\n" +
	"\n" +
	"deleted_by\x18\t \x01(\rH\x01R\tdeletedBy\x88\x01\x01\x12\x1b\n" +
	"\tis_active\x18\n" +
	" \x01(\bR\bisActiveB\r\n" +
	"\v_deleted_atB\r\n" +
	"\v_deleted_by\"g\n" +
	"\rCreateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vrequesterId\x18\x02 \x01(\rR\vrequesterId\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"7\n" +
	"\x0eCreateResponse\x12%\n" +
	"\x04role\x18\x01 \x01(\v2\x11.roles_proto.RoleR\x04role\"\x1c\n" +
	"\n" +
	"GetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"4\n" +
	"\vGetResponse\x12%\n" +
	"\x04role\x18\x01 \x01(\v2\x11.roles_proto.RoleR\x04role\"}\n" +
	"\vListRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\rR\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\rR\x06offset\x12,\n" +
	"\x0finclude_deleted\x18\x03 \x01(\bH\x00R\x0eincludeDeleted\x88\x01\x01B\x12\n" +
	"\x10_include_deleted\"K\n" +
	"\fListResponse\x12%\n" +
	"\x04rows\x18\x01 \x03(\v2\x11.roles_proto.RoleR\x04rows\x12\x14\n" +
	"\x05count\x18\x02 \x01(\rR\x05count\"\x9a\x01\n" +
	"\rUpdateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x12%\n" +
	"\vdescription\x18\x03 \x01(\tH\x01R\vdescription\x88\x01\x01\x12 \n" +
	"\vrequesterId\x18\x04 \x01(\rR\vrequesterIdB\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_description\"7\n" +
	"\x0eUpdateResponse\x12%\n" +
	"\x04role\x18\x01 \x01(\v2\x11.roles_proto.RoleR\x04role\"A\n" +
	"\rDeleteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12 \n" +
	"\vrequesterId\x18\x02 \x01(\rR\vrequesterId\"\x10\n" +
	"\x0eDeleteResponse\"v\n" +
	"\x14AddPermissionRequest\x12\x17\n" +
	"\arole_id\x18\x01 \x01(\rR\x06roleId\x12#\n" +
	"\rpermission_id\x18\x02 \x01(\rR\fpermissionId\x12 \n" +
	"\vrequesterId\x18\x03 \x01(\rR\vrequesterId\">\n" +
	"\x15AddPermissionResponse\x12%\n" +
	"\x04role\x18\x01 \x01(\v2\x11.roles_proto.RoleR\x04role\"y\n" +
	"\x17RemovePermissionRequest\x12\x17\n" +
	"\arole_id\x18\x01 \x01(\rR\x06roleId\x12#\n" +
	"\rpermission_id\x18\x02 \x01(\rR\fpermissionId\x12 \n" +
	"\vrequesterId\x18\x03 \x01(\rR\vrequesterId\"A\n" +
	"\x18RemovePermissionResponse\x12%\n" +
	"\x04role\x18\x01 \x01(\v2\x11.roles_proto.RoleR\x04role\"^\n" +
	"\x15GetPermissionsRequest\x12\x17\n" +
	"\arole_id\x18\x01 \x01(\rR\x06roleId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\rR\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\rR\x06offset\"[\n" +
	"\x16GetPermissionsResponse\x12+\n" +
	"\x04rows\x18\x01 \x03(\v2\x17.roles_proto.PermissionR\x04rows\x12\x14\n" +
	"\x05count\x18\x02 \x01(\rR\x05count2\xe2\x04\n" +
	"\fRolesService\x12A\n" +
	"\x06Create\x12\x1a.roles_proto.CreateRequest\x1a\x1b.roles_proto.CreateResponse\x128\n" +
	"\x03Get\x12\x17.roles_proto.GetRequest\x1a\x18.roles_proto.GetResponse\x12;\n" +
	"\x04List\x12\x18.roles_proto.ListRequest\x1a\x19.roles_proto.ListResponse\x12A\n" +
	"\x06Update\x12\x1a.roles_proto.UpdateRequest\x1a\x1b.roles_proto.UpdateResponse\x12A\n" +
	"\x06Delete\x12\x1a.roles_proto.DeleteRequest\x1a\x1b.roles_proto.DeleteResponse\x12V\n" +
	"\rAddPermission\x12!.roles_proto.AddPermissionRequest\x1a\".roles_proto.AddPermissionResponse\x12_\n" +
	"\x10RemovePermission\x12$.roles_proto.RemovePermissionRequest\x1a%.roles_proto.RemovePermissionResponse\x12Y\n" +
	"\x0eGetPermissions\x12\".roles_proto.GetPermissionsRequest\x1a#.roles_proto.GetPermissionsResponseB Z\x1e./generated_protos/roles_protob\x06proto3"

var (
	file_protos_roles_proto_rawDescOnce sync.Once
	file_protos_roles_proto_rawDescData []byte
)

func file_protos_roles_proto_rawDescGZIP() []byte {
	file_protos_roles_proto_rawDescOnce.Do(func() {
		file_protos_roles_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_roles_proto_rawDesc), len(file_protos_roles_proto_rawDesc)))
	})
	return file_protos_roles_proto_rawDescData
}

var file_protos_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_protos_roles_proto_goTypes = []any{
	(*Permission)(nil),               // 0: roles_proto.Permission
	(*Role)(nil),                     // 1: roles_proto.Role
	(*CreateRequest)(nil),            // 2: roles_proto.CreateRequest
	(*CreateResponse)(nil),           // 3: roles_proto.CreateResponse
	(*GetRequest)(nil),               // 4: roles_proto.GetRequest
	(*GetResponse)(nil),              // 5: roles_proto.GetResponse
	(*ListRequest)(nil),              // 6: roles_proto.ListRequest
	(*ListResponse)(nil),             // 7: roles_proto.ListResponse
	(*UpdateRequest)(nil),            // 8: roles_proto.UpdateRequest
	(*UpdateResponse)(nil),           // 9: roles_proto.UpdateResponse
	(*DeleteRequest)(nil),            // 10: roles_proto.DeleteRequest
	(*DeleteResponse)(nil),           // 11: roles_proto.DeleteResponse
	(*AddPermissionRequest)(nil),     // 12: roles_proto.AddPermissionRequest
	(*AddPermissionResponse)(nil),    // 13: roles_proto.AddPermissionResponse
	(*RemovePermissionRequest)(nil),  // 14: roles_proto.RemovePermissionRequest
	(*RemovePermissionResponse)(nil), // 15: roles_proto.RemovePermissionResponse
	(*GetPermissionsRequest)(nil),    // 16: roles_proto.GetPermissionsRequest
	(*GetPermissionsResponse)(nil),   // 17: roles_proto.GetPermissionsResponse
}
var file_protos_roles_proto_depIdxs = []int32{
	1,  // 0: roles_proto.CreateResponse.role:type_name -> roles_proto.Role
	1,  // 1: roles_proto.GetResponse.role:type_name -> roles_proto.Role
	1,  // 2: roles_proto.ListResponse.rows:type_name -> roles_proto.Role
	1,  // 3: roles_proto.UpdateResponse.role:type_name -> roles_proto.Role
	1,  // 4: roles_proto.AddPermissionResponse.role:type_name -> roles_proto.Role
	1,  // 5: roles_proto.RemovePermissionResponse.role:type_name -> roles_proto.Role
	0,  // 6: roles_proto.GetPermissionsResponse.rows:type_name -> roles_proto.Permission
	2,  // 7: roles_proto.RolesService.Create:input_type -> roles_proto.CreateRequest
	4,  // 8: roles_proto.RolesService.Get:input_type -> roles_proto.GetRequest
	6,  // 9: roles_proto.RolesService.List:input_type -> roles_proto.ListRequest
	8,  // 10: roles_proto.RolesService.Update:input_type -> roles_proto.UpdateRequest
	10, // 11: roles_proto.RolesService.Delete:input_type -> roles_proto.DeleteRequest
	12, // 12: roles_proto.RolesService.AddPermission:input_type -> roles_proto.AddPermissionRequest
	14, // 13: roles_proto.RolesService.RemovePermission:input_type -> roles_proto.RemovePermissionRequest
	16, // 14: roles_proto.RolesService.GetPermissions:input_type -> roles_proto.GetPermissionsRequest
	3,  // 15: roles_proto.RolesService.Create:output_type -> roles_proto.CreateResponse
	5,  // 16: roles_proto.RolesService.Get:output_type -> roles_proto.GetResponse
	7,  // 17: roles_proto.RolesService.List:output_type -> roles_proto.ListResponse
	9,  // 18: roles_proto.RolesService.Update:output_type -> roles_proto.UpdateResponse
	11, // 19: roles_proto.RolesService.Delete:output_type -> roles_proto.DeleteResponse
	13, // 20: roles_proto.RolesService.AddPermission:output_type -> roles_proto.AddPermissionResponse
	15, // 21: roles_proto.RolesService.RemovePermission:output_type -> roles_proto.RemovePermissionResponse
	17, // 22: roles_proto.RolesService.GetPermissions:output_type -> roles_proto.GetPermissionsResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_roles_proto_init() }
func file_protos_roles_proto_init() {
	if File_protos_roles_proto != nil {
		return
	}
	file_protos_roles_proto_msgTypes[0].OneofWrappers = []any{}
	file_protos_roles_proto_msgTypes[1].OneofWrappers = []any{}
	file_protos_roles_proto_msgTypes[6].OneofWrappers = []any{}
	file_protos_roles_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_roles_proto_rawDesc), len(file_protos_roles_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_roles_proto_goTypes,
		DependencyIndexes: file_protos_roles_proto_depIdxs,
		MessageInfos:      file_protos_roles_proto_msgTypes,
	}.Build()
	File_protos_roles_proto = out.File
	file_protos_roles_proto_goTypes = nil
	file_protos_roles_proto_depIdxs = nil
}
