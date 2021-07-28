// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/schema/types/base/v1beta1/properties.proto

package base

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

type properties struct {
	PropertyList []github_com_persistenceOne_persistenceSDK_schema_types.Property `protobuf:"bytes,1,rep,name=PropertyList,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Property" json:"PropertyList"`
}

func (m *properties) Reset()         { *m = properties{} }
func (m *properties) String() string { return proto.CompactTextString(m) }
func (*properties) ProtoMessage()    {}
func (*properties) Descriptor() ([]byte, []int) {
	return fileDescriptor_044448abb714431a, []int{0}
}
func (m *properties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *properties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Properties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *properties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Properties.Merge(m, src)
}
func (m *properties) XXX_Size() int {
	return m.Size()
}
func (m *properties) XXX_DiscardUnknown() {
	xxx_messageInfo_Properties.DiscardUnknown(m)
}

var xxx_messageInfo_Properties proto.InternalMessageInfo

func init() {
	proto.RegisterType((*properties)(nil), "persistence.schema.types.v1beta1.types.Properties")
}

func init() {
	proto.RegisterFile("proto/persistence/schema/types/base/v1beta1/properties.proto", fileDescriptor_044448abb714431a)
}

var fileDescriptor_044448abb714431a = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x2f, 0x4e,
	0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9,
	0x4c, 0x2d, 0xd6, 0x03, 0x6b, 0x13, 0x52, 0x43, 0xd2, 0xa7, 0x07, 0xd1, 0xa7, 0x07, 0xd6, 0xa7,
	0x07, 0xd5, 0x02, 0xe1, 0x49, 0xa9, 0x95, 0x64, 0x64, 0x16, 0xa5, 0xc4, 0x17, 0x24, 0x16, 0x95,
	0x54, 0xea, 0x43, 0x6c, 0x4c, 0xcf, 0x4f, 0xcf, 0x47, 0xb0, 0x20, 0xe6, 0x29, 0x55, 0x70, 0x71,
	0x05, 0xc0, 0xed, 0x10, 0xca, 0xe2, 0xe2, 0x81, 0xf2, 0x2a, 0x7d, 0x32, 0x8b, 0x4b, 0x24, 0x18,
	0x15, 0x98, 0x35, 0x38, 0x9d, 0xdc, 0x4e, 0xdc, 0x93, 0x67, 0xb8, 0x75, 0x4f, 0xde, 0x2e, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0xd9, 0xf9, 0xfe, 0x79, 0xa9, 0xc8, 0xdc,
	0x60, 0x17, 0x6f, 0x14, 0x0f, 0xe9, 0xc1, 0x4c, 0x0c, 0x42, 0x31, 0xdb, 0x49, 0xfb, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0x31, 0xc2, 0x23, 0x89, 0x0d, 0xec, 0x5a,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x13, 0x10, 0x27, 0x3d, 0x01, 0x00, 0x00,
}

func (m *properties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m properties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *properties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PropertyList) > 0 {
		for iNdEx := len(m.PropertyList) - 1; iNdEx >= 0; iNdEx-- {
			i -= m.PropertyList[iNdEx].Size()
			if _, err := m.PropertyList[iNdEx].MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintProperties(dAtA, i, uint64(m.PropertyList[iNdEx].Size()))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProperties(dAtA []byte, offset int, v uint64) int {
	offset -= sovProperties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m properties) Size() (n int) {
	var l int
	_ = l
	if len(m.PropertyList) > 0 {
		for _, s := range m.PropertyList {
			l = s.Size()
			n += 1 + l + sovProperties(uint64(l))
		}
	}
	return n
}

func sovProperties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProperties(x uint64) (n int) {
	return sovProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m properties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProperties
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
			return fmt.Errorf("proto: Properties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Properties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProperties
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
				return ErrInvalidLengthProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PropertyList = append(m.PropertyList, property{})
			if err := m.PropertyList[len(m.PropertyList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err!=nil {return err}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProperties
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
func skipProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProperties
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
					return 0, ErrIntOverflowProperties
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
					return 0, ErrIntOverflowProperties
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
				return 0, ErrInvalidLengthProperties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProperties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProperties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProperties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProperties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProperties = fmt.Errorf("proto: unexpected end of group")
)
