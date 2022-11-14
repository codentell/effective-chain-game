// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modulargmgamerollup/modulargmgamerollup/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateModularFellow struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Score   string `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"`
	X       string `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	Y       string `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *MsgCreateModularFellow) Reset()         { *m = MsgCreateModularFellow{} }
func (m *MsgCreateModularFellow) String() string { return proto.CompactTextString(m) }
func (*MsgCreateModularFellow) ProtoMessage()    {}
func (*MsgCreateModularFellow) Descriptor() ([]byte, []int) {
	return fileDescriptor_66361b525577d6ee, []int{0}
}
func (m *MsgCreateModularFellow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateModularFellow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateModularFellow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateModularFellow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateModularFellow.Merge(m, src)
}
func (m *MsgCreateModularFellow) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateModularFellow) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateModularFellow.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateModularFellow proto.InternalMessageInfo

func (m *MsgCreateModularFellow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateModularFellow) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *MsgCreateModularFellow) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *MsgCreateModularFellow) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

type MsgCreateModularFellowResponse struct {
}

func (m *MsgCreateModularFellowResponse) Reset()         { *m = MsgCreateModularFellowResponse{} }
func (m *MsgCreateModularFellowResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateModularFellowResponse) ProtoMessage()    {}
func (*MsgCreateModularFellowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66361b525577d6ee, []int{1}
}
func (m *MsgCreateModularFellowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateModularFellowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateModularFellowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateModularFellowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateModularFellowResponse.Merge(m, src)
}
func (m *MsgCreateModularFellowResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateModularFellowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateModularFellowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateModularFellowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateModularFellow)(nil), "modular_gm_game_rollup.modulargmgamerollup.MsgCreateModularFellow")
	proto.RegisterType((*MsgCreateModularFellowResponse)(nil), "modular_gm_game_rollup.modulargmgamerollup.MsgCreateModularFellowResponse")
}

func init() {
	proto.RegisterFile("modulargmgamerollup/modulargmgamerollup/tx.proto", fileDescriptor_66361b525577d6ee)
}

var fileDescriptor_66361b525577d6ee = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc8, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2c, 0x4a, 0xcf, 0x4d, 0x4f, 0xcc, 0x4d, 0x2d, 0xca, 0xcf, 0xc9, 0x29, 0x2d, 0xd0,
	0xc7, 0x26, 0x56, 0x52, 0xa1, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0xa4, 0x05, 0x95, 0x8d, 0x4f,
	0xcf, 0x8d, 0x07, 0xc9, 0xc7, 0x43, 0x14, 0xe8, 0x61, 0xd1, 0xa4, 0x94, 0xc2, 0x25, 0xe6, 0x5b,
	0x9c, 0xee, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0xea, 0x0b, 0x91, 0x77, 0x4b, 0xcd, 0xc9, 0xc9, 0x2f,
	0x17, 0x92, 0xe0, 0x62, 0x4f, 0x06, 0x09, 0xe7, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0xc5, 0xc9, 0xf9, 0x45, 0xa9, 0x12, 0x4c, 0x60, 0x71, 0x08,
	0x47, 0x88, 0x87, 0x8b, 0xb1, 0x42, 0x82, 0x19, 0x2c, 0xc2, 0x58, 0x01, 0xe2, 0x55, 0x4a, 0xb0,
	0x40, 0x78, 0x95, 0x4a, 0x0a, 0x5c, 0x72, 0xd8, 0x6d, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x35, 0x5a, 0xcb, 0xc8, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0xb4, 0x94, 0x91, 0x4b, 0x18, 0x9b,
	0x6b, 0x9c, 0xf4, 0x88, 0xf7, 0x94, 0x1e, 0x76, 0xbb, 0xa4, 0xbc, 0x28, 0x37, 0x03, 0xe6, 0x5e,
	0x27, 0x9f, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0xc2, 0x6e, 0x89,
	0x7e, 0x05, 0xf6, 0x48, 0xab, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0x9c, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x32, 0xcd, 0xaa, 0xec, 0x01, 0x00, 0x00,
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
	CreateModularFellow(ctx context.Context, in *MsgCreateModularFellow, opts ...grpc.CallOption) (*MsgCreateModularFellowResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateModularFellow(ctx context.Context, in *MsgCreateModularFellow, opts ...grpc.CallOption) (*MsgCreateModularFellowResponse, error) {
	out := new(MsgCreateModularFellowResponse)
	err := c.cc.Invoke(ctx, "/modular_gm_game_rollup.modulargmgamerollup.Msg/CreateModularFellow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateModularFellow(context.Context, *MsgCreateModularFellow) (*MsgCreateModularFellowResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateModularFellow(ctx context.Context, req *MsgCreateModularFellow) (*MsgCreateModularFellowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModularFellow not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateModularFellow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateModularFellow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateModularFellow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modular_gm_game_rollup.modulargmgamerollup.Msg/CreateModularFellow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateModularFellow(ctx, req.(*MsgCreateModularFellow))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modular_gm_game_rollup.modulargmgamerollup.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModularFellow",
			Handler:    _Msg_CreateModularFellow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modulargmgamerollup/modulargmgamerollup/tx.proto",
}

func (m *MsgCreateModularFellow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateModularFellow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateModularFellow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Y) > 0 {
		i -= len(m.Y)
		copy(dAtA[i:], m.Y)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Y)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarintTx(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateModularFellowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateModularFellowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateModularFellowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateModularFellow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Y)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateModularFellowResponse) Size() (n int) {
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
func (m *MsgCreateModularFellow) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateModularFellow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateModularFellow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
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
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
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
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
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
			m.Y = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateModularFellowResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateModularFellowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateModularFellowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
