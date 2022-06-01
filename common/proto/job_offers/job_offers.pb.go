// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: job_offers/job_offers.proto

package job_offers

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{0}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobOffers []*JobOffer `protobuf:"bytes,1,rep,name=jobOffers,proto3" json:"jobOffers,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetJobOffers() []*JobOffer {
	if x != nil {
		return x.JobOffers
	}
	return nil
}

type JobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position     string  `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Requirements string  `protobuf:"bytes,4,opt,name=requirements,proto3" json:"requirements,omitempty"`
	IsActive     bool    `protobuf:"varint,5,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Followers    []*User `protobuf:"bytes,6,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *JobOffer) Reset() {
	*x = JobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffer) ProtoMessage() {}

func (x *JobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffer.ProtoReflect.Descriptor instead.
func (*JobOffer) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{2}
}

func (x *JobOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *JobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobOffer) GetRequirements() string {
	if x != nil {
		return x.Requirements
	}
	return ""
}

func (x *JobOffer) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *JobOffer) GetFollowers() []*User {
	if x != nil {
		return x.Followers
	}
	return nil
}

type NewJobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position     string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Requirements string `protobuf:"bytes,3,opt,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *NewJobOffer) Reset() {
	*x = NewJobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobOffer) ProtoMessage() {}

func (x *NewJobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJobOffer.ProtoReflect.Descriptor instead.
func (*NewJobOffer) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{3}
}

func (x *NewJobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *NewJobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewJobOffer) GetRequirements() string {
	if x != nil {
		return x.Requirements
	}
	return ""
}

type UpdateJobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position     string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Requirements string `protobuf:"bytes,4,opt,name=requirements,proto3" json:"requirements,omitempty"`
	IsActive     bool   `protobuf:"varint,5,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
}

func (x *UpdateJobOffer) Reset() {
	*x = UpdateJobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobOffer) ProtoMessage() {}

func (x *UpdateJobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobOffer.ProtoReflect.Descriptor instead.
func (*UpdateJobOffer) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJobOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *UpdateJobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateJobOffer) GetRequirements() string {
	if x != nil {
		return x.Requirements
	}
	return ""
}

func (x *UpdateJobOffer) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type JobOfferId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobOfferId) Reset() {
	*x = JobOfferId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOfferId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOfferId) ProtoMessage() {}

func (x *JobOfferId) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOfferId.ProtoReflect.Descriptor instead.
func (*JobOfferId) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{5}
}

func (x *JobOfferId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobOfferId string `protobuf:"bytes,2,opt,name=jobOfferId,proto3" json:"jobOfferId,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{6}
}

func (x *FollowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FollowRequest) GetJobOfferId() string {
	if x != nil {
		return x.JobOfferId
	}
	return ""
}

type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobOfferId string `protobuf:"bytes,2,opt,name=jobOfferId,proto3" json:"jobOfferId,omitempty"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{7}
}

func (x *UnfollowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnfollowRequest) GetJobOfferId() string {
	if x != nil {
		return x.JobOfferId
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[8]
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
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_job_offers_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_job_offers_proto_msgTypes[9]
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
	return file_job_offers_job_offers_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_job_offers_job_offers_proto protoreflect.FileDescriptor

var file_job_offers_job_offers_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x6a, 0x6f, 0x62,
	0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6a, 0x6f,
	0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0xc8,
	0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x6f, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f,
	0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0f, 0x55, 0x6e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xa6, 0x04, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f,
	0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x6a, 0x6f, 0x62,
	0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x1a, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x1a, 0x0b,
	0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x60,
	0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f, 0x6a, 0x6f, 0x62, 0x5f,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x01, 0x2a,
	0x12, 0x66, 0x0a, 0x10, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a,
	0x14, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x6e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x01, 0x2a, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x76, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x6a, 0x69, 0x63, 0x31, 0x39, 0x2f,
	0x58, 0x57, 0x53, 0x2d, 0x54, 0x49, 0x4d, 0x31, 0x35, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_offers_job_offers_proto_rawDescOnce sync.Once
	file_job_offers_job_offers_proto_rawDescData = file_job_offers_job_offers_proto_rawDesc
)

func file_job_offers_job_offers_proto_rawDescGZIP() []byte {
	file_job_offers_job_offers_proto_rawDescOnce.Do(func() {
		file_job_offers_job_offers_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_offers_job_offers_proto_rawDescData)
	})
	return file_job_offers_job_offers_proto_rawDescData
}

var file_job_offers_job_offers_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_job_offers_job_offers_proto_goTypes = []interface{}{
	(*GetAllRequest)(nil),   // 0: job_offers.GetAllRequest
	(*GetAllResponse)(nil),  // 1: job_offers.GetAllResponse
	(*JobOffer)(nil),        // 2: job_offers.JobOffer
	(*NewJobOffer)(nil),     // 3: job_offers.NewJobOffer
	(*UpdateJobOffer)(nil),  // 4: job_offers.UpdateJobOffer
	(*JobOfferId)(nil),      // 5: job_offers.JobOfferId
	(*FollowRequest)(nil),   // 6: job_offers.FollowRequest
	(*UnfollowRequest)(nil), // 7: job_offers.UnfollowRequest
	(*Response)(nil),        // 8: job_offers.Response
	(*User)(nil),            // 9: job_offers.User
}
var file_job_offers_job_offers_proto_depIdxs = []int32{
	2, // 0: job_offers.GetAllResponse.jobOffers:type_name -> job_offers.JobOffer
	9, // 1: job_offers.JobOffer.followers:type_name -> job_offers.User
	0, // 2: job_offers.JobOffersService.GetAll:input_type -> job_offers.GetAllRequest
	5, // 3: job_offers.JobOffersService.Get:input_type -> job_offers.JobOfferId
	3, // 4: job_offers.JobOffersService.Create:input_type -> job_offers.NewJobOffer
	4, // 5: job_offers.JobOffersService.Update:input_type -> job_offers.UpdateJobOffer
	6, // 6: job_offers.JobOffersService.FollowJobOffer:input_type -> job_offers.FollowRequest
	7, // 7: job_offers.JobOffersService.UnfollowJobOffer:input_type -> job_offers.UnfollowRequest
	1, // 8: job_offers.JobOffersService.GetAll:output_type -> job_offers.GetAllResponse
	2, // 9: job_offers.JobOffersService.Get:output_type -> job_offers.JobOffer
	8, // 10: job_offers.JobOffersService.Create:output_type -> job_offers.Response
	8, // 11: job_offers.JobOffersService.Update:output_type -> job_offers.Response
	8, // 12: job_offers.JobOffersService.FollowJobOffer:output_type -> job_offers.Response
	8, // 13: job_offers.JobOffersService.UnfollowJobOffer:output_type -> job_offers.Response
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_job_offers_job_offers_proto_init() }
func file_job_offers_job_offers_proto_init() {
	if File_job_offers_job_offers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_offers_job_offers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_job_offers_job_offers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_job_offers_job_offers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffer); i {
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
		file_job_offers_job_offers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewJobOffer); i {
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
		file_job_offers_job_offers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobOffer); i {
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
		file_job_offers_job_offers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOfferId); i {
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
		file_job_offers_job_offers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_job_offers_job_offers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowRequest); i {
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
		file_job_offers_job_offers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_job_offers_job_offers_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_offers_job_offers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_offers_job_offers_proto_goTypes,
		DependencyIndexes: file_job_offers_job_offers_proto_depIdxs,
		MessageInfos:      file_job_offers_job_offers_proto_msgTypes,
	}.Build()
	File_job_offers_job_offers_proto = out.File
	file_job_offers_job_offers_proto_rawDesc = nil
	file_job_offers_job_offers_proto_goTypes = nil
	file_job_offers_job_offers_proto_depIdxs = nil
}
