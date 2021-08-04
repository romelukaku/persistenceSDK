// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/orders/queries/orders/v1beta1/request.proto

package order

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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

type queryRequest struct {
	OrderID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=orderID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"orderID" valid:"required~required field orderID missing"`
}

func (m *queryRequest) Reset()         { *m = queryRequest{} }
func (m *queryRequest) String() string { return proto.CompactTextString(m) }
func (*queryRequest) ProtoMessage()    {}
func (*queryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a3fbca774d96bb4, []int{0}
}
func (m *queryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *queryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *queryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *queryRequest) XXX_Size() int {
	return m.Size()
}
func (m *queryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*queryRequest)(nil), "persistence.modules.orders.queries.order.v1beta1.queryRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/orders/queries/orders/v1beta1/request.proto", fileDescriptor_6a3fbca774d96bb4)
}

var fileDescriptor_6a3fbca774d96bb4 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x4f, 0x2c,
	0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d,
	0xca, 0x4c, 0x85, 0x73, 0xcb, 0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x8b, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0xf4, 0xc0, 0xda, 0x85, 0x90, 0xf5, 0xeb, 0x41, 0xf4, 0xeb, 0x41, 0x34, 0xe8,
	0x41, 0xf5, 0x43, 0xb8, 0x7a, 0x50, 0xed, 0x52, 0x6a, 0x25, 0x19, 0x99, 0x45, 0x29, 0xf1, 0x05,
	0x89, 0x45, 0x25, 0x95, 0xfa, 0x10, 0x27, 0xa4, 0xe7, 0xa7, 0xe7, 0x23, 0x58, 0x10, 0x83, 0x95,
	0x4a, 0xb8, 0x78, 0x40, 0x06, 0x54, 0x06, 0x41, 0xac, 0x13, 0x8a, 0xe2, 0x62, 0x07, 0x1b, 0xe4,
	0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xe4, 0x70, 0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b,
	0xf2, 0x16, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xc8, 0x8e, 0xf1, 0xcf,
	0x4b, 0x45, 0xe6, 0x06, 0xbb, 0x78, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0x97, 0x54,
	0x16, 0xa4, 0x16, 0xeb, 0x79, 0xba, 0x04, 0xc1, 0x0c, 0xb4, 0xe2, 0xe8, 0x58, 0x20, 0xcf, 0xf0,
	0x62, 0x81, 0x3c, 0x83, 0x93, 0xfd, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0xa9, 0xe6, 0xe6, 0xa7, 0x94, 0xe6, 0x20, 0x82, 0x23, 0x33, 0xaf, 0x24, 0xb5, 0x28, 0x2f, 0x31,
	0x07, 0x35, 0x98, 0x92, 0xd8, 0xc0, 0xae, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x5a,
	0xa9, 0xe2, 0x5e, 0x01, 0x00, 0x00,
}

func (m *queryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *queryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *queryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderID.Size() > 0 {
		i -= m.OrderID.Size()
		if _, err := m.OrderID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRequest(dAtA, i, uint64(m.OrderID.Size()))
		i--
		dAtA[i] = 0xa
	}
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
func (m *queryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OrderID.Size()
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
func (m *queryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: queryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: queryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
