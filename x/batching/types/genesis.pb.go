// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/batching/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the batching module's genesis state.
type GenesisState struct {
	// current_batch_number is the batch number of the most recently-
	// created batch.
	CurrentBatchNumber uint64            `protobuf:"varint,1,opt,name=current_batch_number,json=currentBatchNumber,proto3" json:"current_batch_number,omitempty"`
	Batches            []Batch           `protobuf:"bytes,2,rep,name=batches,proto3" json:"batches"`
	TreeEntries        []TreeEntries     `protobuf:"bytes,3,rep,name=tree_entries,json=treeEntries,proto3" json:"tree_entries"`
	DataResults        []DataResult      `protobuf:"bytes,4,rep,name=data_results,json=dataResults,proto3" json:"data_results"`
	BatchAssignments   []BatchAssignment `protobuf:"bytes,5,rep,name=batch_assignments,json=batchAssignments,proto3" json:"batch_assignments"`
	Params             Params            `protobuf:"bytes,6,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_eccca5d98d3cb479, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCurrentBatchNumber() uint64 {
	if m != nil {
		return m.CurrentBatchNumber
	}
	return 0
}

func (m *GenesisState) GetBatches() []Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *GenesisState) GetTreeEntries() []TreeEntries {
	if m != nil {
		return m.TreeEntries
	}
	return nil
}

func (m *GenesisState) GetDataResults() []DataResult {
	if m != nil {
		return m.DataResults
	}
	return nil
}

func (m *GenesisState) GetBatchAssignments() []BatchAssignment {
	if m != nil {
		return m.BatchAssignments
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// BatchAssignment represents a batch assignment for genesis export
// and import.
type BatchAssignment struct {
	BatchNumber   uint64 `protobuf:"varint,1,opt,name=batch_number,json=batchNumber,proto3" json:"batch_number,omitempty"`
	DataRequestId string `protobuf:"bytes,2,opt,name=data_request_id,json=dataRequestId,proto3" json:"data_request_id,omitempty"`
}

func (m *BatchAssignment) Reset()         { *m = BatchAssignment{} }
func (m *BatchAssignment) String() string { return proto.CompactTextString(m) }
func (*BatchAssignment) ProtoMessage()    {}
func (*BatchAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_eccca5d98d3cb479, []int{1}
}
func (m *BatchAssignment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchAssignment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchAssignment.Merge(m, src)
}
func (m *BatchAssignment) XXX_Size() int {
	return m.Size()
}
func (m *BatchAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_BatchAssignment proto.InternalMessageInfo

func (m *BatchAssignment) GetBatchNumber() uint64 {
	if m != nil {
		return m.BatchNumber
	}
	return 0
}

func (m *BatchAssignment) GetDataRequestId() string {
	if m != nil {
		return m.DataRequestId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sedachain.batching.v1.GenesisState")
	proto.RegisterType((*BatchAssignment)(nil), "sedachain.batching.v1.BatchAssignment")
}

func init() {
	proto.RegisterFile("sedachain/batching/v1/genesis.proto", fileDescriptor_eccca5d98d3cb479)
}

var fileDescriptor_eccca5d98d3cb479 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0x87, 0xed, 0x24, 0x04, 0x31, 0x36, 0x0a, 0x8c, 0x82, 0x64, 0x45, 0x60, 0x9c, 0x80, 0x22,
	0x6f, 0xb0, 0x49, 0xb2, 0x84, 0x0d, 0x11, 0x08, 0x01, 0x02, 0x21, 0xc3, 0xa6, 0x55, 0x25, 0x6b,
	0x6c, 0x3f, 0x39, 0x96, 0xe2, 0x71, 0x3a, 0x33, 0x8e, 0xda, 0x5b, 0xf4, 0x26, 0xbd, 0x46, 0x96,
	0x59, 0x76, 0x55, 0x55, 0xc9, 0x45, 0xaa, 0x8c, 0x1d, 0xa7, 0xad, 0x92, 0xee, 0xec, 0xf7, 0xbe,
	0xdf, 0x37, 0xff, 0x1e, 0x7a, 0xc7, 0x21, 0x22, 0xe1, 0x84, 0x24, 0xd4, 0x0d, 0x88, 0x08, 0x27,
	0x09, 0x8d, 0xdd, 0xf9, 0xc0, 0x8d, 0x81, 0x02, 0x4f, 0xb8, 0x33, 0x63, 0x99, 0xc8, 0xf0, 0xab,
	0x0a, 0x72, 0xb6, 0x90, 0x33, 0x1f, 0x74, 0xda, 0x71, 0x16, 0x67, 0x92, 0x70, 0x37, 0x5f, 0x05,
	0xdc, 0x79, 0xbf, 0xdf, 0x58, 0x05, 0x25, 0xd5, 0xbb, 0xac, 0x23, 0xfd, 0x7b, 0xb1, 0xc8, 0x3f,
	0x41, 0x04, 0xe0, 0x8f, 0xa8, 0x1d, 0xe6, 0x8c, 0x01, 0x15, 0xbe, 0x44, 0x7d, 0x9a, 0xa7, 0x01,
	0x30, 0x43, 0xb5, 0x54, 0xbb, 0xe1, 0xe1, 0xb2, 0x37, 0xde, 0xb4, 0xfe, 0xc8, 0x0e, 0xfe, 0x8c,
	0x9e, 0x4a, 0x12, 0xb8, 0x51, 0xb3, 0xea, 0xb6, 0x36, 0x7c, 0xed, 0xec, 0xdd, 0xa7, 0x23, 0x43,
	0xe3, 0xc6, 0xe2, 0xfa, 0xad, 0xe2, 0x6d, 0x23, 0xf8, 0x17, 0xd2, 0x05, 0x03, 0xf0, 0x81, 0x0a,
	0x96, 0x00, 0x37, 0xea, 0x52, 0xd1, 0x3b, 0xa0, 0xf8, 0xcf, 0x00, 0xbe, 0x15, 0x64, 0x29, 0xd2,
	0xc4, 0xae, 0x84, 0x7f, 0x22, 0x3d, 0x22, 0x82, 0xf8, 0x0c, 0x78, 0x3e, 0x15, 0xdc, 0x68, 0x48,
	0x59, 0xf7, 0x80, 0xec, 0x2b, 0x11, 0xc4, 0x93, 0xe4, 0xd6, 0x15, 0x55, 0x15, 0x8e, 0x8f, 0xd0,
	0xcb, 0xe2, 0x02, 0x08, 0xe7, 0x49, 0x4c, 0x53, 0xa0, 0x82, 0x1b, 0x4f, 0xa4, 0xb0, 0xff, 0xd8,
	0x01, 0xbf, 0x54, 0x78, 0x69, 0x7d, 0x11, 0xdc, 0x2f, 0x73, 0xfc, 0x09, 0x35, 0x67, 0x84, 0x91,
	0x94, 0x1b, 0x4d, 0x4b, 0xb5, 0xb5, 0xe1, 0x9b, 0x03, 0xbe, 0xbf, 0x12, 0x2a, 0x35, 0x65, 0xa4,
	0x77, 0x82, 0x5a, 0x0f, 0xd6, 0xc1, 0x5d, 0xa4, 0xef, 0x79, 0x2b, 0x2d, 0xb8, 0xf3, 0x48, 0x7d,
	0xd4, 0x2a, 0x6f, 0xe6, 0x34, 0x07, 0x2e, 0xfc, 0x24, 0x32, 0x6a, 0x96, 0x6a, 0x3f, 0xf3, 0x9e,
	0x17, 0x67, 0x96, 0xd5, 0x1f, 0xd1, 0xf8, 0xf7, 0x62, 0x65, 0xaa, 0xcb, 0x95, 0xa9, 0xde, 0xac,
	0x4c, 0xf5, 0x62, 0x6d, 0x2a, 0xcb, 0xb5, 0xa9, 0x5c, 0xad, 0x4d, 0xe5, 0x78, 0x14, 0x27, 0x62,
	0x92, 0x07, 0x4e, 0x98, 0xa5, 0xee, 0x66, 0xbb, 0x72, 0x7e, 0xc2, 0x6c, 0x2a, 0x7f, 0x3e, 0x14,
	0x83, 0x76, 0xb6, 0x1b, 0x35, 0x71, 0x3e, 0x03, 0x1e, 0x34, 0x25, 0x35, 0xba, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xed, 0x8e, 0x51, 0x01, 0xdf, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.BatchAssignments) > 0 {
		for iNdEx := len(m.BatchAssignments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchAssignments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DataResults) > 0 {
		for iNdEx := len(m.DataResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TreeEntries) > 0 {
		for iNdEx := len(m.TreeEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TreeEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CurrentBatchNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CurrentBatchNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BatchAssignment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchAssignment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchAssignment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataRequestId) > 0 {
		i -= len(m.DataRequestId)
		copy(dAtA[i:], m.DataRequestId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DataRequestId)))
		i--
		dAtA[i] = 0x12
	}
	if m.BatchNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BatchNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentBatchNumber != 0 {
		n += 1 + sovGenesis(uint64(m.CurrentBatchNumber))
	}
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TreeEntries) > 0 {
		for _, e := range m.TreeEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DataResults) > 0 {
		for _, e := range m.DataResults {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchAssignments) > 0 {
		for _, e := range m.BatchAssignments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *BatchAssignment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchNumber != 0 {
		n += 1 + sovGenesis(uint64(m.BatchNumber))
	}
	l = len(m.DataRequestId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBatchNumber", wireType)
			}
			m.CurrentBatchNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentBatchNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batches = append(m.Batches, Batch{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeEntries = append(m.TreeEntries, TreeEntries{})
			if err := m.TreeEntries[len(m.TreeEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataResults = append(m.DataResults, DataResult{})
			if err := m.DataResults[len(m.DataResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchAssignments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchAssignments = append(m.BatchAssignments, BatchAssignment{})
			if err := m.BatchAssignments[len(m.BatchAssignments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *BatchAssignment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: BatchAssignment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchAssignment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchNumber", wireType)
			}
			m.BatchNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
