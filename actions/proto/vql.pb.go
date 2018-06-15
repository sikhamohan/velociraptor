// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vql.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VQLRequest struct {
	VQL  string `protobuf:"bytes,1,opt,name=VQL" json:"VQL,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *VQLRequest) Reset()                    { *m = VQLRequest{} }
func (m *VQLRequest) String() string            { return proto1.CompactTextString(m) }
func (*VQLRequest) ProtoMessage()               {}
func (*VQLRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *VQLRequest) GetVQL() string {
	if m != nil {
		return m.VQL
	}
	return ""
}

func (m *VQLRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VQLCollectorArgs struct {
	Query []*VQLRequest `protobuf:"bytes,2,rep,name=Query" json:"Query,omitempty"`
}

func (m *VQLCollectorArgs) Reset()                    { *m = VQLCollectorArgs{} }
func (m *VQLCollectorArgs) String() string            { return proto1.CompactTextString(m) }
func (*VQLCollectorArgs) ProtoMessage()               {}
func (*VQLCollectorArgs) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *VQLCollectorArgs) GetQuery() []*VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

type VQLResponse struct {
	Response string      `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
	Columns  []string    `protobuf:"bytes,2,rep,name=Columns" json:"Columns,omitempty"`
	Query    *VQLRequest `protobuf:"bytes,3,opt,name=Query" json:"Query,omitempty"`
}

func (m *VQLResponse) Reset()                    { *m = VQLResponse{} }
func (m *VQLResponse) String() string            { return proto1.CompactTextString(m) }
func (*VQLResponse) ProtoMessage()               {}
func (*VQLResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *VQLResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *VQLResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *VQLResponse) GetQuery() *VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

type ClientInfo struct {
	Info []*VQLResponse `protobuf:"bytes,1,rep,name=info" json:"info,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto1.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ClientInfo) GetInfo() []*VQLResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto1.RegisterType((*VQLRequest)(nil), "proto.VQLRequest")
	proto1.RegisterType((*VQLCollectorArgs)(nil), "proto.VQLCollectorArgs")
	proto1.RegisterType((*VQLResponse)(nil), "proto.VQLResponse")
	proto1.RegisterType((*ClientInfo)(nil), "proto.ClientInfo")
}

func init() { proto1.RegisterFile("vql.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x95, 0xfe, 0xf9, 0x3e, 0x7a, 0xbb, 0x14, 0x0f, 0x28, 0x62, 0x40, 0x56, 0x07, 0x28,
	0x42, 0x4a, 0x44, 0x61, 0xea, 0x46, 0x3b, 0x81, 0xa2, 0xa2, 0x04, 0x94, 0x81, 0x2d, 0xb8, 0xb7,
	0x6d, 0x24, 0xc7, 0xb7, 0xb5, 0x9d, 0xb6, 0xbc, 0x03, 0x8f, 0xc5, 0x93, 0xc0, 0x6b, 0x30, 0x20,
	0x1c, 0x02, 0x0c, 0x4c, 0x3e, 0xbe, 0xf7, 0xf8, 0xe7, 0x73, 0xa0, 0xb3, 0x59, 0xcb, 0x60, 0xa5,
	0xc9, 0x12, 0x6b, 0xbb, 0xe3, 0x70, 0xb4, 0xdd, 0x6e, 0x83, 0x0d, 0x4a, 0x12, 0xf9, 0x0c, 0x77,
	0x81, 0xa0, 0x22, 0x5c, 0x90, 0xcc, 0xd4, 0x22, 0xac, 0x86, 0x3a, 0x5b, 0x59, 0xd2, 0xa1, 0x33,
	0x87, 0x06, 0x8b, 0x4c, 0xd9, 0x5c, 0x54, 0x88, 0xfe, 0x10, 0x20, 0x8d, 0xa3, 0x04, 0xd7, 0x25,
	0x1a, 0xcb, 0x7a, 0xd0, 0x4c, 0xe3, 0xc8, 0xf7, 0xb8, 0x37, 0xe8, 0x24, 0x9f, 0x92, 0x31, 0x68,
	0x4d, 0xb3, 0x02, 0xfd, 0x86, 0x1b, 0x39, 0xdd, 0x57, 0xd0, 0x4b, 0xe3, 0x68, 0x42, 0x52, 0xa2,
	0xb0, 0xa4, 0xaf, 0xf4, 0xc2, 0xb0, 0x07, 0x68, 0xc7, 0x25, 0xea, 0x27, 0xbf, 0xc1, 0x9b, 0x83,
	0xee, 0x70, 0xbf, 0xc2, 0x07, 0x3f, 0xec, 0xf1, 0xf9, 0xeb, 0xfb, 0xdb, 0x8b, 0x77, 0xc6, 0x4e,
	0xef, 0x97, 0xc8, 0xd3, 0x38, 0xe2, 0xeb, 0x12, 0x75, 0x8e, 0x86, 0x5b, 0xe2, 0xb8, 0x43, 0x51,
	0x5a, 0xe4, 0xa4, 0xb8, 0x5d, 0x22, 0x17, 0x32, 0x47, 0x65, 0x83, 0xa4, 0x42, 0xf6, 0x9f, 0x3d,
	0xe8, 0x3a, 0x90, 0x59, 0x91, 0x32, 0xc8, 0x46, 0xb0, 0x57, 0xeb, 0x2a, 0xea, 0xf8, 0xc8, 0xb1,
	0x7d, 0x76, 0x70, 0x73, 0x77, 0x3b, 0xe5, 0xa8, 0x04, 0xcd, 0x70, 0xc6, 0xf5, 0x97, 0x29, 0x48,
	0xbe, 0xfd, 0xcc, 0x87, 0xff, 0x13, 0x92, 0x65, 0xa1, 0x8c, 0x4b, 0xda, 0x49, 0xea, 0x2b, 0x3b,
	0xa9, 0x1b, 0x34, 0xb9, 0xf7, 0x67, 0x83, 0x3a, 0xce, 0x25, 0xc0, 0xc4, 0x05, 0xbc, 0x56, 0x73,
	0x62, 0xc7, 0xd0, 0xca, 0xd5, 0x9c, 0x7c, 0xcf, 0xf5, 0x66, 0xbf, 0x5f, 0x55, 0x5f, 0x26, 0x6e,
	0xff, 0xf8, 0xcf, 0x2d, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0x3f, 0x4c, 0xab, 0xbf,
	0x01, 0x00, 0x00,
}
