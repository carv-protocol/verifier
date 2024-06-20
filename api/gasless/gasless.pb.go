// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: api/gasless/gasless.proto

package gasless

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

type ExplorerSendTxNodeEnterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer       string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	ReplacedNode string `protobuf:"bytes,2,opt,name=replaced_node,json=replacedNode,proto3" json:"replaced_node,omitempty"`
	ExpiredAt    uint64 `protobuf:"varint,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	V            uint32 `protobuf:"varint,4,opt,name=v,proto3" json:"v,omitempty"`
	R            string `protobuf:"bytes,5,opt,name=r,proto3" json:"r,omitempty"`
	S            string `protobuf:"bytes,6,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ExplorerSendTxNodeEnterRequest) Reset() {
	*x = ExplorerSendTxNodeEnterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerSendTxNodeEnterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerSendTxNodeEnterRequest) ProtoMessage() {}

func (x *ExplorerSendTxNodeEnterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerSendTxNodeEnterRequest.ProtoReflect.Descriptor instead.
func (*ExplorerSendTxNodeEnterRequest) Descriptor() ([]byte, []int) {
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{0}
}

func (x *ExplorerSendTxNodeEnterRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *ExplorerSendTxNodeEnterRequest) GetReplacedNode() string {
	if x != nil {
		return x.ReplacedNode
	}
	return ""
}

func (x *ExplorerSendTxNodeEnterRequest) GetExpiredAt() uint64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *ExplorerSendTxNodeEnterRequest) GetV() uint32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *ExplorerSendTxNodeEnterRequest) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *ExplorerSendTxNodeEnterRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type ExplorerSendTxNodeExitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer    string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	ExpiredAt uint64 `protobuf:"varint,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	V         uint32 `protobuf:"varint,3,opt,name=v,proto3" json:"v,omitempty"`
	R         string `protobuf:"bytes,4,opt,name=r,proto3" json:"r,omitempty"`
	S         string `protobuf:"bytes,5,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ExplorerSendTxNodeExitRequest) Reset() {
	*x = ExplorerSendTxNodeExitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerSendTxNodeExitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerSendTxNodeExitRequest) ProtoMessage() {}

func (x *ExplorerSendTxNodeExitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerSendTxNodeExitRequest.ProtoReflect.Descriptor instead.
func (*ExplorerSendTxNodeExitRequest) Descriptor() ([]byte, []int) {
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{1}
}

func (x *ExplorerSendTxNodeExitRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *ExplorerSendTxNodeExitRequest) GetExpiredAt() uint64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *ExplorerSendTxNodeExitRequest) GetV() uint32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *ExplorerSendTxNodeExitRequest) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *ExplorerSendTxNodeExitRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type ExplorerSendTxModifyCommissionRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer         string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	CommissionRate uint32 `protobuf:"varint,2,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	ExpiredAt      uint64 `protobuf:"varint,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	V              uint32 `protobuf:"varint,4,opt,name=v,proto3" json:"v,omitempty"`
	R              string `protobuf:"bytes,5,opt,name=r,proto3" json:"r,omitempty"`
	S              string `protobuf:"bytes,6,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ExplorerSendTxModifyCommissionRateRequest) Reset() {
	*x = ExplorerSendTxModifyCommissionRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerSendTxModifyCommissionRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerSendTxModifyCommissionRateRequest) ProtoMessage() {}

func (x *ExplorerSendTxModifyCommissionRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerSendTxModifyCommissionRateRequest.ProtoReflect.Descriptor instead.
func (*ExplorerSendTxModifyCommissionRateRequest) Descriptor() ([]byte, []int) {
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{2}
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetCommissionRate() uint32 {
	if x != nil {
		return x.CommissionRate
	}
	return 0
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetExpiredAt() uint64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetV() uint32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *ExplorerSendTxModifyCommissionRateRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type ExplorerSendTxSetRewardClaimerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer    string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Claimer   string `protobuf:"bytes,2,opt,name=claimer,proto3" json:"claimer,omitempty"`
	ExpiredAt uint64 `protobuf:"varint,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	V         uint32 `protobuf:"varint,4,opt,name=v,proto3" json:"v,omitempty"`
	R         string `protobuf:"bytes,5,opt,name=r,proto3" json:"r,omitempty"`
	S         string `protobuf:"bytes,6,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ExplorerSendTxSetRewardClaimerRequest) Reset() {
	*x = ExplorerSendTxSetRewardClaimerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerSendTxSetRewardClaimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerSendTxSetRewardClaimerRequest) ProtoMessage() {}

func (x *ExplorerSendTxSetRewardClaimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerSendTxSetRewardClaimerRequest.ProtoReflect.Descriptor instead.
func (*ExplorerSendTxSetRewardClaimerRequest) Descriptor() ([]byte, []int) {
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{3}
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetClaimer() string {
	if x != nil {
		return x.Claimer
	}
	return ""
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetExpiredAt() uint64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetV() uint32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *ExplorerSendTxSetRewardClaimerRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type ExplorerSendTxNodeReportVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer        string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	AttestationId string `protobuf:"bytes,2,opt,name=attestation_id,json=attestationId,proto3" json:"attestation_id,omitempty"`
	Result        uint32 `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Index         uint32 `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	V             uint32 `protobuf:"varint,5,opt,name=v,proto3" json:"v,omitempty"`
	R             string `protobuf:"bytes,6,opt,name=r,proto3" json:"r,omitempty"`
	S             string `protobuf:"bytes,7,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ExplorerSendTxNodeReportVerificationRequest) Reset() {
	*x = ExplorerSendTxNodeReportVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerSendTxNodeReportVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerSendTxNodeReportVerificationRequest) ProtoMessage() {}

func (x *ExplorerSendTxNodeReportVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerSendTxNodeReportVerificationRequest.ProtoReflect.Descriptor instead.
func (*ExplorerSendTxNodeReportVerificationRequest) Descriptor() ([]byte, []int) {
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{4}
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetAttestationId() string {
	if x != nil {
		return x.AttestationId
	}
	return ""
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetV() uint32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *ExplorerSendTxNodeReportVerificationRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gasless_gasless_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_gasless_gasless_proto_msgTypes[5]
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
	return file_api_gasless_gasless_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_api_gasless_gasless_proto protoreflect.FileDescriptor

var file_api_gasless_gasless_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x61,
	0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x1e, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22,
	0x80, 0x01, 0x0a, 0x1d, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x29, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x78, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x76, 0x12, 0x0c,
	0x0a, 0x01, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x25, 0x45,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22,
	0xc4, 0x01, 0x0a, 0x2b, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc1, 0x06, 0x0a,
	0x08, 0x47, 0x61, 0x73, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x17, 0x45, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x78,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x8c, 0x01, 0x0a, 0x16,
	0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4e, 0x6f,
	0x64, 0x65, 0x45, 0x78, 0x69, 0x74, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x78,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x12, 0xb1, 0x01, 0x0a, 0x22, 0x45,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x74, 0x2f,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x78, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x12, 0xa5,
	0x01, 0x0a, 0x1e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x78, 0x53, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65,
	0x72, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x78, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x12, 0xb7, 0x01, 0x0a, 0x24, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x22, 0x33, 0x2f, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x65, 0x74, 0x2f,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x76, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73,
	0x73, 0x3b, 0x67, 0x61, 0x73, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_gasless_gasless_proto_rawDescOnce sync.Once
	file_api_gasless_gasless_proto_rawDescData = file_api_gasless_gasless_proto_rawDesc
)

func file_api_gasless_gasless_proto_rawDescGZIP() []byte {
	file_api_gasless_gasless_proto_rawDescOnce.Do(func() {
		file_api_gasless_gasless_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gasless_gasless_proto_rawDescData)
	})
	return file_api_gasless_gasless_proto_rawDescData
}

var file_api_gasless_gasless_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_gasless_gasless_proto_goTypes = []any{
	(*ExplorerSendTxNodeEnterRequest)(nil),              // 0: api.gasless.ExplorerSendTxNodeEnterRequest
	(*ExplorerSendTxNodeExitRequest)(nil),               // 1: api.gasless.ExplorerSendTxNodeExitRequest
	(*ExplorerSendTxModifyCommissionRateRequest)(nil),   // 2: api.gasless.ExplorerSendTxModifyCommissionRateRequest
	(*ExplorerSendTxSetRewardClaimerRequest)(nil),       // 3: api.gasless.ExplorerSendTxSetRewardClaimerRequest
	(*ExplorerSendTxNodeReportVerificationRequest)(nil), // 4: api.gasless.ExplorerSendTxNodeReportVerificationRequest
	(*Response)(nil), // 5: api.gasless.Response
}
var file_api_gasless_gasless_proto_depIdxs = []int32{
	0, // 0: api.gasless.Gassless.ExplorerSendTxNodeEnter:input_type -> api.gasless.ExplorerSendTxNodeEnterRequest
	1, // 1: api.gasless.Gassless.ExplorerSendTxNodeExit:input_type -> api.gasless.ExplorerSendTxNodeExitRequest
	2, // 2: api.gasless.Gassless.ExplorerSendTxModifyCommissionRate:input_type -> api.gasless.ExplorerSendTxModifyCommissionRateRequest
	3, // 3: api.gasless.Gassless.ExplorerSendTxSetRewardClaimer:input_type -> api.gasless.ExplorerSendTxSetRewardClaimerRequest
	4, // 4: api.gasless.Gassless.ExplorerSendTxNodeReportVerification:input_type -> api.gasless.ExplorerSendTxNodeReportVerificationRequest
	5, // 5: api.gasless.Gassless.ExplorerSendTxNodeEnter:output_type -> api.gasless.Response
	5, // 6: api.gasless.Gassless.ExplorerSendTxNodeExit:output_type -> api.gasless.Response
	5, // 7: api.gasless.Gassless.ExplorerSendTxModifyCommissionRate:output_type -> api.gasless.Response
	5, // 8: api.gasless.Gassless.ExplorerSendTxSetRewardClaimer:output_type -> api.gasless.Response
	5, // 9: api.gasless.Gassless.ExplorerSendTxNodeReportVerification:output_type -> api.gasless.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_gasless_gasless_proto_init() }
func file_api_gasless_gasless_proto_init() {
	if File_api_gasless_gasless_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gasless_gasless_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExplorerSendTxNodeEnterRequest); i {
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
		file_api_gasless_gasless_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExplorerSendTxNodeExitRequest); i {
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
		file_api_gasless_gasless_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExplorerSendTxModifyCommissionRateRequest); i {
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
		file_api_gasless_gasless_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExplorerSendTxSetRewardClaimerRequest); i {
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
		file_api_gasless_gasless_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ExplorerSendTxNodeReportVerificationRequest); i {
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
		file_api_gasless_gasless_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_gasless_gasless_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gasless_gasless_proto_goTypes,
		DependencyIndexes: file_api_gasless_gasless_proto_depIdxs,
		MessageInfos:      file_api_gasless_gasless_proto_msgTypes,
	}.Build()
	File_api_gasless_gasless_proto = out.File
	file_api_gasless_gasless_proto_rawDesc = nil
	file_api_gasless_gasless_proto_goTypes = nil
	file_api_gasless_gasless_proto_depIdxs = nil
}
