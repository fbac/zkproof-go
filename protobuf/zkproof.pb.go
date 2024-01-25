// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: protobuf/zkproof.proto

package protobuf

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Y1   int64  `protobuf:"varint,2,opt,name=y1,proto3" json:"y1,omitempty"`
	Y2   int64  `protobuf:"varint,3,opt,name=y2,proto3" json:"y2,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RegisterRequest) GetY1() int64 {
	if x != nil {
		return x.Y1
	}
	return 0
}

func (x *RegisterRequest) GetY2() int64 {
	if x != nil {
		return x.Y2
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{1}
}

type AuthenticationChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	R1   int64  `protobuf:"varint,2,opt,name=r1,proto3" json:"r1,omitempty"`
	R2   int64  `protobuf:"varint,3,opt,name=r2,proto3" json:"r2,omitempty"`
}

func (x *AuthenticationChallengeRequest) Reset() {
	*x = AuthenticationChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationChallengeRequest) ProtoMessage() {}

func (x *AuthenticationChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationChallengeRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationChallengeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationChallengeRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthenticationChallengeRequest) GetR1() int64 {
	if x != nil {
		return x.R1
	}
	return 0
}

func (x *AuthenticationChallengeRequest) GetR2() int64 {
	if x != nil {
		return x.R2
	}
	return 0
}

type AuthenticationChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	C      int64  `protobuf:"varint,2,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *AuthenticationChallengeResponse) Reset() {
	*x = AuthenticationChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationChallengeResponse) ProtoMessage() {}

func (x *AuthenticationChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationChallengeResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationChallengeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationChallengeResponse) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *AuthenticationChallengeResponse) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

type AuthenticationAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	S      int64  `protobuf:"varint,2,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *AuthenticationAnswerRequest) Reset() {
	*x = AuthenticationAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationAnswerRequest) ProtoMessage() {}

func (x *AuthenticationAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationAnswerRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationAnswerRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{4}
}

func (x *AuthenticationAnswerRequest) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *AuthenticationAnswerRequest) GetS() int64 {
	if x != nil {
		return x.S
	}
	return 0
}

type AuthenticationAnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *AuthenticationAnswerResponse) Reset() {
	*x = AuthenticationAnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_zkproof_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationAnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationAnswerResponse) ProtoMessage() {}

func (x *AuthenticationAnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_zkproof_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationAnswerResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationAnswerResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_zkproof_proto_rawDescGZIP(), []int{5}
}

func (x *AuthenticationAnswerResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_protobuf_zkproof_proto protoreflect.FileDescriptor

var file_protobuf_zkproof_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x7a, 0x6b, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x79, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x79, 0x32, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x32, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a,
	0x1e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x72, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x72, 0x32, 0x22, 0x48, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12,
	0x0c, 0x0a, 0x01, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x63, 0x22, 0x44, 0x0a,
	0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x73, 0x22, 0x3d, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x32, 0xac, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x76, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x12, 0x28, 0x2e, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x7a, 0x6b,
	0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x7a, 0x6b, 0x70, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x62, 0x61, 0x63, 0x2f, 0x7a, 0x6b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobuf_zkproof_proto_rawDescOnce sync.Once
	file_protobuf_zkproof_proto_rawDescData = file_protobuf_zkproof_proto_rawDesc
)

func file_protobuf_zkproof_proto_rawDescGZIP() []byte {
	file_protobuf_zkproof_proto_rawDescOnce.Do(func() {
		file_protobuf_zkproof_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_zkproof_proto_rawDescData)
	})
	return file_protobuf_zkproof_proto_rawDescData
}

var file_protobuf_zkproof_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_zkproof_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),                 // 0: zkp_auth.RegisterRequest
	(*RegisterResponse)(nil),                // 1: zkp_auth.RegisterResponse
	(*AuthenticationChallengeRequest)(nil),  // 2: zkp_auth.AuthenticationChallengeRequest
	(*AuthenticationChallengeResponse)(nil), // 3: zkp_auth.AuthenticationChallengeResponse
	(*AuthenticationAnswerRequest)(nil),     // 4: zkp_auth.AuthenticationAnswerRequest
	(*AuthenticationAnswerResponse)(nil),    // 5: zkp_auth.AuthenticationAnswerResponse
}
var file_protobuf_zkproof_proto_depIdxs = []int32{
	0, // 0: zkp_auth.Auth.Register:input_type -> zkp_auth.RegisterRequest
	2, // 1: zkp_auth.Auth.CreateAuthenticationChallenge:input_type -> zkp_auth.AuthenticationChallengeRequest
	4, // 2: zkp_auth.Auth.VerifyAuthentication:input_type -> zkp_auth.AuthenticationAnswerRequest
	1, // 3: zkp_auth.Auth.Register:output_type -> zkp_auth.RegisterResponse
	3, // 4: zkp_auth.Auth.CreateAuthenticationChallenge:output_type -> zkp_auth.AuthenticationChallengeResponse
	5, // 5: zkp_auth.Auth.VerifyAuthentication:output_type -> zkp_auth.AuthenticationAnswerResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_zkproof_proto_init() }
func file_protobuf_zkproof_proto_init() {
	if File_protobuf_zkproof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_zkproof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_protobuf_zkproof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_protobuf_zkproof_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationChallengeRequest); i {
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
		file_protobuf_zkproof_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationChallengeResponse); i {
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
		file_protobuf_zkproof_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationAnswerRequest); i {
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
		file_protobuf_zkproof_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationAnswerResponse); i {
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
			RawDescriptor: file_protobuf_zkproof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_zkproof_proto_goTypes,
		DependencyIndexes: file_protobuf_zkproof_proto_depIdxs,
		MessageInfos:      file_protobuf_zkproof_proto_msgTypes,
	}.Build()
	File_protobuf_zkproof_proto = out.File
	file_protobuf_zkproof_proto_rawDesc = nil
	file_protobuf_zkproof_proto_goTypes = nil
	file_protobuf_zkproof_proto_depIdxs = nil
}
