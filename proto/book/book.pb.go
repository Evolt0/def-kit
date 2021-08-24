// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/book/book.proto

package book

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 包含id的请求
type BookRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRequest) Reset()         { *m = BookRequest{} }
func (m *BookRequest) String() string { return proto.CompactTextString(m) }
func (*BookRequest) ProtoMessage()    {}
func (*BookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98814f1514a537b6, []int{0}
}
func (m *BookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRequest.Merge(m, src)
}
func (m *BookRequest) XXX_Size() int {
	return m.Size()
}
func (m *BookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookRequest proto.InternalMessageInfo

func (m *BookRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 包含名称的响应消息
type BookResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookResponse) Reset()         { *m = BookResponse{} }
func (m *BookResponse) String() string { return proto.CompactTextString(m) }
func (*BookResponse) ProtoMessage()    {}
func (*BookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98814f1514a537b6, []int{1}
}
func (m *BookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookResponse.Merge(m, src)
}
func (m *BookResponse) XXX_Size() int {
	return m.Size()
}
func (m *BookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookResponse proto.InternalMessageInfo

func (m *BookResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*BookRequest)(nil), "book.BookRequest")
	proto.RegisterType((*BookResponse)(nil), "book.BookResponse")
}

func init() { proto.RegisterFile("proto/book/book.proto", fileDescriptor_98814f1514a537b6) }

var fileDescriptor_98814f1514a537b6 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xca, 0xcf, 0xcf, 0x06, 0x13, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0x24,
	0xcb, 0xc5, 0xed, 0x94, 0x9f, 0x9f, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xc4, 0xc7,
	0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x94, 0x99, 0xa2, 0xa4, 0xc4,
	0xc5, 0x03, 0x91, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x05, 0xab, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0x5c, 0xb9, 0xd8, 0x41, 0x6a, 0xdc, 0x4a, 0xf3,
	0x84, 0xac, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x40, 0x3c, 0xcf, 0xbc, 0xb4, 0x7c, 0xa7, 0x4a, 0x4f,
	0x17, 0x21, 0x41, 0x3d, 0xb0, 0x9d, 0x48, 0x96, 0x48, 0x09, 0x21, 0x0b, 0x41, 0x0c, 0x56, 0x62,
	0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67,
	0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x3b, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x61,
	0x51, 0xcc, 0xc1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookFunClient is the client API for BookFun service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookFunClient interface {
	GetBookInfoByID(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
}

type bookFunClient struct {
	cc *grpc.ClientConn
}

func NewBookFunClient(cc *grpc.ClientConn) BookFunClient {
	return &bookFunClient{cc}
}

func (c *bookFunClient) GetBookInfoByID(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/book.BookFun/GetBookInfoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookFunServer is the server API for BookFun service.
type BookFunServer interface {
	GetBookInfoByID(context.Context, *BookRequest) (*BookResponse, error)
}

// UnimplementedBookFunServer can be embedded to have forward compatible implementations.
type UnimplementedBookFunServer struct {
}

func (*UnimplementedBookFunServer) GetBookInfoByID(ctx context.Context, req *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfoByID not implemented")
}

func RegisterBookFunServer(s *grpc.Server, srv BookFunServer) {
	s.RegisterService(&_BookFun_serviceDesc, srv)
}

func _BookFun_GetBookInfoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookFunServer).GetBookInfoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookFun/GetBookInfoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookFunServer).GetBookInfoByID(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookFun_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookFun",
	HandlerType: (*BookFunServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfoByID",
			Handler:    _BookFun_GetBookInfoByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book/book.proto",
}

func (m *BookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintBook(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBook(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBook(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBook(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBook(x uint64) (n int) {
	return sovBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BookRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
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
			return fmt.Errorf("proto: BookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BookResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
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
			return fmt.Errorf("proto: BookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
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
				return ErrInvalidLengthBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBook
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
					return 0, ErrIntOverflowBook
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
					return 0, ErrIntOverflowBook
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
				return 0, ErrInvalidLengthBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBook = fmt.Errorf("proto: unexpected end of group")
)
