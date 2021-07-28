// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/cosmos/base/v1beta1/basereq.proto

package test_types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	"net/http"
	"strings"
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

type BaseReq struct {
	From          string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Memo          string         `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	ChainID       string         `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AccountNumber uint64         `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64         `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	TimeoutHeight uint64         `protobuf:"varint,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	Fees          types.Coins    `protobuf:"bytes,7,rep,name=fees,proto3" json:"fees"`
	GasPrices     types.DecCoins `protobuf:"bytes,8,rep,name=gas_prices,json=gasPrices,proto3" json:"gas_prices"`
	Gas           string         `protobuf:"bytes,9,opt,name=gas,proto3" json:"gas,omitempty"`
	GasAdjustment string         `protobuf:"bytes,10,opt,name=gas_adjustment,json=gasAdjustment,proto3" json:"gas_adjustment,omitempty"`
	Simulate      bool           `protobuf:"varint,11,opt,name=simulate,proto3" json:"simulate,omitempty"`
}

//additionally added
func NewBaseReq(
	from, memo, chainID string, gas, gasAdjustment string, accNumber, seq uint64,
	fees types.Coins, gasPrices types.DecCoins, simulate bool,
) BaseReq {
	return BaseReq{
		From:          strings.TrimSpace(from),
		Memo:          strings.TrimSpace(memo),
		ChainID:       strings.TrimSpace(chainID),
		Fees:          fees,
		GasPrices:     gasPrices,
		Gas:           strings.TrimSpace(gas),
		GasAdjustment: strings.TrimSpace(gasAdjustment),
		AccountNumber: accNumber,
		Sequence:      seq,
		Simulate:      simulate,
	}
}

func (br BaseReq) ValidateBasic(w http.ResponseWriter) bool {
	if !br.Simulate {
		switch {
		case len(br.ChainID) == 0:
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "chain-id required but not specified")
			return false

		case !br.Fees.IsZero() && !br.GasPrices.IsZero():
			// both fees and gas prices were provided
			rest.WriteErrorResponse(w, http.StatusBadRequest, "cannot provide both fees and gas prices")
			return false

		case !br.Fees.IsValid() && !br.GasPrices.IsValid():
			// neither fees or gas prices were provided
			rest.WriteErrorResponse(w, http.StatusPaymentRequired, "invalid fees or gas prices provided")
			return false
		}
	}

	if _, err := types.AccAddressFromBech32(br.From); err != nil || len(br.From) == 0 {
		rest.WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("invalid from address: %s", br.From))
		return false
	}

	return true
}

func (br BaseReq) Sanitize() BaseReq {
	return NewBaseReq(
		br.From, br.Memo, br.ChainID, br.Gas, br.GasAdjustment,
		br.AccountNumber, br.Sequence, br.Fees, br.GasPrices, br.Simulate,
	)
}

//auto generated

func (m *BaseReq) Reset()         { *m = BaseReq{} }
func (m *BaseReq) String() string { return proto.CompactTextString(m) }
func (*BaseReq) ProtoMessage()    {}
func (*BaseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_57919bc98f071adc, []int{0}
}
func (m *BaseReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseReq.Merge(m, src)
}
func (m *BaseReq) XXX_Size() int {
	return m.Size()
}
func (m *BaseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseReq.DiscardUnknown(m)
}

var xxx_messageInfo_BaseReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseReq)(nil), "cosmos.base.v1beta1.BaseReq")
}

func init() {
	proto.RegisterFile("proto/cosmos/base/v1beta1/basereq.proto", fileDescriptor_57919bc98f071adc)
}

var fileDescriptor_57919bc98f071adc = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x8f, 0x93, 0x40,
	0x1c, 0xc5, 0x41, 0x70, 0x4b, 0x67, 0xb3, 0x46, 0xc7, 0xcb, 0x6c, 0x63, 0x28, 0xd1, 0xb8, 0x21,
	0x31, 0x81, 0xac, 0x7b, 0xf3, 0xb6, 0xd5, 0x83, 0x5e, 0x8c, 0xe1, 0xe8, 0x85, 0x0c, 0xc3, 0x7f,
	0x61, 0x8c, 0xc3, 0x50, 0xfe, 0x83, 0x49, 0xbf, 0x81, 0x17, 0x13, 0x3f, 0x42, 0x3f, 0x4e, 0x8f,
	0x3d, 0x7a, 0x32, 0xa6, 0xbd, 0xf8, 0x31, 0x0c, 0x43, 0x53, 0x3d, 0xd4, 0xdb, 0x7b, 0x3f, 0xde,
	0x7b, 0x30, 0x0c, 0xb9, 0x6a, 0x3b, 0x6d, 0x74, 0x2a, 0xb4, 0xd2, 0x98, 0x16, 0x1c, 0x21, 0xfd,
	0x72, 0x5d, 0x80, 0xe1, 0xd7, 0xd6, 0x74, 0xb0, 0x4c, 0x6c, 0x80, 0x3e, 0x16, 0x1a, 0x95, 0xc6,
	0x64, 0xa0, 0xc9, 0x21, 0x32, 0xbb, 0x32, 0xb5, 0xec, 0xca, 0xbc, 0xe5, 0x9d, 0x59, 0xa5, 0xe3,
	0x50, 0xa5, 0x2b, 0xfd, 0x57, 0x8d, 0xe5, 0xd9, 0xb3, 0xff, 0xbe, 0x44, 0x68, 0xd9, 0x8c, 0xa1,
	0xa7, 0xdf, 0x3c, 0x32, 0x59, 0x70, 0x84, 0x0c, 0x96, 0x94, 0x12, 0xff, 0xae, 0xd3, 0x8a, 0xb9,
	0x91, 0x1b, 0x4f, 0x33, 0xab, 0x07, 0xa6, 0x40, 0x69, 0x76, 0x6f, 0x64, 0x83, 0xa6, 0x97, 0x24,
	0x10, 0x35, 0x97, 0x4d, 0x2e, 0x4b, 0xe6, 0x59, 0x3e, 0xb1, 0xfe, 0x5d, 0x49, 0x9f, 0x93, 0x07,
	0x5c, 0x08, 0xdd, 0x37, 0x26, 0x6f, 0x7a, 0x55, 0x40, 0xc7, 0xfc, 0xc8, 0x8d, 0xfd, 0xec, 0xe2,
	0x40, 0xdf, 0x5b, 0x48, 0x67, 0x24, 0x40, 0x58, 0xf6, 0xd0, 0x08, 0x60, 0xf7, 0x6d, 0xe0, 0xe8,
	0x87, 0x09, 0x23, 0x15, 0xe8, 0xde, 0xe4, 0x35, 0xc8, 0xaa, 0x36, 0xec, 0x6c, 0x9c, 0x38, 0xd0,
	0xb7, 0x16, 0xd2, 0x1b, 0xe2, 0xdf, 0x01, 0x20, 0x9b, 0x44, 0x5e, 0x7c, 0xfe, 0xf2, 0x32, 0x39,
	0xf1, 0xa7, 0x92, 0xd7, 0x5a, 0x36, 0x0b, 0x7f, 0xf3, 0x73, 0xee, 0x64, 0x36, 0x4c, 0x6f, 0x09,
	0xa9, 0x38, 0xe6, 0x6d, 0x27, 0x05, 0x20, 0x0b, 0x6c, 0xf5, 0xc9, 0xc9, 0xea, 0x1b, 0x10, 0xff,
	0xb4, 0xa7, 0x15, 0xc7, 0x0f, 0xb6, 0x44, 0x1f, 0x12, 0xaf, 0xe2, 0xc8, 0xa6, 0xf6, 0xdc, 0x83,
	0x1c, 0x3e, 0x78, 0x18, 0xe5, 0xe5, 0xa7, 0x1e, 0x8d, 0x82, 0xc6, 0x30, 0x62, 0x1f, 0x5e, 0x54,
	0x1c, 0x6f, 0x8f, 0xd0, 0x9e, 0x59, 0xaa, 0xfe, 0x33, 0x37, 0xc0, 0xce, 0x23, 0x37, 0x0e, 0xb2,
	0xa3, 0x7f, 0x15, 0x7c, 0x5d, 0xcf, 0x9d, 0xdf, 0xeb, 0xb9, 0xb3, 0x78, 0xb1, 0xd9, 0x85, 0xee,
	0x76, 0x17, 0xba, 0xbf, 0x76, 0xa1, 0xfb, 0x7d, 0x1f, 0x3a, 0xdb, 0x7d, 0xe8, 0xfc, 0xd8, 0x87,
	0xce, 0xc7, 0x47, 0x28, 0x6a, 0x50, 0x3c, 0x35, 0x80, 0x26, 0x37, 0xab, 0x16, 0xb0, 0x38, 0xb3,
	0x77, 0x78, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x9b, 0x38, 0x31, 0x4f, 0x02, 0x00, 0x00,
}

func (m *BaseReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Simulate {
		i--
		if m.Simulate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.GasAdjustment) > 0 {
		i -= len(m.GasAdjustment)
		copy(dAtA[i:], m.GasAdjustment)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.GasAdjustment)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Gas) > 0 {
		i -= len(m.Gas)
		copy(dAtA[i:], m.Gas)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.Gas)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.GasPrices) > 0 {
		for iNdEx := len(m.GasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBasereq(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBasereq(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.Sequence != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if m.AccountNumber != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBasereq(dAtA []byte, offset int, v uint64) int {
	offset -= sovBasereq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovBasereq(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovBasereq(uint64(m.Sequence))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovBasereq(uint64(m.TimeoutHeight))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovBasereq(uint64(l))
		}
	}
	if len(m.GasPrices) > 0 {
		for _, e := range m.GasPrices {
			l = e.Size()
			n += 1 + l + sovBasereq(uint64(l))
		}
	}
	l = len(m.Gas)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.GasAdjustment)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	if m.Simulate {
		n += 2
	}
	return n
}

func sovBasereq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBasereq(x uint64) (n int) {
	return sovBasereq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasereq
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
			return fmt.Errorf("proto: BaseReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = append(m.GasPrices, types.DecCoin{})
			if err := m.GasPrices[len(m.GasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAdjustment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasAdjustment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Simulate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
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
			m.Simulate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBasereq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasereq
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
func skipBasereq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBasereq
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
					return 0, ErrIntOverflowBasereq
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
					return 0, ErrIntOverflowBasereq
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
				return 0, ErrInvalidLengthBasereq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBasereq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBasereq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBasereq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBasereq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBasereq = fmt.Errorf("proto: unexpected end of group")
)
