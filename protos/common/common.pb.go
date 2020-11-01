// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common // import "github.com/hyperledger/fabric/protos/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_NOT_IMPLEMENTED          Status = 501
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"NOT_IMPLEMENTED":          501,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{0}
}

type HeaderType int32

const (
	HeaderType_MESSAGE              HeaderType = 0
	HeaderType_CONFIG               HeaderType = 1
	HeaderType_CONFIG_UPDATE        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION HeaderType = 3
	HeaderType_ORDERER_TRANSACTION  HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO    HeaderType = 5
	HeaderType_CHAINCODE_PACKAGE    HeaderType = 6
	HeaderType_PEER_ADMIN_OPERATION HeaderType = 8
	HeaderType_TOKEN_TRANSACTION    HeaderType = 9
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",
	2: "CONFIG_UPDATE",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
	6: "CHAINCODE_PACKAGE",
	8: "PEER_ADMIN_OPERATION",
	9: "TOKEN_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":              0,
	"CONFIG":               1,
	"CONFIG_UPDATE":        2,
	"ENDORSER_TRANSACTION": 3,
	"ORDERER_TRANSACTION":  4,
	"DELIVER_SEEK_INFO":    5,
	"CHAINCODE_PACKAGE":    6,
	"PEER_ADMIN_OPERATION": 8,
	"TOKEN_TRANSACTION":    9,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{1}
}

// This enum enlists indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIG         BlockMetadataIndex = 1
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
	BlockMetadataIndex_ORDERER             BlockMetadataIndex = 3
	BlockMetadataIndex_COMMIT_HASH         BlockMetadataIndex = 4
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIG",
	2: "TRANSACTIONS_FILTER",
	3: "ORDERER",
	4: "COMMIT_HASH",
}
var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIG":         1,
	"TRANSACTIONS_FILTER": 2,
	"ORDERER":             3,
	"COMMIT_HASH":         4,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}
func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{2}
}

// LastConfig is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfig struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastConfig) Reset()         { *m = LastConfig{} }
func (m *LastConfig) String() string { return proto.CompactTextString(m) }
func (*LastConfig) ProtoMessage()    {}
func (*LastConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{0}
}
func (m *LastConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastConfig.Unmarshal(m, b)
}
func (m *LastConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastConfig.Marshal(b, m, deterministic)
}
func (dst *LastConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastConfig.Merge(dst, src)
}
func (m *LastConfig) XXX_Size() int {
	return xxx_messageInfo_LastConfig.Size(m)
}
func (m *LastConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LastConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LastConfig proto.InternalMessageInfo

func (m *LastConfig) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value                []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures           []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader      []byte   `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataSignature) Reset()         { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()    {}
func (*MetadataSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{2}
}
func (m *MetadataSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataSignature.Unmarshal(m, b)
}
func (m *MetadataSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataSignature.Marshal(b, m, deterministic)
}
func (dst *MetadataSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataSignature.Merge(dst, src)
}
func (m *MetadataSignature) XXX_Size() int {
	return xxx_messageInfo_MetadataSignature.Size(m)
}
func (m *MetadataSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataSignature proto.InternalMessageInfo

func (m *MetadataSignature) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *MetadataSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Header struct {
	ChannelHeader        []byte   `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	SignatureHeader      []byte   `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{3}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetChannelHeader() []byte {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
	// If mutual TLS is employed, this represents
	// the hash of the client's TLS certificate
	TlsCertHash          []byte   `protobuf:"bytes,8,opt,name=tls_cert_hash,json=tlsCertHash,proto3" json:"tls_cert_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelHeader) Reset()         { *m = ChannelHeader{} }
func (m *ChannelHeader) String() string { return proto.CompactTextString(m) }
func (*ChannelHeader) ProtoMessage()    {}
func (*ChannelHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{4}
}
func (m *ChannelHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelHeader.Unmarshal(m, b)
}
func (m *ChannelHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelHeader.Marshal(b, m, deterministic)
}
func (dst *ChannelHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelHeader.Merge(dst, src)
}
func (m *ChannelHeader) XXX_Size() int {
	return xxx_messageInfo_ChannelHeader.Size(m)
}
func (m *ChannelHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelHeader proto.InternalMessageInfo

func (m *ChannelHeader) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ChannelHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ChannelHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChannelHeader) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelHeader) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChannelHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ChannelHeader) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ChannelHeader) GetTlsCertHash() []byte {
	if m != nil {
		return m.TlsCertHash
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, a marshaled msp.SerializedIdentity
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureHeader) Reset()         { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()    {}
func (*SignatureHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{5}
}
func (m *SignatureHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureHeader.Unmarshal(m, b)
}
func (m *SignatureHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureHeader.Marshal(b, m, deterministic)
}
func (dst *SignatureHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureHeader.Merge(dst, src)
}
func (m *SignatureHeader) XXX_Size() int {
	return xxx_messageInfo_SignatureHeader.Size(m)
}
func (m *SignatureHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureHeader proto.InternalMessageInfo

func (m *SignatureHeader) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *SignatureHeader) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{6}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{7}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *BlockData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Metadata             *BlockMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{8}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BlockHeader is the element of the block which forms the block chain
// The block header is hashed using the configured chain hashing algorithm
// over the ASN.1 encoding of the BlockHeader
type BlockHeader struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	PreviousHash         []byte   `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash             []byte   `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{9}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

type BlockData struct {
	Data                 [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{10}
}
func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockData.Unmarshal(m, b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
}
func (dst *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(dst, src)
}
func (m *BlockData) XXX_Size() int {
	return xxx_messageInfo_BlockData.Size(m)
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockMetadata struct {
	Metadata             [][]byte             `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	BlockTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{11}
}
func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMetadata.Unmarshal(m, b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
}
func (dst *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(dst, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return xxx_messageInfo_BlockMetadata.Size(m)
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetMetadata() [][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *BlockMetadata) GetBlockTime() *timestamp.Timestamp {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

// OrdererBlockMetadata defines metadata that is set by the ordering service.
type OrdererBlockMetadata struct {
	LastConfig           *LastConfig `protobuf:"bytes,1,opt,name=last_config,json=lastConfig,proto3" json:"last_config,omitempty"`
	ConsenterMetadata    []byte      `protobuf:"bytes,2,opt,name=consenter_metadata,json=consenterMetadata,proto3" json:"consenter_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrdererBlockMetadata) Reset()         { *m = OrdererBlockMetadata{} }
func (m *OrdererBlockMetadata) String() string { return proto.CompactTextString(m) }
func (*OrdererBlockMetadata) ProtoMessage()    {}
func (*OrdererBlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_50dcafdf8a2f7496, []int{12}
}
func (m *OrdererBlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdererBlockMetadata.Unmarshal(m, b)
}
func (m *OrdererBlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdererBlockMetadata.Marshal(b, m, deterministic)
}
func (dst *OrdererBlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdererBlockMetadata.Merge(dst, src)
}
func (m *OrdererBlockMetadata) XXX_Size() int {
	return xxx_messageInfo_OrdererBlockMetadata.Size(m)
}
func (m *OrdererBlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdererBlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OrdererBlockMetadata proto.InternalMessageInfo

func (m *OrdererBlockMetadata) GetLastConfig() *LastConfig {
	if m != nil {
		return m.LastConfig
	}
	return nil
}

func (m *OrdererBlockMetadata) GetConsenterMetadata() []byte {
	if m != nil {
		return m.ConsenterMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*LastConfig)(nil), "common.LastConfig")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChannelHeader)(nil), "common.ChannelHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterType((*OrdererBlockMetadata)(nil), "common.OrdererBlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_50dcafdf8a2f7496) }

var fileDescriptor_common_50dcafdf8a2f7496 = []byte{
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0xc7, 0x37, 0x71, 0x5e, 0x8f, 0x9b, 0xd6, 0x9d, 0x76, 0x9f, 0xf5, 0x53, 0x58, 0x6d, 0x65,
	0x58, 0x54, 0xba, 0x22, 0x15, 0xdd, 0x1b, 0xb8, 0x74, 0xec, 0x69, 0x6b, 0x35, 0xb1, 0xc3, 0xd8,
	0x59, 0xc4, 0x82, 0x64, 0x39, 0xc9, 0x34, 0x89, 0x70, 0xec, 0xc8, 0x9e, 0x54, 0x2d, 0xb7, 0xdc,
	0x23, 0x24, 0xb8, 0xe5, 0xbb, 0x70, 0x89, 0xf8, 0x3c, 0x20, 0x6e, 0xd1, 0x78, 0x6c, 0x37, 0x29,
	0x2b, 0xb8, 0xf3, 0x39, 0xe7, 0x37, 0xe7, 0xfc, 0xe7, 0x9c, 0x93, 0x09, 0xec, 0x4c, 0xe2, 0xe5,
	0x32, 0x8e, 0xba, 0xab, 0x24, 0x66, 0x31, 0x6a, 0x08, 0xeb, 0xe8, 0xc5, 0x2c, 0x8e, 0x67, 0x21,
	0x3d, 0xcb, 0xbc, 0xe3, 0xf5, 0xcd, 0x19, 0x5b, 0x2c, 0x69, 0xca, 0x82, 0xe5, 0x4a, 0x80, 0x9a,
	0x06, 0xd0, 0x0f, 0x52, 0x66, 0xc4, 0xd1, 0xcd, 0x62, 0x86, 0x0e, 0xa1, 0xbe, 0x88, 0xa6, 0xf4,
	0x4e, 0xad, 0x1c, 0x57, 0x4e, 0x6a, 0x44, 0x18, 0xda, 0xd7, 0xd0, 0x1a, 0x50, 0x16, 0x4c, 0x03,
	0x16, 0x70, 0xe2, 0x36, 0x08, 0xd7, 0x34, 0x23, 0x76, 0x88, 0x30, 0xd0, 0xe7, 0x00, 0xe9, 0x62,
	0x16, 0x05, 0x6c, 0x9d, 0xd0, 0x54, 0xad, 0x1e, 0x4b, 0x27, 0xf2, 0xf9, 0xff, 0xbb, 0xb9, 0xa2,
	0xe2, 0xac, 0x5b, 0x10, 0x64, 0x03, 0xd6, 0xbe, 0x81, 0xfd, 0x7f, 0x00, 0xe8, 0x63, 0x50, 0x4a,
	0xc4, 0x9f, 0xd3, 0x60, 0x4a, 0x93, 0xbc, 0xe0, 0x5e, 0xe9, 0xbf, 0xca, 0xdc, 0xe8, 0x7d, 0x68,
	0x97, 0x2e, 0xb5, 0x9a, 0x31, 0x0f, 0x0e, 0xed, 0x2d, 0x34, 0x72, 0xee, 0x25, 0xec, 0x4e, 0xe6,
	0x41, 0x14, 0xd1, 0x70, 0x3b, 0x61, 0x27, 0xf7, 0xe6, 0xd8, 0xbb, 0x2a, 0x57, 0xdf, 0x59, 0x59,
	0xfb, 0xbe, 0x0a, 0x1d, 0x63, 0xeb, 0x30, 0x82, 0x1a, 0xbb, 0x5f, 0x89, 0xde, 0xd4, 0x49, 0xf6,
	0x8d, 0x54, 0x68, 0xde, 0xd2, 0x24, 0x5d, 0xc4, 0x51, 0x96, 0xa7, 0x4e, 0x0a, 0x13, 0x7d, 0x06,
	0xed, 0x72, 0x1a, 0xaa, 0x74, 0x5c, 0x39, 0x91, 0xcf, 0x8f, 0xba, 0x62, 0x5e, 0xdd, 0x62, 0x5e,
	0x5d, 0xaf, 0x20, 0xc8, 0x03, 0x8c, 0x9e, 0x03, 0x14, 0x77, 0x59, 0x4c, 0xd5, 0xda, 0x71, 0xe5,
	0xa4, 0x4d, 0xda, 0xb9, 0xc7, 0x9a, 0xa2, 0x03, 0xa8, 0xb3, 0x3b, 0x1e, 0xa9, 0x67, 0x91, 0x1a,
	0xbb, 0xb3, 0xa6, 0x7c, 0x70, 0x74, 0x15, 0x4f, 0xe6, 0x6a, 0x43, 0x8c, 0x36, 0x33, 0x78, 0xf7,
	0xe8, 0x1d, 0xa3, 0x51, 0xa6, 0xaf, 0x29, 0xba, 0x57, 0x3a, 0x90, 0x06, 0x1d, 0x16, 0xa6, 0xfe,
	0x84, 0x26, 0xcc, 0x9f, 0x07, 0xe9, 0x5c, 0x6d, 0x65, 0x84, 0xcc, 0xc2, 0xd4, 0xa0, 0x09, 0xbb,
	0x0a, 0xd2, 0xb9, 0xa6, 0xc3, 0x9e, 0xfb, 0x68, 0x24, 0x2a, 0x34, 0x27, 0x09, 0x0d, 0x58, 0x5c,
	0xf4, 0xb8, 0x30, 0xb9, 0x88, 0x28, 0x8e, 0x26, 0xc5, 0xa0, 0x84, 0xa1, 0x61, 0x68, 0x0e, 0x83,
	0xfb, 0x30, 0x0e, 0xa6, 0xe8, 0x23, 0x68, 0x6c, 0x4c, 0x47, 0x3e, 0xdf, 0x2d, 0x96, 0x48, 0xa4,
	0x26, 0x79, 0x94, 0x77, 0x9a, 0x6f, 0x4c, 0x9e, 0x27, 0xfb, 0xd6, 0x7a, 0xd0, 0xc2, 0xd1, 0x2d,
	0x0d, 0x63, 0xd1, 0xf5, 0x95, 0x48, 0x59, 0x48, 0xc8, 0xcd, 0xff, 0xd8, 0x97, 0x1f, 0x2a, 0x50,
	0xef, 0x85, 0xf1, 0xe4, 0x5b, 0xf4, 0xea, 0x91, 0x92, 0x83, 0x42, 0x49, 0x16, 0x7e, 0x24, 0xe7,
	0xe5, 0x86, 0x1c, 0xf9, 0x7c, 0x7f, 0x0b, 0x35, 0x03, 0x16, 0x08, 0x85, 0xe8, 0x53, 0x68, 0x2d,
	0xf3, 0x5d, 0xcf, 0x07, 0xfe, 0x74, 0x0b, 0x2d, 0x7e, 0x08, 0xa4, 0xc4, 0xb4, 0x19, 0xc8, 0x1b,
	0x05, 0xd1, 0xff, 0xa0, 0x11, 0xad, 0x97, 0xe3, 0x5c, 0x55, 0x8d, 0xe4, 0x16, 0xfa, 0x00, 0x3a,
	0xab, 0x84, 0xde, 0x2e, 0xe2, 0x75, 0x2a, 0x26, 0x25, 0x6e, 0xb6, 0x53, 0x38, 0xf9, 0xa8, 0xd0,
	0x7b, 0xd0, 0xe6, 0x39, 0x05, 0x20, 0x65, 0x40, 0x8b, 0x3b, 0xb2, 0x39, 0xbe, 0x80, 0x76, 0x29,
	0xb7, 0x6c, 0x6f, 0xe5, 0x58, 0x2a, 0xdb, 0x4b, 0xa1, 0xb3, 0x25, 0x12, 0x1d, 0x6d, 0xdc, 0x46,
	0x80, 0xa5, 0xcd, 0x77, 0x7b, 0xcc, 0x61, 0xbe, 0xbe, 0x79, 0x57, 0xfe, 0x75, 0xb7, 0x4b, 0x58,
	0xfb, 0x0e, 0x0e, 0x9d, 0x64, 0x4a, 0x13, 0x9a, 0x6c, 0x57, 0x7b, 0x0d, 0x72, 0x18, 0xa4, 0xcc,
	0x9f, 0x64, 0x2f, 0x55, 0x3e, 0x14, 0x54, 0xb4, 0xef, 0xe1, 0x0d, 0x23, 0x10, 0x3e, 0xbc, 0x67,
	0x9f, 0x00, 0x9a, 0xc4, 0x51, 0x4a, 0x23, 0x46, 0x13, 0xbf, 0x14, 0x2b, 0x7a, 0xb3, 0x5f, 0x46,
	0x8a, 0x1a, 0xa7, 0xbf, 0x56, 0xa0, 0xe1, 0xb2, 0x80, 0xad, 0x53, 0x24, 0x43, 0x73, 0x64, 0x5f,
	0xdb, 0xce, 0x97, 0xb6, 0xf2, 0x04, 0xed, 0x40, 0xd3, 0x1d, 0x19, 0x06, 0x76, 0x5d, 0xe5, 0xb7,
	0x0a, 0x52, 0x40, 0xee, 0xe9, 0xa6, 0x4f, 0xf0, 0x17, 0x23, 0xec, 0x7a, 0xca, 0x8f, 0x12, 0xda,
	0x85, 0xf6, 0x85, 0x43, 0x7a, 0x96, 0x69, 0x62, 0x5b, 0xf9, 0x29, 0xb3, 0x6d, 0xc7, 0xf3, 0x2f,
	0x9c, 0x91, 0x6d, 0x2a, 0x3f, 0x4b, 0xe8, 0x39, 0xa8, 0x39, 0xed, 0x63, 0xdb, 0xb3, 0xbc, 0xaf,
	0x7c, 0xcf, 0x71, 0xfc, 0xbe, 0x4e, 0x2e, 0xb1, 0xf2, 0x8b, 0x84, 0x8e, 0xe0, 0xa9, 0x65, 0x7b,
	0x98, 0xd8, 0x7a, 0xdf, 0x77, 0x31, 0x79, 0x83, 0x89, 0x8f, 0x09, 0x71, 0x88, 0xf2, 0x87, 0x84,
	0x0e, 0x61, 0x8f, 0xa7, 0xb2, 0x06, 0xc3, 0x3e, 0x1e, 0x60, 0xdb, 0xc3, 0xa6, 0xf2, 0xa7, 0x84,
	0x54, 0x38, 0xe0, 0xa0, 0x65, 0x60, 0x7f, 0x64, 0xeb, 0x6f, 0x74, 0xab, 0xaf, 0xf7, 0xfa, 0x58,
	0xf9, 0x4b, 0x3a, 0xfd, 0xbd, 0x02, 0x20, 0x76, 0xc5, 0xe3, 0xaf, 0x8f, 0x0c, 0xcd, 0x01, 0x76,
	0x5d, 0xfd, 0x12, 0x2b, 0x4f, 0x10, 0x40, 0xc3, 0x70, 0xec, 0x0b, 0xeb, 0x52, 0xa9, 0xa0, 0x7d,
	0xe8, 0x88, 0x6f, 0x7f, 0x34, 0x34, 0x75, 0x0f, 0x2b, 0x55, 0xa4, 0xc2, 0x21, 0xb6, 0x4d, 0x87,
	0xb8, 0x98, 0xf8, 0x1e, 0xd1, 0x6d, 0x57, 0x37, 0x3c, 0xcb, 0xb1, 0x15, 0x09, 0x3d, 0x83, 0x03,
	0x87, 0x98, 0x98, 0x3c, 0x0a, 0xd4, 0xd0, 0x53, 0xd8, 0x37, 0x71, 0xdf, 0xe2, 0x8a, 0x5d, 0x8c,
	0xaf, 0x7d, 0xcb, 0xbe, 0x70, 0x94, 0x3a, 0x77, 0x1b, 0x57, 0xba, 0x65, 0x1b, 0x8e, 0x89, 0xfd,
	0xa1, 0x6e, 0x5c, 0xf3, 0xfa, 0x0d, 0x5e, 0x60, 0x88, 0x31, 0xf1, 0x75, 0x73, 0x60, 0xd9, 0xbe,
	0x33, 0xc4, 0x44, 0xcf, 0xf2, 0xb4, 0xf8, 0x01, 0xcf, 0xb9, 0xc6, 0xf6, 0x56, 0xfa, 0xf6, 0x69,
	0x08, 0x68, 0x6b, 0x09, 0x2c, 0xfe, 0x77, 0x84, 0x76, 0x01, 0x5c, 0xeb, 0xd2, 0xd6, 0xbd, 0x11,
	0xc1, 0xae, 0xf2, 0x04, 0xed, 0x81, 0xdc, 0xd7, 0x5d, 0xcf, 0x2f, 0xef, 0xf6, 0x0c, 0x0e, 0x36,
	0xf2, 0xb8, 0xfe, 0x85, 0xd5, 0xf7, 0x30, 0x51, 0xaa, 0xbc, 0x1b, 0xf9, 0x3d, 0x14, 0x89, 0x1f,
	0x33, 0x9c, 0xc1, 0xc0, 0xf2, 0xfc, 0x2b, 0xdd, 0xbd, 0x52, 0x6a, 0x3d, 0x17, 0x3e, 0x8c, 0x93,
	0x59, 0x77, 0x7e, 0xbf, 0xa2, 0x49, 0x48, 0xa7, 0x33, 0x9a, 0x74, 0x6f, 0x82, 0x71, 0xb2, 0x98,
	0x88, 0x8d, 0x4d, 0xf3, 0x5d, 0x7b, 0xfb, 0x6a, 0xb6, 0x60, 0xf3, 0xf5, 0x98, 0x9b, 0x67, 0x1b,
	0xf0, 0x99, 0x80, 0xc5, 0x5f, 0x6d, 0x7a, 0x26, 0xe0, 0x71, 0x23, 0x33, 0x5f, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xc7, 0xe4, 0x8e, 0x9f, 0x07, 0x00, 0x00,
}
