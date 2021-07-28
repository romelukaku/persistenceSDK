// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/orders/transactions/immediate/v1beta1/tx.proto

package immediate

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type message struct {
	From                    github_com_cosmos_cosmos_sdk_types.AccAddress                        `protobuf:"bytes,1,opt,name=From,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from" valid:"required~required field from missing"`
	FromID                  github_com_persistenceOne_persistenceSDK_schema_types.ID             `protobuf:"bytes,2,opt,name=FromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"fromID" valid:"required~required field fromID missing"`
	ClassificationID        github_com_persistenceOne_persistenceSDK_schema_types.ID             `protobuf:"bytes,3,opt,name=ClassificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classificationID" valid:"required~required field classificationID missing"`
	MakerOwnableID          github_com_persistenceOne_persistenceSDK_schema_types.ID             `protobuf:"bytes,4,opt,name=MakerOwnableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"makerOwnableID" valid:"required~required field makerOwnableID missing"`
	TakerOwnableID          github_com_persistenceOne_persistenceSDK_schema_types.ID             `protobuf:"bytes,5,opt,name=TakerOwnableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"takerOwnableID" valid:"required~required field takerOwnableID missing"`
	ExpiresIn               github_com_persistenceOne_persistenceSDK_schema_types.Height         `protobuf:"bytes,6,opt,name=ExpiresIn,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Height" json:"expiresIn" valid:"required~required field expiresIn missing"`
	MakerOwnableSplit       github_com_cosmos_cosmos_sdk_types.Dec                               `protobuf:"bytes,7,opt,name=MakerOwnableSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing"`
	TakerOwnableSplit       github_com_cosmos_cosmos_sdk_types.Dec                               `protobuf:"bytes,8,opt,name=TakerOwnableSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing"`
	ImmutableMetaProperties github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,9,opt,name=ImmutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing"`
	ImmutableProperties     github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,10,opt,name=ImmutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"immutableProperties" valid:"required~required field immutableProperties missing"`
	MutableMetaProperties   github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,11,opt,name=MutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing"`
	MutableProperties       github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,12,opt,name=MutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_94b91c4bf6fd560f, []int{0}
}
func (m *message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *message) XXX_Size() int {
	return m.Size()
}
func (m *message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*message)(nil), "persistence.assets.orders.transactions.immediate.v1beta1.message")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/orders/transactions/immediate/v1beta1/tx.proto", fileDescriptor_94b91c4bf6fd560f)
}

var fileDescriptor_94b91c4bf6fd560f = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd4, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xc0, 0xf1, 0xdd, 0xdf, 0xaf, 0xa6, 0xcd, 0x28, 0x62, 0x47, 0xc4, 0xc5, 0xc3, 0x46, 0x3c,
	0x14, 0x2f, 0x9d, 0xa1, 0x88, 0x50, 0xc4, 0x43, 0x5b, 0x57, 0xcd, 0xa2, 0xa1, 0x92, 0xf6, 0x20,
	0x22, 0xc8, 0x64, 0xf7, 0x31, 0x19, 0xba, 0xb3, 0xb3, 0xcc, 0x33, 0xa9, 0xad, 0x47, 0x0f, 0xe2,
	0xd1, 0x17, 0xe0, 0xa1, 0x2f, 0xa7, 0xc7, 0x1e, 0xc5, 0x43, 0x91, 0xe4, 0xe2, 0xcb, 0x90, 0x6c,
	0x96, 0x74, 0xe3, 0xa6, 0x20, 0x21, 0x9e, 0x36, 0x7f, 0xe0, 0xf3, 0x7d, 0x9e, 0x61, 0x19, 0xf2,
	0x3c, 0x33, 0xda, 0x6a, 0x9e, 0x81, 0x41, 0x89, 0x16, 0xd2, 0x08, 0xb8, 0x40, 0x04, 0x8b, 0x5c,
	0x9b, 0x18, 0x0c, 0x72, 0x6b, 0x44, 0x8a, 0x22, 0xb2, 0x52, 0xa7, 0xc8, 0xa5, 0x52, 0x10, 0x4b,
	0x61, 0x81, 0x1f, 0x6e, 0x74, 0xc0, 0x8a, 0x0d, 0x6e, 0x8f, 0x58, 0x2e, 0xd0, 0xcd, 0x12, 0xc1,
	0xc6, 0x04, 0x1b, 0x13, 0xac, 0x4c, 0xb0, 0x09, 0xc1, 0x0a, 0xe2, 0xce, 0x9a, 0xed, 0x49, 0x13,
	0xbf, 0xcb, 0x84, 0xb1, 0xc7, 0x7c, 0x3c, 0x4e, 0x57, 0x77, 0xf5, 0xc5, 0xa7, 0x71, 0xe1, 0xde,
	0x37, 0x42, 0x96, 0x15, 0x20, 0x8a, 0x2e, 0xd0, 0x90, 0x2c, 0x3d, 0x33, 0x5a, 0x79, 0xee, 0x5d,
	0xf7, 0x7e, 0x7d, 0xe7, 0xe1, 0xe9, 0x79, 0xc3, 0xf9, 0x71, 0xde, 0x58, 0xef, 0x4a, 0xdb, 0xeb,
	0x77, 0x58, 0xa4, 0x15, 0x8f, 0x34, 0x2a, 0x8d, 0xc5, 0x63, 0x1d, 0xe3, 0x03, 0x6e, 0x8f, 0x33,
	0x40, 0xb6, 0x1d, 0x45, 0xdb, 0x71, 0x6c, 0x00, 0xb1, 0x9d, 0x13, 0xf4, 0x35, 0xa9, 0x8d, 0x9e,
	0x61, 0xe0, 0xfd, 0x97, 0x63, 0x5b, 0x05, 0xb6, 0x59, 0xc2, 0x4a, 0xbb, 0xed, 0xa6, 0x50, 0xfe,
	0xba, 0x17, 0xbc, 0xe0, 0x18, 0xf5, 0x40, 0x89, 0xc2, 0x0f, 0x83, 0x76, 0xe1, 0xd1, 0x84, 0xdc,
	0x78, 0x92, 0x08, 0x44, 0xf9, 0x5e, 0x46, 0x62, 0xb4, 0x7d, 0x18, 0x78, 0xff, 0x2f, 0xa8, 0x51,
	0x91, 0x69, 0x8f, 0x5c, 0x6f, 0x89, 0x03, 0x30, 0xbb, 0x1f, 0x52, 0xd1, 0x49, 0x20, 0x0c, 0xbc,
	0xa5, 0x05, 0xb5, 0xfe, 0x70, 0x47, 0xa5, 0xfd, 0xe9, 0xd2, 0x95, 0x45, 0x95, 0xa6, 0x5d, 0xda,
	0x21, 0xf5, 0xa7, 0x47, 0x99, 0x34, 0x80, 0x61, 0xea, 0xd5, 0xf2, 0x48, 0x50, 0x44, 0x1e, 0xcf,
	0x17, 0x69, 0x82, 0xec, 0xf6, 0x6c, 0xfb, 0x82, 0xa5, 0x6f, 0xc9, 0x6a, 0x79, 0xbf, 0xbd, 0x2c,
	0x91, 0xd6, 0x5b, 0xce, 0x5b, 0xac, 0x68, 0xad, 0xfd, 0xc5, 0x7b, 0x15, 0x40, 0xd4, 0xae, 0x42,
	0x23, 0x7d, 0xbf, 0xa2, 0xaf, 0xcc, 0xa7, 0x57, 0x20, 0xfa, 0xd9, 0x25, 0xb7, 0x43, 0xa5, 0xfa,
	0x76, 0xf4, 0x53, 0x0b, 0xac, 0x78, 0x65, 0x74, 0x06, 0xc6, 0x4a, 0x40, 0xaf, 0x9e, 0x47, 0x5e,
	0x16, 0x91, 0x60, 0xbe, 0xe3, 0x9a, 0x36, 0xdb, 0x97, 0xc5, 0xe8, 0x47, 0x72, 0x73, 0xf2, 0x57,
	0x69, 0x06, 0x92, 0xcf, 0xd0, 0x2c, 0x66, 0xd8, 0x9a, 0x6f, 0x86, 0x52, 0x7f, 0x56, 0x84, 0x7e,
	0x72, 0xc9, 0xad, 0xd6, 0xcc, 0x23, 0xb8, 0xfa, 0x0f, 0x8e, 0x60, 0x76, 0x8a, 0x1e, 0x92, 0xd5,
	0x56, 0x65, 0xfd, 0x6b, 0x0b, 0x5e, 0xbf, 0x9a, 0x78, 0xb4, 0xf2, 0xe5, 0xa4, 0xe1, 0xfc, 0x3a,
	0x69, 0x38, 0x3b, 0xcd, 0xd3, 0x81, 0xef, 0x9e, 0x0d, 0x7c, 0xf7, 0xe7, 0xc0, 0x77, 0xbf, 0x0e,
	0x7d, 0xe7, 0x6c, 0xe8, 0x3b, 0xdf, 0x87, 0xbe, 0xf3, 0x86, 0x29, 0x1d, 0xf7, 0x13, 0x98, 0x5c,
	0xe9, 0x32, 0xb5, 0x60, 0x52, 0x91, 0x5c, 0x72, 0xb7, 0x77, 0x6a, 0xf9, 0x7d, 0xfb, 0xe0, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x77, 0x57, 0xb0, 0x1c, 0x06, 0x00, 0x00,
}

func (m *message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MutableProperties.Size() > 0 {
		i -= m.MutableProperties.Size()
		if _, err := m.MutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.MutableProperties.Size()))
		i--
		dAtA[i] = 0x62
	}
	if m.MutableMetaProperties.Size() > 0 {
		i -= m.MutableMetaProperties.Size()
		if _, err := m.MutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.MutableMetaProperties.Size()))
		i--
		dAtA[i] = 0x5a
	}
	if m.ImmutableProperties.Size() > 0 {
		i -= m.ImmutableProperties.Size()
		if _, err := m.ImmutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.ImmutableProperties.Size()))
		i--
		dAtA[i] = 0x52
	}
	if m.ImmutableMetaProperties.Size() > 0 {
		i -= m.ImmutableMetaProperties.Size()
		if _, err := m.ImmutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.ImmutableMetaProperties.Size()))
		i--
		dAtA[i] = 0x4a
	}
	if m.TakerOwnableSplit.Size() > 0 {
		i -= m.TakerOwnableSplit.Size()
		if _, err := m.TakerOwnableSplit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.TakerOwnableSplit.Size()))
		i--
		dAtA[i] = 0x42
	}
	if m.MakerOwnableSplit.Size() > 0 {
		i -= m.MakerOwnableSplit.Size()
		if _, err := m.MakerOwnableSplit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.MakerOwnableSplit.Size()))
		i--
		dAtA[i] = 0x3a
	}
	if m.ExpiresIn.Size() > 0 {
		i -= m.ExpiresIn.Size()
		if _, err := m.ExpiresIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.ExpiresIn.Size()))
		i--
		dAtA[i] = 0x32
	}
	if m.TakerOwnableID.Size() > 0 {
		i -= m.TakerOwnableID.Size()
		if _, err := m.TakerOwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.TakerOwnableID.Size()))
		i--
		dAtA[i] = 0x2a
	}
	if m.MakerOwnableID.Size() > 0 {
		i -= m.MakerOwnableID.Size()
		if _, err := m.MakerOwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.MakerOwnableID.Size()))
		i--
		dAtA[i] = 0x22
	}
	if m.ClassificationID.Size() > 0 {
		i -= m.ClassificationID.Size()
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.ClassificationID.Size()))
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID.Size() > 0 {
		i -= m.FromID.Size()
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.FromID.Size()))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.FromID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.ClassificationID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MakerOwnableID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TakerOwnableID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.ExpiresIn.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MakerOwnableSplit.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TakerOwnableSplit.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.ImmutableMetaProperties.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.ImmutableProperties.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MutableMetaProperties.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MutableProperties.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpiresIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableSplit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerOwnableSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerOwnableSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
