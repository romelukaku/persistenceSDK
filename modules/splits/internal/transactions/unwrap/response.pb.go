// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence/assets/splits/transactions/unwrap/v1beta1/response.proto

package unwrap

import (
	"errors"
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

type transactionResponse struct {
	success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	error   error `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
}

func (m *transactionResponse) Reset()         { *m = transactionResponse{} }
func (m *transactionResponse) String() string { return proto.CompactTextString(m) }
func (*transactionResponse) ProtoMessage()    {}
func (*transactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0619cd869489ec0, []int{0}
}
func (m *transactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *transactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *transactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *transactionResponse) GetSuccess() bool {
	if m != nil {
		return m.success
	}
	return false
}

func init() {
	proto.RegisterType((*transactionResponse)(nil), "persistence.assets.splits.transactions.unwrap.v1beta1.transactionResponse")
}

func init() {
	proto.RegisterFile("persistence/assets/splits/transactions/unwrap/v1beta1/response.proto", fileDescriptor_a0619cd869489ec0)
}

var fileDescriptor_a0619cd869489ec0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0x2f,
	0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc,
	0xcf, 0x2b, 0xd6, 0x2f, 0xcd, 0x2b, 0x2f, 0x4a, 0x2c, 0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x32, 0x45, 0x32, 0x45, 0x0f, 0x62, 0x8a, 0x1e, 0xc4, 0x14, 0x3d, 0x64, 0x53, 0xf4, 0x20,
	0xa6, 0xe8, 0x41, 0x4d, 0x91, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0x1b, 0xa2,
	0x0f, 0xe1, 0x40, 0x4c, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x87, 0x88, 0x83, 0x58, 0x10, 0x51,
	0xa5, 0x42, 0x2e, 0x61, 0x24, 0xf3, 0x82, 0xa0, 0x8e, 0x10, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d,
	0x4e, 0x4e, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x85, 0x94, 0xb9,
	0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x78, 0x4f, 0xdc,
	0x93, 0x67, 0xb8, 0x75, 0x4f, 0x1e, 0x22, 0x18, 0x04, 0xa1, 0xac, 0xa4, 0x5f, 0x2c, 0x90, 0x67,
	0xb8, 0xb4, 0x45, 0x57, 0x38, 0x04, 0xd3, 0x6c, 0x27, 0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0xd2, 0xce, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x85, 0x87, 0x57, 0x66, 0x5e,
	0x49, 0x6a, 0x51, 0x5e, 0x62, 0x0e, 0xb6, 0x80, 0x4b, 0x62, 0x03, 0x7b, 0xc0, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x91, 0x6a, 0xeb, 0x70, 0x01, 0x00, 0x00,
}

func (m *transactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.error.Error()) > 0 {
		i -= len(m.error.Error())
		copy(dAtA[i:], m.error.Error())
		i = encodeVarintResponse(dAtA, i, uint64(len(m.error.Error())))
		i--
		dAtA[i] = 0x12
	}
	if m.success {
		i--
		if m.success {
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
func (m *transactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.success {
		n += 2
	}
	l = len(m.error.Error())
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
func (m *transactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: transactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.success = bool(v != 0)
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
			m.error = errors.New(string(dAtA[iNdEx:postIndex]))
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
