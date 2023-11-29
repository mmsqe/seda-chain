// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/randomness/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgNewSeed struct {
	Seed     string `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Pi       string `protobuf:"bytes,2,opt,name=pi,proto3" json:"pi,omitempty"`
	Proposer string `protobuf:"bytes,3,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (m *MsgNewSeed) Reset()         { *m = MsgNewSeed{} }
func (m *MsgNewSeed) String() string { return proto.CompactTextString(m) }
func (*MsgNewSeed) ProtoMessage()    {}
func (*MsgNewSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_9575b460ec9dfc32, []int{0}
}
func (m *MsgNewSeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewSeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewSeed.Merge(m, src)
}
func (m *MsgNewSeed) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewSeed.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewSeed proto.InternalMessageInfo

func (m *MsgNewSeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *MsgNewSeed) GetPi() string {
	if m != nil {
		return m.Pi
	}
	return ""
}

func (m *MsgNewSeed) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

type MsgNewSeedResponse struct {
}

func (m *MsgNewSeedResponse) Reset()         { *m = MsgNewSeedResponse{} }
func (m *MsgNewSeedResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewSeedResponse) ProtoMessage()    {}
func (*MsgNewSeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9575b460ec9dfc32, []int{1}
}
func (m *MsgNewSeedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewSeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewSeedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewSeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewSeedResponse.Merge(m, src)
}
func (m *MsgNewSeedResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewSeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewSeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewSeedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgNewSeed)(nil), "sedachain.randomness.v1.MsgNewSeed")
	proto.RegisterType((*MsgNewSeedResponse)(nil), "sedachain.randomness.v1.MsgNewSeedResponse")
}

func init() { proto.RegisterFile("sedachain/randomness/v1/tx.proto", fileDescriptor_9575b460ec9dfc32) }

var fileDescriptor_9575b460ec9dfc32 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x4d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4a, 0xcc, 0x4b, 0xc9, 0xcf, 0xcd, 0x4b, 0x2d, 0x2e,
	0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xab,
	0xd0, 0x43, 0xa8, 0xd0, 0x2b, 0x33, 0x94, 0x12, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xcf,
	0x2d, 0x4e, 0x07, 0x69, 0xc8, 0x2d, 0x4e, 0x87, 0xe8, 0x50, 0x8a, 0xe6, 0xe2, 0xf2, 0x2d, 0x4e,
	0xf7, 0x4b, 0x2d, 0x0f, 0x4e, 0x4d, 0x4d, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x4e, 0x4d, 0x4d, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xf8, 0xb8, 0x98, 0x0a, 0x32, 0x25, 0x98,
	0xc0, 0x22, 0x4c, 0x05, 0x99, 0x42, 0x52, 0x5c, 0x1c, 0x05, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0xa9,
	0x45, 0x12, 0xcc, 0x60, 0x51, 0x38, 0xdf, 0x8a, 0xb7, 0xe9, 0xf9, 0x06, 0x2d, 0x38, 0x57, 0x49,
	0x84, 0x4b, 0x08, 0x61, 0x78, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x51, 0x3e, 0x17,
	0xb3, 0x6f, 0x71, 0xba, 0x50, 0x06, 0x97, 0x48, 0x70, 0x49, 0x7e, 0x51, 0xaa, 0x4b, 0x62, 0x49,
	0x62, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x49, 0x78, 0x62, 0x71, 0xae, 0x90, 0xb2, 0x1e, 0x0e,
	0x4f, 0xe8, 0x21, 0xcc, 0x92, 0xd2, 0x26, 0x42, 0x11, 0xcc, 0x42, 0x27, 0xff, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x07, 0x19, 0x08, 0x0e, 0x93, 0xe4, 0xfc, 0x1c, 0x30, 0x47, 0x17, 0x12, 0xd4,
	0x15, 0xc8, 0x81, 0x5d, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x56, 0x67, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x09, 0xda, 0x22, 0x00, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	StoreDataRequestWasm(ctx context.Context, in *MsgNewSeed, opts ...grpc.CallOption) (*MsgNewSeedResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StoreDataRequestWasm(ctx context.Context, in *MsgNewSeed, opts ...grpc.CallOption) (*MsgNewSeedResponse, error) {
	out := new(MsgNewSeedResponse)
	err := c.cc.Invoke(ctx, "/sedachain.randomness.v1.Msg/StoreDataRequestWasm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	StoreDataRequestWasm(context.Context, *MsgNewSeed) (*MsgNewSeedResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StoreDataRequestWasm(ctx context.Context, req *MsgNewSeed) (*MsgNewSeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDataRequestWasm not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StoreDataRequestWasm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewSeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreDataRequestWasm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sedachain.randomness.v1.Msg/StoreDataRequestWasm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreDataRequestWasm(ctx, req.(*MsgNewSeed))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sedachain.randomness.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreDataRequestWasm",
			Handler:    _Msg_StoreDataRequestWasm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sedachain/randomness/v1/tx.proto",
}

func (m *MsgNewSeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewSeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewSeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Pi) > 0 {
		i -= len(m.Pi)
		copy(dAtA[i:], m.Pi)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Pi)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewSeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewSeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewSeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgNewSeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Pi)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewSeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgNewSeed) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewSeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewSeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
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
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pi", wireType)
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
			m.Pi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
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
			m.Proposer = string(dAtA[iNdEx:postIndex])
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
func (m *MsgNewSeedResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewSeedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewSeedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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