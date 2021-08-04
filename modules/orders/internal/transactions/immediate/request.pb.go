// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/orders/transactions/immediate/v1beta1/request.proto

package immediate

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type transactionRequest struct {
	BaseReq                 test_types.BaseReq `protobuf:"bytes,1,opt,name=baseReq,proto3" json:"baseReq"`
	FromID                  string             `protobuf:"bytes,2,opt,name=fromID,proto3" json:"fromID,omitempty" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ClassificationID        string             `protobuf:"bytes,3,opt,name=classificationID,proto3" json:"classificationID,omitempty" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	MakerOwnableID          string             `protobuf:"bytes,4,opt,name=makerOwnableID,proto3" json:"makerOwnableID,omitempty" valid:"required~required field makerOwnableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field makerOwnableID"`
	TakerOwnableID          string             `protobuf:"bytes,5,opt,name=takerOwnableID,proto3" json:"takerOwnableID,omitempty" valid:"required~required field takerOwnableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field takerOwnableID"`
	ExpiresIn               int64              `protobuf:"varint,6,opt,name=expiresIn,proto3" json:"expiresIn,omitempty" valid:"required~required field expiresIn missing, matches(^[0-9]+$)~invalid field expiresIn"`
	MakerOwnableSplit       string             `protobuf:"bytes,7,opt,name=makerOwnableSplit,proto3" json:"makerOwnableSplit,omitempty" valid:"required~required field makerOwnableSplit missing, matches(^[0-9.]+$)~invalid field makerOwnableSplit"`
	TakerOwnableSplit       string             `protobuf:"bytes,8,opt,name=takerOwnableSplit,proto3" json:"takerOwnableSplit,omitempty" valid:"required~required field takerOwnableSplit missing, matches(^[0-9.]+$)~invalid field takerOwnableSplit"`
	ImmutableMetaProperties string             `protobuf:"bytes,9,opt,name=immutableMetaProperties,proto3" json:"immutableMetaProperties,omitempty" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string             `protobuf:"bytes,10,opt,name=immutableProperties,proto3" json:"immutableProperties,omitempty" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string             `protobuf:"bytes,11,opt,name=mutableMetaProperties,proto3" json:"mutableMetaProperties,omitempty" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string             `protobuf:"bytes,12,opt,name=mutableProperties,proto3" json:"mutableProperties,omitempty" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

func (m *transactionRequest) Reset()         { *m = transactionRequest{} }
func (m *transactionRequest) String() string { return proto.CompactTextString(m) }
func (*transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_444eb686900ed9ee, []int{0}
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
	proto.RegisterType((*transactionRequest)(nil), "persistence.modules.orders.transactions.immediate.v1beta1.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/orders/transactions/immediate/v1beta1/request.proto", fileDescriptor_444eb686900ed9ee)
}

var fileDescriptor_444eb686900ed9ee = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x13, 0x56, 0xba, 0xd5, 0x43, 0x08, 0x8c, 0x80, 0x68, 0x42, 0x69, 0xc5, 0x61, 0xaa,
	0x10, 0xb2, 0x19, 0xe3, 0x30, 0x71, 0x2c, 0x45, 0x5a, 0x24, 0x2a, 0x50, 0xd8, 0x89, 0x0b, 0x72,
	0x93, 0x77, 0xc3, 0x22, 0x89, 0x53, 0xbf, 0x2e, 0x7f, 0xbe, 0x01, 0x47, 0x3e, 0xc2, 0x3e, 0x0d,
	0xda, 0x71, 0x47, 0x4e, 0x08, 0xb5, 0x17, 0x3e, 0x06, 0x8a, 0xdd, 0xb5, 0xd9, 0xd2, 0xdd, 0xe2,
	0xf7, 0xf9, 0x3d, 0xcf, 0xfb, 0x58, 0x8a, 0xc9, 0x9b, 0x52, 0x2b, 0xa3, 0x78, 0x09, 0x1a, 0x25,
	0x1a, 0x28, 0x12, 0xe0, 0x02, 0x11, 0x0c, 0x72, 0xa5, 0x53, 0xd0, 0xc8, 0x8d, 0x16, 0x05, 0x8a,
	0xc4, 0x48, 0x55, 0x20, 0x97, 0x79, 0x0e, 0xa9, 0x14, 0x06, 0xf8, 0x97, 0xbd, 0x31, 0x18, 0xb1,
	0xc7, 0x35, 0x4c, 0xa6, 0x80, 0x86, 0xd9, 0x18, 0x7a, 0x50, 0xcb, 0x61, 0x2e, 0x87, 0xb9, 0x1c,
	0x56, 0xcf, 0x61, 0xcb, 0x1c, 0xb6, 0xc8, 0xd9, 0xd9, 0x35, 0x9f, 0xa4, 0x4e, 0x3f, 0x96, 0x42,
	0x9b, 0xef, 0xdc, 0x75, 0x3a, 0x51, 0x27, 0x6a, 0xf5, 0xe5, 0x36, 0xec, 0xec, 0x37, 0xfb, 0x26,
	0x0a, 0x73, 0x85, 0x7c, 0x2c, 0x70, 0x55, 0xab, 0x3a, 0x68, 0x98, 0x38, 0xd3, 0xe3, 0x5f, 0x2d,
	0x42, 0x6b, 0xfb, 0x63, 0xd7, 0x99, 0x1e, 0x92, 0xcd, 0x8a, 0x8b, 0x61, 0x12, 0xf8, 0x3d, 0xbf,
	0xbf, 0xfd, 0xbc, 0xcf, 0xea, 0xfd, 0x5d, 0x2e, 0xab, 0x90, 0x8b, 0x9a, 0x6c, 0xe0, 0xf8, 0x41,
	0xeb, 0xec, 0x4f, 0xd7, 0x8b, 0x2f, 0xec, 0xf4, 0x01, 0x69, 0x1f, 0x6b, 0x95, 0x47, 0xc3, 0xe0,
	0x46, 0xcf, 0xef, 0x77, 0xe2, 0xc5, 0x89, 0x3e, 0x21, 0x77, 0x5e, 0x65, 0x02, 0x51, 0x1e, 0xcb,
	0x44, 0x54, 0xab, 0xa3, 0x61, 0xb0, 0x61, 0x89, 0xc6, 0x9c, 0xee, 0x92, 0xdb, 0x23, 0xf1, 0x19,
	0xf4, 0xdb, 0xaf, 0x85, 0x18, 0x67, 0x10, 0x0d, 0x83, 0x96, 0x25, 0xaf, 0x4c, 0x2b, 0xee, 0xe8,
	0x32, 0x77, 0xd3, 0x71, 0x97, 0xa7, 0xf4, 0x11, 0xe9, 0xbc, 0xfe, 0x56, 0x4a, 0x0d, 0x18, 0x15,
	0x41, 0xbb, 0xe7, 0xf7, 0x37, 0xe2, 0xd5, 0x80, 0x3e, 0x25, 0x77, 0xeb, 0xb9, 0xef, 0xcb, 0x4c,
	0x9a, 0x60, 0xd3, 0x06, 0x35, 0x85, 0x8a, 0x3e, 0x6a, 0xd0, 0x5b, 0x8e, 0x6e, 0x08, 0xf4, 0x80,
	0x3c, 0x8c, 0xf2, 0x7c, 0x6a, 0xaa, 0xc9, 0x08, 0x8c, 0x78, 0xa7, 0x55, 0x09, 0xda, 0x48, 0xc0,
	0xa0, 0x63, 0x3d, 0xd7, 0xc9, 0xf4, 0x19, 0xb9, 0xb7, 0x94, 0x6a, 0x2e, 0x62, 0x5d, 0xeb, 0x24,
	0xfa, 0x82, 0xdc, 0x1f, 0xad, 0xdd, 0xb4, 0x6d, 0x3d, 0xeb, 0x45, 0x7b, 0xfb, 0xc6, 0x96, 0x5b,
	0x8b, 0xdb, 0x5f, 0x15, 0x5e, 0x6e, 0xfd, 0x38, 0xed, 0x7a, 0xff, 0x4e, 0xbb, 0xde, 0xe0, 0xf0,
	0x6c, 0x16, 0xfa, 0xe7, 0xb3, 0xd0, 0xff, 0x3b, 0x0b, 0xfd, 0x9f, 0xf3, 0xd0, 0x3b, 0x9f, 0x87,
	0xde, 0xef, 0x79, 0xe8, 0x7d, 0x60, 0xb9, 0x4a, 0xa7, 0x19, 0x2c, 0x9f, 0x8d, 0x2c, 0x0c, 0xe8,
	0x42, 0x64, 0xd7, 0xbc, 0x9f, 0x71, 0xdb, 0xfe, 0x99, 0xfb, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0x37, 0x5f, 0xdc, 0x80, 0x03, 0x00, 0x00,
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
	if len(m.MutableProperties) > 0 {
		i -= len(m.MutableProperties)
		copy(dAtA[i:], m.MutableProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MutableProperties)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.MutableMetaProperties) > 0 {
		i -= len(m.MutableMetaProperties)
		copy(dAtA[i:], m.MutableMetaProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MutableMetaProperties)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ImmutableProperties) > 0 {
		i -= len(m.ImmutableProperties)
		copy(dAtA[i:], m.ImmutableProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ImmutableProperties)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ImmutableMetaProperties) > 0 {
		i -= len(m.ImmutableMetaProperties)
		copy(dAtA[i:], m.ImmutableMetaProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ImmutableMetaProperties)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.TakerOwnableSplit) > 0 {
		i -= len(m.TakerOwnableSplit)
		copy(dAtA[i:], m.TakerOwnableSplit)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.TakerOwnableSplit)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MakerOwnableSplit) > 0 {
		i -= len(m.MakerOwnableSplit)
		copy(dAtA[i:], m.MakerOwnableSplit)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MakerOwnableSplit)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ExpiresIn != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.ExpiresIn))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TakerOwnableID) > 0 {
		i -= len(m.TakerOwnableID)
		copy(dAtA[i:], m.TakerOwnableID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.TakerOwnableID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MakerOwnableID) > 0 {
		i -= len(m.MakerOwnableID)
		copy(dAtA[i:], m.MakerOwnableID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MakerOwnableID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClassificationID) > 0 {
		i -= len(m.ClassificationID)
		copy(dAtA[i:], m.ClassificationID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ClassificationID)))
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
	l = len(m.ClassificationID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.MakerOwnableID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.TakerOwnableID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.ExpiresIn != 0 {
		n += 1 + sovRequest(uint64(m.ExpiresIn))
	}
	l = len(m.MakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.TakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ImmutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ImmutableProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.MutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.MutableProperties)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			m.ClassificationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
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
			m.MakerOwnableID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
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
			m.TakerOwnableID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresIn", wireType)
			}
			m.ExpiresIn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresIn |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableSplit", wireType)
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
			m.MakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
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
			m.TakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
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
			m.ImmutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
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
			m.ImmutableProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			m.MutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			m.MutableProperties = string(dAtA[iNdEx:postIndex])
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
