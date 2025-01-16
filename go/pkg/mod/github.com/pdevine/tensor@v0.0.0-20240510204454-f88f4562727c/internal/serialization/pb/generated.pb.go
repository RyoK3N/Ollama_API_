// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gorgonia.org/tensor/internal/serialization/pb/generated.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		gorgonia.org/tensor/internal/serialization/pb/generated.proto

	It has these top-level messages:
		AP
		Dense
		MaskedDense
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var Triangle_name = map[int32]string{
	0: "NOT_TRIANGLE",
	1: "UPPER",
	2: "LOWER",
	3: "SYMMETRIC",
}
var Triangle_value = map[string]int32{
	"NOT_TRIANGLE": 0,
	"UPPER":        1,
	"LOWER":        2,
	"SYMMETRIC":    3,
}

func (Triangle) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *AP) Reset()                    { *m = AP{} }
func (m *AP) String() string            { return proto.CompactTextString(m) }
func (*AP) ProtoMessage()               {}
func (*AP) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Dense) Reset()                    { *m = Dense{} }
func (m *Dense) String() string            { return proto.CompactTextString(m) }
func (*Dense) ProtoMessage()               {}
func (*Dense) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *MaskedDense) Reset()                    { *m = MaskedDense{} }
func (m *MaskedDense) String() string            { return proto.CompactTextString(m) }
func (*MaskedDense) ProtoMessage()               {}
func (*MaskedDense) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*AP)(nil), "gorgonia.org.tensor.internal.serialization.pb.AP")
	proto.RegisterType((*Dense)(nil), "gorgonia.org.tensor.internal.serialization.pb.Dense")
	proto.RegisterType((*MaskedDense)(nil), "gorgonia.org.tensor.internal.serialization.pb.MaskedDense")
	proto.RegisterEnum("gorgonia.org.tensor.internal.serialization.pb.Triangle", Triangle_name, Triangle_value)
}
func (m *AP) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AP) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Shape) > 0 {
		dAtA2 := make([]byte, len(m.Shape)*10)
		var j1 int
		for _, num1 := range m.Shape {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Strides) > 0 {
		dAtA4 := make([]byte, len(m.Strides)*10)
		var j3 int
		for _, num1 := range m.Strides {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if m.O != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.O))
	}
	if m.T != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.T))
	}
	return i, nil
}

func (m *Dense) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dense) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Shape) > 0 {
		dAtA6 := make([]byte, len(m.Shape)*10)
		var j5 int
		for _, num1 := range m.Shape {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j5))
		i += copy(dAtA[i:], dAtA6[:j5])
	}
	if len(m.Strides) > 0 {
		dAtA8 := make([]byte, len(m.Strides)*10)
		var j7 int
		for _, num1 := range m.Strides {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA8[j7] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j7++
			}
			dAtA8[j7] = uint8(num)
			j7++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j7))
		i += copy(dAtA[i:], dAtA8[:j7])
	}
	if m.O != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.O))
	}
	if m.T != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.T))
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *MaskedDense) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MaskedDense) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Shape) > 0 {
		dAtA10 := make([]byte, len(m.Shape)*10)
		var j9 int
		for _, num1 := range m.Shape {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA10[j9] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j9++
			}
			dAtA10[j9] = uint8(num)
			j9++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j9))
		i += copy(dAtA[i:], dAtA10[:j9])
	}
	if len(m.Strides) > 0 {
		dAtA12 := make([]byte, len(m.Strides)*10)
		var j11 int
		for _, num1 := range m.Strides {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA12[j11] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j11++
			}
			dAtA12[j11] = uint8(num)
			j11++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j11))
		i += copy(dAtA[i:], dAtA12[:j11])
	}
	if m.O != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.O))
	}
	if m.T != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.T))
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Mask) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Mask)))
		for _, b := range m.Mask {
			if b {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i++
		}
	}
	if len(m.MaskIsSoft) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.MaskIsSoft)))
		for _, b := range m.MaskIsSoft {
			if b {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i++
		}
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AP) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.Shape) > 0 {
		l = 0
		for _, e := range m.Shape {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if len(m.Strides) > 0 {
		l = 0
		for _, e := range m.Strides {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if m.O != 0 {
		n += 1 + sovGenerated(uint64(m.O))
	}
	if m.T != 0 {
		n += 1 + sovGenerated(uint64(m.T))
	}
	return n
}

func (m *Dense) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.Shape) > 0 {
		l = 0
		for _, e := range m.Shape {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if len(m.Strides) > 0 {
		l = 0
		for _, e := range m.Strides {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if m.O != 0 {
		n += 1 + sovGenerated(uint64(m.O))
	}
	if m.T != 0 {
		n += 1 + sovGenerated(uint64(m.T))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *MaskedDense) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.Shape) > 0 {
		l = 0
		for _, e := range m.Shape {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if len(m.Strides) > 0 {
		l = 0
		for _, e := range m.Strides {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	if m.O != 0 {
		n += 1 + sovGenerated(uint64(m.O))
	}
	if m.T != 0 {
		n += 1 + sovGenerated(uint64(m.T))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Mask) > 0 {
		n += 1 + sovGenerated(uint64(len(m.Mask))) + len(m.Mask)*1
	}
	if len(m.MaskIsSoft) > 0 {
		n += 1 + sovGenerated(uint64(len(m.MaskIsSoft))) + len(m.MaskIsSoft)*1
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Shape = append(m.Shape, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Shape = append(m.Shape, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Strides = append(m.Strides, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Strides = append(m.Strides, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Strides", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field O", wireType)
			}
			m.O = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.O |= (DataOrder(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			m.T = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.T |= (Triangle(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Dense) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dense: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dense: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Shape = append(m.Shape, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Shape = append(m.Shape, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Strides = append(m.Strides, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Strides = append(m.Strides, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Strides", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field O", wireType)
			}
			m.O = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.O |= (DataOrder(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			m.T = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.T |= (Triangle(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *MaskedDense) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MaskedDense: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaskedDense: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Shape = append(m.Shape, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Shape = append(m.Shape, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Strides = append(m.Strides, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Strides = append(m.Strides, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Strides", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field O", wireType)
			}
			m.O = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.O |= (DataOrder(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			m.T = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.T |= (Triangle(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Mask = append(m.Mask, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Mask = append(m.Mask, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Mask", wireType)
			}
		case 8:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MaskIsSoft = append(m.MaskIsSoft, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MaskIsSoft = append(m.MaskIsSoft, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MaskIsSoft", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/pdevine/tensor/internal/serialization/pb/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x77, 0x92, 0x6e, 0x9b, 0x4c, 0x13, 0x0d, 0x43, 0x0f, 0x43, 0x94, 0xcd, 0xd8, 0xd3,
	0x22, 0x74, 0x17, 0xf4, 0x20, 0x08, 0x1e, 0x5a, 0x1b, 0x24, 0x90, 0x7f, 0x4c, 0x52, 0x44, 0x2f,
	0x61, 0xb6, 0x3b, 0xd9, 0x0e, 0x4d, 0x76, 0x96, 0x99, 0x29, 0x52, 0xef, 0x42, 0xcd, 0x27, 0xf0,
	0x12, 0xa8, 0xda, 0x83, 0x1f, 0xc3, 0xa3, 0x17, 0xc1, 0x4f, 0x20, 0x92, 0x7e, 0x01, 0xcf, 0x9e,
	0x64, 0x27, 0x44, 0xe2, 0xd1, 0x9b, 0x3d, 0xcd, 0xf3, 0xfc, 0x66, 0x9e, 0x77, 0xde, 0x97, 0x61,
	0xe0, 0x93, 0x44, 0xaa, 0x44, 0xa6, 0x82, 0x05, 0x52, 0x25, 0xa1, 0xe1, 0xa9, 0x96, 0x2a, 0x14,
	0xa9, 0xe1, 0x2a, 0x65, 0x93, 0x50, 0x73, 0x25, 0xd8, 0x44, 0xbc, 0x66, 0x46, 0xc8, 0x34, 0xcc,
	0xa2, 0x30, 0xe1, 0x29, 0x57, 0xcc, 0xf0, 0x38, 0xc8, 0x94, 0x34, 0x12, 0xed, 0xad, 0xc7, 0x83,
	0x65, 0x3c, 0x58, 0xc5, 0x83, 0xbf, 0xe2, 0x41, 0x16, 0xd5, 0xf7, 0x12, 0x61, 0x4e, 0xce, 0xa2,
	0xe0, 0x58, 0x4e, 0xc3, 0x44, 0x26, 0x32, 0xb4, 0x55, 0xa2, 0xb3, 0xb1, 0x75, 0xd6, 0x58, 0xb5,
	0xac, 0xbe, 0xfb, 0x01, 0xc0, 0xc2, 0x7e, 0x1f, 0xed, 0x40, 0x57, 0x9f, 0xb0, 0x8c, 0x63, 0x40,
	0x8a, 0xbe, 0x4b, 0x97, 0x06, 0x61, 0xb8, 0xa5, 0x8d, 0x12, 0x31, 0xd7, 0xb8, 0x60, 0xf9, 0xca,
	0xa2, 0x3b, 0x10, 0x48, 0x5c, 0x24, 0xc0, 0xaf, 0x1e, 0x54, 0x7f, 0x7d, 0x6f, 0x94, 0x0f, 0x99,
	0x61, 0x3d, 0x15, 0x73, 0x45, 0x81, 0x44, 0x4d, 0x08, 0x0c, 0xde, 0x20, 0xc0, 0xbf, 0xf5, 0xe0,
	0x51, 0xf0, 0x4f, 0xdd, 0x07, 0x43, 0x25, 0x58, 0x9a, 0x4c, 0x38, 0x05, 0xe6, 0x71, 0xe9, 0xe2,
	0xb2, 0xe1, 0xfc, 0x7c, 0xdf, 0x70, 0x76, 0xbf, 0x02, 0xe8, 0x1e, 0xf2, 0x54, 0xf3, 0xff, 0xb1,
	0x4f, 0x84, 0xe0, 0x86, 0x39, 0xcf, 0x38, 0x76, 0x09, 0xf0, 0xcb, 0xd4, 0xea, 0x9c, 0xc5, 0xcc,
	0x30, 0xbc, 0x49, 0x80, 0x5f, 0xa1, 0x56, 0xaf, 0xcd, 0xf3, 0xb6, 0x00, 0xb7, 0x3b, 0x4c, 0x9f,
	0xf2, 0xf8, 0xc6, 0x4f, 0x95, 0xb3, 0x29, 0xd3, 0xa7, 0x78, 0x8b, 0x14, 0xfd, 0x12, 0xb5, 0x1a,
	0x11, 0x58, 0xc9, 0xd7, 0x91, 0xd0, 0x23, 0x2d, 0xc7, 0x06, 0x97, 0xec, 0x1e, 0xcc, 0x59, 0x4b,
	0x0f, 0xe4, 0x78, 0xed, 0x6d, 0xef, 0xbf, 0x01, 0xb0, 0xb4, 0xba, 0x17, 0xdd, 0x83, 0x95, 0x6e,
	0x6f, 0x38, 0x1a, 0xd2, 0xd6, 0x7e, 0xf7, 0x59, 0xbb, 0x59, 0x73, 0xea, 0xb7, 0x67, 0x73, 0xb2,
	0xdd, 0x95, 0xe6, 0xcf, 0x91, 0x1d, 0xe8, 0x1e, 0xf5, 0xfb, 0x4d, 0x5a, 0x03, 0xf5, 0xf2, 0x6c,
	0x4e, 0xdc, 0xa3, 0x2c, 0xe3, 0x2a, 0xa7, 0xed, 0xde, 0xf3, 0x26, 0xad, 0x15, 0x96, 0xb4, 0x2d,
	0x5f, 0x71, 0x85, 0xee, 0xc2, 0xf2, 0xe0, 0x45, 0xa7, 0xd3, 0x1c, 0xd2, 0xd6, 0xd3, 0x5a, 0xb1,
	0x5e, 0x9d, 0xcd, 0x49, 0x79, 0x70, 0x3e, 0x9d, 0x72, 0xa3, 0xc4, 0x71, 0xbd, 0x72, 0xf1, 0xd1,
	0x73, 0x3e, 0x5d, 0x79, 0xce, 0xe7, 0x2b, 0xcf, 0x39, 0xc0, 0x5f, 0x16, 0x1e, 0xf8, 0xb6, 0xf0,
	0xc0, 0x8f, 0x85, 0xe7, 0xbc, 0xbb, 0xf6, 0x9c, 0xcb, 0x6b, 0x0f, 0xbc, 0x2c, 0x64, 0x51, 0xb4,
	0x69, 0x7f, 0xca, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xff, 0xbb, 0x8f, 0xc8, 0x03,
	0x00, 0x00,
}
