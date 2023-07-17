// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/nonce.proto

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

// Nonce functions both as a retrieval key for receiving messages and a counter for sending messages
type Nonce struct {
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *Nonce) Reset()         { *m = Nonce{} }
func (m *Nonce) String() string { return proto.CompactTextString(m) }
func (*Nonce) ProtoMessage()    {}
func (*Nonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_94414231e4aaff86, []int{0}
}
func (m *Nonce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nonce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nonce.Merge(m, src)
}
func (m *Nonce) XXX_Size() int {
	return m.Size()
}
func (m *Nonce) XXX_DiscardUnknown() {
	xxx_messageInfo_Nonce.DiscardUnknown(m)
}

var xxx_messageInfo_Nonce proto.InternalMessageInfo

func (m *Nonce) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*Nonce)(nil), "circle.cctp.v1.Nonce")
}

func init() { proto.RegisterFile("circle/cctp/v1/nonce.proto", fileDescriptor_94414231e4aaff86) }

var fileDescriptor_94414231e4aaff86 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0x2c, 0x4a,
	0xce, 0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0xcf, 0xcb, 0xcf, 0x4b, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xc8, 0xe9, 0x81, 0xe4, 0xf4, 0xca, 0x0c,
	0x95, 0x64, 0xb9, 0x58, 0xfd, 0x40, 0xd2, 0x42, 0x22, 0x5c, 0xac, 0x60, 0x75, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x2c, 0x41, 0x10, 0x8e, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0xcc,
	0x4c, 0xcb, 0xcc, 0xd3, 0xcf, 0xcb, 0x4f, 0xca, 0x49, 0xd5, 0x05, 0x5b, 0x5c, 0x01, 0xb1, 0xbf,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x6c, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x69, 0x1f, 0xb0, 0x9b, 0x00, 0x00, 0x00,
}

func (m *Nonce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nonce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nonce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintNonce(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNonce(dAtA []byte, offset int, v uint64) int {
	offset -= sovNonce(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Nonce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovNonce(uint64(m.Nonce))
	}
	return n
}

func sovNonce(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNonce(x uint64) (n int) {
	return sovNonce(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Nonce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonce
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
			return fmt.Errorf("proto: Nonce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nonce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonce
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
func skipNonce(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNonce
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
					return 0, ErrIntOverflowNonce
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
					return 0, ErrIntOverflowNonce
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
				return 0, ErrInvalidLengthNonce
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNonce
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNonce
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNonce        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNonce          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNonce = fmt.Errorf("proto: unexpected end of group")
)
