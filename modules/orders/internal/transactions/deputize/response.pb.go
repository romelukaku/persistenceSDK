// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/orders/transactions/deputize/v1beta1/response.proto

package deputize

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

type transactionResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   error `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
}

func (m *transactionResponse) Reset()         { *m = transactionResponse{} }
func (m *transactionResponse) String() string { return proto.CompactTextString(m) }
func (*transactionResponse) ProtoMessage()    {}
func (*transactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25a11e62e306dfa4, []int{0}
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
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*transactionResponse)(nil), "persistence.modules.orders.transactions.deputize.v1beta1.transactionResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/orders/transactions/deputize/v1beta1/response.proto", fileDescriptor_25a11e62e306dfa4)
}

var fileDescriptor_25a11e62e306dfa4 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xbd, 0x4a, 0x03, 0x41,
	0x14, 0x05, 0xe0, 0x1d, 0xf1, 0x77, 0xc1, 0x66, 0x6d, 0x16, 0x8b, 0xd9, 0xa0, 0x20, 0x69, 0x9c,
	0x21, 0x58, 0x08, 0x96, 0x69, 0x6c, 0xac, 0xb6, 0x4c, 0x23, 0x93, 0xdd, 0x4b, 0x1c, 0x88, 0x33,
	0xc3, 0xbd, 0x77, 0x05, 0x7d, 0x0a, 0x1f, 0xc1, 0xc7, 0x49, 0x99, 0x52, 0x2c, 0x82, 0xec, 0x36,
	0x3e, 0x86, 0x90, 0x49, 0x74, 0x8b, 0x54, 0xf3, 0x53, 0x7c, 0xe7, 0x70, 0xd2, 0x87, 0x80, 0x9e,
	0xbd, 0x0e, 0x80, 0x64, 0x89, 0xc1, 0x55, 0xa0, 0x0d, 0x11, 0x30, 0x69, 0x8f, 0x35, 0x20, 0x69,
	0x46, 0xe3, 0xc8, 0x54, 0x6c, 0xbd, 0x23, 0x5d, 0x43, 0x68, 0xd8, 0xbe, 0x81, 0x7e, 0x19, 0x4d,
	0x81, 0xcd, 0x48, 0x23, 0x50, 0xf0, 0x8e, 0x40, 0xad, 0x99, 0xec, 0xb6, 0xe7, 0xa8, 0xe8, 0xa8,
	0xe8, 0xa8, 0xbe, 0xa3, 0xb6, 0x8e, 0xda, 0x38, 0xe7, 0x57, 0xfc, 0x64, 0xb1, 0x7e, 0x0c, 0x06,
	0xf9, 0x55, 0xc7, 0x4a, 0x33, 0x3f, 0xf3, 0xff, 0xb7, 0x18, 0x70, 0x31, 0x49, 0xcf, 0x7a, 0x50,
	0xb9, 0x49, 0xcf, 0xf2, 0xf4, 0x88, 0x9a, 0xaa, 0x02, 0xa2, 0x5c, 0x0c, 0xc4, 0xf0, 0xb8, 0xdc,
	0x3e, 0xb3, 0xcb, 0xf4, 0x00, 0x10, 0x3d, 0xe6, 0x7b, 0x03, 0x31, 0x3c, 0x19, 0x9f, 0x2e, 0x56,
	0x45, 0xf2, 0xb5, 0x2a, 0xe2, 0x67, 0x19, 0x8f, 0xbb, 0xfd, 0x9f, 0x8f, 0x22, 0x19, 0xdf, 0x2f,
	0x5a, 0x29, 0x96, 0xad, 0x14, 0xdf, 0xad, 0x14, 0xef, 0x9d, 0x4c, 0x96, 0x9d, 0x4c, 0x3e, 0x3b,
	0x99, 0x4c, 0xae, 0x9f, 0x7d, 0xdd, 0xcc, 0xe1, 0x6f, 0x13, 0xeb, 0x18, 0xd0, 0x99, 0xf9, 0xee,
	0x71, 0xa6, 0x87, 0xeb, 0xae, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x34, 0x59, 0xd9,
	0x5c, 0x01, 0x00, 0x00,
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
func (m *transactionResponse) Size() (n int) {
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
