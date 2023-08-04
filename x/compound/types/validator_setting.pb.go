// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compound/validator_setting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ValidatorSetting struct {
	ValidatorAddress  string                                 `protobuf:"bytes,1,opt,name=validatorAddress,proto3" json:"validatorAddress,omitempty"`
	PercentToCompound github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=percentToCompound,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"percentToCompound"`
}

func (m *ValidatorSetting) Reset()         { *m = ValidatorSetting{} }
func (m *ValidatorSetting) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetting) ProtoMessage()    {}
func (*ValidatorSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_0015669316fce38e, []int{0}
}
func (m *ValidatorSetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetting.Merge(m, src)
}
func (m *ValidatorSetting) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetting proto.InternalMessageInfo

func (m *ValidatorSetting) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidatorSetting)(nil), "temporal.compound.ValidatorSetting")
}

func init() {
	proto.RegisterFile("temporal/compound/validator_setting.proto", fileDescriptor_0015669316fce38e)
}

var fileDescriptor_0015669316fce38e = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xcd, 0x4b, 0xd1, 0x2f, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x8a, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x29, 0xd5, 0x83, 0x29, 0x95, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x4a, 0x73, 0x18, 0xb9, 0x04, 0xc2, 0x60,
	0x86, 0x04, 0x43, 0xcc, 0x10, 0xd2, 0xe2, 0x12, 0x80, 0x1b, 0xec, 0x98, 0x92, 0x52, 0x94, 0x5a,
	0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x21, 0x2e, 0x14, 0xc3, 0x25, 0x58, 0x90,
	0x5a, 0x94, 0x9c, 0x9a, 0x57, 0x12, 0x92, 0xef, 0x0c, 0xb5, 0x4b, 0x82, 0x09, 0xa4, 0xd8, 0x49,
	0xef, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0x40,
	0x4e, 0xd1, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x86, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x9e, 0x79, 0x25, 0x41, 0x98, 0x06, 0x39, 0xf9, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x11, 0x92, 0xa1, 0x30, 0xcf, 0xea, 0x56, 0xe5,
	0xe7, 0xa5, 0xc2, 0x79, 0xfa, 0x15, 0x88, 0x70, 0x02, 0x5b, 0x92, 0xc4, 0x06, 0xf6, 0xb3, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x02, 0x49, 0x41, 0x17, 0x49, 0x01, 0x00, 0x00,
}

func (m *ValidatorSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PercentToCompound.Size()
		i -= size
		if _, err := m.PercentToCompound.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidatorSetting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintValidatorSetting(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidatorSetting(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidatorSetting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovValidatorSetting(uint64(l))
	}
	l = m.PercentToCompound.Size()
	n += 1 + l + sovValidatorSetting(uint64(l))
	return n
}

func sovValidatorSetting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidatorSetting(x uint64) (n int) {
	return sovValidatorSetting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidatorSetting
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
			return fmt.Errorf("proto: ValidatorSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorSetting
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
				return ErrInvalidLengthValidatorSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PercentToCompound", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorSetting
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
				return ErrInvalidLengthValidatorSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PercentToCompound.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidatorSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidatorSetting
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
func skipValidatorSetting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidatorSetting
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
					return 0, ErrIntOverflowValidatorSetting
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
					return 0, ErrIntOverflowValidatorSetting
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
				return 0, ErrInvalidLengthValidatorSetting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidatorSetting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidatorSetting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidatorSetting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidatorSetting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidatorSetting = fmt.Errorf("proto: unexpected end of group")
)
