// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/orders/transactions/cancel/v1beta1/response.proto

package cancel

import (
	"errors"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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
// A compilation Error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type transactionResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Error   error `protobuf:"bytes,2,opt,name=Error,proto3,customtype=Error" json:"Error"`
}

func (m *transactionResponse) Reset()         { *m = transactionResponse{} }
func (m *transactionResponse) String() string { return proto.CompactTextString(m) }
func (*transactionResponse) ProtoMessage()    {}
func (*transactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afaa7053bd79219b, []int{0}
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
	proto.RegisterType((*transactionResponse)(nil), "persistence.modules.orders.transactions.cancel.v1beta1.transactionResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/orders/transactions/cancel/v1beta1/response.proto", fileDescriptor_afaa7053bd79219b)
}

var fileDescriptor_afaa7053bd79219b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0x03, 0x41,
	0x14, 0x85, 0x67, 0xc4, 0xdf, 0x05, 0x9b, 0xd8, 0x2c, 0x29, 0x66, 0x83, 0x82, 0x04, 0x02, 0x33,
	0x04, 0xb1, 0xb1, 0x0c, 0xd8, 0x58, 0x6e, 0x99, 0x26, 0x4c, 0x66, 0x2f, 0x71, 0x61, 0x33, 0x77,
	0xb9, 0x77, 0x22, 0xf8, 0x16, 0x3e, 0x82, 0x8f, 0x93, 0x32, 0xa5, 0x58, 0x04, 0xd9, 0x6d, 0x7c,
	0x0c, 0x21, 0x13, 0x75, 0x41, 0xab, 0x99, 0x73, 0x39, 0x7c, 0xe7, 0x70, 0x92, 0x87, 0x9a, 0x30,
	0xa0, 0xa9, 0x81, 0xb8, 0xe4, 0x00, 0xde, 0x81, 0xb1, 0xcc, 0x10, 0xd8, 0x20, 0x15, 0x40, 0x6c,
	0x02, 0x59, 0xcf, 0xd6, 0x85, 0x12, 0x3d, 0x1b, 0x67, 0xbd, 0x83, 0xca, 0x3c, 0x8d, 0xe7, 0x10,
	0xec, 0xd8, 0x10, 0x70, 0x8d, 0x9e, 0x41, 0xef, 0x20, 0xbd, 0xdb, 0x0e, 0x45, 0x47, 0x8a, 0x8e,
	0x14, 0xdd, 0xa5, 0xe8, 0x48, 0xd1, 0x7b, 0x4a, 0x7f, 0x14, 0x1e, 0x4b, 0x2a, 0x66, 0xb5, 0xa5,
	0xf0, 0x6c, 0x62, 0x1d, 0x87, 0xbc, 0x44, 0x9e, 0x75, 0x45, 0xcc, 0xe8, 0x5f, 0xff, 0x35, 0x2f,
	0x70, 0x81, 0xbf, 0xbf, 0xe8, 0xbb, 0x9c, 0x26, 0x17, 0x9d, 0xcc, 0x7c, 0x5f, 0xb4, 0x97, 0x26,
	0x27, 0xbc, 0x72, 0x0e, 0x98, 0x53, 0x39, 0x90, 0xc3, 0xd3, 0xfc, 0x5b, 0xf6, 0xae, 0x92, 0x23,
	0x20, 0x42, 0x4a, 0x0f, 0x06, 0x72, 0x78, 0x36, 0x39, 0x5f, 0x6f, 0x33, 0xf1, 0xbe, 0xcd, 0xe2,
	0x31, 0x8f, 0xcf, 0xdd, 0xe1, 0xe7, 0x6b, 0x26, 0x26, 0xf7, 0xeb, 0x46, 0xc9, 0x4d, 0xa3, 0xe4,
	0x47, 0xa3, 0xe4, 0x4b, 0xab, 0xc4, 0xa6, 0x55, 0xe2, 0xad, 0x55, 0x62, 0x3a, 0x5a, 0x62, 0xb1,
	0xaa, 0xe0, 0x67, 0xbc, 0xd2, 0x07, 0x20, 0x6f, 0xab, 0xff, 0x56, 0x9c, 0x1f, 0xef, 0x9a, 0xde,
	0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x9b, 0xb3, 0x82, 0x83, 0x01, 0x00, 0x00,
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
