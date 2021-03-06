// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/transfer_gateway/transfer_gateway.proto

/*
Package transfer_gateway is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/transfer_gateway/transfer_gateway.proto

It has these top-level messages:
	TransferGatewayState
	TransferGatewayOracleState
	TransferGatewayWithdrawalReceipt
	TransferGatewayPendingWithdrawalSummary
	TransferGatewayAccount
	TransferGatewayTokenDeposited
	TransferGatewayTokenWithdrawn
	TransferGatewayMainnetEvent
	TransferGatewayTokenWithdrawalSigned
	TransferGatewayInitRequest
	TransferGatewayAddOracleRequest
	TransferGatewayRemoveOracleRequest
	TransferGatewayGetOraclesRequest
	TransferGatewayGetOraclesResponse
	TransferGatewayProcessEventBatchRequest
	TransferGatewayStateRequest
	TransferGatewayStateResponse
	TransferGatewayWithdrawERC721Request
	TransferGatewayWithdrawalReceiptRequest
	TransferGatewayWithdrawalReceiptResponse
	TransferGatewayConfirmWithdrawalReceiptRequest
	TransferGatewayPendingWithdrawalsRequest
	TransferGatewayPendingWithdrawalsResponse
*/
package transfer_gateway

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TransferGatewayTokenKind int32

const (
	TransferGatewayTokenKind_ETH    TransferGatewayTokenKind = 0
	TransferGatewayTokenKind_ERC20  TransferGatewayTokenKind = 1
	TransferGatewayTokenKind_ERC721 TransferGatewayTokenKind = 2
)

var TransferGatewayTokenKind_name = map[int32]string{
	0: "ETH",
	1: "ERC20",
	2: "ERC721",
}
var TransferGatewayTokenKind_value = map[string]int32{
	"ETH":    0,
	"ERC20":  1,
	"ERC721": 2,
}

func (x TransferGatewayTokenKind) String() string {
	return proto.EnumName(TransferGatewayTokenKind_name, int32(x))
}
func (TransferGatewayTokenKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{0}
}

type TransferGatewayState struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	// Last Mainnet block processed by the Transfer Gateway contract
	LastMainnetBlockNum uint64 `protobuf:"varint,2,opt,name=last_mainnet_block_num,json=lastMainnetBlockNum,proto3" json:"last_mainnet_block_num,omitempty"`
	// Token owners that have initiated (but have not as yet completed) a withdrawal to Mainnet.
	TokenWithdrawers []*types.Address `protobuf:"bytes,3,rep,name=token_withdrawers,json=tokenWithdrawers" json:"token_withdrawers,omitempty"`
}

func (m *TransferGatewayState) Reset()         { *m = TransferGatewayState{} }
func (m *TransferGatewayState) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayState) ProtoMessage()    {}
func (*TransferGatewayState) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{0}
}

func (m *TransferGatewayState) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TransferGatewayState) GetLastMainnetBlockNum() uint64 {
	if m != nil {
		return m.LastMainnetBlockNum
	}
	return 0
}

func (m *TransferGatewayState) GetTokenWithdrawers() []*types.Address {
	if m != nil {
		return m.TokenWithdrawers
	}
	return nil
}

type TransferGatewayOracleState struct {
	// DAppChain address of the Oracle
	Address *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *TransferGatewayOracleState) Reset()         { *m = TransferGatewayOracleState{} }
func (m *TransferGatewayOracleState) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayOracleState) ProtoMessage()    {}
func (*TransferGatewayOracleState) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{1}
}

func (m *TransferGatewayOracleState) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type TransferGatewayWithdrawalReceipt struct {
	// Mainnet address of token owner
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// Mainnet address of token contract
	TokenContract *types.Address           `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract" json:"token_contract,omitempty"`
	TokenKind     TransferGatewayTokenKind `protobuf:"varint,3,opt,name=token_kind,json=tokenKind,proto3,enum=TransferGatewayTokenKind" json:"token_kind,omitempty"`
	// // ERC721 token ID, or amount of ERC20/ETH
	Value           *types.BigUInt `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	WithdrawalNonce uint64         `protobuf:"varint,5,opt,name=withdrawal_nonce,json=withdrawalNonce,proto3" json:"withdrawal_nonce,omitempty"`
	// Signature generated by the Oracle that confirmed the withdrawal
	OracleSignature []byte `protobuf:"bytes,6,opt,name=oracle_signature,json=oracleSignature,proto3" json:"oracle_signature,omitempty"`
}

func (m *TransferGatewayWithdrawalReceipt) Reset()         { *m = TransferGatewayWithdrawalReceipt{} }
func (m *TransferGatewayWithdrawalReceipt) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayWithdrawalReceipt) ProtoMessage()    {}
func (*TransferGatewayWithdrawalReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{2}
}

func (m *TransferGatewayWithdrawalReceipt) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayWithdrawalReceipt) GetTokenContract() *types.Address {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *TransferGatewayWithdrawalReceipt) GetTokenKind() TransferGatewayTokenKind {
	if m != nil {
		return m.TokenKind
	}
	return TransferGatewayTokenKind_ETH
}

func (m *TransferGatewayWithdrawalReceipt) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TransferGatewayWithdrawalReceipt) GetWithdrawalNonce() uint64 {
	if m != nil {
		return m.WithdrawalNonce
	}
	return 0
}

func (m *TransferGatewayWithdrawalReceipt) GetOracleSignature() []byte {
	if m != nil {
		return m.OracleSignature
	}
	return nil
}

type TransferGatewayPendingWithdrawalSummary struct {
	// DAppChain address of token owner
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// Hash of the unsigned withdrawal receipt
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *TransferGatewayPendingWithdrawalSummary) Reset() {
	*m = TransferGatewayPendingWithdrawalSummary{}
}
func (m *TransferGatewayPendingWithdrawalSummary) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayPendingWithdrawalSummary) ProtoMessage()    {}
func (*TransferGatewayPendingWithdrawalSummary) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{3}
}

func (m *TransferGatewayPendingWithdrawalSummary) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayPendingWithdrawalSummary) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type TransferGatewayAccount struct {
	Owner             *types.Address                    `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	WithdrawalNonce   uint64                            `protobuf:"varint,2,opt,name=withdrawal_nonce,json=withdrawalNonce,proto3" json:"withdrawal_nonce,omitempty"`
	WithdrawalReceipt *TransferGatewayWithdrawalReceipt `protobuf:"bytes,3,opt,name=withdrawal_receipt,json=withdrawalReceipt" json:"withdrawal_receipt,omitempty"`
}

func (m *TransferGatewayAccount) Reset()         { *m = TransferGatewayAccount{} }
func (m *TransferGatewayAccount) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayAccount) ProtoMessage()    {}
func (*TransferGatewayAccount) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{4}
}

func (m *TransferGatewayAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TransferGatewayAccount) GetWithdrawalNonce() uint64 {
	if m != nil {
		return m.WithdrawalNonce
	}
	return 0
}

func (m *TransferGatewayAccount) GetWithdrawalReceipt() *TransferGatewayWithdrawalReceipt {
	if m != nil {
		return m.WithdrawalReceipt
	}
	return nil
}

// Token Deposit (ETH/ERC20/ERC721) made into the Mainnet Gateway
type TransferGatewayTokenDeposited struct {
	// Mainnet address of token owner
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// Mainnet address of token contract, blank if ETH was deposited
	TokenContract *types.Address           `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract" json:"token_contract,omitempty"`
	TokenKind     TransferGatewayTokenKind `protobuf:"varint,3,opt,name=token_kind,json=tokenKind,proto3,enum=TransferGatewayTokenKind" json:"token_kind,omitempty"`
	// ERC721 token ID, or amount of ERC20/ETH
	Value *types.BigUInt `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *TransferGatewayTokenDeposited) Reset()         { *m = TransferGatewayTokenDeposited{} }
func (m *TransferGatewayTokenDeposited) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayTokenDeposited) ProtoMessage()    {}
func (*TransferGatewayTokenDeposited) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{5}
}

func (m *TransferGatewayTokenDeposited) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayTokenDeposited) GetTokenContract() *types.Address {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *TransferGatewayTokenDeposited) GetTokenKind() TransferGatewayTokenKind {
	if m != nil {
		return m.TokenKind
	}
	return TransferGatewayTokenKind_ETH
}

func (m *TransferGatewayTokenDeposited) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

// Withdrawal from Mainnet Transfer Gateway
type TransferGatewayTokenWithdrawn struct {
	// Mainnet address of token owner
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// Mainnet address of token contract, blank if ETH was withdrawn
	TokenContract *types.Address           `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract" json:"token_contract,omitempty"`
	TokenKind     TransferGatewayTokenKind `protobuf:"varint,3,opt,name=token_kind,json=tokenKind,proto3,enum=TransferGatewayTokenKind" json:"token_kind,omitempty"`
	// ERC721 token ID, or amount of ERC20/ETH
	Value *types.BigUInt `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *TransferGatewayTokenWithdrawn) Reset()         { *m = TransferGatewayTokenWithdrawn{} }
func (m *TransferGatewayTokenWithdrawn) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayTokenWithdrawn) ProtoMessage()    {}
func (*TransferGatewayTokenWithdrawn) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{6}
}

func (m *TransferGatewayTokenWithdrawn) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayTokenWithdrawn) GetTokenContract() *types.Address {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *TransferGatewayTokenWithdrawn) GetTokenKind() TransferGatewayTokenKind {
	if m != nil {
		return m.TokenKind
	}
	return TransferGatewayTokenKind_ETH
}

func (m *TransferGatewayTokenWithdrawn) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

type TransferGatewayMainnetEvent struct {
	EthBlock uint64 `protobuf:"varint,1,opt,name=eth_block,json=ethBlock,proto3" json:"eth_block,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*TransferGatewayMainnetEvent_Deposit
	//	*TransferGatewayMainnetEvent_Withdrawal
	Payload isTransferGatewayMainnetEvent_Payload `protobuf_oneof:"payload"`
}

func (m *TransferGatewayMainnetEvent) Reset()         { *m = TransferGatewayMainnetEvent{} }
func (m *TransferGatewayMainnetEvent) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayMainnetEvent) ProtoMessage()    {}
func (*TransferGatewayMainnetEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{7}
}

type isTransferGatewayMainnetEvent_Payload interface {
	isTransferGatewayMainnetEvent_Payload()
}

type TransferGatewayMainnetEvent_Deposit struct {
	Deposit *TransferGatewayTokenDeposited `protobuf:"bytes,2,opt,name=deposit,oneof"`
}
type TransferGatewayMainnetEvent_Withdrawal struct {
	Withdrawal *TransferGatewayTokenWithdrawn `protobuf:"bytes,3,opt,name=withdrawal,oneof"`
}

func (*TransferGatewayMainnetEvent_Deposit) isTransferGatewayMainnetEvent_Payload()    {}
func (*TransferGatewayMainnetEvent_Withdrawal) isTransferGatewayMainnetEvent_Payload() {}

func (m *TransferGatewayMainnetEvent) GetPayload() isTransferGatewayMainnetEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransferGatewayMainnetEvent) GetEthBlock() uint64 {
	if m != nil {
		return m.EthBlock
	}
	return 0
}

func (m *TransferGatewayMainnetEvent) GetDeposit() *TransferGatewayTokenDeposited {
	if x, ok := m.GetPayload().(*TransferGatewayMainnetEvent_Deposit); ok {
		return x.Deposit
	}
	return nil
}

func (m *TransferGatewayMainnetEvent) GetWithdrawal() *TransferGatewayTokenWithdrawn {
	if x, ok := m.GetPayload().(*TransferGatewayMainnetEvent_Withdrawal); ok {
		return x.Withdrawal
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransferGatewayMainnetEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransferGatewayMainnetEvent_OneofMarshaler, _TransferGatewayMainnetEvent_OneofUnmarshaler, _TransferGatewayMainnetEvent_OneofSizer, []interface{}{
		(*TransferGatewayMainnetEvent_Deposit)(nil),
		(*TransferGatewayMainnetEvent_Withdrawal)(nil),
	}
}

func _TransferGatewayMainnetEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransferGatewayMainnetEvent)
	// payload
	switch x := m.Payload.(type) {
	case *TransferGatewayMainnetEvent_Deposit:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Deposit); err != nil {
			return err
		}
	case *TransferGatewayMainnetEvent_Withdrawal:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdrawal); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransferGatewayMainnetEvent.Payload has unexpected type %T", x)
	}
	return nil
}

func _TransferGatewayMainnetEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransferGatewayMainnetEvent)
	switch tag {
	case 2: // payload.deposit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferGatewayTokenDeposited)
		err := b.DecodeMessage(msg)
		m.Payload = &TransferGatewayMainnetEvent_Deposit{msg}
		return true, err
	case 3: // payload.withdrawal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferGatewayTokenWithdrawn)
		err := b.DecodeMessage(msg)
		m.Payload = &TransferGatewayMainnetEvent_Withdrawal{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransferGatewayMainnetEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransferGatewayMainnetEvent)
	// payload
	switch x := m.Payload.(type) {
	case *TransferGatewayMainnetEvent_Deposit:
		s := proto.Size(x.Deposit)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransferGatewayMainnetEvent_Withdrawal:
		s := proto.Size(x.Withdrawal)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Emitted when a withdrawal to the Mainnet Gateway has been signed.
// This event contains all the required data to complete a toke withdrawal via the Mainnet Gateway.
type TransferGatewayTokenWithdrawalSigned struct {
	// Mainnet address of token owner
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// Mainnet address of token contract, blank if ETH
	TokenContract *types.Address           `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract" json:"token_contract,omitempty"`
	TokenKind     TransferGatewayTokenKind `protobuf:"varint,3,opt,name=token_kind,json=tokenKind,proto3,enum=TransferGatewayTokenKind" json:"token_kind,omitempty"`
	// ERC721 token ID, or amount of ERC20/ETH
	Value *types.BigUInt `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	// Oracle signature
	Sig []byte `protobuf:"bytes,5,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *TransferGatewayTokenWithdrawalSigned) Reset()         { *m = TransferGatewayTokenWithdrawalSigned{} }
func (m *TransferGatewayTokenWithdrawalSigned) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayTokenWithdrawalSigned) ProtoMessage()    {}
func (*TransferGatewayTokenWithdrawalSigned) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{8}
}

func (m *TransferGatewayTokenWithdrawalSigned) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayTokenWithdrawalSigned) GetTokenContract() *types.Address {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *TransferGatewayTokenWithdrawalSigned) GetTokenKind() TransferGatewayTokenKind {
	if m != nil {
		return m.TokenKind
	}
	return TransferGatewayTokenKind_ETH
}

func (m *TransferGatewayTokenWithdrawalSigned) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TransferGatewayTokenWithdrawalSigned) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type TransferGatewayInitRequest struct {
	// Only the owner will be allowed to add/remove oracles
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	// List of oracles that the Gateway should accept data from
	Oracles []*types.Address `protobuf:"bytes,2,rep,name=oracles" json:"oracles,omitempty"`
	// Initial value for TransferGatewayState.last_mainnet_block_num,
	// Oracles will start looking for relevant Mainnet events from this block onwards.
	// Should be set to the Mainnet block number that immediately preceeded the block containing
	// the tx that deployed the Mainnet Gateway contract.
	FirstMainnetBlockNum uint64 `protobuf:"varint,3,opt,name=first_mainnet_block_num,json=firstMainnetBlockNum,proto3" json:"first_mainnet_block_num,omitempty"`
}

func (m *TransferGatewayInitRequest) Reset()         { *m = TransferGatewayInitRequest{} }
func (m *TransferGatewayInitRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayInitRequest) ProtoMessage()    {}
func (*TransferGatewayInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{9}
}

func (m *TransferGatewayInitRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TransferGatewayInitRequest) GetOracles() []*types.Address {
	if m != nil {
		return m.Oracles
	}
	return nil
}

func (m *TransferGatewayInitRequest) GetFirstMainnetBlockNum() uint64 {
	if m != nil {
		return m.FirstMainnetBlockNum
	}
	return 0
}

type TransferGatewayAddOracleRequest struct {
	Oracle *types.Address `protobuf:"bytes,1,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *TransferGatewayAddOracleRequest) Reset()         { *m = TransferGatewayAddOracleRequest{} }
func (m *TransferGatewayAddOracleRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayAddOracleRequest) ProtoMessage()    {}
func (*TransferGatewayAddOracleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{10}
}

func (m *TransferGatewayAddOracleRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type TransferGatewayRemoveOracleRequest struct {
	Oracle *types.Address `protobuf:"bytes,1,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *TransferGatewayRemoveOracleRequest) Reset()         { *m = TransferGatewayRemoveOracleRequest{} }
func (m *TransferGatewayRemoveOracleRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayRemoveOracleRequest) ProtoMessage()    {}
func (*TransferGatewayRemoveOracleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{11}
}

func (m *TransferGatewayRemoveOracleRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type TransferGatewayGetOraclesRequest struct {
}

func (m *TransferGatewayGetOraclesRequest) Reset()         { *m = TransferGatewayGetOraclesRequest{} }
func (m *TransferGatewayGetOraclesRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayGetOraclesRequest) ProtoMessage()    {}
func (*TransferGatewayGetOraclesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{12}
}

type TransferGatewayGetOraclesResponse struct {
	Oracles []*TransferGatewayOracleState `protobuf:"bytes,1,rep,name=oracles" json:"oracles,omitempty"`
}

func (m *TransferGatewayGetOraclesResponse) Reset()         { *m = TransferGatewayGetOraclesResponse{} }
func (m *TransferGatewayGetOraclesResponse) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayGetOraclesResponse) ProtoMessage()    {}
func (*TransferGatewayGetOraclesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{13}
}

func (m *TransferGatewayGetOraclesResponse) GetOracles() []*TransferGatewayOracleState {
	if m != nil {
		return m.Oracles
	}
	return nil
}

// Batch of events from the Gateway on Ethereum mainnet
type TransferGatewayProcessEventBatchRequest struct {
	Events []*TransferGatewayMainnetEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *TransferGatewayProcessEventBatchRequest) Reset() {
	*m = TransferGatewayProcessEventBatchRequest{}
}
func (m *TransferGatewayProcessEventBatchRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayProcessEventBatchRequest) ProtoMessage()    {}
func (*TransferGatewayProcessEventBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{14}
}

func (m *TransferGatewayProcessEventBatchRequest) GetEvents() []*TransferGatewayMainnetEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type TransferGatewayStateRequest struct {
}

func (m *TransferGatewayStateRequest) Reset()         { *m = TransferGatewayStateRequest{} }
func (m *TransferGatewayStateRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayStateRequest) ProtoMessage()    {}
func (*TransferGatewayStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{15}
}

type TransferGatewayStateResponse struct {
	State *TransferGatewayState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (m *TransferGatewayStateResponse) Reset()         { *m = TransferGatewayStateResponse{} }
func (m *TransferGatewayStateResponse) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayStateResponse) ProtoMessage()    {}
func (*TransferGatewayStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{16}
}

func (m *TransferGatewayStateResponse) GetState() *TransferGatewayState {
	if m != nil {
		return m.State
	}
	return nil
}

type TransferGatewayWithdrawERC721Request struct {
	// ID of ERC721 token
	TokenId *types.BigUInt `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	// DAppChain address of ERC721 contract token belongs to
	TokenContract *types.Address `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract" json:"token_contract,omitempty"`
}

func (m *TransferGatewayWithdrawERC721Request) Reset()         { *m = TransferGatewayWithdrawERC721Request{} }
func (m *TransferGatewayWithdrawERC721Request) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayWithdrawERC721Request) ProtoMessage()    {}
func (*TransferGatewayWithdrawERC721Request) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{17}
}

func (m *TransferGatewayWithdrawERC721Request) GetTokenId() *types.BigUInt {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *TransferGatewayWithdrawERC721Request) GetTokenContract() *types.Address {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

type TransferGatewayWithdrawalReceiptRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
}

func (m *TransferGatewayWithdrawalReceiptRequest) Reset() {
	*m = TransferGatewayWithdrawalReceiptRequest{}
}
func (m *TransferGatewayWithdrawalReceiptRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayWithdrawalReceiptRequest) ProtoMessage()    {}
func (*TransferGatewayWithdrawalReceiptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{18}
}

func (m *TransferGatewayWithdrawalReceiptRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type TransferGatewayWithdrawalReceiptResponse struct {
	Receipt *TransferGatewayWithdrawalReceipt `protobuf:"bytes,1,opt,name=receipt" json:"receipt,omitempty"`
}

func (m *TransferGatewayWithdrawalReceiptResponse) Reset() {
	*m = TransferGatewayWithdrawalReceiptResponse{}
}
func (m *TransferGatewayWithdrawalReceiptResponse) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayWithdrawalReceiptResponse) ProtoMessage()    {}
func (*TransferGatewayWithdrawalReceiptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{19}
}

func (m *TransferGatewayWithdrawalReceiptResponse) GetReceipt() *TransferGatewayWithdrawalReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

type TransferGatewayConfirmWithdrawalReceiptRequest struct {
	// DAppChain address of the entity attempting to make the withdrawal
	TokenOwner *types.Address `protobuf:"bytes,1,opt,name=token_owner,json=tokenOwner" json:"token_owner,omitempty"`
	// 66-byte hash of the withdrawal hash
	OracleSignature []byte `protobuf:"bytes,2,opt,name=oracle_signature,json=oracleSignature,proto3" json:"oracle_signature,omitempty"`
	// 32-byte hash of the withdrawal details
	WithdrawalHash []byte `protobuf:"bytes,3,opt,name=withdrawal_hash,json=withdrawalHash,proto3" json:"withdrawal_hash,omitempty"`
}

func (m *TransferGatewayConfirmWithdrawalReceiptRequest) Reset() {
	*m = TransferGatewayConfirmWithdrawalReceiptRequest{}
}
func (m *TransferGatewayConfirmWithdrawalReceiptRequest) String() string {
	return proto.CompactTextString(m)
}
func (*TransferGatewayConfirmWithdrawalReceiptRequest) ProtoMessage() {}
func (*TransferGatewayConfirmWithdrawalReceiptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{20}
}

func (m *TransferGatewayConfirmWithdrawalReceiptRequest) GetTokenOwner() *types.Address {
	if m != nil {
		return m.TokenOwner
	}
	return nil
}

func (m *TransferGatewayConfirmWithdrawalReceiptRequest) GetOracleSignature() []byte {
	if m != nil {
		return m.OracleSignature
	}
	return nil
}

func (m *TransferGatewayConfirmWithdrawalReceiptRequest) GetWithdrawalHash() []byte {
	if m != nil {
		return m.WithdrawalHash
	}
	return nil
}

type TransferGatewayPendingWithdrawalsRequest struct {
}

func (m *TransferGatewayPendingWithdrawalsRequest) Reset() {
	*m = TransferGatewayPendingWithdrawalsRequest{}
}
func (m *TransferGatewayPendingWithdrawalsRequest) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayPendingWithdrawalsRequest) ProtoMessage()    {}
func (*TransferGatewayPendingWithdrawalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{21}
}

type TransferGatewayPendingWithdrawalsResponse struct {
	Withdrawals []*TransferGatewayPendingWithdrawalSummary `protobuf:"bytes,1,rep,name=withdrawals" json:"withdrawals,omitempty"`
}

func (m *TransferGatewayPendingWithdrawalsResponse) Reset() {
	*m = TransferGatewayPendingWithdrawalsResponse{}
}
func (m *TransferGatewayPendingWithdrawalsResponse) String() string { return proto.CompactTextString(m) }
func (*TransferGatewayPendingWithdrawalsResponse) ProtoMessage()    {}
func (*TransferGatewayPendingWithdrawalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferGateway, []int{22}
}

func (m *TransferGatewayPendingWithdrawalsResponse) GetWithdrawals() []*TransferGatewayPendingWithdrawalSummary {
	if m != nil {
		return m.Withdrawals
	}
	return nil
}

func init() {
	proto.RegisterType((*TransferGatewayState)(nil), "TransferGatewayState")
	proto.RegisterType((*TransferGatewayOracleState)(nil), "TransferGatewayOracleState")
	proto.RegisterType((*TransferGatewayWithdrawalReceipt)(nil), "TransferGatewayWithdrawalReceipt")
	proto.RegisterType((*TransferGatewayPendingWithdrawalSummary)(nil), "TransferGatewayPendingWithdrawalSummary")
	proto.RegisterType((*TransferGatewayAccount)(nil), "TransferGatewayAccount")
	proto.RegisterType((*TransferGatewayTokenDeposited)(nil), "TransferGatewayTokenDeposited")
	proto.RegisterType((*TransferGatewayTokenWithdrawn)(nil), "TransferGatewayTokenWithdrawn")
	proto.RegisterType((*TransferGatewayMainnetEvent)(nil), "TransferGatewayMainnetEvent")
	proto.RegisterType((*TransferGatewayTokenWithdrawalSigned)(nil), "TransferGatewayTokenWithdrawalSigned")
	proto.RegisterType((*TransferGatewayInitRequest)(nil), "TransferGatewayInitRequest")
	proto.RegisterType((*TransferGatewayAddOracleRequest)(nil), "TransferGatewayAddOracleRequest")
	proto.RegisterType((*TransferGatewayRemoveOracleRequest)(nil), "TransferGatewayRemoveOracleRequest")
	proto.RegisterType((*TransferGatewayGetOraclesRequest)(nil), "TransferGatewayGetOraclesRequest")
	proto.RegisterType((*TransferGatewayGetOraclesResponse)(nil), "TransferGatewayGetOraclesResponse")
	proto.RegisterType((*TransferGatewayProcessEventBatchRequest)(nil), "TransferGatewayProcessEventBatchRequest")
	proto.RegisterType((*TransferGatewayStateRequest)(nil), "TransferGatewayStateRequest")
	proto.RegisterType((*TransferGatewayStateResponse)(nil), "TransferGatewayStateResponse")
	proto.RegisterType((*TransferGatewayWithdrawERC721Request)(nil), "TransferGatewayWithdrawERC721Request")
	proto.RegisterType((*TransferGatewayWithdrawalReceiptRequest)(nil), "TransferGatewayWithdrawalReceiptRequest")
	proto.RegisterType((*TransferGatewayWithdrawalReceiptResponse)(nil), "TransferGatewayWithdrawalReceiptResponse")
	proto.RegisterType((*TransferGatewayConfirmWithdrawalReceiptRequest)(nil), "TransferGatewayConfirmWithdrawalReceiptRequest")
	proto.RegisterType((*TransferGatewayPendingWithdrawalsRequest)(nil), "TransferGatewayPendingWithdrawalsRequest")
	proto.RegisterType((*TransferGatewayPendingWithdrawalsResponse)(nil), "TransferGatewayPendingWithdrawalsResponse")
	proto.RegisterEnum("TransferGatewayTokenKind", TransferGatewayTokenKind_name, TransferGatewayTokenKind_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/transfer_gateway/transfer_gateway.proto", fileDescriptorTransferGateway)
}

var fileDescriptorTransferGateway = []byte{
	// 935 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xce, 0xda, 0x89, 0x9d, 0x9c, 0x84, 0xd4, 0x1d, 0x4a, 0x31, 0x4d, 0x5b, 0xdc, 0x01, 0xa9,
	0x4e, 0x11, 0xeb, 0xe2, 0x12, 0x81, 0xca, 0x4d, 0x1b, 0x13, 0x9a, 0x50, 0xd1, 0x56, 0x93, 0x20,
	0x24, 0x6e, 0x56, 0x93, 0xdd, 0x89, 0x3d, 0xca, 0xee, 0x8c, 0xd9, 0x99, 0x8d, 0x15, 0x89, 0xf7,
	0xe0, 0x9e, 0x07, 0x80, 0x2b, 0x9e, 0x80, 0x67, 0xe0, 0x35, 0x78, 0x06, 0xb4, 0x33, 0xbb, 0xb6,
	0xd9, 0xdd, 0xc4, 0xc9, 0x65, 0x6e, 0xac, 0xf5, 0xf9, 0x9b, 0x73, 0xbe, 0xf3, 0x0b, 0x47, 0x43,
	0xae, 0x47, 0xc9, 0xb1, 0xeb, 0xcb, 0xa8, 0x17, 0x4a, 0x19, 0x09, 0xa6, 0x27, 0x32, 0x3e, 0xed,
	0x0d, 0xe5, 0xe7, 0xe9, 0xdf, 0xde, 0x71, 0xc2, 0x43, 0xcd, 0x45, 0x4f, 0x9f, 0x8f, 0x99, 0xea,
	0xe9, 0x98, 0x0a, 0x75, 0xc2, 0x62, 0x6f, 0x48, 0x35, 0x9b, 0xd0, 0xf3, 0x12, 0xc1, 0x1d, 0xc7,
	0x52, 0xcb, 0x7b, 0x4f, 0x17, 0x58, 0xcd, 0xac, 0xa5, 0xbf, 0x56, 0x03, 0xff, 0xee, 0xc0, 0x9d,
	0xa3, 0xcc, 0xd8, 0x2b, 0x6b, 0xeb, 0x50, 0x53, 0xcd, 0xd0, 0x43, 0x58, 0x91, 0x13, 0xc1, 0xe2,
	0xb6, 0xd3, 0x71, 0xba, 0xeb, 0xfd, 0x55, 0xf7, 0x65, 0x10, 0xc4, 0x4c, 0x29, 0x62, 0xc9, 0xe8,
	0x19, 0xdc, 0x0d, 0xa9, 0xd2, 0x5e, 0x44, 0xb9, 0x10, 0x4c, 0x7b, 0xc7, 0xa1, 0xf4, 0x4f, 0x3d,
	0x91, 0x44, 0xed, 0x5a, 0xc7, 0xe9, 0x2e, 0x93, 0xf7, 0x53, 0xee, 0x0f, 0x96, 0xb9, 0x9b, 0xf2,
	0xde, 0x24, 0x11, 0xda, 0x81, 0xdb, 0x5a, 0x9e, 0x32, 0xe1, 0x4d, 0xb8, 0x1e, 0x05, 0x31, 0x9d,
	0xb0, 0x58, 0xb5, 0xeb, 0x9d, 0xfa, 0xff, 0x1e, 0x68, 0x19, 0x91, 0x9f, 0x66, 0x12, 0xf8, 0x05,
	0xdc, 0x2b, 0xf8, 0xf8, 0x36, 0xa6, 0x7e, 0xc8, 0xac, 0xa7, 0x18, 0x9a, 0xd4, 0xaa, 0x96, 0x7c,
	0xcd, 0x19, 0xf8, 0x8f, 0x1a, 0x74, 0x0a, 0x26, 0xf2, 0x07, 0x68, 0x48, 0x98, 0xcf, 0xf8, 0x58,
	0xa3, 0x6d, 0x58, 0xb7, 0xde, 0x55, 0x07, 0x0e, 0x86, 0xf9, 0xd6, 0x44, 0xdf, 0x83, 0x4d, 0x2b,
	0xea, 0x4b, 0xa1, 0x63, 0xea, 0x6b, 0x13, 0xf5, 0xbc, 0xf4, 0x7b, 0x86, 0x3f, 0xc8, 0xd8, 0xe8,
	0x6b, 0xb0, 0xea, 0xde, 0x29, 0x17, 0x41, 0xbb, 0xde, 0x71, 0xba, 0x9b, 0xfd, 0x8f, 0xdc, 0x82,
	0x4b, 0x47, 0xa9, 0xc4, 0x6b, 0x2e, 0x02, 0xb2, 0xa6, 0xf3, 0xcf, 0x34, 0x11, 0x67, 0x34, 0x4c,
	0x58, 0x7b, 0x39, 0x7b, 0x61, 0x97, 0x0f, 0x7f, 0x3c, 0x10, 0x9a, 0x58, 0x32, 0xda, 0x86, 0xd6,
	0x64, 0x1a, 0x8a, 0x27, 0xa4, 0xf0, 0x59, 0x7b, 0xc5, 0xa4, 0xe0, 0xd6, 0x8c, 0xfe, 0x26, 0x25,
	0xa7, 0xa2, 0xd2, 0x00, 0xe7, 0x29, 0x3e, 0x14, 0x54, 0x27, 0x31, 0x6b, 0x37, 0x3a, 0x4e, 0x77,
	0x83, 0xdc, 0xb2, 0xf4, 0xc3, 0x9c, 0x8c, 0x47, 0xf0, 0xb8, 0xe0, 0xdc, 0x3b, 0x26, 0x02, 0x2e,
	0x86, 0x33, 0xd8, 0x0e, 0x93, 0x28, 0xa2, 0xf1, 0xf9, 0x75, 0x60, 0x43, 0xb0, 0x3c, 0xa2, 0x6a,
	0x64, 0xc0, 0xda, 0x20, 0xe6, 0x1b, 0xff, 0xe5, 0xc0, 0xdd, 0xc2, 0x53, 0x2f, 0x7d, 0x5f, 0x26,
	0x42, 0x2f, 0xac, 0xc1, 0xaa, 0xd0, 0x6b, 0xd5, 0xa1, 0xbf, 0x03, 0x34, 0x27, 0x1a, 0xdb, 0x8c,
	0x9b, 0x3c, 0xac, 0xf7, 0x1f, 0xb9, 0x8b, 0x4a, 0x83, 0xdc, 0x9e, 0x14, 0x49, 0xf8, 0x1f, 0x07,
	0x1e, 0x54, 0xe5, 0xef, 0x5b, 0x36, 0x96, 0x8a, 0x6b, 0x16, 0xdc, 0xcc, 0x7a, 0xba, 0x30, 0xae,
	0x1c, 0x14, 0x71, 0x43, 0xe3, 0xfa, 0xdb, 0x81, 0xad, 0x82, 0x9d, 0x6c, 0x3c, 0xed, 0x9d, 0x31,
	0xa1, 0xd1, 0x16, 0xac, 0x31, 0x3d, 0xb2, 0x73, 0xcc, 0xc4, 0xb4, 0x4c, 0x56, 0x99, 0x1e, 0x99,
	0xd9, 0x85, 0x9e, 0x43, 0x33, 0xb0, 0x79, 0xcd, 0x02, 0x78, 0xe8, 0x5e, 0x9a, 0xfb, 0xfd, 0x25,
	0x92, 0x2b, 0xa0, 0x17, 0x00, 0xb3, 0xea, 0xc9, 0x4a, 0xae, 0x5a, 0x7d, 0x0a, 0xf1, 0xfe, 0x12,
	0x99, 0xd3, 0xd9, 0x5d, 0x83, 0xe6, 0x98, 0x9e, 0x87, 0x92, 0x06, 0xf8, 0x5f, 0x07, 0x3e, 0xbd,
	0x4c, 0x95, 0x86, 0x69, 0x0f, 0xdf, 0xd4, 0xe2, 0x43, 0x2d, 0xa8, 0x2b, 0x3e, 0x34, 0xf3, 0x6b,
	0x83, 0xa4, 0x9f, 0xf8, 0x37, 0xa7, 0x34, 0xfc, 0x0f, 0x04, 0xd7, 0x84, 0xfd, 0x92, 0x30, 0xb5,
	0x78, 0x44, 0x60, 0x68, 0xda, 0xd1, 0xa6, 0xda, 0xb5, 0xc2, 0x9e, 0xc9, 0x19, 0x68, 0x07, 0x3e,
	0x3c, 0xe1, 0x71, 0xe5, 0x2e, 0xab, 0x9b, 0x3a, 0xb8, 0x63, 0xd8, 0x85, 0x65, 0x86, 0x07, 0xf0,
	0x71, 0x71, 0x6e, 0x05, 0x81, 0x5d, 0x4c, 0xb9, 0x77, 0x1d, 0x68, 0xd8, 0x47, 0x4a, 0xee, 0x65,
	0x74, 0xfc, 0x1d, 0xe0, 0x82, 0x11, 0xc2, 0x22, 0x79, 0xc6, 0xae, 0x6b, 0x07, 0x97, 0xf6, 0xdb,
	0x2b, 0xa6, 0xad, 0x11, 0x95, 0x59, 0xc1, 0x3f, 0xc3, 0xa3, 0x4b, 0x64, 0xd4, 0x58, 0x0a, 0xc5,
	0xd0, 0xce, 0x0c, 0x30, 0xc7, 0x00, 0xb6, 0xe5, 0x5e, 0xbc, 0x7b, 0xa7, 0x18, 0x62, 0xaf, 0xbc,
	0x2f, 0x62, 0xe9, 0x33, 0xa5, 0x4c, 0x73, 0xed, 0x52, 0xed, 0x8f, 0xf2, 0x60, 0xbe, 0x84, 0x06,
	0x4b, 0x89, 0xf9, 0x03, 0xf7, 0xdd, 0x4b, 0xda, 0x92, 0x64, 0xb2, 0xf8, 0x41, 0xa9, 0x7b, 0xad,
	0x07, 0x59, 0x6c, 0xaf, 0xe1, 0x7e, 0x35, 0x3b, 0x0b, 0xeb, 0x33, 0x58, 0x51, 0x29, 0x21, 0x03,
	0xf0, 0x03, 0xb7, 0x52, 0xda, 0xca, 0xe0, 0x5f, 0x4b, 0x3d, 0x96, 0xb7, 0xd7, 0x1e, 0x19, 0x7c,
	0xd5, 0xff, 0x22, 0x8f, 0xe4, 0x13, 0x58, 0xb5, 0x7d, 0xc0, 0x83, 0x69, 0x62, 0xf2, 0x82, 0x6e,
	0x1a, 0xce, 0x41, 0x70, 0xed, 0xee, 0xc2, 0x07, 0x25, 0x28, 0xcb, 0xfb, 0xe8, 0x6a, 0xd5, 0x8f,
	0x87, 0xd0, 0x5d, 0x6c, 0x2a, 0x43, 0xe8, 0x1b, 0x68, 0xe6, 0x6b, 0xd1, 0xb9, 0xea, 0x5a, 0xcc,
	0x35, 0xf0, 0x9f, 0x0e, 0x14, 0xa5, 0x07, 0x52, 0x9c, 0xf0, 0x38, 0xba, 0xd0, 0xf7, 0x6b, 0x0c,
	0xa8, 0xaa, 0xbb, 0xa5, 0x56, 0x79, 0xb7, 0xa0, 0xc7, 0x30, 0xb7, 0xfa, 0x3d, 0x73, 0x6c, 0xd4,
	0x8d, 0xe4, 0xe6, 0x8c, 0xbc, 0x9f, 0x9e, 0x1d, 0x4f, 0x4a, 0xd0, 0x94, 0x0e, 0x9c, 0x69, 0xe3,
	0x4c, 0x60, 0xfb, 0x0a, 0xb2, 0x19, 0x8e, 0xdf, 0xc3, 0xfa, 0xec, 0xa9, 0xbc, 0xc6, 0xbb, 0xee,
	0x15, 0xaf, 0x29, 0x32, 0xaf, 0xfc, 0xe4, 0x39, 0xb4, 0x2f, 0x9a, 0xaa, 0xa8, 0x09, 0xf5, 0xbd,
	0xa3, 0xfd, 0xd6, 0x12, 0x5a, 0x83, 0x95, 0x3d, 0x32, 0xe8, 0x3f, 0x6d, 0x39, 0x08, 0xa0, 0x61,
	0x2b, 0xb4, 0x55, 0x3b, 0x6e, 0x98, 0x03, 0xff, 0xd9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x38, 0x74, 0xf0, 0x6a, 0x0c, 0x00, 0x00,
}
