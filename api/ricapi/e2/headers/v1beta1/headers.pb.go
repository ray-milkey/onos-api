// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/ricapi/e2/headers/v1beta1/headers.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EncodingType determines encoding type for the response messages
type EncodingType int32

const (
	EncodingType_PROTO    EncodingType = 0
	EncodingType_ASN1_PER EncodingType = 1
	EncodingType_ASN1_XER EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "PROTO",
	1: "ASN1_PER",
	2: "ASN1_XER",
}

var EncodingType_value = map[string]int32{
	"PROTO":    0,
	"ASN1_PER": 1,
	"ASN1_XER": 2,
}

func (x EncodingType) String() string {
	return proto.EnumName(EncodingType_name, int32(x))
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{0}
}

// ResponseStatus
type ResponseStatus int32

const (
	ResponseStatus_FAILED    ResponseStatus = 0
	ResponseStatus_SUCCEEDED ResponseStatus = 1
)

var ResponseStatus_name = map[int32]string{
	0: "FAILED",
	1: "SUCCEEDED",
}

var ResponseStatus_value = map[string]int32{
	"FAILED":    0,
	"SUCCEEDED": 1,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{1}
}

type ServiceModelInfo struct {
	ServiceModelId       string   `protobuf:"bytes,1,opt,name=service_model_id,json=serviceModelId,proto3" json:"service_model_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceModelInfo) Reset()         { *m = ServiceModelInfo{} }
func (m *ServiceModelInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceModelInfo) ProtoMessage()    {}
func (*ServiceModelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{0}
}
func (m *ServiceModelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceModelInfo.Unmarshal(m, b)
}
func (m *ServiceModelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceModelInfo.Marshal(b, m, deterministic)
}
func (m *ServiceModelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceModelInfo.Merge(m, src)
}
func (m *ServiceModelInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceModelInfo.Size(m)
}
func (m *ServiceModelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceModelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceModelInfo proto.InternalMessageInfo

func (m *ServiceModelInfo) GetServiceModelId() string {
	if m != nil {
		return m.ServiceModelId
	}
	return ""
}

// RequestHeader a common request header for all requests including encoding type, client/xApp/session info, ordering info, etc
type RequestHeader struct {
	EncodingType         EncodingType      `protobuf:"varint,1,opt,name=encoding_type,json=encodingType,proto3,enum=ricapi.e2.headers.v1beta1.EncodingType" json:"encoding_type,omitempty"`
	ServiceModelInfo     *ServiceModelInfo `protobuf:"bytes,2,opt,name=service_model_info,json=serviceModelInfo,proto3" json:"service_model_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{1}
}
func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetEncodingType() EncodingType {
	if m != nil {
		return m.EncodingType
	}
	return EncodingType_PROTO
}

func (m *RequestHeader) GetServiceModelInfo() *ServiceModelInfo {
	if m != nil {
		return m.ServiceModelInfo
	}
	return nil
}

// ResponseHeader a common response header for all responses including encoding type, client/xApp/session info, ordering info, etc
type ResponseHeader struct {
	EncodingType         EncodingType      `protobuf:"varint,1,opt,name=encoding_type,json=encodingType,proto3,enum=ricapi.e2.headers.v1beta1.EncodingType" json:"encoding_type,omitempty"`
	ServiceModelInfo     *ServiceModelInfo `protobuf:"bytes,2,opt,name=service_model_info,json=serviceModelInfo,proto3" json:"service_model_info,omitempty"`
	ResponseStatus       ResponseStatus    `protobuf:"varint,3,opt,name=response_status,json=responseStatus,proto3,enum=ricapi.e2.headers.v1beta1.ResponseStatus" json:"response_status,omitempty"`
	IndicationHeader     []byte            `protobuf:"bytes,4,opt,name=indication_header,json=indicationHeader,proto3" json:"indication_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{2}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetEncodingType() EncodingType {
	if m != nil {
		return m.EncodingType
	}
	return EncodingType_PROTO
}

func (m *ResponseHeader) GetServiceModelInfo() *ServiceModelInfo {
	if m != nil {
		return m.ServiceModelInfo
	}
	return nil
}

func (m *ResponseHeader) GetResponseStatus() ResponseStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return ResponseStatus_FAILED
}

func (m *ResponseHeader) GetIndicationHeader() []byte {
	if m != nil {
		return m.IndicationHeader
	}
	return nil
}

func init() {
	proto.RegisterEnum("ricapi.e2.headers.v1beta1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("ricapi.e2.headers.v1beta1.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*ServiceModelInfo)(nil), "ricapi.e2.headers.v1beta1.ServiceModelInfo")
	proto.RegisterType((*RequestHeader)(nil), "ricapi.e2.headers.v1beta1.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "ricapi.e2.headers.v1beta1.ResponseHeader")
}

func init() {
	proto.RegisterFile("api/ricapi/e2/headers/v1beta1/headers.proto", fileDescriptor_6f29ad563ec8c20e)
}

var fileDescriptor_6f29ad563ec8c20e = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0x26, 0x6c, 0x43, 0xc3, 0x0b, 0x99, 0xe7, 0x53, 0x76, 0x43, 0x5c, 0x96, 0x11, 0x2d, 0x11,
	0x99, 0x76, 0xdb, 0x34, 0x31, 0xc8, 0x54, 0x24, 0x5a, 0x90, 0x43, 0xa5, 0xb6, 0x97, 0x28, 0x24,
	0x0f, 0x70, 0x55, 0xec, 0x34, 0x36, 0x48, 0xfc, 0xb3, 0xaa, 0xbf, 0xae, 0x22, 0x09, 0x85, 0x22,
	0x95, 0x7b, 0x4f, 0xf6, 0xf7, 0xde, 0xf7, 0x9e, 0xbf, 0xef, 0xe9, 0x19, 0xd9, 0x51, 0xca, 0xdc,
	0x8c, 0xc5, 0xdb, 0x03, 0x3c, 0x77, 0x01, 0x51, 0x02, 0x99, 0x74, 0xd7, 0x9d, 0x29, 0xa8, 0xa8,
	0xb3, 0xc3, 0x4e, 0x9a, 0x09, 0x25, 0xc8, 0xd7, 0x82, 0xe8, 0x80, 0xe7, 0xec, 0x12, 0x25, 0xb1,
	0xf5, 0x1b, 0xe1, 0x00, 0xb2, 0x35, 0x8b, 0xe1, 0x5c, 0x24, 0x70, 0x37, 0xe0, 0x33, 0x41, 0x2c,
	0x84, 0x65, 0x11, 0x0b, 0x97, 0xdb, 0x60, 0xc8, 0x12, 0x53, 0x6b, 0x6a, 0x56, 0x9d, 0x1a, 0xf2,
	0x90, 0x9b, 0xb4, 0x1e, 0x34, 0xd4, 0xa0, 0x70, 0xbf, 0x02, 0xa9, 0xce, 0xf2, 0xc6, 0x64, 0x88,
	0x1a, 0xc0, 0x63, 0x91, 0x30, 0x3e, 0x0f, 0xd5, 0x26, 0x85, 0xbc, 0xd0, 0xf0, 0xbe, 0x39, 0xaf,
	0x4a, 0x70, 0xfc, 0x92, 0x3f, 0xd9, 0xa4, 0x40, 0x75, 0x38, 0x40, 0xe4, 0x1a, 0x91, 0x23, 0x25,
	0x7c, 0x26, 0xcc, 0x6a, 0x53, 0xb3, 0x3e, 0x79, 0xf6, 0x89, 0x96, 0xc7, 0x96, 0x28, 0x96, 0x47,
	0x91, 0xd6, 0x63, 0x15, 0x19, 0x14, 0x64, 0x2a, 0xb8, 0x84, 0x37, 0xa6, 0x9d, 0x50, 0xf4, 0x39,
	0x2b, 0xa5, 0x87, 0x52, 0x45, 0x6a, 0x25, 0xcd, 0x77, 0xb9, 0xd4, 0xef, 0x27, 0xfa, 0xee, 0xcc,
	0x06, 0x79, 0x01, 0x35, 0xb2, 0x17, 0x98, 0xd8, 0xe8, 0x0b, 0xe3, 0x09, 0x8b, 0x23, 0xc5, 0x04,
	0x0f, 0x8b, 0x62, 0xf3, 0x7d, 0x53, 0xb3, 0x74, 0x8a, 0xf7, 0x89, 0x62, 0x52, 0xed, 0x5f, 0x48,
	0x3f, 0x74, 0x4e, 0xea, 0xe8, 0xc3, 0x98, 0x8e, 0x26, 0x23, 0x5c, 0x21, 0x3a, 0xfa, 0xd8, 0x0d,
	0x2e, 0x3a, 0xe1, 0xd8, 0xa7, 0x58, 0x7b, 0x46, 0x57, 0x3e, 0xc5, 0xd5, 0xb6, 0xbd, 0x1f, 0x79,
	0xf9, 0x2a, 0x42, 0xb5, 0xff, 0xdd, 0xc1, 0xd0, 0xef, 0xe3, 0x0a, 0x69, 0xa0, 0x7a, 0x70, 0xd9,
	0xeb, 0xf9, 0x7e, 0xdf, 0xef, 0x63, 0xed, 0xdf, 0xdf, 0x9b, 0x3f, 0x73, 0xa6, 0x16, 0xab, 0xa9,
	0x13, 0x8b, 0xa5, 0x2b, 0xb8, 0x90, 0x69, 0x26, 0x6e, 0x21, 0x56, 0xf9, 0xfd, 0xc7, 0x76, 0xf1,
	0x4f, 0xfe, 0x81, 0x69, 0x2d, 0x5f, 0xfe, 0x9f, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0x27,
	0xfe, 0xf4, 0x2b, 0x03, 0x00, 0x00,
}
