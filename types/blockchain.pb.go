// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package types // import "github.com/aergoio/aergo/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxType int32

const (
	TxType_NORMAL     TxType = 0
	TxType_GOVERNANCE TxType = 1
)

var TxType_name = map[int32]string{
	0: "NORMAL",
	1: "GOVERNANCE",
}
var TxType_value = map[string]int32{
	"NORMAL":     0,
	"GOVERNANCE": 1,
}

func (x TxType) String() string {
	return proto.EnumName(TxType_name, int32(x))
}
func (TxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{0}
}

type Block struct {
	Hash                 []byte       `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header               *BlockHeader `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{0}
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

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlockHeader struct {
	PrevBlockHash        []byte   `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	BlockNo              uint64   `protobuf:"varint,2,opt,name=blockNo,proto3" json:"blockNo,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlocksRootHash       []byte   `protobuf:"bytes,4,opt,name=blocksRootHash,proto3" json:"blocksRootHash,omitempty"`
	TxsRootHash          []byte   `protobuf:"bytes,5,opt,name=txsRootHash,proto3" json:"txsRootHash,omitempty"`
	ReceiptsRootHash     []byte   `protobuf:"bytes,6,opt,name=receiptsRootHash,proto3" json:"receiptsRootHash,omitempty"`
	Confirms             uint64   `protobuf:"varint,7,opt,name=confirms,proto3" json:"confirms,omitempty"`
	PubKey               []byte   `protobuf:"bytes,8,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
	CoinbaseAccount      []byte   `protobuf:"bytes,10,opt,name=coinbaseAccount,proto3" json:"coinbaseAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{1}
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

func (m *BlockHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeader) GetBlockNo() uint64 {
	if m != nil {
		return m.BlockNo
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetBlocksRootHash() []byte {
	if m != nil {
		return m.BlocksRootHash
	}
	return nil
}

func (m *BlockHeader) GetTxsRootHash() []byte {
	if m != nil {
		return m.TxsRootHash
	}
	return nil
}

func (m *BlockHeader) GetReceiptsRootHash() []byte {
	if m != nil {
		return m.ReceiptsRootHash
	}
	return nil
}

func (m *BlockHeader) GetConfirms() uint64 {
	if m != nil {
		return m.Confirms
	}
	return 0
}

func (m *BlockHeader) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetCoinbaseAccount() []byte {
	if m != nil {
		return m.CoinbaseAccount
	}
	return nil
}

type BlockBody struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{2}
}
func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (dst *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(dst, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxList struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxList) Reset()         { *m = TxList{} }
func (m *TxList) String() string { return proto.CompactTextString(m) }
func (*TxList) ProtoMessage()    {}
func (*TxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{3}
}
func (m *TxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxList.Unmarshal(m, b)
}
func (m *TxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxList.Marshal(b, m, deterministic)
}
func (dst *TxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxList.Merge(dst, src)
}
func (m *TxList) XXX_Size() int {
	return xxx_messageInfo_TxList.Size(m)
}
func (m *TxList) XXX_DiscardUnknown() {
	xxx_messageInfo_TxList.DiscardUnknown(m)
}

var xxx_messageInfo_TxList proto.InternalMessageInfo

func (m *TxList) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Tx struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Body                 *TxBody  `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{4}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Tx) GetBody() *TxBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type TxBody struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Account              []byte   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Recipient            []byte   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Limit                uint64   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Price                uint64   `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Type                 TxType   `protobuf:"varint,8,opt,name=type,proto3,enum=types.TxType" json:"type,omitempty"`
	Sign                 []byte   `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxBody) Reset()         { *m = TxBody{} }
func (m *TxBody) String() string { return proto.CompactTextString(m) }
func (*TxBody) ProtoMessage()    {}
func (*TxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{5}
}
func (m *TxBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxBody.Unmarshal(m, b)
}
func (m *TxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxBody.Marshal(b, m, deterministic)
}
func (dst *TxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxBody.Merge(dst, src)
}
func (m *TxBody) XXX_Size() int {
	return xxx_messageInfo_TxBody.Size(m)
}
func (m *TxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TxBody proto.InternalMessageInfo

func (m *TxBody) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxBody) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *TxBody) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxBody) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxBody) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxBody) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TxBody) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxBody) GetType() TxType {
	if m != nil {
		return m.Type
	}
	return TxType_NORMAL
}

func (m *TxBody) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type TxIdx struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Idx                  int32    `protobuf:"varint,2,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxIdx) Reset()         { *m = TxIdx{} }
func (m *TxIdx) String() string { return proto.CompactTextString(m) }
func (*TxIdx) ProtoMessage()    {}
func (*TxIdx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{6}
}
func (m *TxIdx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxIdx.Unmarshal(m, b)
}
func (m *TxIdx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxIdx.Marshal(b, m, deterministic)
}
func (dst *TxIdx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxIdx.Merge(dst, src)
}
func (m *TxIdx) XXX_Size() int {
	return xxx_messageInfo_TxIdx.Size(m)
}
func (m *TxIdx) XXX_DiscardUnknown() {
	xxx_messageInfo_TxIdx.DiscardUnknown(m)
}

var xxx_messageInfo_TxIdx proto.InternalMessageInfo

func (m *TxIdx) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TxIdx) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type TxInBlock struct {
	TxIdx                *TxIdx   `protobuf:"bytes,1,opt,name=txIdx,proto3" json:"txIdx,omitempty"`
	Tx                   *Tx      `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxInBlock) Reset()         { *m = TxInBlock{} }
func (m *TxInBlock) String() string { return proto.CompactTextString(m) }
func (*TxInBlock) ProtoMessage()    {}
func (*TxInBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{7}
}
func (m *TxInBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInBlock.Unmarshal(m, b)
}
func (m *TxInBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInBlock.Marshal(b, m, deterministic)
}
func (dst *TxInBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInBlock.Merge(dst, src)
}
func (m *TxInBlock) XXX_Size() int {
	return xxx_messageInfo_TxInBlock.Size(m)
}
func (m *TxInBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInBlock.DiscardUnknown(m)
}

var xxx_messageInfo_TxInBlock proto.InternalMessageInfo

func (m *TxInBlock) GetTxIdx() *TxIdx {
	if m != nil {
		return m.TxIdx
	}
	return nil
}

func (m *TxInBlock) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

type State struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              uint64   `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	CodeHash             []byte   `protobuf:"bytes,3,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	StorageRoot          []byte   `protobuf:"bytes,4,opt,name=storageRoot,proto3" json:"storageRoot,omitempty"`
	SqlRecoveryPoint     uint64   `protobuf:"varint,5,opt,name=sqlRecoveryPoint,proto3" json:"sqlRecoveryPoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{8}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *State) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *State) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *State) GetStorageRoot() []byte {
	if m != nil {
		return m.StorageRoot
	}
	return nil
}

func (m *State) GetSqlRecoveryPoint() uint64 {
	if m != nil {
		return m.SqlRecoveryPoint
	}
	return 0
}

type StateProof struct {
	State                *State   `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
	Inclusion            bool     `protobuf:"varint,2,opt,name=inclusion,proto3" json:"inclusion,omitempty"`
	ProofKey             []byte   `protobuf:"bytes,3,opt,name=proofKey,proto3" json:"proofKey,omitempty"`
	ProofVal             []byte   `protobuf:"bytes,4,opt,name=proofVal,proto3" json:"proofVal,omitempty"`
	AuditPath            [][]byte `protobuf:"bytes,5,rep,name=auditPath,proto3" json:"auditPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateProof) Reset()         { *m = StateProof{} }
func (m *StateProof) String() string { return proto.CompactTextString(m) }
func (*StateProof) ProtoMessage()    {}
func (*StateProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{9}
}
func (m *StateProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateProof.Unmarshal(m, b)
}
func (m *StateProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateProof.Marshal(b, m, deterministic)
}
func (dst *StateProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateProof.Merge(dst, src)
}
func (m *StateProof) XXX_Size() int {
	return xxx_messageInfo_StateProof.Size(m)
}
func (m *StateProof) XXX_DiscardUnknown() {
	xxx_messageInfo_StateProof.DiscardUnknown(m)
}

var xxx_messageInfo_StateProof proto.InternalMessageInfo

func (m *StateProof) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *StateProof) GetInclusion() bool {
	if m != nil {
		return m.Inclusion
	}
	return false
}

func (m *StateProof) GetProofKey() []byte {
	if m != nil {
		return m.ProofKey
	}
	return nil
}

func (m *StateProof) GetProofVal() []byte {
	if m != nil {
		return m.ProofVal
	}
	return nil
}

func (m *StateProof) GetAuditPath() [][]byte {
	if m != nil {
		return m.AuditPath
	}
	return nil
}

type Receipt struct {
	ContractAddress      []byte   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Ret                  string   `protobuf:"bytes,3,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{10}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (dst *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(dst, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *Receipt) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Receipt) GetRet() string {
	if m != nil {
		return m.Ret
	}
	return ""
}

type Vote struct {
	Candidate            []byte   `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{11}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (dst *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(dst, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetCandidate() []byte {
	if m != nil {
		return m.Candidate
	}
	return nil
}

func (m *Vote) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteList struct {
	Votes                []*Vote  `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteList) Reset()         { *m = VoteList{} }
func (m *VoteList) String() string { return proto.CompactTextString(m) }
func (*VoteList) ProtoMessage()    {}
func (*VoteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{12}
}
func (m *VoteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteList.Unmarshal(m, b)
}
func (m *VoteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteList.Marshal(b, m, deterministic)
}
func (dst *VoteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteList.Merge(dst, src)
}
func (m *VoteList) XXX_Size() int {
	return xxx_messageInfo_VoteList.Size(m)
}
func (m *VoteList) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteList.DiscardUnknown(m)
}

var xxx_messageInfo_VoteList proto.InternalMessageInfo

func (m *VoteList) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type FnArgument struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FnArgument) Reset()         { *m = FnArgument{} }
func (m *FnArgument) String() string { return proto.CompactTextString(m) }
func (*FnArgument) ProtoMessage()    {}
func (*FnArgument) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{13}
}
func (m *FnArgument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FnArgument.Unmarshal(m, b)
}
func (m *FnArgument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FnArgument.Marshal(b, m, deterministic)
}
func (dst *FnArgument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FnArgument.Merge(dst, src)
}
func (m *FnArgument) XXX_Size() int {
	return xxx_messageInfo_FnArgument.Size(m)
}
func (m *FnArgument) XXX_DiscardUnknown() {
	xxx_messageInfo_FnArgument.DiscardUnknown(m)
}

var xxx_messageInfo_FnArgument proto.InternalMessageInfo

func (m *FnArgument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Function struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arguments            []*FnArgument `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{14}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (dst *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(dst, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetArguments() []*FnArgument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type ABI struct {
	Version              string      `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Language             string      `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Functions            []*Function `protobuf:"bytes,3,rep,name=functions,proto3" json:"functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ABI) Reset()         { *m = ABI{} }
func (m *ABI) String() string { return proto.CompactTextString(m) }
func (*ABI) ProtoMessage()    {}
func (*ABI) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{15}
}
func (m *ABI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABI.Unmarshal(m, b)
}
func (m *ABI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABI.Marshal(b, m, deterministic)
}
func (dst *ABI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABI.Merge(dst, src)
}
func (m *ABI) XXX_Size() int {
	return xxx_messageInfo_ABI.Size(m)
}
func (m *ABI) XXX_DiscardUnknown() {
	xxx_messageInfo_ABI.DiscardUnknown(m)
}

var xxx_messageInfo_ABI proto.InternalMessageInfo

func (m *ABI) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ABI) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *ABI) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

type Query struct {
	ContractAddress      []byte   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Queryinfo            []byte   `protobuf:"bytes,2,opt,name=queryinfo,proto3" json:"queryinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_32d62bb02894f90f, []int{16}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *Query) GetQueryinfo() []byte {
	if m != nil {
		return m.Queryinfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*BlockHeader)(nil), "types.BlockHeader")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*TxList)(nil), "types.TxList")
	proto.RegisterType((*Tx)(nil), "types.Tx")
	proto.RegisterType((*TxBody)(nil), "types.TxBody")
	proto.RegisterType((*TxIdx)(nil), "types.TxIdx")
	proto.RegisterType((*TxInBlock)(nil), "types.TxInBlock")
	proto.RegisterType((*State)(nil), "types.State")
	proto.RegisterType((*StateProof)(nil), "types.StateProof")
	proto.RegisterType((*Receipt)(nil), "types.Receipt")
	proto.RegisterType((*Vote)(nil), "types.Vote")
	proto.RegisterType((*VoteList)(nil), "types.VoteList")
	proto.RegisterType((*FnArgument)(nil), "types.FnArgument")
	proto.RegisterType((*Function)(nil), "types.Function")
	proto.RegisterType((*ABI)(nil), "types.ABI")
	proto.RegisterType((*Query)(nil), "types.Query")
	proto.RegisterEnum("types.TxType", TxType_name, TxType_value)
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_32d62bb02894f90f) }

var fileDescriptor_blockchain_32d62bb02894f90f = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x27, 0x89, 0x9d, 0xc6, 0x93, 0x5e, 0x2f, 0xac, 0x10, 0x32, 0x50, 0xa1, 0x9c, 0x55, 0x50,
	0x54, 0xe9, 0x5a, 0xa9, 0x3c, 0xf0, 0x00, 0x2f, 0x29, 0xba, 0x83, 0xc2, 0xd1, 0x96, 0x25, 0xea,
	0x03, 0x12, 0x0f, 0x1b, 0x7b, 0x9b, 0x2c, 0x24, 0xbb, 0x3e, 0x7b, 0x5d, 0x39, 0x1f, 0x07, 0x89,
	0x6f, 0xc4, 0xc7, 0xe0, 0x4b, 0xa0, 0x99, 0xdd, 0xc4, 0x26, 0x2d, 0x48, 0xf7, 0x94, 0x9d, 0xdf,
	0xfc, 0x9f, 0xdf, 0x8c, 0x03, 0xa3, 0xf9, 0xca, 0xa4, 0xbf, 0xa7, 0x4b, 0xa1, 0xf4, 0x59, 0x5e,
	0x18, 0x6b, 0x58, 0x68, 0x37, 0xb9, 0x2c, 0x93, 0x35, 0x84, 0x97, 0xa8, 0x62, 0x0c, 0x82, 0xa5,
	0x28, 0x97, 0x71, 0x67, 0xdc, 0x99, 0x1c, 0x72, 0x7a, 0xb3, 0x53, 0xe8, 0x2f, 0xa5, 0xc8, 0x64,
	0x11, 0x77, 0xc7, 0x9d, 0xc9, 0xf0, 0x82, 0x9d, 0x91, 0xd3, 0x19, 0x79, 0x7c, 0x47, 0x1a, 0xee,
	0x2d, 0xd8, 0x09, 0x04, 0x73, 0x93, 0x6d, 0xe2, 0x1e, 0x59, 0x8e, 0xda, 0x96, 0x97, 0x26, 0xdb,
	0x70, 0xd2, 0x26, 0x7f, 0x75, 0x61, 0xd8, 0xf2, 0x66, 0x27, 0xf0, 0x2c, 0x2f, 0xe4, 0x83, 0x83,
	0x9a, 0xf4, 0xff, 0x06, 0x59, 0x0c, 0x07, 0x54, 0xff, 0xb5, 0xa1, 0x42, 0x02, 0xbe, 0x15, 0xd9,
	0x31, 0x44, 0x56, 0xad, 0x65, 0x69, 0xc5, 0x3a, 0xa7, 0xd4, 0x3d, 0xde, 0x00, 0xec, 0x73, 0x38,
	0x22, 0xc3, 0x92, 0x1b, 0x63, 0x29, 0x7c, 0x40, 0xe1, 0xf7, 0x50, 0x36, 0x86, 0xa1, 0xad, 0x1b,
	0xa3, 0x90, 0x8c, 0xda, 0x10, 0x3b, 0x85, 0x51, 0x21, 0x53, 0xa9, 0x72, 0xdb, 0x98, 0xf5, 0xc9,
	0xec, 0x11, 0xce, 0x3e, 0x86, 0x41, 0x6a, 0xf4, 0xbd, 0x2a, 0xd6, 0x65, 0x7c, 0x40, 0xe5, 0xee,
	0x64, 0xf6, 0x21, 0xf4, 0xf3, 0x6a, 0xfe, 0x83, 0xdc, 0xc4, 0x03, 0xf2, 0xf6, 0x12, 0x4e, 0xbf,
	0x54, 0x0b, 0x1d, 0x47, 0x6e, 0xfa, 0xf8, 0x66, 0x13, 0x78, 0x9e, 0x1a, 0xa5, 0xe7, 0xa2, 0x94,
	0xd3, 0x34, 0x35, 0x95, 0xb6, 0x31, 0x90, 0x7a, 0x1f, 0x4e, 0x26, 0x10, 0xed, 0x06, 0xcd, 0x3e,
	0x81, 0x9e, 0xad, 0xcb, 0xb8, 0x33, 0xee, 0x4d, 0x86, 0x17, 0x91, 0xe7, 0x61, 0x56, 0x73, 0x44,
	0x93, 0xcf, 0xa0, 0x3f, 0xab, 0xdf, 0xa8, 0xd2, 0xfe, 0xbf, 0xd9, 0x57, 0xd0, 0x9d, 0xd5, 0x4f,
	0xae, 0xc4, 0x0b, 0x4f, 0xb3, 0x5b, 0x88, 0x67, 0x3b, 0xbf, 0x16, 0xc7, 0x7f, 0x77, 0x30, 0x09,
	0xd5, 0xf2, 0x01, 0x84, 0xda, 0xe8, 0x54, 0x52, 0x88, 0x80, 0x3b, 0x01, 0xe9, 0x14, 0xbe, 0xa1,
	0x2e, 0x85, 0xde, 0x8a, 0x48, 0x67, 0x21, 0x53, 0x95, 0x2b, 0xa9, 0x2d, 0xd1, 0x79, 0xc8, 0x1b,
	0x00, 0x87, 0x27, 0xd6, 0xe4, 0x16, 0x50, 0x38, 0x2f, 0x61, 0xbc, 0x5c, 0x6c, 0x56, 0x46, 0x64,
	0x9e, 0xba, 0xad, 0x88, 0xf9, 0x57, 0x6a, 0xad, 0x2c, 0x71, 0x15, 0x70, 0x27, 0x20, 0x9a, 0x17,
	0x2a, 0x95, 0x9e, 0x1d, 0x27, 0x60, 0x67, 0xd8, 0x0c, 0x11, 0x73, 0xd4, 0xea, 0x6c, 0xb6, 0xc9,
	0x25, 0x27, 0xd5, 0x53, 0x2c, 0x25, 0x5f, 0x42, 0x38, 0xab, 0xaf, 0xb2, 0x1a, 0x6b, 0x9f, 0xef,
	0xad, 0x71, 0x03, 0xb0, 0x11, 0xf4, 0x54, 0x56, 0x53, 0xbf, 0x21, 0xc7, 0x67, 0xf2, 0x3d, 0x44,
	0xb3, 0xfa, 0x4a, 0xbb, 0xeb, 0x4b, 0x20, 0xb4, 0x18, 0x85, 0x1c, 0x87, 0x17, 0x87, 0xbb, 0xec,
	0x57, 0x59, 0xcd, 0x9d, 0x8a, 0x7d, 0x04, 0x5d, 0x5b, 0xfb, 0xc1, 0xb7, 0x08, 0xeb, 0xda, 0x3a,
	0xf9, 0xa3, 0x03, 0xe1, 0xcf, 0x56, 0x58, 0xf9, 0xdf, 0x13, 0x9f, 0x8b, 0x95, 0x40, 0x7c, 0x7b,
	0x40, 0x4e, 0x74, 0xcb, 0x9a, 0x49, 0x2a, 0xda, 0x0d, 0x7c, 0x27, 0xe3, 0x59, 0x94, 0xd6, 0x14,
	0x62, 0x21, 0x71, 0xb7, 0xfd, 0xed, 0xb4, 0x21, 0x3c, 0x8b, 0xf2, 0xed, 0x8a, 0xcb, 0xd4, 0x3c,
	0xc8, 0x62, 0x73, 0x6b, 0x94, 0xb6, 0x44, 0x41, 0xc0, 0x1f, 0xe1, 0xc9, 0x9f, 0x1d, 0x00, 0xaa,
	0xf1, 0xb6, 0x30, 0xe6, 0x1e, 0x3b, 0x26, 0x69, 0xaf, 0x63, 0xc2, 0xb8, 0x6f, 0xe6, 0x18, 0x22,
	0xa5, 0xd3, 0x55, 0x55, 0x2a, 0xa3, 0xa9, 0xf0, 0x01, 0x6f, 0x00, 0x2c, 0x3d, 0xc7, 0x50, 0x78,
	0x4d, 0xbe, 0xf4, 0xad, 0xbc, 0xd3, 0xdd, 0x89, 0x95, 0xaf, 0x7b, 0x27, 0x63, 0x54, 0x51, 0x65,
	0xca, 0xde, 0x0a, 0x8b, 0xb7, 0xde, 0x43, 0xa2, 0x76, 0x40, 0xf2, 0x2b, 0x1c, 0x70, 0x77, 0xd1,
	0xee, 0x00, 0xb5, 0x2d, 0x44, 0x6a, 0xa7, 0x59, 0x56, 0xc8, 0xb2, 0xf4, 0xbc, 0xee, 0xc3, 0xb8,
	0x99, 0xa5, 0x15, 0xb6, 0x2a, 0xa9, 0xca, 0x88, 0x7b, 0x09, 0x59, 0x2f, 0xa4, 0xdb, 0xe4, 0x88,
	0xe3, 0x33, 0xf9, 0x1a, 0x82, 0x3b, 0xe3, 0x5a, 0x4b, 0x85, 0xce, 0x54, 0xb6, 0x1d, 0xc1, 0x21,
	0x6f, 0x80, 0xd6, 0xa6, 0x77, 0xdb, 0x9b, 0x9e, 0xbc, 0x84, 0x01, 0x7a, 0xd3, 0x01, 0xbf, 0x80,
	0xf0, 0xc1, 0x58, 0xb9, 0x3d, 0xe1, 0xa1, 0x1f, 0x20, 0xea, 0xb9, 0xd3, 0x24, 0x63, 0x80, 0xd7,
	0x7a, 0x5a, 0x2c, 0xaa, 0x35, 0x9e, 0x0f, 0x83, 0x40, 0x8b, 0xb5, 0xcb, 0x16, 0x71, 0x7a, 0x27,
	0x37, 0x30, 0x78, 0x5d, 0xe9, 0xd4, 0xe2, 0x3c, 0x9f, 0xd0, 0xb3, 0x73, 0x88, 0x84, 0xf7, 0xc7,
	0xde, 0x30, 0xd1, 0xfb, 0x3e, 0x51, 0x13, 0x99, 0x37, 0x36, 0xc9, 0x6f, 0xd0, 0x9b, 0x5e, 0x5e,
	0xe1, 0xc2, 0x3d, 0xc8, 0x82, 0x78, 0x73, 0xe1, 0xb6, 0x22, 0x32, 0xb3, 0x12, 0x7a, 0x51, 0x89,
	0x85, 0xf4, 0xc3, 0xda, 0xc9, 0xec, 0x25, 0x44, 0xf7, 0xbe, 0x9a, 0x32, 0xee, 0x51, 0xb6, 0xe7,
	0xdb, 0x6c, 0x1e, 0xe7, 0x8d, 0x45, 0x72, 0x03, 0xe1, 0x4f, 0x95, 0x2c, 0x36, 0xef, 0x40, 0xd4,
	0x31, 0x44, 0x6f, 0xd1, 0x45, 0xe9, 0x7b, 0xe3, 0x3f, 0x3e, 0x0d, 0x70, 0x7a, 0x82, 0x1f, 0x2e,
	0xbc, 0x77, 0x06, 0xd0, 0xbf, 0xbe, 0xe1, 0x3f, 0x4e, 0xdf, 0x8c, 0xde, 0x63, 0x47, 0x00, 0xdf,
	0xde, 0xdc, 0xbd, 0xe2, 0xd7, 0xd3, 0xeb, 0x6f, 0x5e, 0x8d, 0x3a, 0x97, 0xe3, 0x5f, 0x3e, 0x5d,
	0x28, 0xbb, 0xac, 0xe6, 0x67, 0xa9, 0x59, 0x9f, 0x0b, 0x59, 0x2c, 0x8c, 0x32, 0xee, 0xf7, 0x9c,
	0x8a, 0x9d, 0xf7, 0xe9, 0x2f, 0xf6, 0x8b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0xa0, 0x55,
	0x84, 0x76, 0x07, 0x00, 0x00,
}
