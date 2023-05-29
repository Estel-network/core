// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: igmf/otc/module_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ModuleInfo struct {
	ModuleIndex uint64 `protobuf:"varint,1,opt,name=moduleIndex,proto3" json:"moduleIndex,omitempty"`
	ServiceFee  int32  `protobuf:"varint,2,opt,name=serviceFee,proto3" json:"serviceFee,omitempty"`
}

func (m *ModuleInfo) Reset()         { *m = ModuleInfo{} }
func (m *ModuleInfo) String() string { return proto.CompactTextString(m) }
func (*ModuleInfo) ProtoMessage()    {}
func (*ModuleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e894eab564a41e8, []int{0}
}
func (m *ModuleInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleInfo.Merge(m, src)
}
func (m *ModuleInfo) XXX_Size() int {
	return m.Size()
}
func (m *ModuleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleInfo proto.InternalMessageInfo

func (m *ModuleInfo) GetModuleIndex() uint64 {
	if m != nil {
		return m.ModuleIndex
	}
	return 0
}

func (m *ModuleInfo) GetServiceFee() int32 {
	if m != nil {
		return m.ServiceFee
	}
	return 0
}

func init() {
	proto.RegisterType((*ModuleInfo)(nil), "igmf.otc.ModuleInfo")
}

func init() { proto.RegisterFile("igmf/otc/module_info.proto", fileDescriptor_8e894eab564a41e8) }

var fileDescriptor_8e894eab564a41e8 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4c, 0xcf, 0x4d,
	0xd3, 0xcf, 0x2f, 0x49, 0xd6, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x8d, 0xcf, 0xcc, 0x4b, 0xcb,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe9, 0xe5, 0x97, 0x24, 0x2b, 0xf9,
	0x71, 0x71, 0xf9, 0x82, 0xa5, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0x14, 0xb8, 0xb8, 0x73, 0xa1, 0xbc,
	0x94, 0xd4, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x64, 0x21, 0x21, 0x39, 0x2e, 0xae,
	0xe2, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xb7, 0xd4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x24, 0x11, 0x27, 0xad, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12,
	0x00, 0xbb, 0xa7, 0x02, 0xec, 0xa2, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x63, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x3c, 0xef, 0x8b, 0xaa, 0x00, 0x00, 0x00,
}

func (m *ModuleInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ServiceFee != 0 {
		i = encodeVarintModuleInfo(dAtA, i, uint64(m.ServiceFee))
		i--
		dAtA[i] = 0x10
	}
	if m.ModuleIndex != 0 {
		i = encodeVarintModuleInfo(dAtA, i, uint64(m.ModuleIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModuleInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovModuleInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ModuleInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ModuleIndex != 0 {
		n += 1 + sovModuleInfo(uint64(m.ModuleIndex))
	}
	if m.ServiceFee != 0 {
		n += 1 + sovModuleInfo(uint64(m.ServiceFee))
	}
	return n
}

func sovModuleInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModuleInfo(x uint64) (n int) {
	return sovModuleInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ModuleInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleInfo
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
			return fmt.Errorf("proto: ModuleInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleIndex", wireType)
			}
			m.ModuleIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModuleIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFee", wireType)
			}
			m.ServiceFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceFee |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModuleInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleInfo
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
func skipModuleInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModuleInfo
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
					return 0, ErrIntOverflowModuleInfo
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
					return 0, ErrIntOverflowModuleInfo
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
				return 0, ErrInvalidLengthModuleInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModuleInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModuleInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModuleInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModuleInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModuleInfo = fmt.Errorf("proto: unexpected end of group")
)