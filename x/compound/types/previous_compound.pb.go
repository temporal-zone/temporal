// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compound/previous_compound.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PreviousCompound struct {
	Delegator string    `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Timestamp time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *PreviousCompound) Reset()         { *m = PreviousCompound{} }
func (m *PreviousCompound) String() string { return proto.CompactTextString(m) }
func (*PreviousCompound) ProtoMessage()    {}
func (*PreviousCompound) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a0f2847d682b66c, []int{0}
}
func (m *PreviousCompound) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousCompound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousCompound.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousCompound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousCompound.Merge(m, src)
}
func (m *PreviousCompound) XXX_Size() int {
	return m.Size()
}
func (m *PreviousCompound) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousCompound.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousCompound proto.InternalMessageInfo

func (m *PreviousCompound) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *PreviousCompound) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*PreviousCompound)(nil), "temporal.compound.PreviousCompound")
}

func init() {
	proto.RegisterFile("temporal/compound/previous_compound.proto", fileDescriptor_2a0f2847d682b66c)
}

var fileDescriptor_2a0f2847d682b66c = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xcd, 0x4b, 0xd1, 0x2f, 0x28,
	0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0x87, 0x89, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0x09, 0xc2, 0x94, 0xea, 0xc1, 0x24, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20,
	0x16, 0x44, 0xa1, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a,
	0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x00, 0x55, 0x20, 0x99, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0xd1, 0x09, 0xe1, 0x40, 0xa5, 0xc4, 0x21, 0x3c, 0xfd, 0xdc, 0xe2,
	0x74, 0xfd, 0x32, 0x43, 0x10, 0x05, 0x91, 0x50, 0x9a, 0xc7, 0xc8, 0x25, 0x10, 0x00, 0x75, 0x99,
	0x33, 0xd4, 0x7e, 0x21, 0x33, 0x2e, 0xce, 0x94, 0xd4, 0x9c, 0xd4, 0xf4, 0xc4, 0x92, 0xfc, 0x22,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x89, 0x4b, 0x5b, 0x74, 0x45, 0xa0, 0x46, 0x3a, 0xa6,
	0xa4, 0x14, 0xa5, 0x16, 0x17, 0x07, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x07, 0x21, 0x94, 0x0a, 0x39,
	0x71, 0x71, 0xc2, 0xdd, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x71, 0xb5,
	0x1e, 0xcc, 0xd5, 0x7a, 0x21, 0x30, 0x15, 0x4e, 0x1c, 0x27, 0xee, 0xc9, 0x33, 0x4c, 0xb8, 0x2f,
	0xcf, 0x18, 0x84, 0xd0, 0x66, 0xc5, 0xd7, 0xf4, 0x7c, 0x83, 0x16, 0xc2, 0x4c, 0x27, 0x9f, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0x02, 0x05, 0x9b, 0x3e, 0x2c, 0x0c, 0x75, 0xab, 0xf2, 0xf3, 0x52, 0xe1, 0x3c, 0xfd, 0x0a, 0x44,
	0xf0, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x9d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0x60, 0xbd, 0x2a, 0xa0, 0x01, 0x00, 0x00,
}

func (m *PreviousCompound) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousCompound) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousCompound) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPreviousCompound(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintPreviousCompound(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousCompound(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousCompound(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousCompound) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovPreviousCompound(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovPreviousCompound(uint64(l))
	return n
}

func sovPreviousCompound(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousCompound(x uint64) (n int) {
	return sovPreviousCompound(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousCompound) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousCompound
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
			return fmt.Errorf("proto: PreviousCompound: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousCompound: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousCompound
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
				return ErrInvalidLengthPreviousCompound
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPreviousCompound
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousCompound
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
				return ErrInvalidLengthPreviousCompound
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPreviousCompound
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousCompound(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPreviousCompound
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPreviousCompound(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousCompound
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
					return 0, ErrIntOverflowPreviousCompound
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
					return 0, ErrIntOverflowPreviousCompound
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
				return 0, ErrInvalidLengthPreviousCompound
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreviousCompound
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreviousCompound
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreviousCompound        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousCompound          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreviousCompound = fmt.Errorf("proto: unexpected end of group")
)