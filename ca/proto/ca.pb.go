// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: ca.proto

package proto

import (
	proto "github.com/letsencrypt/boulder/core/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IssueCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Csr            []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	RegistrationID int64  `protobuf:"varint,2,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	OrderID        int64  `protobuf:"varint,3,opt,name=orderID,proto3" json:"orderID,omitempty"`
	IssuerNameID   int64  `protobuf:"varint,4,opt,name=issuerNameID,proto3" json:"issuerNameID,omitempty"`
	TypeIdentifier string `protobuf:"bytes,5,opt,name=typeIdentifier,proto3" json:"typeIdentifier,omitempty"`
}

func (x *IssueCertificateRequest) Reset() {
	*x = IssueCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCertificateRequest) ProtoMessage() {}

func (x *IssueCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCertificateRequest.ProtoReflect.Descriptor instead.
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{0}
}

func (x *IssueCertificateRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

func (x *IssueCertificateRequest) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *IssueCertificateRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *IssueCertificateRequest) GetIssuerNameID() int64 {
	if x != nil {
		return x.IssuerNameID
	}
	return 0
}

func (x *IssueCertificateRequest) GetTypeIdentifier() string {
	if x != nil {
		return x.TypeIdentifier
	}
	return ""
}

type IssuePrecertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DER []byte `protobuf:"bytes,1,opt,name=DER,proto3" json:"DER,omitempty"`
}

func (x *IssuePrecertificateResponse) Reset() {
	*x = IssuePrecertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuePrecertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuePrecertificateResponse) ProtoMessage() {}

func (x *IssuePrecertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuePrecertificateResponse.ProtoReflect.Descriptor instead.
func (*IssuePrecertificateResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{1}
}

func (x *IssuePrecertificateResponse) GetDER() []byte {
	if x != nil {
		return x.DER
	}
	return nil
}

type IssueCertificateForPrecertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DER            []byte   `protobuf:"bytes,1,opt,name=DER,proto3" json:"DER,omitempty"`
	SCTs           [][]byte `protobuf:"bytes,2,rep,name=SCTs,proto3" json:"SCTs,omitempty"`
	RegistrationID int64    `protobuf:"varint,3,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	OrderID        int64    `protobuf:"varint,4,opt,name=orderID,proto3" json:"orderID,omitempty"`
	TypeIdentifier string   `protobuf:"bytes,5,opt,name=typeIdentifier,proto3" json:"typeIdentifier,omitempty"`
}

func (x *IssueCertificateForPrecertificateRequest) Reset() {
	*x = IssueCertificateForPrecertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCertificateForPrecertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCertificateForPrecertificateRequest) ProtoMessage() {}

func (x *IssueCertificateForPrecertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCertificateForPrecertificateRequest.ProtoReflect.Descriptor instead.
func (*IssueCertificateForPrecertificateRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{2}
}

func (x *IssueCertificateForPrecertificateRequest) GetDER() []byte {
	if x != nil {
		return x.DER
	}
	return nil
}

func (x *IssueCertificateForPrecertificateRequest) GetSCTs() [][]byte {
	if x != nil {
		return x.SCTs
	}
	return nil
}

func (x *IssueCertificateForPrecertificateRequest) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *IssueCertificateForPrecertificateRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *IssueCertificateForPrecertificateRequest) GetTypeIdentifier() string {
	if x != nil {
		return x.TypeIdentifier
	}
	return ""
}

// Exactly one of certDER or [serial and issuerID] must be set.
type GenerateOCSPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Next unused field number: 8
	Status      string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Reason      int32                  `protobuf:"varint,3,opt,name=reason,proto3" json:"reason,omitempty"`
	RevokedAtNS int64                  `protobuf:"varint,4,opt,name=revokedAtNS,proto3" json:"revokedAtNS,omitempty"` // Unix timestamp (nanoseconds)
	RevokedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=revokedAt,proto3" json:"revokedAt,omitempty"`
	Serial      string                 `protobuf:"bytes,5,opt,name=serial,proto3" json:"serial,omitempty"`
	IssuerID    int64                  `protobuf:"varint,6,opt,name=issuerID,proto3" json:"issuerID,omitempty"`
}

func (x *GenerateOCSPRequest) Reset() {
	*x = GenerateOCSPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOCSPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOCSPRequest) ProtoMessage() {}

func (x *GenerateOCSPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOCSPRequest.ProtoReflect.Descriptor instead.
func (*GenerateOCSPRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateOCSPRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GenerateOCSPRequest) GetReason() int32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

func (x *GenerateOCSPRequest) GetRevokedAtNS() int64 {
	if x != nil {
		return x.RevokedAtNS
	}
	return 0
}

func (x *GenerateOCSPRequest) GetRevokedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RevokedAt
	}
	return nil
}

func (x *GenerateOCSPRequest) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *GenerateOCSPRequest) GetIssuerID() int64 {
	if x != nil {
		return x.IssuerID
	}
	return 0
}

type OCSPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OCSPResponse) Reset() {
	*x = OCSPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCSPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCSPResponse) ProtoMessage() {}

func (x *OCSPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCSPResponse.ProtoReflect.Descriptor instead.
func (*OCSPResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{4}
}

func (x *OCSPResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type GenerateCRLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*GenerateCRLRequest_Metadata
	//	*GenerateCRLRequest_Entry
	Payload isGenerateCRLRequest_Payload `protobuf_oneof:"payload"`
}

func (x *GenerateCRLRequest) Reset() {
	*x = GenerateCRLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCRLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCRLRequest) ProtoMessage() {}

func (x *GenerateCRLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCRLRequest.ProtoReflect.Descriptor instead.
func (*GenerateCRLRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{5}
}

func (m *GenerateCRLRequest) GetPayload() isGenerateCRLRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GenerateCRLRequest) GetMetadata() *CRLMetadata {
	if x, ok := x.GetPayload().(*GenerateCRLRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *GenerateCRLRequest) GetEntry() *proto.CRLEntry {
	if x, ok := x.GetPayload().(*GenerateCRLRequest_Entry); ok {
		return x.Entry
	}
	return nil
}

type isGenerateCRLRequest_Payload interface {
	isGenerateCRLRequest_Payload()
}

type GenerateCRLRequest_Metadata struct {
	Metadata *CRLMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type GenerateCRLRequest_Entry struct {
	Entry *proto.CRLEntry `protobuf:"bytes,2,opt,name=entry,proto3,oneof"`
}

func (*GenerateCRLRequest_Metadata) isGenerateCRLRequest_Payload() {}

func (*GenerateCRLRequest_Entry) isGenerateCRLRequest_Payload() {}

type CRLMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Next unused field number: 5
	IssuerNameID int64                  `protobuf:"varint,1,opt,name=issuerNameID,proto3" json:"issuerNameID,omitempty"`
	ThisUpdateNS int64                  `protobuf:"varint,2,opt,name=thisUpdateNS,proto3" json:"thisUpdateNS,omitempty"` // Unix timestamp (nanoseconds), also used for CRLNumber.
	ThisUpdate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=thisUpdate,proto3" json:"thisUpdate,omitempty"`
	ShardIdx     int64                  `protobuf:"varint,3,opt,name=shardIdx,proto3" json:"shardIdx,omitempty"`
}

func (x *CRLMetadata) Reset() {
	*x = CRLMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRLMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRLMetadata) ProtoMessage() {}

func (x *CRLMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRLMetadata.ProtoReflect.Descriptor instead.
func (*CRLMetadata) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{6}
}

func (x *CRLMetadata) GetIssuerNameID() int64 {
	if x != nil {
		return x.IssuerNameID
	}
	return 0
}

func (x *CRLMetadata) GetThisUpdateNS() int64 {
	if x != nil {
		return x.ThisUpdateNS
	}
	return 0
}

func (x *CRLMetadata) GetThisUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.ThisUpdate
	}
	return nil
}

func (x *CRLMetadata) GetShardIdx() int64 {
	if x != nil {
		return x.ShardIdx
	}
	return 0
}

type GenerateCRLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *GenerateCRLResponse) Reset() {
	*x = GenerateCRLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCRLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCRLResponse) ProtoMessage() {}

func (x *GenerateCRLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCRLResponse.ProtoReflect.Descriptor instead.
func (*GenerateCRLResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateCRLResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_ca_proto protoreflect.FileDescriptor

var file_ca_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x61, 0x1a, 0x15,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x63, 0x73, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x1b, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x50, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x45, 0x52,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x44, 0x45, 0x52, 0x22, 0x92, 0x01, 0x0a, 0x28,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x50, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x45, 0x52, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x44, 0x45, 0x52, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x43,
	0x54, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x43, 0x54, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x22, 0xd5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x43, 0x53,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x64, 0x41, 0x74, 0x4e, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x4e, 0x53, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x0c, 0x4f, 0x43, 0x53, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x43, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x61, 0x2e, 0x43, 0x52, 0x4c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x52, 0x4c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xad, 0x01, 0x0a,
	0x0b, 0x43, 0x52, 0x4c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x53,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x53, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x78, 0x22, 0x2b, 0x0a, 0x13,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xd5, 0x01, 0x0a, 0x14, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x55, 0x0a, 0x13, 0x49, 0x73, 0x73, 0x75, 0x65, 0x50, 0x72, 0x65, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x50, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x21, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72,
	0x50, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x63, 0x61, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22,
	0x00, 0x32, 0x4c, 0x0a, 0x0d, 0x4f, 0x43, 0x53, 0x50, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x43,
	0x53, 0x50, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x4f, 0x43, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x61,
	0x2e, 0x4f, 0x43, 0x53, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0x54, 0x0a, 0x0c, 0x43, 0x52, 0x4c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x44, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x52, 0x4c, 0x12, 0x16,
	0x2e, 0x63, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x74, 0x73, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f,
	0x62, 0x6f, 0x75, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ca_proto_rawDescOnce sync.Once
	file_ca_proto_rawDescData = file_ca_proto_rawDesc
)

func file_ca_proto_rawDescGZIP() []byte {
	file_ca_proto_rawDescOnce.Do(func() {
		file_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_ca_proto_rawDescData)
	})
	return file_ca_proto_rawDescData
}

var file_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ca_proto_goTypes = []interface{}{
	(*IssueCertificateRequest)(nil),                  // 0: ca.IssueCertificateRequest
	(*IssuePrecertificateResponse)(nil),              // 1: ca.IssuePrecertificateResponse
	(*IssueCertificateForPrecertificateRequest)(nil), // 2: ca.IssueCertificateForPrecertificateRequest
	(*GenerateOCSPRequest)(nil),                      // 3: ca.GenerateOCSPRequest
	(*OCSPResponse)(nil),                             // 4: ca.OCSPResponse
	(*GenerateCRLRequest)(nil),                       // 5: ca.GenerateCRLRequest
	(*CRLMetadata)(nil),                              // 6: ca.CRLMetadata
	(*GenerateCRLResponse)(nil),                      // 7: ca.GenerateCRLResponse
	(*timestamppb.Timestamp)(nil),                    // 8: google.protobuf.Timestamp
	(*proto.CRLEntry)(nil),                           // 9: core.CRLEntry
	(*proto.Certificate)(nil),                        // 10: core.Certificate
}
var file_ca_proto_depIdxs = []int32{
	8,  // 0: ca.GenerateOCSPRequest.revokedAt:type_name -> google.protobuf.Timestamp
	6,  // 1: ca.GenerateCRLRequest.metadata:type_name -> ca.CRLMetadata
	9,  // 2: ca.GenerateCRLRequest.entry:type_name -> core.CRLEntry
	8,  // 3: ca.CRLMetadata.thisUpdate:type_name -> google.protobuf.Timestamp
	0,  // 4: ca.CertificateAuthority.IssuePrecertificate:input_type -> ca.IssueCertificateRequest
	2,  // 5: ca.CertificateAuthority.IssueCertificateForPrecertificate:input_type -> ca.IssueCertificateForPrecertificateRequest
	3,  // 6: ca.OCSPGenerator.GenerateOCSP:input_type -> ca.GenerateOCSPRequest
	5,  // 7: ca.CRLGenerator.GenerateCRL:input_type -> ca.GenerateCRLRequest
	1,  // 8: ca.CertificateAuthority.IssuePrecertificate:output_type -> ca.IssuePrecertificateResponse
	10, // 9: ca.CertificateAuthority.IssueCertificateForPrecertificate:output_type -> core.Certificate
	4,  // 10: ca.OCSPGenerator.GenerateOCSP:output_type -> ca.OCSPResponse
	7,  // 11: ca.CRLGenerator.GenerateCRL:output_type -> ca.GenerateCRLResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_ca_proto_init() }
func file_ca_proto_init() {
	if File_ca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCertificateRequest); i {
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
		file_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuePrecertificateResponse); i {
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
		file_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCertificateForPrecertificateRequest); i {
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
		file_ca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOCSPRequest); i {
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
		file_ca_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCSPResponse); i {
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
		file_ca_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCRLRequest); i {
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
		file_ca_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CRLMetadata); i {
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
		file_ca_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCRLResponse); i {
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
	file_ca_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GenerateCRLRequest_Metadata)(nil),
		(*GenerateCRLRequest_Entry)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_ca_proto_goTypes,
		DependencyIndexes: file_ca_proto_depIdxs,
		MessageInfos:      file_ca_proto_msgTypes,
	}.Build()
	File_ca_proto = out.File
	file_ca_proto_rawDesc = nil
	file_ca_proto_goTypes = nil
	file_ca_proto_depIdxs = nil
}
