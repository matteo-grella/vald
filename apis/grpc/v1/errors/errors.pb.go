//
// Copyright (C) 2019-2020 Vdaas.org Vald team ( kpango, rinx, kmrmt )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errors

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

type Errors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Errors) Reset()         { *m = Errors{} }
func (m *Errors) String() string { return proto.CompactTextString(m) }
func (*Errors) ProtoMessage()    {}
func (*Errors) Descriptor() ([]byte, []int) {
	return fileDescriptor_3da6f8f1f37746ee, []int{0}
}
func (m *Errors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Errors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Errors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Errors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Errors.Merge(m, src)
}
func (m *Errors) XXX_Size() int {
	return m.Size()
}
func (m *Errors) XXX_DiscardUnknown() {
	xxx_messageInfo_Errors.DiscardUnknown(m)
}

var xxx_messageInfo_Errors proto.InternalMessageInfo

type Errors_RPC struct {
	Type                 string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Details              []string      `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	Instance             string        `protobuf:"bytes,4,opt,name=instance,proto3" json:"instance,omitempty"`
	Status               int64         `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Error                string        `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Roots                []*Errors_RPC `protobuf:"bytes,7,rep,name=roots,proto3" json:"roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Errors_RPC) Reset()         { *m = Errors_RPC{} }
func (m *Errors_RPC) String() string { return proto.CompactTextString(m) }
func (*Errors_RPC) ProtoMessage()    {}
func (*Errors_RPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_3da6f8f1f37746ee, []int{0, 0}
}
func (m *Errors_RPC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Errors_RPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Errors_RPC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Errors_RPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Errors_RPC.Merge(m, src)
}
func (m *Errors_RPC) XXX_Size() int {
	return m.Size()
}
func (m *Errors_RPC) XXX_DiscardUnknown() {
	xxx_messageInfo_Errors_RPC.DiscardUnknown(m)
}

var xxx_messageInfo_Errors_RPC proto.InternalMessageInfo

func (m *Errors_RPC) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Errors_RPC) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Errors_RPC) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Errors_RPC) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Errors_RPC) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Errors_RPC) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Errors_RPC) GetRoots() []*Errors_RPC {
	if m != nil {
		return m.Roots
	}
	return nil
}

func init() {
	proto.RegisterType((*Errors)(nil), "errors.Errors")
	proto.RegisterType((*Errors_RPC)(nil), "errors.Errors.RPC")
}

func init() { proto.RegisterFile("apis/proto/v1/errors/errors.proto", fileDescriptor_3da6f8f1f37746ee) }

var fileDescriptor_3da6f8f1f37746ee = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xc9, 0x37, 0xed, 0xf4, 0xeb, 0x75, 0x23, 0x41, 0x24, 0x14, 0x29, 0xa3, 0xab, 0x71,
	0xd1, 0x0c, 0xd5, 0x27, 0xb0, 0xc5, 0x7d, 0xc9, 0xc2, 0x85, 0xbb, 0x74, 0x12, 0xc6, 0xc0, 0x74,
	0x12, 0x92, 0x34, 0xd8, 0x47, 0x73, 0xef, 0xc2, 0xa5, 0x8f, 0x20, 0xf3, 0x24, 0xd2, 0x64, 0xfc,
	0x03, 0xae, 0x72, 0xce, 0x25, 0xbf, 0x43, 0x72, 0x2e, 0x5c, 0x72, 0xa3, 0x5c, 0x65, 0xac, 0xf6,
	0xba, 0x0a, 0xcb, 0x4a, 0x5a, 0xab, 0xad, 0x1b, 0x0e, 0x1a, 0xc7, 0x38, 0x4f, 0x6e, 0x76, 0xd7,
	0x28, 0xff, 0xb4, 0xdf, 0xd2, 0x5a, 0xef, 0x2a, 0xd9, 0x05, 0x7d, 0x30, 0x56, 0x3f, 0x1f, 0x12,
	0x5b, 0x2f, 0x1a, 0xd9, 0x2d, 0x02, 0x6f, 0x95, 0xe0, 0x5e, 0x56, 0x7f, 0x44, 0x8a, 0xba, 0x7a,
	0x45, 0x90, 0xdf, 0xa7, 0xb4, 0x17, 0x04, 0x19, 0xdb, 0xac, 0x31, 0x86, 0x91, 0x3f, 0x18, 0x49,
	0x50, 0x81, 0xca, 0x29, 0x8b, 0x1a, 0x9f, 0x42, 0xb6, 0x73, 0x0d, 0xf9, 0x17, 0x47, 0x47, 0x89,
	0x09, 0x4c, 0x84, 0xf4, 0x5c, 0xb5, 0x8e, 0x64, 0x45, 0x56, 0x4e, 0xd9, 0x97, 0xc5, 0x33, 0xf8,
	0xaf, 0x3a, 0xe7, 0x79, 0x57, 0x4b, 0x32, 0x8a, 0xc0, 0xb7, 0xc7, 0xe7, 0x90, 0x3b, 0xcf, 0xfd,
	0xde, 0x91, 0x71, 0x81, 0xca, 0x8c, 0x0d, 0x0e, 0x9f, 0xc1, 0x38, 0xfe, 0x89, 0xe4, 0x11, 0x48,
	0x06, 0x97, 0x30, 0xb6, 0x5a, 0x7b, 0x47, 0x26, 0x45, 0x56, 0x9e, 0xdc, 0x60, 0x3a, 0xb4, 0x90,
	0x1e, 0x4c, 0xd9, 0x66, 0xcd, 0xd2, 0x85, 0x55, 0xfd, 0xd6, 0xcf, 0xd1, 0x7b, 0x3f, 0x47, 0x1f,
	0xfd, 0x1c, 0xc1, 0x85, 0xb6, 0x0d, 0x0d, 0x82, 0x73, 0x47, 0x03, 0x6f, 0x05, 0xe5, 0x46, 0xd1,
	0xb0, 0x1c, 0xe8, 0x15, 0x3c, 0xf0, 0x56, 0xa4, 0x88, 0x0d, 0x7a, 0xbc, 0xfe, 0xd5, 0x61, 0x44,
	0x8e, 0x0d, 0x89, 0x2a, 0x2e, 0xa1, 0xb1, 0xa6, 0xfe, 0xd9, 0xc1, 0x36, 0x8f, 0x95, 0xdd, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xde, 0x2b, 0xc1, 0xbd, 0xa2, 0x01, 0x00, 0x00,
}

func (m *Errors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Errors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Errors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *Errors_RPC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Errors_RPC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Errors_RPC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Roots) > 0 {
		for iNdEx := len(m.Roots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Roots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintErrors(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Instance) > 0 {
		i -= len(m.Instance)
		copy(dAtA[i:], m.Instance)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Instance)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Details) > 0 {
		for iNdEx := len(m.Details) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Details[iNdEx])
			copy(dAtA[i:], m.Details[iNdEx])
			i = encodeVarintErrors(dAtA, i, uint64(len(m.Details[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Errors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Errors_RPC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if len(m.Details) > 0 {
		for _, s := range m.Details {
			l = len(s)
			n += 1 + l + sovErrors(uint64(l))
		}
	}
	l = len(m.Instance)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovErrors(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if len(m.Roots) > 0 {
		for _, e := range m.Roots {
			l = e.Size()
			n += 1 + l + sovErrors(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Errors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Errors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Errors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Errors_RPC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RPC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RPC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = append(m.Details, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roots = append(m.Roots, &Errors_RPC{})
			if err := m.Roots[len(m.Roots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)