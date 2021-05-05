// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.5
// source: protos/service/messenger_service.grpc.proto

package service

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

type CreatedByType int32

const (
	CreatedByType_JOBSEEKER CreatedByType = 0
	CreatedByType_EMPLOYER  CreatedByType = 1
)

// Enum value maps for CreatedByType.
var (
	CreatedByType_name = map[int32]string{
		0: "JOBSEEKER",
		1: "EMPLOYER",
	}
	CreatedByType_value = map[string]int32{
		"JOBSEEKER": 0,
		"EMPLOYER":  1,
	}
)

func (x CreatedByType) Enum() *CreatedByType {
	p := new(CreatedByType)
	*p = x
	return p
}

func (x CreatedByType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreatedByType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_service_messenger_service_grpc_proto_enumTypes[0].Descriptor()
}

func (CreatedByType) Type() protoreflect.EnumType {
	return &file_protos_service_messenger_service_grpc_proto_enumTypes[0]
}

func (x CreatedByType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreatedByType.Descriptor instead.
func (CreatedByType) EnumDescriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{0}
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployerId          string        `protobuf:"bytes,1,opt,name=employer_id,json=employerId,proto3" json:"employer_id,omitempty"`
	UserId              string        `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	JobId               string        `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	JobseekerId         string        `protobuf:"bytes,4,opt,name=jobseeker_id,json=jobseekerId,proto3" json:"jobseeker_id,omitempty"`
	Message             string        `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	ApplicationStatusId int32         `protobuf:"varint,6,opt,name=application_status_id,json=applicationStatusId,proto3" json:"application_status_id,omitempty"`
	ApplicationStatus   string        `protobuf:"bytes,7,opt,name=application_status,json=applicationStatus,proto3" json:"application_status,omitempty"`
	CreatedByType       CreatedByType `protobuf:"varint,8,opt,name=created_by_type,json=createdByType,proto3,enum=protos.service.CreatedByType" json:"created_by_type,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRequest) GetEmployerId() string {
	if x != nil {
		return x.EmployerId
	}
	return ""
}

func (x *CreateMessageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateMessageRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *CreateMessageRequest) GetJobseekerId() string {
	if x != nil {
		return x.JobseekerId
	}
	return ""
}

func (x *CreateMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateMessageRequest) GetApplicationStatusId() int32 {
	if x != nil {
		return x.ApplicationStatusId
	}
	return 0
}

func (x *CreateMessageRequest) GetApplicationStatus() string {
	if x != nil {
		return x.ApplicationStatus
	}
	return ""
}

func (x *CreateMessageRequest) GetCreatedByType() CreatedByType {
	if x != nil {
		return x.CreatedByType
	}
	return CreatedByType_JOBSEEKER
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	RoomId            string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	EmployerObjectId  string `protobuf:"bytes,3,opt,name=employer_object_id,json=employerObjectId,proto3" json:"employer_object_id,omitempty"`
	JobseekerObjectId string `protobuf:"bytes,4,opt,name=jobseeker_object_id,json=jobseekerObjectId,proto3" json:"jobseeker_object_id,omitempty"`
	MessageObjectId   string `protobuf:"bytes,5,opt,name=message_object_id,json=messageObjectId,proto3" json:"message_object_id,omitempty"`
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateMessageResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateMessageResponse) GetEmployerObjectId() string {
	if x != nil {
		return x.EmployerObjectId
	}
	return ""
}

func (x *CreateMessageResponse) GetJobseekerObjectId() string {
	if x != nil {
		return x.JobseekerObjectId
	}
	return ""
}

func (x *CreateMessageResponse) GetMessageObjectId() string {
	if x != nil {
		return x.MessageObjectId
	}
	return ""
}

type AddMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId               string        `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	UserId              string        `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message             string        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	ApplicationStatusId int32         `protobuf:"varint,4,opt,name=application_status_id,json=applicationStatusId,proto3" json:"application_status_id,omitempty"`
	ApplicationStatus   string        `protobuf:"bytes,5,opt,name=application_status,json=applicationStatus,proto3" json:"application_status,omitempty"`
	CreatedBy           string        `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedByType       CreatedByType `protobuf:"varint,7,opt,name=createdByType,proto3,enum=protos.service.CreatedByType" json:"createdByType,omitempty"`
}

func (x *AddMessageRequest) Reset() {
	*x = AddMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageRequest) ProtoMessage() {}

func (x *AddMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageRequest.ProtoReflect.Descriptor instead.
func (*AddMessageRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *AddMessageRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *AddMessageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddMessageRequest) GetApplicationStatusId() int32 {
	if x != nil {
		return x.ApplicationStatusId
	}
	return 0
}

func (x *AddMessageRequest) GetApplicationStatus() string {
	if x != nil {
		return x.ApplicationStatus
	}
	return ""
}

func (x *AddMessageRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *AddMessageRequest) GetCreatedByType() CreatedByType {
	if x != nil {
		return x.CreatedByType
	}
	return CreatedByType_JOBSEEKER
}

type AddMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *AddMessageResponse) Reset() {
	*x = AddMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageResponse) ProtoMessage() {}

func (x *AddMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageResponse.ProtoReflect.Descriptor instead.
func (*AddMessageResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *AddMessageResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddMessageResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type AddMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddMessageRequests []*AddMessageRequest `protobuf:"bytes,1,rep,name=add_message_requests,json=addMessageRequests,proto3" json:"add_message_requests,omitempty"`
}

func (x *AddMessagesRequest) Reset() {
	*x = AddMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessagesRequest) ProtoMessage() {}

func (x *AddMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessagesRequest.ProtoReflect.Descriptor instead.
func (*AddMessagesRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *AddMessagesRequest) GetAddMessageRequests() []*AddMessageRequest {
	if x != nil {
		return x.AddMessageRequests
	}
	return nil
}

type AddMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	MessageIds []string `protobuf:"bytes,2,rep,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
}

func (x *AddMessagesResponse) Reset() {
	*x = AddMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessagesResponse) ProtoMessage() {}

func (x *AddMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_messenger_service_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessagesResponse.ProtoReflect.Descriptor instead.
func (*AddMessagesResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_messenger_service_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *AddMessagesResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddMessagesResponse) GetMessageIds() []string {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

var File_protos_service_messenger_service_grpc_proto protoreflect.FileDescriptor

var file_protos_service_messenger_service_grpc_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xce, 0x02,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x62, 0x73, 0x65,
	0x65, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a,
	0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd2,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6a, 0x6f, 0x62, 0x73, 0x65,
	0x65, 0x6b, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0xa4, 0x02, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a,
	0x14, 0x61, 0x64, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x12,
	0x61, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x4e, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x73, 0x2a, 0x2c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x4f, 0x42, 0x53, 0x45, 0x45, 0x4b, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4d, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x52, 0x10, 0x01,
	0x32, 0x9d, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x61, 0x6e, 0x70, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_service_messenger_service_grpc_proto_rawDescOnce sync.Once
	file_protos_service_messenger_service_grpc_proto_rawDescData = file_protos_service_messenger_service_grpc_proto_rawDesc
)

func file_protos_service_messenger_service_grpc_proto_rawDescGZIP() []byte {
	file_protos_service_messenger_service_grpc_proto_rawDescOnce.Do(func() {
		file_protos_service_messenger_service_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_service_messenger_service_grpc_proto_rawDescData)
	})
	return file_protos_service_messenger_service_grpc_proto_rawDescData
}

var file_protos_service_messenger_service_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_service_messenger_service_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_service_messenger_service_grpc_proto_goTypes = []interface{}{
	(CreatedByType)(0),            // 0: protos.service.CreatedByType
	(*CreateMessageRequest)(nil),  // 1: protos.service.CreateMessageRequest
	(*CreateMessageResponse)(nil), // 2: protos.service.CreateMessageResponse
	(*AddMessageRequest)(nil),     // 3: protos.service.AddMessageRequest
	(*AddMessageResponse)(nil),    // 4: protos.service.AddMessageResponse
	(*AddMessagesRequest)(nil),    // 5: protos.service.AddMessagesRequest
	(*AddMessagesResponse)(nil),   // 6: protos.service.AddMessagesResponse
}
var file_protos_service_messenger_service_grpc_proto_depIdxs = []int32{
	0, // 0: protos.service.CreateMessageRequest.created_by_type:type_name -> protos.service.CreatedByType
	0, // 1: protos.service.AddMessageRequest.createdByType:type_name -> protos.service.CreatedByType
	3, // 2: protos.service.AddMessagesRequest.add_message_requests:type_name -> protos.service.AddMessageRequest
	1, // 3: protos.service.MessengerService.CreateMessage:input_type -> protos.service.CreateMessageRequest
	3, // 4: protos.service.MessengerService.AddMessage:input_type -> protos.service.AddMessageRequest
	5, // 5: protos.service.MessengerService.AddMessages:input_type -> protos.service.AddMessagesRequest
	2, // 6: protos.service.MessengerService.CreateMessage:output_type -> protos.service.CreateMessageResponse
	4, // 7: protos.service.MessengerService.AddMessage:output_type -> protos.service.AddMessageResponse
	6, // 8: protos.service.MessengerService.AddMessages:output_type -> protos.service.AddMessagesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_service_messenger_service_grpc_proto_init() }
func file_protos_service_messenger_service_grpc_proto_init() {
	if File_protos_service_messenger_service_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_service_messenger_service_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
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
		file_protos_service_messenger_service_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageResponse); i {
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
		file_protos_service_messenger_service_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessageRequest); i {
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
		file_protos_service_messenger_service_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessageResponse); i {
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
		file_protos_service_messenger_service_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessagesRequest); i {
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
		file_protos_service_messenger_service_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMessagesResponse); i {
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
			RawDescriptor: file_protos_service_messenger_service_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_messenger_service_grpc_proto_goTypes,
		DependencyIndexes: file_protos_service_messenger_service_grpc_proto_depIdxs,
		EnumInfos:         file_protos_service_messenger_service_grpc_proto_enumTypes,
		MessageInfos:      file_protos_service_messenger_service_grpc_proto_msgTypes,
	}.Build()
	File_protos_service_messenger_service_grpc_proto = out.File
	file_protos_service_messenger_service_grpc_proto_rawDesc = nil
	file_protos_service_messenger_service_grpc_proto_goTypes = nil
	file_protos_service_messenger_service_grpc_proto_depIdxs = nil
}
