// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/record/user_instruction.proto

package types

import (
	fmt "fmt"
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

type UserInstruction struct {
	LocalAddress    string    `protobuf:"bytes,1,opt,name=localAddress,proto3" json:"localAddress,omitempty"`
	RemoteAddress   string    `protobuf:"bytes,2,opt,name=remoteAddress,proto3" json:"remoteAddress,omitempty"`
	ChainId         string    `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Frequency       int64     `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Created         time.Time `protobuf:"bytes,5,opt,name=created,proto3,stdtime" json:"created"`
	Expires         time.Time `protobuf:"bytes,6,opt,name=expires,proto3,stdtime" json:"expires"`
	Instruction     string    `protobuf:"bytes,7,opt,name=instruction,proto3" json:"instruction,omitempty"`
	StrategyId      int64     `protobuf:"varint,8,opt,name=strategyId,proto3" json:"strategyId,omitempty"`
	ContractAddress string    `protobuf:"bytes,9,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (m *UserInstruction) Reset()         { *m = UserInstruction{} }
func (m *UserInstruction) String() string { return proto.CompactTextString(m) }
func (*UserInstruction) ProtoMessage()    {}
func (*UserInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c8d50a40dcaef3, []int{0}
}
func (m *UserInstruction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserInstruction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInstruction.Merge(m, src)
}
func (m *UserInstruction) XXX_Size() int {
	return m.Size()
}
func (m *UserInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_UserInstruction proto.InternalMessageInfo

func (m *UserInstruction) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *UserInstruction) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *UserInstruction) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *UserInstruction) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *UserInstruction) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *UserInstruction) GetExpires() time.Time {
	if m != nil {
		return m.Expires
	}
	return time.Time{}
}

func (m *UserInstruction) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *UserInstruction) GetStrategyId() int64 {
	if m != nil {
		return m.StrategyId
	}
	return 0
}

func (m *UserInstruction) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInstruction)(nil), "temporal.record.UserInstruction")
}

func init() {
	proto.RegisterFile("temporal/record/user_instruction.proto", fileDescriptor_b5c8d50a40dcaef3)
}

var fileDescriptor_b5c8d50a40dcaef3 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0xc6, 0xe3, 0x0b, 0x97, 0x3f, 0xe6, 0x5e, 0x21, 0x59, 0x1d, 0x2c, 0x54, 0x25, 0x11, 0x42,
	0x15, 0x4b, 0x93, 0xaa, 0xdd, 0x2b, 0x95, 0x8d, 0x8e, 0x51, 0xbb, 0x74, 0xa9, 0x82, 0x73, 0x08,
	0x91, 0x92, 0x38, 0xb5, 0x1d, 0x09, 0xfa, 0x0a, 0x5d, 0x78, 0x2c, 0x46, 0xc6, 0x4e, 0xb4, 0x82,
	0x17, 0xa9, 0x92, 0xe0, 0x02, 0xdd, 0xba, 0xf9, 0x7c, 0xe7, 0xf7, 0xd9, 0xfe, 0x8e, 0x8d, 0x2f,
	0x14, 0x24, 0x19, 0x17, 0x7e, 0xec, 0x0a, 0x60, 0x5c, 0x04, 0x6e, 0x2e, 0x41, 0x3c, 0x47, 0xa9,
	0x54, 0x22, 0x67, 0x2a, 0xe2, 0xa9, 0x93, 0x09, 0xae, 0x38, 0xe9, 0x6a, 0xce, 0xa9, 0xb8, 0xde,
	0x59, 0xc8, 0x43, 0x5e, 0xf6, 0xdc, 0x62, 0x55, 0x61, 0x3d, 0x2b, 0xe4, 0x3c, 0x8c, 0xc1, 0x2d,
	0xab, 0x49, 0x3e, 0x75, 0x55, 0x94, 0x80, 0x54, 0x7e, 0x92, 0x55, 0x40, 0xff, 0xad, 0x86, 0xbb,
	0x8f, 0x12, 0xc4, 0xf8, 0x70, 0x02, 0xe9, 0xe3, 0x7f, 0x31, 0x67, 0x7e, 0x7c, 0x17, 0x04, 0x02,
	0xa4, 0xa4, 0xc8, 0x46, 0xc3, 0xb6, 0x77, 0xa2, 0x91, 0x01, 0xfe, 0x2f, 0x20, 0xe1, 0x0a, 0x34,
	0xf4, 0xa7, 0x84, 0x4e, 0x45, 0x42, 0x71, 0x93, 0xcd, 0xfc, 0x28, 0x1d, 0x07, 0xb4, 0x56, 0xf6,
	0x75, 0x49, 0xce, 0x71, 0x7b, 0x2a, 0xe0, 0x25, 0x87, 0x94, 0x2d, 0x68, 0xdd, 0x46, 0xc3, 0x9a,
	0x77, 0x10, 0xc8, 0x2d, 0x6e, 0x32, 0x01, 0xbe, 0x82, 0x80, 0xfe, 0xb5, 0xd1, 0xb0, 0x73, 0xdd,
	0x73, 0xaa, 0x20, 0x8e, 0x0e, 0xe2, 0x3c, 0xe8, 0x20, 0xa3, 0xd6, 0x6a, 0x63, 0x19, 0xcb, 0x0f,
	0x0b, 0x79, 0xda, 0x54, 0xf8, 0x61, 0x9e, 0x45, 0x02, 0x24, 0x6d, 0xfc, 0xc6, 0xbf, 0x37, 0x11,
	0x1b, 0x77, 0x8e, 0x46, 0x4e, 0x9b, 0xe5, 0xdd, 0x8f, 0x25, 0x32, 0xc0, 0x58, 0x2a, 0xe1, 0x2b,
	0x08, 0x17, 0xe3, 0x80, 0xb6, 0x8a, 0x00, 0xa3, 0xfa, 0x6a, 0x63, 0x21, 0xef, 0x48, 0x27, 0x0e,
	0xee, 0x32, 0x9e, 0x2a, 0xe1, 0x33, 0xa5, 0xe7, 0xd4, 0x2e, 0xf6, 0xda, 0xa3, 0x3f, 0x9b, 0xa3,
	0xfb, 0xd5, 0xd6, 0x44, 0xeb, 0xad, 0x89, 0x3e, 0xb7, 0x26, 0x5a, 0xee, 0x4c, 0x63, 0xbd, 0x33,
	0x8d, 0xf7, 0x9d, 0x69, 0x3c, 0x5d, 0x85, 0x91, 0x9a, 0xe5, 0x13, 0x87, 0xf1, 0xc4, 0xd5, 0x4f,
	0x7f, 0xf9, 0xca, 0x53, 0xf8, 0xae, 0xdc, 0xb9, 0xfe, 0x32, 0x6a, 0x91, 0x81, 0x9c, 0x34, 0xca,
	0xa8, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x43, 0x9d, 0x76, 0x52, 0x02, 0x00, 0x00,
}

func (m *UserInstruction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserInstruction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserInstruction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintUserInstruction(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if m.StrategyId != 0 {
		i = encodeVarintUserInstruction(dAtA, i, uint64(m.StrategyId))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Instruction) > 0 {
		i -= len(m.Instruction)
		copy(dAtA[i:], m.Instruction)
		i = encodeVarintUserInstruction(dAtA, i, uint64(len(m.Instruction)))
		i--
		dAtA[i] = 0x3a
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expires, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expires):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintUserInstruction(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Created):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintUserInstruction(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if m.Frequency != 0 {
		i = encodeVarintUserInstruction(dAtA, i, uint64(m.Frequency))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintUserInstruction(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RemoteAddress) > 0 {
		i -= len(m.RemoteAddress)
		copy(dAtA[i:], m.RemoteAddress)
		i = encodeVarintUserInstruction(dAtA, i, uint64(len(m.RemoteAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LocalAddress) > 0 {
		i -= len(m.LocalAddress)
		copy(dAtA[i:], m.LocalAddress)
		i = encodeVarintUserInstruction(dAtA, i, uint64(len(m.LocalAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserInstruction(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserInstruction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserInstruction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LocalAddress)
	if l > 0 {
		n += 1 + l + sovUserInstruction(uint64(l))
	}
	l = len(m.RemoteAddress)
	if l > 0 {
		n += 1 + l + sovUserInstruction(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovUserInstruction(uint64(l))
	}
	if m.Frequency != 0 {
		n += 1 + sovUserInstruction(uint64(m.Frequency))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovUserInstruction(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expires)
	n += 1 + l + sovUserInstruction(uint64(l))
	l = len(m.Instruction)
	if l > 0 {
		n += 1 + l + sovUserInstruction(uint64(l))
	}
	if m.StrategyId != 0 {
		n += 1 + sovUserInstruction(uint64(m.StrategyId))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovUserInstruction(uint64(l))
	}
	return n
}

func sovUserInstruction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserInstruction(x uint64) (n int) {
	return sovUserInstruction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserInstruction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserInstruction
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
			return fmt.Errorf("proto: UserInstruction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserInstruction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			m.Frequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frequency |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Expires, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instruction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instruction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrategyId", wireType)
			}
			m.StrategyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StrategyId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserInstruction
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
				return ErrInvalidLengthUserInstruction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserInstruction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserInstruction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserInstruction
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
func skipUserInstruction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserInstruction
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
					return 0, ErrIntOverflowUserInstruction
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
					return 0, ErrIntOverflowUserInstruction
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
				return 0, ErrInvalidLengthUserInstruction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserInstruction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserInstruction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserInstruction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserInstruction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserInstruction = fmt.Errorf("proto: unexpected end of group")
)
