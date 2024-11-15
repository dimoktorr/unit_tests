// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: api/v1/example_service.proto

package v1

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Examples []*Example `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetExamples() []*Example {
	if x != nil {
		return x.Examples
	}
	return nil
}

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{2}
}

func (x *Example) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Example) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Example) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// v1.0.0
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// v1.1.0
type User1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*User1_Login
	//	*User1_Email
	Identifier isUser1_Identifier `protobuf_oneof:"identifier"`
	Password   string             `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User1) Reset() {
	*x = User1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User1) ProtoMessage() {}

func (x *User1) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User1.ProtoReflect.Descriptor instead.
func (*User1) Descriptor() ([]byte, []int) {
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{4}
}

func (m *User1) GetIdentifier() isUser1_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *User1) GetLogin() string {
	if x, ok := x.GetIdentifier().(*User1_Login); ok {
		return x.Login
	}
	return ""
}

func (x *User1) GetEmail() string {
	if x, ok := x.GetIdentifier().(*User1_Email); ok {
		return x.Email
	}
	return ""
}

func (x *User1) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type isUser1_Identifier interface {
	isUser1_Identifier()
}

type User1_Login struct {
	Login string `protobuf:"bytes,1,opt,name=login,proto3,oneof"`
}

type User1_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

func (*User1_Login) isUser1_Identifier() {}

func (*User1_Email) isUser1_Identifier() {}

// v1.1.0
type User2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Info     *Info  `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *User2) Reset() {
	*x = User2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User2) ProtoMessage() {}

func (x *User2) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User2.ProtoReflect.Descriptor instead.
func (*User2) Descriptor() ([]byte, []int) {
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{5}
}

func (x *User2) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User2) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User2) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_example_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_example_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_api_v1_example_service_proto_rawDescGZIP(), []int{6}
}

func (x *Info) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Info) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_api_v1_example_service_proto protoreflect.FileDescriptor

var file_api_v1_example_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x61, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x31, 0x12, 0x16, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x22, 0x68, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0x63,
	0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x6d, 0x6f, 0x6b, 0x74, 0x6f, 0x72, 0x72, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_example_service_proto_rawDescOnce sync.Once
	file_api_v1_example_service_proto_rawDescData = file_api_v1_example_service_proto_rawDesc
)

func file_api_v1_example_service_proto_rawDescGZIP() []byte {
	file_api_v1_example_service_proto_rawDescOnce.Do(func() {
		file_api_v1_example_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_example_service_proto_rawDescData)
	})
	return file_api_v1_example_service_proto_rawDescData
}

var file_api_v1_example_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_example_service_proto_goTypes = []any{
	(*GetRequest)(nil),  // 0: example.grpc.api.v1.GetRequest
	(*GetResponse)(nil), // 1: example.grpc.api.v1.GetResponse
	(*Example)(nil),     // 2: example.grpc.api.v1.Example
	(*User)(nil),        // 3: example.grpc.api.v1.User
	(*User1)(nil),       // 4: example.grpc.api.v1.User1
	(*User2)(nil),       // 5: example.grpc.api.v1.User2
	(*Info)(nil),        // 6: example.grpc.api.v1.Info
}
var file_api_v1_example_service_proto_depIdxs = []int32{
	2, // 0: example.grpc.api.v1.GetResponse.examples:type_name -> example.grpc.api.v1.Example
	6, // 1: example.grpc.api.v1.User2.info:type_name -> example.grpc.api.v1.Info
	0, // 2: example.grpc.api.v1.ExampleService.GetExample:input_type -> example.grpc.api.v1.GetRequest
	1, // 3: example.grpc.api.v1.ExampleService.GetExample:output_type -> example.grpc.api.v1.GetResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_example_service_proto_init() }
func file_api_v1_example_service_proto_init() {
	if File_api_v1_example_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_example_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_api_v1_example_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_api_v1_example_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Example); i {
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
		file_api_v1_example_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_api_v1_example_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*User1); i {
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
		file_api_v1_example_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*User2); i {
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
		file_api_v1_example_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Info); i {
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
	file_api_v1_example_service_proto_msgTypes[4].OneofWrappers = []any{
		(*User1_Login)(nil),
		(*User1_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_example_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_example_service_proto_goTypes,
		DependencyIndexes: file_api_v1_example_service_proto_depIdxs,
		MessageInfos:      file_api_v1_example_service_proto_msgTypes,
	}.Build()
	File_api_v1_example_service_proto = out.File
	file_api_v1_example_service_proto_rawDesc = nil
	file_api_v1_example_service_proto_goTypes = nil
	file_api_v1_example_service_proto_depIdxs = nil
}
