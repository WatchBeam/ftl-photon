// Code generated by protoc-gen-go.
// source: IngestConnection.proto
// DO NOT EDIT!

/*
Package Photon_Commands is a generated protocol buffer package.

It is generated from these files:
	IngestConnection.proto

It has these top-level messages:
	PhotonWrapper
	Connect
	Connect_Response
*/
package Photon_Commands

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
	StatusCodes_AUDIO_SSRC_COLLISION   StatusCodes = 403
	StatusCodes_VIDEO_SSRC_COLLISION   StatusCodes = 404
	StatusCodes_INVALID_STREAM_KEY     StatusCodes = 405
	StatusCodes_CHANNEL_IN_USE         StatusCodes = 406
	StatusCodes_REGION_UNSUPPORTED     StatusCodes = 407
	StatusCodes_NO_MEDIA_TIMEOUT       StatusCodes = 408
	StatusCodes_INTERNAL_SERVER_ERROR  StatusCodes = 500
	StatusCodes_INTERNAL_COMMAND_ERROR StatusCodes = 501
)

var StatusCodes_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "PING",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "OLD_VERSION",
	403: "AUDIO_SSRC_COLLISION",
	404: "VIDEO_SSRC_COLLISION",
	405: "INVALID_STREAM_KEY",
	406: "CHANNEL_IN_USE",
	407: "REGION_UNSUPPORTED",
	408: "NO_MEDIA_TIMEOUT",
	500: "INTERNAL_SERVER_ERROR",
	501: "INTERNAL_COMMAND_ERROR",
}
var StatusCodes_value = map[string]int32{
	"UNKNOWN":                0,
	"OK":                     200,
	"PING":                   201,
	"BAD_REQUEST":            400,
	"UNAUTHORIZED":           401,
	"OLD_VERSION":            402,
	"AUDIO_SSRC_COLLISION":   403,
	"VIDEO_SSRC_COLLISION":   404,
	"INVALID_STREAM_KEY":     405,
	"CHANNEL_IN_USE":         406,
	"REGION_UNSUPPORTED":     407,
	"NO_MEDIA_TIMEOUT":       408,
	"INTERNAL_SERVER_ERROR":  500,
	"INTERNAL_COMMAND_ERROR": 501,
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

func init() { proto.RegisterFile("IngestConnection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0x06, 0xe0, 0x4e, 0xa7, 0x6a, 0xbe, 0x9e, 0x7c, 0x34, 0xc6, 0x34, 0x25, 0x2d, 0x2c, 0xa2,
	0xb0, 0x89, 0xba, 0x98, 0x8a, 0xb2, 0x65, 0x33, 0xcc, 0x58, 0x8d, 0x95, 0x89, 0x1d, 0x3c, 0x3f,
	0x15, 0x6c, 0xac, 0x69, 0x6a, 0x42, 0xa4, 0xc4, 0x8e, 0x66, 0xa6, 0x48, 0x59, 0xb0, 0xe3, 0x02,
	0xf8, 0x87, 0xeb, 0xe0, 0x2a, 0xe0, 0x7e, 0x60, 0x8f, 0x32, 0x49, 0xa0, 0x4a, 0xbb, 0x62, 0x69,
	0xbd, 0xcf, 0xf1, 0x7b, 0x64, 0x19, 0xf6, 0xa9, 0x1e, 0xaa, 0xbc, 0xf0, 0x8c, 0xd6, 0x6a, 0x50,
	0x8c, 0x8c, 0x76, 0xa6, 0x99, 0x29, 0x0c, 0xae, 0xf5, 0x5f, 0x9a, 0xc2, 0x68, 0xc7, 0x33, 0x93,
	0x49, 0xaa, 0x2f, 0xf2, 0xc3, 0x83, 0xa1, 0x31, 0xc3, 0xb1, 0x3a, 0x2e, 0xe3, 0xf3, 0xcb, 0x17,
	0xc7, 0xa9, 0x9e, 0x2d, 0x6c, 0xeb, 0x35, 0xdc, 0x5a, 0xe8, 0xb3, 0x2c, 0x9d, 0x4e, 0x55, 0x86,
	0x1f, 0x03, 0x84, 0x45, 0x5a, 0x5c, 0xe6, 0x9e, 0xb9, 0x50, 0x0d, 0xab, 0x69, 0xb5, 0x77, 0x4f,
	0xee, 0x3b, 0x6b, 0x37, 0x3a, 0x7f, 0x49, 0x2e, 0xae, 0x78, 0xec, 0x40, 0x65, 0x69, 0x1a, 0xb5,
	0xa6, 0xd5, 0xae, 0x9e, 0xec, 0x39, 0x8b, 0x6e, 0x67, 0xd5, 0xed, 0xb8, 0x7a, 0x26, 0x56, 0xa8,
	0x95, 0xce, 0x7d, 0xb9, 0x3e, 0x4e, 0xa0, 0xee, 0x8d, 0x47, 0x4a, 0x17, 0xfd, 0xb9, 0x1c, 0x98,
	0x71, 0xa2, 0xb2, 0x7c, 0x64, 0xf4, 0x72, 0x87, 0xe6, 0xb5, 0x1d, 0xd6, 0x9c, 0xb8, 0x79, 0xbc,
	0xf5, 0xc6, 0x02, 0xb4, 0xec, 0x90, 0x42, 0xe5, 0x53, 0xa3, 0x73, 0x35, 0x2f, 0x0b, 0x55, 0xf6,
	0x4a, 0x65, 0xff, 0x5c, 0x76, 0xe3, 0x38, 0x6e, 0x40, 0xa5, 0x33, 0x49, 0x07, 0x5d, 0x35, 0x6b,
	0xd8, 0x4d, 0xab, 0xbd, 0x23, 0x56, 0xc7, 0xa3, 0x07, 0x50, 0x5b, 0xc7, 0xff, 0xc1, 0x16, 0xe3,
	0x8c, 0xa0, 0x0d, 0xbc, 0x0d, 0x9b, 0xc9, 0x43, 0x64, 0x1d, 0x7d, 0xdb, 0x84, 0xea, 0x95, 0xa7,
	0xc5, 0x55, 0xa8, 0xc4, 0xac, 0xcb, 0xf8, 0x19, 0x43, 0x1b, 0xb8, 0x02, 0x9b, 0xbc, 0x8b, 0xbe,
	0x5b, 0x78, 0x07, 0xb6, 0xfa, 0x94, 0x9d, 0xa2, 0x1f, 0x16, 0x46, 0x50, 0x7d, 0xe2, 0xfa, 0x52,
	0x90, 0xa7, 0x31, 0x09, 0x23, 0xf4, 0xd6, 0xc6, 0xb7, 0xe1, 0xff, 0x98, 0xb9, 0x71, 0xd4, 0xe1,
	0x82, 0x3e, 0x27, 0x3e, 0x7a, 0x67, 0xcf, 0x11, 0x0f, 0x7c, 0x99, 0x10, 0x11, 0x52, 0xce, 0xd0,
	0x7b, 0x1b, 0x1f, 0xc0, 0x9e, 0x1b, 0xfb, 0x94, 0xcb, 0x30, 0x14, 0x9e, 0xf4, 0x78, 0x10, 0xd0,
	0x32, 0xfa, 0x50, 0x46, 0x09, 0xf5, 0xc9, 0xb5, 0xe8, 0xa3, 0x8d, 0xef, 0x02, 0xa6, 0x2c, 0x71,
	0x03, 0xea, 0xcb, 0x30, 0x12, 0xc4, 0xed, 0xc9, 0x2e, 0x79, 0x86, 0x3e, 0xd9, 0xf8, 0x0e, 0xec,
	0x7a, 0x1d, 0x97, 0x31, 0x12, 0x48, 0xca, 0x64, 0x1c, 0x12, 0xf4, 0xb9, 0xd4, 0x82, 0x9c, 0x52,
	0xce, 0x64, 0xcc, 0xc2, 0xb8, 0xdf, 0xe7, 0x22, 0x22, 0x3e, 0xfa, 0x62, 0xe3, 0x3a, 0x20, 0xc6,
	0x65, 0x8f, 0xf8, 0xd4, 0x95, 0x11, 0xed, 0x11, 0x1e, 0x47, 0xe8, 0xab, 0x8d, 0x0f, 0xa1, 0x4e,
	0x59, 0x44, 0x04, 0x73, 0x03, 0x19, 0x12, 0x91, 0x10, 0x21, 0x89, 0x10, 0x5c, 0xa0, 0x9f, 0x36,
	0xbe, 0x07, 0xfb, 0x7f, 0x32, 0x8f, 0xf7, 0x7a, 0x2e, 0xf3, 0x97, 0xe1, 0x2f, 0xfb, 0x7c, 0xbb,
	0xfc, 0x5a, 0x8f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x1c, 0x4d, 0x75, 0x0f, 0x03, 0x00,
	0x00,
}