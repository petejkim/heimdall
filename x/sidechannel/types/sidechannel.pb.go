// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sidechannel/v1beta1/sidechannel.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/abci/types"
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

type PreviousValidators struct {
	Height     uint64             `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Validators []*types.Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *PreviousValidators) Reset()         { *m = PreviousValidators{} }
func (m *PreviousValidators) String() string { return proto.CompactTextString(m) }
func (*PreviousValidators) ProtoMessage()    {}
func (*PreviousValidators) Descriptor() ([]byte, []int) {
	return fileDescriptor_af090c9ef441e24d, []int{0}
}
func (m *PreviousValidators) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousValidators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousValidators.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousValidators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousValidators.Merge(m, src)
}
func (m *PreviousValidators) XXX_Size() int {
	return m.Size()
}
func (m *PreviousValidators) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousValidators.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousValidators proto.InternalMessageInfo

func (m *PreviousValidators) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PreviousValidators) GetValidators() []*types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func init() {
	proto.RegisterType((*PreviousValidators)(nil), "heimdall.sidechannel.v1beta1.PreviousValidators")
}

func init() {
	proto.RegisterFile("sidechannel/v1beta1/sidechannel.proto", fileDescriptor_af090c9ef441e24d)
}

var fileDescriptor_af090c9ef441e24d = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xce, 0x4c, 0x49,
	0x4d, 0xce, 0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x47, 0x12, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc9, 0x48, 0xcd, 0xcc, 0x4d, 0x49,
	0xcc, 0xc9, 0xd1, 0x43, 0x96, 0x83, 0xaa, 0x97, 0x92, 0x2e, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca,
	0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x4c, 0x4a, 0xce, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x68,
	0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x52, 0x06, 0x97,
	0x50, 0x40, 0x51, 0x6a, 0x59, 0x66, 0x7e, 0x69, 0x71, 0x58, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49,
	0x7e, 0x51, 0xb1, 0x90, 0x18, 0x17, 0x5b, 0x46, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x4b, 0x10, 0x94, 0x27, 0x64, 0xc5, 0xc5, 0x55, 0x06, 0x57, 0x25, 0xc1, 0xa4, 0xc0,
	0xac, 0xc1, 0x6d, 0x24, 0xa5, 0x87, 0xb0, 0x55, 0x0f, 0x64, 0xab, 0x1e, 0xdc, 0xa0, 0x20, 0x24,
	0xd5, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9b, 0x58, 0x92, 0x99, 0x9c, 0x97,
	0x5a, 0x52, 0x9e, 0x5f, 0x94, 0xad, 0x0f, 0xf3, 0xac, 0x7e, 0x05, 0x72, 0x50, 0x40, 0x7c, 0x95,
	0xc4, 0x06, 0xf6, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0x2d, 0xdc, 0x12, 0x3a, 0x01,
	0x00, 0x00,
}

func (m *PreviousValidators) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousValidators) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousValidators) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSidechannel(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintSidechannel(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSidechannel(dAtA []byte, offset int, v uint64) int {
	offset -= sovSidechannel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousValidators) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovSidechannel(uint64(m.Height))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovSidechannel(uint64(l))
		}
	}
	return n
}

func sovSidechannel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSidechannel(x uint64) (n int) {
	return sovSidechannel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousValidators) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSidechannel
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
			return fmt.Errorf("proto: PreviousValidators: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousValidators: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSidechannel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSidechannel
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
				return ErrInvalidLengthSidechannel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSidechannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &types.Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSidechannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSidechannel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSidechannel
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
func skipSidechannel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSidechannel
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
					return 0, ErrIntOverflowSidechannel
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
					return 0, ErrIntOverflowSidechannel
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
				return 0, ErrInvalidLengthSidechannel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSidechannel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSidechannel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSidechannel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSidechannel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSidechannel = fmt.Errorf("proto: unexpected end of group")
)
