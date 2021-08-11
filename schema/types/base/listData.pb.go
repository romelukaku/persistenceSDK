// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/schema/types/base/v1beta1/listData.proto

package base

import (
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

type listData struct {
	Value sortedDataList `protobuf:"bytes,1,opt,name=value,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.sortedDataList" json:"value"`
}

func (m *listData) Reset()      { *m = listData{} }
func (*listData) ProtoMessage() {}
func (*listData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a573801bdcf5730d, []int{0}
}
func (m *listData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *listData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *listData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListData.Merge(m, src)
}
func (m *listData) XXX_Size() int {
	return m.Size()
}
func (m *listData) XXX_DiscardUnknown() {
	xxx_messageInfo_ListData.DiscardUnknown(m)
}

var xxx_messageInfo_ListData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*listData)(nil), "persistence.schema.types.base.v1beta1.listData")
}

func init() {
	proto.RegisterFile("proto/persistence/schema/types/base/v1beta1/listData.proto", fileDescriptor_a573801bdcf5730d)
}

var fileDescriptor_a573801bdcf5730d = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x2f, 0x4e,
	0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0xc9, 0x2c, 0x2e, 0x71, 0x49, 0x2c, 0x49,
	0xd4, 0x03, 0x6b, 0x12, 0x52, 0x45, 0xd2, 0xa5, 0x07, 0xd1, 0xa5, 0x07, 0xd6, 0xa5, 0x07, 0xd2,
	0xa5, 0x07, 0xd5, 0x25, 0xa5, 0x56, 0x92, 0x91, 0x59, 0x94, 0x12, 0x5f, 0x90, 0x58, 0x54, 0x52,
	0xa9, 0x0f, 0xb1, 0x2e, 0x3d, 0x3f, 0x3d, 0x1f, 0xc1, 0x82, 0x18, 0xa7, 0x54, 0xcc, 0xc5, 0x01,
	0xb3, 0x40, 0x28, 0x9d, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xd3, 0x29, 0xf0, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x3d, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0x91, 0x9d, 0xec, 0x9f, 0x97, 0x8a, 0xcc, 0x0d, 0x76, 0xf1, 0xc6,
	0xf4, 0x84, 0x5e, 0x71, 0x7e, 0x51, 0x49, 0x6a, 0x0a, 0xc8, 0x74, 0x9f, 0xcc, 0xe2, 0x92, 0x20,
	0x88, 0xf9, 0x4e, 0xda, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x88,
	0x61, 0x44, 0x12, 0x1b, 0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xbb, 0x90,
	0xc2, 0x35, 0x01, 0x00, 0x00,
}

func (m *listData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m listData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *listData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		for index := len(m.Value) - 1; index >= 0; index-- {
			{
				size, err := m.Value[index].MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListData(dAtA, i, uint64(len(m.Value)))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintListData(dAtA []byte, offset int, v uint64) int {
	offset -= sovListData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m listData) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovListData(uint64(l))
	}
	return n
}

func sovListData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListData(x uint64) (n int) {
	return sovListData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m listData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListData
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
			return fmt.Errorf("proto: listData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: listData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListData
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
				return ErrInvalidLengthListData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			//TODO : Check id stringData is okay to be used here.
			m.Value = append(m.Value, stringData{})
			if err := m.Value[len(m.Value)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListData
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
func skipListData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListData
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
					return 0, ErrIntOverflowListData
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
					return 0, ErrIntOverflowListData
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
				return 0, ErrInvalidLengthListData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListData = fmt.Errorf("proto: unexpected end of group")
)
