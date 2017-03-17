// Code generated by protoc-gen-go.
// source: PhotonCommands.proto
// DO NOT EDIT!

/*
Package photongo is a generated protocol buffer package.

It is generated from these files:
	PhotonCommands.proto

It has these top-level messages:
	PhotonWrapper
	Connect
	Connect_Response
*/
package photongo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolVersion int32

const (
	ProtocolVersion_NONE ProtocolVersion = 0
	ProtocolVersion_V1   ProtocolVersion = 1
)

var ProtocolVersion_name = map[int32]string{
	0: "NONE",
	1: "V1",
}
var ProtocolVersion_value = map[string]int32{
	"NONE": 0,
	"V1":   1,
}

func (x ProtocolVersion) String() string {
	return proto.EnumName(ProtocolVersion_name, int32(x))
}
func (ProtocolVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusCodes int32

const (
	StatusCodes_UNKNOWN                StatusCodes = 0
	StatusCodes_OK                     StatusCodes = 200
	StatusCodes_PING                   StatusCodes = 201
	StatusCodes_BAD_REQUEST            StatusCodes = 400
	StatusCodes_UNAUTHORIZED           StatusCodes = 401
	StatusCodes_OLD_VERSION            StatusCodes = 402
	StatusCodes_NO_RESPONSE            StatusCodes = 403
	StatusCodes_AUDIO_SSRC_COLLISION   StatusCodes = 404
	StatusCodes_VIDEO_SSRC_COLLISION   StatusCodes = 405
	StatusCodes_INVALID_STREAM_KEY     StatusCodes = 406
	StatusCodes_CHANNEL_IN_USE         StatusCodes = 407
	StatusCodes_REGION_UNSUPPORTED     StatusCodes = 408
	StatusCodes_NO_MEDIA_TIMEOUT       StatusCodes = 409
	StatusCodes_INTERNAL_SERVER_ERROR  StatusCodes = 500
	StatusCodes_INTERNAL_COMMAND_ERROR StatusCodes = 501
	StatusCodes_INTERNAL_LOCAL_ERROR   StatusCodes = 502
)

var StatusCodes_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "PING",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "OLD_VERSION",
	403: "NO_RESPONSE",
	404: "AUDIO_SSRC_COLLISION",
	405: "VIDEO_SSRC_COLLISION",
	406: "INVALID_STREAM_KEY",
	407: "CHANNEL_IN_USE",
	408: "REGION_UNSUPPORTED",
	409: "NO_MEDIA_TIMEOUT",
	500: "INTERNAL_SERVER_ERROR",
	501: "INTERNAL_COMMAND_ERROR",
	502: "INTERNAL_LOCAL_ERROR",
}
var StatusCodes_value = map[string]int32{
	"UNKNOWN":                0,
	"OK":                     200,
	"PING":                   201,
	"BAD_REQUEST":            400,
	"UNAUTHORIZED":           401,
	"OLD_VERSION":            402,
	"NO_RESPONSE":            403,
	"AUDIO_SSRC_COLLISION":   404,
	"VIDEO_SSRC_COLLISION":   405,
	"INVALID_STREAM_KEY":     406,
	"CHANNEL_IN_USE":         407,
	"REGION_UNSUPPORTED":     408,
	"NO_MEDIA_TIMEOUT":       409,
	"INTERNAL_SERVER_ERROR":  500,
	"INTERNAL_COMMAND_ERROR": 501,
	"INTERNAL_LOCAL_ERROR":   502,
}

func (x StatusCodes) String() string {
	return proto.EnumName(StatusCodes_name, int32(x))
}
func (StatusCodes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PhotonWrapper struct {
	StatusCode StatusCodes          `protobuf:"varint,1,opt,name=StatusCode,enum=Photon.Commands.StatusCodes" json:"StatusCode,omitempty"`
	Command    *google_protobuf.Any `protobuf:"bytes,15,opt,name=Command" json:"Command,omitempty"`
}

func (m *PhotonWrapper) Reset()                    { *m = PhotonWrapper{} }
func (m *PhotonWrapper) String() string            { return proto.CompactTextString(m) }
func (*PhotonWrapper) ProtoMessage()               {}
func (*PhotonWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PhotonWrapper) GetStatusCode() StatusCodes {
	if m != nil {
		return m.StatusCode
	}
	return StatusCodes_UNKNOWN
}

func (m *PhotonWrapper) GetCommand() *google_protobuf.Any {
	if m != nil {
		return m.Command
	}
	return nil
}

type Connect struct {
	ClientProtocolVersion ProtocolVersion `protobuf:"varint,1,opt,name=ClientProtocolVersion,enum=Photon.Commands.ProtocolVersion" json:"ClientProtocolVersion,omitempty"`
}

func (m *Connect) Reset()                    { *m = Connect{} }
func (m *Connect) String() string            { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()               {}
func (*Connect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Connect) GetClientProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.ClientProtocolVersion
	}
	return ProtocolVersion_NONE
}

type Connect_Response struct {
	ServerProtocolVersion ProtocolVersion `protobuf:"varint,1,opt,name=ServerProtocolVersion,enum=Photon.Commands.ProtocolVersion" json:"ServerProtocolVersion,omitempty"`
	HmacKey               string          `protobuf:"bytes,3,opt,name=HmacKey" json:"HmacKey,omitempty"`
}

func (m *Connect_Response) Reset()                    { *m = Connect_Response{} }
func (m *Connect_Response) String() string            { return proto.CompactTextString(m) }
func (*Connect_Response) ProtoMessage()               {}
func (*Connect_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Connect_Response) GetServerProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.ServerProtocolVersion
	}
	return ProtocolVersion_NONE
}

func (m *Connect_Response) GetHmacKey() string {
	if m != nil {
		return m.HmacKey
	}
	return ""
}

func init() {
	proto.RegisterType((*PhotonWrapper)(nil), "Photon.Commands.PhotonWrapper")
	proto.RegisterType((*Connect)(nil), "Photon.Commands.Connect")
	proto.RegisterType((*Connect_Response)(nil), "Photon.Commands.Connect_Response")
	proto.RegisterEnum("Photon.Commands.ProtocolVersion", ProtocolVersion_name, ProtocolVersion_value)
	proto.RegisterEnum("Photon.Commands.StatusCodes", StatusCodes_name, StatusCodes_value)
}

func init() { proto.RegisterFile("PhotonCommands.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x97, 0x65, 0x5a, 0xb7, 0xa7, 0xb0, 0x19, 0xd3, 0x42, 0x37, 0x38, 0x54, 0xe5, 0x52,
	0xed, 0x90, 0x89, 0x71, 0xe5, 0x92, 0x25, 0xd6, 0x6a, 0x35, 0xb5, 0x83, 0x9d, 0x64, 0x62, 0x17,
	0x2b, 0xeb, 0x42, 0x99, 0xd4, 0xc6, 0x55, 0x92, 0x21, 0xf5, 0xc0, 0x8d, 0x0f, 0xc0, 0xfb, 0xcb,
	0x37, 0x82, 0x4f, 0xc1, 0x97, 0x00, 0xce, 0xa8, 0x69, 0x3b, 0xa6, 0x6e, 0x27, 0x8e, 0x79, 0x7e,
	0xbf, 0x27, 0xff, 0xbf, 0x6c, 0x43, 0xcd, 0x7f, 0xa9, 0x0b, 0x9d, 0x3a, 0x7a, 0x34, 0x8a, 0xd3,
	0xb3, 0xdc, 0x1a, 0x67, 0xba, 0xd0, 0x78, 0x7b, 0x36, 0xb5, 0x16, 0xe3, 0xdd, 0x9d, 0x81, 0xd6,
	0x83, 0x61, 0xb2, 0x5f, 0xe2, 0xd3, 0x8b, 0x17, 0xfb, 0x71, 0x3a, 0x99, 0xb9, 0xad, 0xd7, 0x70,
	0x7b, 0x66, 0x1f, 0x67, 0xf1, 0x78, 0x9c, 0x64, 0xf8, 0x29, 0x80, 0x2c, 0xe2, 0xe2, 0x22, 0x77,
	0xf4, 0x59, 0xd2, 0x30, 0x9a, 0x46, 0x7b, 0xeb, 0xe0, 0xa1, 0xb5, 0xf4, 0x47, 0xeb, 0x9f, 0x92,
	0x8b, 0x2b, 0x3e, 0xb6, 0xa0, 0x32, 0x77, 0x1a, 0xdb, 0x4d, 0xa3, 0x5d, 0x3d, 0xa8, 0x59, 0xb3,
	0x6c, 0x6b, 0x91, 0x6d, 0xd9, 0xe9, 0x44, 0x2c, 0xa4, 0x56, 0x3c, 0xf5, 0xd3, 0x34, 0xe9, 0x17,
	0x38, 0x82, 0xba, 0x33, 0x3c, 0x4f, 0xd2, 0xc2, 0x9f, 0x9a, 0x7d, 0x3d, 0x8c, 0x92, 0x2c, 0x3f,
	0xd7, 0xe9, 0xbc, 0x43, 0xf3, 0x5a, 0x87, 0x25, 0x4f, 0xdc, 0xbc, 0xde, 0x7a, 0x63, 0x00, 0x9a,
	0x67, 0x28, 0x91, 0xe4, 0x63, 0x9d, 0xe6, 0xc9, 0x34, 0x4c, 0x26, 0xd9, 0xab, 0x24, 0xfb, 0xef,
	0xb0, 0x1b, 0xd7, 0x71, 0x03, 0x2a, 0x9d, 0x51, 0xdc, 0xef, 0x26, 0x93, 0x86, 0xd9, 0x34, 0xda,
	0x9b, 0x62, 0xf1, 0xb9, 0xf7, 0x08, 0xb6, 0x97, 0xe5, 0x0d, 0x58, 0x63, 0x9c, 0x11, 0xb4, 0x82,
	0xd7, 0x61, 0x35, 0x7a, 0x8c, 0x8c, 0xbd, 0x9f, 0xab, 0x50, 0xbd, 0x72, 0xb4, 0xb8, 0x0a, 0x95,
	0x90, 0x75, 0x19, 0x3f, 0x66, 0x68, 0x05, 0x57, 0x60, 0x95, 0x77, 0xd1, 0x77, 0x03, 0x6f, 0xc2,
	0x9a, 0x4f, 0xd9, 0x11, 0xfa, 0x61, 0x60, 0x04, 0xd5, 0x43, 0xdb, 0x55, 0x82, 0x3c, 0x0b, 0x89,
	0x0c, 0xd0, 0x5b, 0x13, 0xdf, 0x81, 0x5b, 0x21, 0xb3, 0xc3, 0xa0, 0xc3, 0x05, 0x3d, 0x21, 0x2e,
	0x7a, 0x67, 0x4e, 0x25, 0xee, 0xb9, 0x2a, 0x22, 0x42, 0x52, 0xce, 0xd0, 0xfb, 0x72, 0xc2, 0xb8,
	0x12, 0x44, 0xfa, 0x9c, 0x49, 0x82, 0x3e, 0x98, 0x78, 0x07, 0x6a, 0x76, 0xe8, 0x52, 0xae, 0xa4,
	0x14, 0x8e, 0x72, 0xb8, 0xe7, 0xd1, 0x52, 0xfe, 0x58, 0xa2, 0x88, 0xba, 0xe4, 0x1a, 0xfa, 0x64,
	0xe2, 0xfb, 0x80, 0x29, 0x8b, 0x6c, 0x8f, 0xba, 0x4a, 0x06, 0x82, 0xd8, 0x3d, 0xd5, 0x25, 0xcf,
	0xd1, 0x67, 0x13, 0xdf, 0x85, 0x2d, 0xa7, 0x63, 0x33, 0x46, 0x3c, 0x45, 0x99, 0x0a, 0x25, 0x41,
	0x5f, 0x4a, 0x5b, 0x90, 0x23, 0xca, 0x99, 0x0a, 0x99, 0x0c, 0x7d, 0x9f, 0x8b, 0x80, 0xb8, 0xe8,
	0xab, 0x89, 0xeb, 0x80, 0x18, 0x57, 0x3d, 0xe2, 0x52, 0x5b, 0x05, 0xb4, 0x47, 0x78, 0x18, 0xa0,
	0x6f, 0x26, 0xde, 0x85, 0x3a, 0x65, 0x01, 0x11, 0xcc, 0xf6, 0x94, 0x24, 0x22, 0x22, 0x42, 0x11,
	0x21, 0xb8, 0x40, 0xbf, 0x4c, 0xfc, 0x00, 0xee, 0x5d, 0x32, 0x87, 0xf7, 0x7a, 0x36, 0x73, 0xe7,
	0xf0, 0x77, 0xd9, 0xf8, 0x12, 0x7a, 0xdc, 0xb1, 0xbd, 0x39, 0xfa, 0x63, 0x1e, 0xc2, 0xc9, 0xc6,
	0xb8, 0xbc, 0xda, 0x81, 0x3e, 0x5d, 0x2f, 0xdf, 0xe4, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x03, 0x8d, 0xef, 0x46, 0x03, 0x00, 0x00,
}
