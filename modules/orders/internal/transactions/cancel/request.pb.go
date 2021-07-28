// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/orders/transactions/cancel/v1beta1/request.proto

package cancel

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
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

type transactionRequest struct {
	BaseReq test_types.BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"baseReq" `
	FromID  string             `protobuf:"bytes,2,opt,name=fromID,proto3" json:"fromID,omitempty" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	OrderID string             `protobuf:"bytes,3,opt,name=orderID,proto3" json:"orderID,omitempty" valid:"required~required field orderID missing, matches(^[A-Za-z0-9-_=.|*]+$)~invalid field orderID"`
}

func (m *transactionRequest) Reset()         { *m = transactionRequest{} }
func (m *transactionRequest) String() string { return proto.CompactTextString(m) }
func (*transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf8f11d8e270c6e, []int{0}
}
func (m *transactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *transactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *transactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*transactionRequest)(nil), "persistence.assets.orders.transactions.cancel.v1beta1.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/orders/transactions/cancel/v1beta1/request.proto", fileDescriptor_8bf8f11d8e270c6e)
}

var fileDescriptor_8bf8f11d8e270c6e = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4f, 0x3a, 0x31,
	0x14, 0xc7, 0xdb, 0xdf, 0xcf, 0x00, 0x9e, 0xdb, 0x0d, 0xe6, 0xc2, 0x50, 0x88, 0x83, 0x21, 0x31,
	0x69, 0x83, 0xc4, 0xc5, 0x91, 0xe0, 0x70, 0xeb, 0x8d, 0x2e, 0xa4, 0x77, 0x3c, 0xf1, 0x12, 0x68,
	0xb9, 0xf7, 0x8a, 0x89, 0xff, 0x81, 0xa3, 0xbb, 0x0b, 0x7f, 0x0e, 0x23, 0xa3, 0x93, 0x31, 0xb0,
	0xf8, 0x67, 0x98, 0xbb, 0x02, 0x5e, 0xa2, 0x5b, 0xbf, 0xe9, 0xeb, 0xa7, 0x9f, 0x97, 0x6f, 0x10,
	0x2f, 0xd0, 0x3a, 0xab, 0x16, 0x80, 0x94, 0x93, 0x03, 0x93, 0x81, 0xd2, 0x44, 0xe0, 0x48, 0x59,
	0x9c, 0x00, 0x92, 0x72, 0xa8, 0x0d, 0xe9, 0xcc, 0xe5, 0xd6, 0x90, 0xca, 0xb4, 0xc9, 0x60, 0xa6,
	0x9e, 0xfa, 0x29, 0x38, 0xdd, 0x57, 0x08, 0xc5, 0x12, 0xc8, 0xc9, 0x8a, 0x11, 0xde, 0xd4, 0x20,
	0xd2, 0x43, 0xa4, 0x87, 0xc8, 0x3a, 0x44, 0x7a, 0x88, 0xdc, 0x43, 0xda, 0x97, 0xee, 0x31, 0xc7,
	0xc9, 0x78, 0xa1, 0xd1, 0x3d, 0x2b, 0x6f, 0x33, 0xb5, 0x53, 0xfb, 0x73, 0xf2, 0xf8, 0xf6, 0xe0,
	0xb7, 0x69, 0x66, 0x69, 0x6e, 0x49, 0xa5, 0x9a, 0xe0, 0xe8, 0x54, 0x06, 0x84, 0xc2, 0x3f, 0xba,
	0x78, 0xe3, 0x41, 0x58, 0xfb, 0x3c, 0xf1, 0xc2, 0x61, 0x1c, 0xb4, 0xca, 0xb9, 0x31, 0x42, 0x11,
	0xf1, 0x2e, 0xef, 0x9d, 0x5d, 0xf7, 0x64, 0xdd, 0xde, 0x83, 0x65, 0x39, 0x73, 0xf0, 0x94, 0x43,
	0x4d, 0x90, 0x40, 0x31, 0x3c, 0x59, 0x7f, 0x74, 0x58, 0xd2, 0x4c, 0x7d, 0x0c, 0xcf, 0x83, 0xc6,
	0x03, 0xda, 0x79, 0x3c, 0x8a, 0xfe, 0x75, 0x79, 0xef, 0x34, 0xd9, 0xa7, 0x30, 0x0a, 0x9a, 0xd5,
	0xf2, 0xf1, 0x28, 0xfa, 0x5f, 0x5d, 0x1c, 0xe2, 0x6d, 0xeb, 0x65, 0xd5, 0x61, 0x5f, 0xab, 0x0e,
	0x1b, 0xde, 0xad, 0xb7, 0x82, 0x6f, 0xb6, 0x82, 0x7f, 0x6e, 0x05, 0x7f, 0xdd, 0x09, 0xb6, 0xd9,
	0x09, 0xf6, 0xbe, 0x13, 0xec, 0xfe, 0x6a, 0x6e, 0x27, 0xcb, 0x19, 0x1c, 0x5b, 0xc8, 0x8d, 0x03,
	0x34, 0x7a, 0xf6, 0x57, 0x1d, 0x69, 0xa3, 0xda, 0x75, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x42,
	0xc0, 0x52, 0x39, 0xcc, 0x01, 0x00, 0x00,
}

func (m *transactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OrderID) > 0 {
		i -= len(m.OrderID)
		copy(dAtA[i:], m.OrderID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.OrderID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromID) > 0 {
		i -= len(m.FromID)
		copy(dAtA[i:], m.FromID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.FromID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BaseReq.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *transactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseReq.Size()
	n += 1 + l + sovRequest(uint64(l))
	l = len(m.FromID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.OrderID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: transactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
