// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/shared_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [SharedCriterionService.GetSharedCriterion][google.ads.googleads.v1.services.SharedCriterionService.GetSharedCriterion].
type GetSharedCriterionRequest struct {
	// The resource name of the shared criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSharedCriterionRequest) Reset()         { *m = GetSharedCriterionRequest{} }
func (m *GetSharedCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSharedCriterionRequest) ProtoMessage()    {}
func (*GetSharedCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b35fbbce441bac, []int{0}
}

func (m *GetSharedCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSharedCriterionRequest.Unmarshal(m, b)
}
func (m *GetSharedCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSharedCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetSharedCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSharedCriterionRequest.Merge(m, src)
}
func (m *GetSharedCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSharedCriterionRequest.Size(m)
}
func (m *GetSharedCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSharedCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSharedCriterionRequest proto.InternalMessageInfo

func (m *GetSharedCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [SharedCriterionService.MutateSharedCriteria][google.ads.googleads.v1.services.SharedCriterionService.MutateSharedCriteria].
type MutateSharedCriteriaRequest struct {
	// The ID of the customer whose shared criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual shared criteria.
	Operations []*SharedCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateSharedCriteriaRequest) Reset()         { *m = MutateSharedCriteriaRequest{} }
func (m *MutateSharedCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriteriaRequest) ProtoMessage()    {}
func (*MutateSharedCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b35fbbce441bac, []int{1}
}

func (m *MutateSharedCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateSharedCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriteriaRequest.Merge(m, src)
}
func (m *MutateSharedCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Size(m)
}
func (m *MutateSharedCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriteriaRequest proto.InternalMessageInfo

func (m *MutateSharedCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateSharedCriteriaRequest) GetOperations() []*SharedCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateSharedCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateSharedCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an shared criterion.
type SharedCriterionOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*SharedCriterionOperation_Create
	//	*SharedCriterionOperation_Remove
	Operation            isSharedCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *SharedCriterionOperation) Reset()         { *m = SharedCriterionOperation{} }
func (m *SharedCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*SharedCriterionOperation) ProtoMessage()    {}
func (*SharedCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b35fbbce441bac, []int{2}
}

func (m *SharedCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedCriterionOperation.Unmarshal(m, b)
}
func (m *SharedCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedCriterionOperation.Marshal(b, m, deterministic)
}
func (m *SharedCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedCriterionOperation.Merge(m, src)
}
func (m *SharedCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_SharedCriterionOperation.Size(m)
}
func (m *SharedCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SharedCriterionOperation proto.InternalMessageInfo

type isSharedCriterionOperation_Operation interface {
	isSharedCriterionOperation_Operation()
}

type SharedCriterionOperation_Create struct {
	Create *resources.SharedCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedCriterionOperation_Create) isSharedCriterionOperation_Operation() {}

func (*SharedCriterionOperation_Remove) isSharedCriterionOperation_Operation() {}

func (m *SharedCriterionOperation) GetOperation() isSharedCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *SharedCriterionOperation) GetCreate() *resources.SharedCriterion {
	if x, ok := m.GetOperation().(*SharedCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *SharedCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*SharedCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedCriterionOperation_Create)(nil),
		(*SharedCriterionOperation_Remove)(nil),
	}
}

// Response message for a shared criterion mutate.
type MutateSharedCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateSharedCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateSharedCriteriaResponse) Reset()         { *m = MutateSharedCriteriaResponse{} }
func (m *MutateSharedCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriteriaResponse) ProtoMessage()    {}
func (*MutateSharedCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b35fbbce441bac, []int{3}
}

func (m *MutateSharedCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateSharedCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriteriaResponse.Merge(m, src)
}
func (m *MutateSharedCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Size(m)
}
func (m *MutateSharedCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriteriaResponse proto.InternalMessageInfo

func (m *MutateSharedCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateSharedCriteriaResponse) GetResults() []*MutateSharedCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the shared criterion mutate.
type MutateSharedCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateSharedCriterionResult) Reset()         { *m = MutateSharedCriterionResult{} }
func (m *MutateSharedCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriterionResult) ProtoMessage()    {}
func (*MutateSharedCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b35fbbce441bac, []int{4}
}

func (m *MutateSharedCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriterionResult.Unmarshal(m, b)
}
func (m *MutateSharedCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriterionResult.Merge(m, src)
}
func (m *MutateSharedCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriterionResult.Size(m)
}
func (m *MutateSharedCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriterionResult proto.InternalMessageInfo

func (m *MutateSharedCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSharedCriterionRequest)(nil), "google.ads.googleads.v1.services.GetSharedCriterionRequest")
	proto.RegisterType((*MutateSharedCriteriaRequest)(nil), "google.ads.googleads.v1.services.MutateSharedCriteriaRequest")
	proto.RegisterType((*SharedCriterionOperation)(nil), "google.ads.googleads.v1.services.SharedCriterionOperation")
	proto.RegisterType((*MutateSharedCriteriaResponse)(nil), "google.ads.googleads.v1.services.MutateSharedCriteriaResponse")
	proto.RegisterType((*MutateSharedCriterionResult)(nil), "google.ads.googleads.v1.services.MutateSharedCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/shared_criterion_service.proto", fileDescriptor_82b35fbbce441bac)
}

var fileDescriptor_82b35fbbce441bac = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0xfe, 0x65, 0xf7, 0x47, 0xb5, 0xb3, 0x55, 0x61, 0xfc, 0x17, 0xb7, 0x05, 0x97, 0x58, 0xb0,
	0xec, 0xc5, 0xc4, 0xdd, 0xde, 0x94, 0x94, 0x56, 0xb3, 0x62, 0x5b, 0x41, 0x6d, 0x49, 0xa1, 0x42,
	0x59, 0x08, 0x63, 0x32, 0xc6, 0x40, 0x92, 0x89, 0x33, 0x93, 0x85, 0x52, 0x7a, 0xd3, 0x17, 0xf0,
	0xc2, 0x37, 0xf0, 0xd2, 0x97, 0xe8, 0x7d, 0x6f, 0xbc, 0x10, 0xdf, 0xc0, 0x0b, 0xf1, 0x29, 0x24,
	0x99, 0xcc, 0xb6, 0x8d, 0x59, 0x56, 0xf6, 0xee, 0xe4, 0x9c, 0x33, 0xdf, 0xf9, 0xbe, 0xf3, 0x27,
	0xe0, 0x69, 0x40, 0x69, 0x10, 0x11, 0x13, 0xfb, 0xdc, 0x94, 0x66, 0x6e, 0x8d, 0x7a, 0x26, 0x27,
	0x6c, 0x14, 0x7a, 0x84, 0x9b, 0xfc, 0x03, 0x66, 0xc4, 0x77, 0x3d, 0x16, 0x0a, 0xc2, 0x42, 0x9a,
	0xb8, 0x65, 0x04, 0xa5, 0x8c, 0x0a, 0x0a, 0x3b, 0xf2, 0x15, 0xc2, 0x3e, 0x47, 0x63, 0x00, 0x34,
	0xea, 0x21, 0x05, 0xd0, 0x5e, 0x9b, 0x54, 0x82, 0x11, 0x4e, 0x33, 0x56, 0x57, 0x43, 0x62, 0xb7,
	0x97, 0xd4, 0xcb, 0x34, 0x34, 0x71, 0x92, 0x50, 0x81, 0x45, 0x48, 0x13, 0x5e, 0x46, 0xef, 0x97,
	0x51, 0x96, 0x7a, 0x26, 0x17, 0x58, 0x64, 0xd5, 0x40, 0xfe, 0xcc, 0x8b, 0x42, 0x92, 0x08, 0x19,
	0x30, 0x9e, 0x81, 0x07, 0xdb, 0x44, 0xec, 0x17, 0xc5, 0x9e, 0xab, 0x5a, 0x0e, 0xf9, 0x98, 0x11,
	0x2e, 0xe0, 0x23, 0x70, 0x43, 0x11, 0x72, 0x13, 0x1c, 0x13, 0x5d, 0xeb, 0x68, 0x2b, 0xf3, 0xce,
	0x82, 0x72, 0xbe, 0xc1, 0x31, 0x31, 0x7e, 0x69, 0x60, 0xf1, 0x75, 0x26, 0xb0, 0x20, 0x57, 0x50,
	0xb0, 0x02, 0x79, 0x08, 0x5a, 0x5e, 0xc6, 0x05, 0x8d, 0x09, 0x73, 0x43, 0xbf, 0x84, 0x00, 0xca,
	0xf5, 0xd2, 0x87, 0x87, 0x00, 0xd0, 0x94, 0x30, 0x29, 0x44, 0x6f, 0x74, 0x9a, 0x2b, 0xad, 0xbe,
	0x85, 0xa6, 0xf5, 0x10, 0x55, 0x38, 0xef, 0x2a, 0x08, 0xe7, 0x12, 0x1a, 0x7c, 0x0c, 0x6e, 0xa5,
	0x98, 0x89, 0x10, 0x47, 0xee, 0x7b, 0x1c, 0x46, 0x19, 0x23, 0x7a, 0xb3, 0xa3, 0xad, 0x5c, 0x77,
	0x6e, 0x96, 0xee, 0x2d, 0xe9, 0xcd, 0xa5, 0x8e, 0x70, 0x14, 0xfa, 0x58, 0x10, 0x97, 0x26, 0xd1,
	0x91, 0xfe, 0x7f, 0x91, 0xb6, 0xa0, 0x9c, 0xbb, 0x49, 0x74, 0x64, 0x7c, 0xd2, 0x80, 0x3e, 0xa9,
	0x2c, 0x7c, 0x05, 0xe6, 0x3c, 0x46, 0xb0, 0x90, 0x5d, 0x6a, 0xf5, 0xfb, 0x13, 0x25, 0x8c, 0x87,
	0x5c, 0xd5, 0xb0, 0xf3, 0x9f, 0x53, 0x62, 0x40, 0x1d, 0xcc, 0x31, 0x12, 0xd3, 0x91, 0xe4, 0x3b,
	0x9f, 0x47, 0xe4, 0xf7, 0xa0, 0x05, 0xe6, 0xc7, 0x02, 0x8d, 0x33, 0x0d, 0x2c, 0xd5, 0x37, 0x9f,
	0xa7, 0x34, 0xe1, 0x04, 0x6e, 0x81, 0xbb, 0x95, 0x06, 0xb8, 0x84, 0x31, 0xca, 0x0a, 0xd8, 0x56,
	0x1f, 0x2a, 0x92, 0x2c, 0xf5, 0xd0, 0x7e, 0xb1, 0x31, 0xce, 0xed, 0xab, 0xad, 0x79, 0x91, 0xa7,
	0xc3, 0xb7, 0xe0, 0x1a, 0x23, 0x3c, 0x8b, 0x84, 0x9a, 0xd0, 0xc6, 0xf4, 0x09, 0xd5, 0x10, 0xcb,
	0x77, 0x2b, 0x47, 0x71, 0x14, 0x9a, 0x31, 0xa8, 0xdd, 0x1e, 0x95, 0xf7, 0x4f, 0x2b, 0xd8, 0xff,
	0xd6, 0x04, 0xf7, 0x2a, 0xcf, 0xf7, 0x25, 0x09, 0x78, 0xa6, 0x01, 0xf8, 0xf7, 0x82, 0xc3, 0xf5,
	0xe9, 0xec, 0x27, 0x9e, 0x45, 0x7b, 0x86, 0xc9, 0x1a, 0x6b, 0xa7, 0xdf, 0x7f, 0x7e, 0x6e, 0xf4,
	0xe1, 0x93, 0xfc, 0xca, 0x8f, 0xaf, 0x48, 0xda, 0x50, 0xb7, 0xc0, 0xcd, 0x6e, 0x79, 0xf6, 0x6a,
	0x8c, 0x66, 0xf7, 0x04, 0xfe, 0xd0, 0xc0, 0x9d, 0xba, 0x11, 0xc3, 0xd9, 0x26, 0xa0, 0xee, 0xb2,
	0xbd, 0x39, 0xeb, 0x73, 0xb9, 0x59, 0xc6, 0x66, 0xa1, 0x68, 0xcd, 0x58, 0xcd, 0x15, 0x5d, 0x48,
	0x38, 0xbe, 0x74, 0xec, 0x1b, 0xdd, 0x93, 0x8a, 0x20, 0x2b, 0x2e, 0x20, 0x2d, 0xad, 0xdb, 0x5e,
	0x3c, 0xb7, 0xf5, 0x8b, 0xb2, 0xa5, 0x95, 0x86, 0x1c, 0x79, 0x34, 0x1e, 0x9c, 0x36, 0xc0, 0xb2,
	0x47, 0xe3, 0xa9, 0x14, 0x07, 0x8b, 0xf5, 0x73, 0xdf, 0xcb, 0x7f, 0x6e, 0x7b, 0xda, 0xe1, 0x4e,
	0x09, 0x10, 0xd0, 0x08, 0x27, 0x01, 0xa2, 0x2c, 0x30, 0x03, 0x92, 0x14, 0xbf, 0x3e, 0xf3, 0xa2,
	0xe4, 0xe4, 0x5f, 0xfd, 0xba, 0x32, 0xbe, 0x34, 0x9a, 0xdb, 0xb6, 0xfd, 0xb5, 0xd1, 0xd9, 0x96,
	0x80, 0xb6, 0xcf, 0x91, 0x34, 0x73, 0xeb, 0xa0, 0x87, 0xca, 0xc2, 0xfc, 0x5c, 0xa5, 0x0c, 0x6d,
	0x9f, 0x0f, 0xc7, 0x29, 0xc3, 0x83, 0xde, 0x50, 0xa5, 0xfc, 0x6e, 0x2c, 0x4b, 0xbf, 0x65, 0xd9,
	0x3e, 0xb7, 0xac, 0x71, 0x92, 0x65, 0x1d, 0xf4, 0x2c, 0x4b, 0xa5, 0xbd, 0x9b, 0x2b, 0x78, 0xae,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x96, 0xe2, 0x57, 0x91, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SharedCriterionServiceClient is the client API for SharedCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedCriterionServiceClient interface {
	// Returns the requested shared criterion in full detail.
	GetSharedCriterion(ctx context.Context, in *GetSharedCriterionRequest, opts ...grpc.CallOption) (*resources.SharedCriterion, error)
	// Creates or removes shared criteria. Operation statuses are returned.
	MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error)
}

type sharedCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedCriterionServiceClient(cc grpc.ClientConnInterface) SharedCriterionServiceClient {
	return &sharedCriterionServiceClient{cc}
}

func (c *sharedCriterionServiceClient) GetSharedCriterion(ctx context.Context, in *GetSharedCriterionRequest, opts ...grpc.CallOption) (*resources.SharedCriterion, error) {
	out := new(resources.SharedCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.SharedCriterionService/GetSharedCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedCriterionServiceClient) MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error) {
	out := new(MutateSharedCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.SharedCriterionService/MutateSharedCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedCriterionServiceServer is the server API for SharedCriterionService service.
type SharedCriterionServiceServer interface {
	// Returns the requested shared criterion in full detail.
	GetSharedCriterion(context.Context, *GetSharedCriterionRequest) (*resources.SharedCriterion, error)
	// Creates or removes shared criteria. Operation statuses are returned.
	MutateSharedCriteria(context.Context, *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error)
}

// UnimplementedSharedCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSharedCriterionServiceServer struct {
}

func (*UnimplementedSharedCriterionServiceServer) GetSharedCriterion(ctx context.Context, req *GetSharedCriterionRequest) (*resources.SharedCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetSharedCriterion not implemented")
}
func (*UnimplementedSharedCriterionServiceServer) MutateSharedCriteria(ctx context.Context, req *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateSharedCriteria not implemented")
}

func RegisterSharedCriterionServiceServer(s *grpc.Server, srv SharedCriterionServiceServer) {
	s.RegisterService(&_SharedCriterionService_serviceDesc, srv)
}

func _SharedCriterionService_GetSharedCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSharedCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedCriterionServiceServer).GetSharedCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.SharedCriterionService/GetSharedCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedCriterionServiceServer).GetSharedCriterion(ctx, req.(*GetSharedCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedCriterionService_MutateSharedCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.SharedCriterionService/MutateSharedCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, req.(*MutateSharedCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.SharedCriterionService",
	HandlerType: (*SharedCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSharedCriterion",
			Handler:    _SharedCriterionService_GetSharedCriterion_Handler,
		},
		{
			MethodName: "MutateSharedCriteria",
			Handler:    _SharedCriterionService_MutateSharedCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/shared_criterion_service.proto",
}