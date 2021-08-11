// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/modules/metas/auxiliaries/supplement/v1beta1/request.proto

package supplement

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type auxiliaryRequest struct {
	PropertyList []github_com_persistenceOne_persistenceSDK_schema_types.Property `protobuf:"bytes,1,rep,name=propertyList,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Property" json:"propertyList"`
}

func (m *auxiliaryRequest) Reset()         { *m = auxiliaryRequest{} }
func (m *auxiliaryRequest) String() string { return proto.CompactTextString(m) }
func (*auxiliaryRequest) ProtoMessage()    {}
func (*auxiliaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e6188ae0a861cf7, []int{0}
}
func (m *auxiliaryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *auxiliaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuxiliaryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *auxiliaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuxiliaryRequest.Merge(m, src)
}
func (m *auxiliaryRequest) XXX_Size() int {
	return m.Size()
}
func (m *auxiliaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuxiliaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuxiliaryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*auxiliaryRequest)(nil), "persistence.modules.metas.auxiliaries.supplement.v1beta1.auxiliaryRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/modules/metas/auxiliaries/supplement/v1beta1/request.proto", fileDescriptor_5e6188ae0a861cf7)
}

var fileDescriptor_5e6188ae0a861cf7 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xb1, 0x4a, 0xf4, 0x40,
	0x10, 0x07, 0xf0, 0x0d, 0x1f, 0x7c, 0x68, 0xb0, 0x90, 0xab, 0xc4, 0x62, 0x23, 0x22, 0x62, 0xb5,
	0xcb, 0x61, 0x21, 0x58, 0x5c, 0x71, 0x88, 0xcd, 0x09, 0xca, 0xd9, 0xd9, 0xc8, 0x26, 0x37, 0x24,
	0x2b, 0x49, 0x76, 0xdd, 0x99, 0x15, 0xf3, 0x06, 0x57, 0xfa, 0x08, 0xf7, 0x38, 0x57, 0x5e, 0x29,
	0x16, 0x87, 0x24, 0x8d, 0x8f, 0x21, 0x24, 0x27, 0x17, 0x1b, 0xbb, 0x99, 0xe6, 0xf7, 0x9f, 0xff,
	0x84, 0x13, 0xeb, 0x0c, 0x19, 0x69, 0xc1, 0xa1, 0x46, 0x82, 0x32, 0x01, 0xa9, 0x10, 0x81, 0x50,
	0x16, 0x40, 0x0a, 0xa5, 0xf2, 0xaf, 0x3a, 0xd7, 0xca, 0x69, 0x40, 0x89, 0xde, 0xda, 0x1c, 0x0a,
	0x28, 0x49, 0xbe, 0x0c, 0x63, 0x20, 0x35, 0x94, 0x0e, 0x9e, 0x3d, 0x20, 0x89, 0x56, 0x19, 0x5c,
	0xf4, 0x18, 0xd1, 0x31, 0xa2, 0x65, 0x44, 0x8f, 0x11, 0x5b, 0x46, 0x6c, 0x98, 0xc3, 0x53, 0xca,
	0xb4, 0x9b, 0x3d, 0x5a, 0xe5, 0xa8, 0x92, 0xdd, 0x45, 0xa9, 0x49, 0xcd, 0x76, 0xea, 0x02, 0x8e,
	0xe7, 0x41, 0xb8, 0xff, 0x43, 0x55, 0xd3, 0x2e, 0x7b, 0xf0, 0x14, 0xee, 0x59, 0x67, 0x2c, 0x38,
	0xaa, 0x6e, 0x34, 0xd2, 0x41, 0x70, 0xf4, 0xef, 0x6c, 0x77, 0x7c, 0xbd, 0x5c, 0x47, 0xec, 0x63,
	0x1d, 0x8d, 0x52, 0x4d, 0x99, 0x8f, 0x45, 0x62, 0x8a, 0x7e, 0xcb, 0xdb, 0x12, 0xfa, 0xeb, 0xfd,
	0xd5, 0x44, 0x62, 0x92, 0x41, 0xa1, 0x24, 0x55, 0x16, 0x50, 0xdc, 0x6d, 0xc4, 0xe9, 0x2f, 0xfb,
	0x72, 0x67, 0xbe, 0x88, 0xd8, 0xd7, 0x22, 0x62, 0xe3, 0xd1, 0xb2, 0xe6, 0xc1, 0xaa, 0xe6, 0xc1,
	0x67, 0xcd, 0x83, 0xb7, 0x86, 0xb3, 0x55, 0xc3, 0xd9, 0x7b, 0xc3, 0xd9, 0xc3, 0x49, 0x61, 0x66,
	0x3e, 0x87, 0xbf, 0x3f, 0x18, 0xff, 0x6f, 0x1b, 0x9d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x36, 0xfb, 0x72, 0x81, 0x01, 0x00, 0x00,
}

func (m *auxiliaryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *auxiliaryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *auxiliaryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintRequest(dAtA, i, uint64(m.PropertyList[iNdEx].Size()))
			i--
			dAtA[i] = 0xa
		}
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
func (m *auxiliaryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PropertyList) > 0 {
		for _, s := range m.PropertyList {
			l = s.Size()
			n += 1 + l + sovRequest(uint64(l))
		}
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *auxiliaryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: auxiliaryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: auxiliaryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
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
			m.PropertyList = append(m.PropertyList, base.Property{})
			if err := m.PropertyList[len(m.PropertyList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
