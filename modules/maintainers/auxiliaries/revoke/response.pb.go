// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/maintainers/auxiliaries/revoke/v1beta1/response.proto

package revoke

import (
	"errors"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type auxiliaryResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   error `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
}

func (m *auxiliaryResponse) Reset()         { *m = auxiliaryResponse{} }
func (m *auxiliaryResponse) String() string { return proto.CompactTextString(m) }
func (*auxiliaryResponse) ProtoMessage()    {}
func (*auxiliaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8ef3b7b14d62cd, []int{0}
}
func (m *auxiliaryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *auxiliaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuxiliaryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *auxiliaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuxiliaryResponse.Merge(m, src)
}
func (m *auxiliaryResponse) XXX_Size() int {
	return m.Size()
}
func (m *auxiliaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuxiliaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuxiliaryResponse proto.InternalMessageInfo

func (m *auxiliaryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*auxiliaryResponse)(nil), "persistence.modules.maintainers.auxiliaries.revoke.v1beta1.auxiliaryResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/maintainers/auxiliaries/revoke/v1beta1/response.proto", fileDescriptor_ab8ef3b7b14d62cd)
}

var fileDescriptor_ab8ef3b7b14d62cd = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x06, 0xf0, 0x5d, 0xf1, 0x6f, 0xc0, 0xc2, 0x54, 0xc1, 0x62, 0x73, 0x28, 0x48, 0xaa, 0x5d,
	0x0e, 0x3b, 0xed, 0xf2, 0x04, 0x92, 0x4a, 0x6c, 0x64, 0x2f, 0x19, 0xce, 0x85, 0xbb, 0xec, 0x32,
	0xb3, 0x11, 0xcf, 0xa7, 0xf0, 0x11, 0x7c, 0x9c, 0x2b, 0xaf, 0x14, 0x8b, 0x43, 0x92, 0xc6, 0xc7,
	0x10, 0x6e, 0x13, 0x4c, 0x77, 0xd5, 0xee, 0x4c, 0xf1, 0xfb, 0x3e, 0x26, 0x7a, 0x70, 0x68, 0xbd,
	0x55, 0x0e, 0x90, 0x0c, 0x79, 0xa8, 0x4b, 0x50, 0x9a, 0x08, 0x3c, 0xa9, 0xa5, 0x36, 0xb5, 0xd7,
	0xa6, 0x06, 0x24, 0xa5, 0x9b, 0x37, 0xb3, 0x30, 0x1a, 0x0d, 0x90, 0xaa, 0xc0, 0x35, 0xde, 0xbc,
	0x83, 0x7a, 0x9d, 0xce, 0xc0, 0xeb, 0xa9, 0x42, 0x20, 0x67, 0x6b, 0x02, 0xb9, 0xa3, 0xe2, 0xfb,
	0x91, 0x25, 0x83, 0x25, 0x47, 0x96, 0x1c, 0x59, 0x72, 0xb0, 0x64, 0x6f, 0x5d, 0xde, 0xf8, 0x17,
	0x83, 0xd5, 0xb3, 0xd3, 0xe8, 0x57, 0x2a, 0x54, 0x9b, 0xdb, 0xb9, 0xfd, 0xff, 0x85, 0x90, 0xab,
	0xc7, 0xe8, 0x62, 0x70, 0x56, 0x45, 0x9f, 0x1f, 0x27, 0xd1, 0x09, 0x35, 0x65, 0x09, 0x44, 0x09,
	0x9f, 0xf0, 0xec, 0xb4, 0x18, 0xc6, 0xf8, 0x3a, 0x3a, 0x02, 0x44, 0x8b, 0xc9, 0xc1, 0x84, 0x67,
	0x67, 0xf9, 0xf9, 0x7a, 0x9b, 0xb2, 0xef, 0x6d, 0x1a, 0x96, 0x45, 0x78, 0xee, 0x0e, 0x7f, 0x3f,
	0x53, 0x96, 0xe7, 0xeb, 0x56, 0xf0, 0x4d, 0x2b, 0xf8, 0x4f, 0x2b, 0xf8, 0x47, 0x27, 0xd8, 0xa6,
	0x13, 0xec, 0xab, 0x13, 0xec, 0x29, 0x5b, 0xda, 0xaa, 0x59, 0xc0, 0xfe, 0xcb, 0xcc, 0x8e, 0x77,
	0x25, 0x6f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x46, 0x6c, 0x66, 0x5d, 0x01, 0x00, 0x00,
}

func (m *auxiliaryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *auxiliaryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *auxiliaryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error.Error()) > 0 {
		i -= len(m.Error.Error())
		copy(dAtA[i:], m.Error.Error())
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Error.Error())))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *auxiliaryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error.Error())
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *auxiliaryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: auxiliaryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: auxiliaryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = error(errors.New(string(dAtA[iNdEx:postIndex])))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponse
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
func skipResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
				return 0, ErrInvalidLengthResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponse = fmt.Errorf("proto: unexpected end of group")
)
