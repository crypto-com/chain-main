// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supply/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// SupplyRequest is the request type for the Query/TotalSupply RPC
// method.
type SupplyRequest struct {
}

func (m *SupplyRequest) Reset()         { *m = SupplyRequest{} }
func (m *SupplyRequest) String() string { return proto.CompactTextString(m) }
func (*SupplyRequest) ProtoMessage()    {}
func (*SupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cd1add7eb523c6, []int{0}
}
func (m *SupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyRequest.Merge(m, src)
}
func (m *SupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *SupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyRequest proto.InternalMessageInfo

// SupplyResponse is the response type for the Query/TotalSupply RPC
// method
type SupplyResponse struct {
	// supply is the supply of the coins
	Supply github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`
}

func (m *SupplyResponse) Reset()         { *m = SupplyResponse{} }
func (m *SupplyResponse) String() string { return proto.CompactTextString(m) }
func (*SupplyResponse) ProtoMessage()    {}
func (*SupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cd1add7eb523c6, []int{1}
}
func (m *SupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyResponse.Merge(m, src)
}
func (m *SupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *SupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyResponse proto.InternalMessageInfo

func (m *SupplyResponse) GetSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Supply
	}
	return nil
}

func init() {
	proto.RegisterType((*SupplyRequest)(nil), "chainmain.supply.v1.SupplyRequest")
	proto.RegisterType((*SupplyResponse)(nil), "chainmain.supply.v1.SupplyResponse")
}

func init() { proto.RegisterFile("supply/v1/query.proto", fileDescriptor_64cd1add7eb523c6) }

var fileDescriptor_64cd1add7eb523c6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0x93, 0x5e, 0xdd, 0x0e, 0xe9, 0xfd, 0x23, 0xe5, 0x5e, 0x24, 0x08, 0x25, 0x45, 0xe9,
	0xd2, 0x25, 0x36, 0x29, 0xe2, 0x05, 0xca, 0x8a, 0x84, 0x28, 0x4c, 0x6c, 0x4e, 0x6a, 0xa5, 0x16,
	0xa9, 0x4f, 0x1a, 0x3b, 0x15, 0x41, 0x2c, 0xf0, 0x04, 0x48, 0xbc, 0x05, 0x4f, 0xd2, 0xb1, 0x12,
	0x0b, 0x13, 0xa0, 0x96, 0x07, 0x41, 0xb1, 0x53, 0xfe, 0x48, 0x15, 0x13, 0x93, 0x8f, 0x8e, 0xcf,
	0xf9, 0xbe, 0x9f, 0x3f, 0x5b, 0x6b, 0x22, 0x4f, 0xd3, 0xa4, 0xc0, 0x93, 0x00, 0x8f, 0x73, 0x9a,
	0x15, 0x28, 0xcd, 0x40, 0x82, 0xfd, 0x2f, 0x1a, 0x12, 0xc6, 0x47, 0x84, 0x71, 0xa4, 0x07, 0xd0,
	0x24, 0x70, 0xdc, 0x08, 0xc4, 0x08, 0x04, 0x0e, 0x89, 0xa0, 0x78, 0x12, 0x84, 0x54, 0x92, 0x00,
	0x47, 0xc0, 0xb8, 0x5e, 0x72, 0xfe, 0xc7, 0x10, 0x83, 0x2a, 0x71, 0x59, 0x55, 0xdd, 0x66, 0x0c,
	0x10, 0x27, 0x14, 0x93, 0x94, 0x61, 0xc2, 0x39, 0x48, 0x22, 0x19, 0x70, 0xa1, 0x6f, 0xbd, 0xbf,
	0xd6, 0xef, 0x63, 0x65, 0xd0, 0xa7, 0xe3, 0x9c, 0x0a, 0xe9, 0xe5, 0xd6, 0x9f, 0x65, 0x43, 0xa4,
	0xc0, 0x05, 0xb5, 0x23, 0xab, 0xae, 0x19, 0xd6, 0xcd, 0xed, 0x1f, 0x9d, 0x46, 0x77, 0x03, 0x69,
	0x0e, 0x54, 0x72, 0xa0, 0x8a, 0x03, 0xed, 0x03, 0xe3, 0xbd, 0x9d, 0xe9, 0x63, 0xcb, 0xb8, 0x7b,
	0x6a, 0x75, 0x62, 0x26, 0x87, 0x79, 0x88, 0x22, 0x18, 0xe1, 0x0a, 0x5a, 0x1f, 0xbe, 0x18, 0x9c,
	0x61, 0x59, 0xa4, 0x54, 0xa8, 0x05, 0xd1, 0xaf, 0xa4, 0xbb, 0x57, 0x35, 0xeb, 0xe7, 0x51, 0x19,
	0x80, 0x7d, 0x61, 0x35, 0x4e, 0x40, 0x92, 0x44, 0x53, 0xd8, 0x1e, 0x5a, 0x11, 0x05, 0xfa, 0xc4,
	0xec, 0xb4, 0xbf, 0x9c, 0xd1, 0xcf, 0xf0, 0xbc, 0xeb, 0xfb, 0x97, 0xdb, 0x5a, 0xd3, 0x76, 0xf0,
	0xdb, 0x30, 0x7e, 0x0f, 0x5f, 0x96, 0x96, 0xf6, 0xa5, 0xf5, 0xeb, 0x80, 0x8d, 0x73, 0x36, 0xf8,
	0x6e, 0xf3, 0xb6, 0x32, 0xdf, 0xb2, 0x37, 0x57, 0x9a, 0x27, 0xca, 0xb3, 0x77, 0x38, 0x9d, 0xbb,
	0xe6, 0x6c, 0xee, 0x9a, 0xcf, 0x73, 0xd7, 0xbc, 0x59, 0xb8, 0xc6, 0x6c, 0xe1, 0x1a, 0x0f, 0x0b,
	0xd7, 0x38, 0xdd, 0xfb, 0x98, 0x67, 0x56, 0xa4, 0x12, 0x7c, 0xc8, 0x62, 0x5f, 0x69, 0x69, 0x45,
	0x5f, 0x49, 0x9e, 0x2f, 0x45, 0x55, 0xc4, 0x61, 0x5d, 0xfd, 0xf1, 0xee, 0x6b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x1d, 0xb8, 0x1f, 0x65, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// TotalSupply queries the total supply of all coins.
	TotalSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
	// LiquidSupply queries the liquid supply of all coins.
	LiquidSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TotalSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/chainmain.supply.v1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/chainmain.supply.v1.Query/LiquidSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// TotalSupply queries the total supply of all coins.
	TotalSupply(context.Context, *SupplyRequest) (*SupplyResponse, error)
	// LiquidSupply queries the liquid supply of all coins.
	LiquidSupply(context.Context, *SupplyRequest) (*SupplyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) LiquidSupply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidSupply not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.supply.v1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.supply.v1.Query/LiquidSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidSupply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainmain.supply.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "LiquidSupply",
			Handler:    _Query_LiquidSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supply/v1/query.proto",
}

func (m *SupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Supply) > 0 {
		for iNdEx := len(m.Supply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Supply) > 0 {
		for _, e := range m.Supply {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SupplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *SupplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supply = append(m.Supply, types.Coin{})
			if err := m.Supply[len(m.Supply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
