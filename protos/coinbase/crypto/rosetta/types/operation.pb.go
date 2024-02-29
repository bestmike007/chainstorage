// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: coinbase/crypto/rosetta/types/operation.proto

// The stable release for rosetta types

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Operations contain all balance-changing information within a transaction. They are always
// one-sided (only affect 1 AccountIdentifier) and can succeed or fail independently from a
// Transaction. Operations are used both to represent on-chain data (Data API) and to
// construct new transactions (Construction API), creating a standard interface for reading
// and writing to blockchains.
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation_identifier uniquely identifies an operation within a transaction.
	OperationIdentifier *OperationIdentifier `protobuf:"bytes,1,opt,name=operation_identifier,json=operationIdentifier,proto3" json:"operation_identifier,omitempty"`
	// Restrict referenced related_operations to identifier indices " the current
	// operation_identifier.index. This ensures there exists a clear DAG-structure of relations.
	// Since operations are one-sided, one could imagine relating operations in a single transfer
	// or linking operations in a call tree.
	RelatedOperations []*OperationIdentifier `protobuf:"bytes,2,rep,name=related_operations,json=relatedOperations,proto3" json:"related_operations,omitempty"`
	// Type is the network-specific type of the operation. Ensure that any type that can be returned
	// here is also specified in the NetworkOptionsResponse. This can be very useful to downstream
	// consumers that parse all block data.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Status is the network-specific status of the operation. Status is not defined on the transaction
	// object because blockchains with smart contracts may have transactions that partially apply (some
	// operations are successful and some are not). Blockchains with atomic transactions (all operations
	// succeed or all operations fail) will have the same status for each operation. On-chain operations
	// (operations retrieved in the /block and /block/transaction endpoints) MUST have a populated status
	// field (anything on-chain must have succeeded or failed). However, operations provided during
	// transaction construction (often times called "intent" in the documentation) MUST NOT have a
	// populated status field (operations yet to be included on-chain have not yet succeeded or failed).
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// The account_identifier uniquely identifies an account within a network. All fields in the
	// account_identifier are utilized to determine this uniqueness (including the metadata field,
	// if populated).
	Account *AccountIdentifier `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
	// Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.
	Amount *Amount `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	// CoinChange is used to represent a change in state of a some coin identified by a coin_identifier.
	// This object is part of the Operation model and must be populated for UTXO-based blockchains.
	// Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and
	// UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate
	// this model).
	CoinChange *CoinChange `protobuf:"bytes,7,opt,name=coin_change,json=coinChange,proto3" json:"coin_change,omitempty"`
	// Metadata stores any protocol specific information regarding the operation
	Metadata map[string]*anypb.Any `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetOperationIdentifier() *OperationIdentifier {
	if x != nil {
		return x.OperationIdentifier
	}
	return nil
}

func (x *Operation) GetRelatedOperations() []*OperationIdentifier {
	if x != nil {
		return x.RelatedOperations
	}
	return nil
}

func (x *Operation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Operation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Operation) GetAccount() *AccountIdentifier {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Operation) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Operation) GetCoinChange() *CoinChange {
	if x != nil {
		return x.CoinChange
	}
	return nil
}

func (x *Operation) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// OperationIdentifier uniquely identifies an operation within a transaction.
type OperationIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation index is used to ensure each operation has a unique identifier within a transaction.
	// This index is only relative to the transaction and NOT GLOBAL. The operations in each transaction
	// should start from index 0. To clarify, there may not be any notion of an operation index in the
	// blockchain being described.
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// Some blockchains specify an operation index that is essential for client use. For example, Bitcoin
	// uses a network_index to identify which UTXO was used in a transaction. network_index should not be
	// populated if there is no notion of an operation index in a blockchain (typically most
	// account-based blockchains).
	// Note that this this should be >= 0
	NetworkIndex int64 `protobuf:"varint,2,opt,name=network_index,json=networkIndex,proto3" json:"network_index,omitempty"`
}

func (x *OperationIdentifier) Reset() {
	*x = OperationIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationIdentifier) ProtoMessage() {}

func (x *OperationIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationIdentifier.ProtoReflect.Descriptor instead.
func (*OperationIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_operation_proto_rawDescGZIP(), []int{1}
}

func (x *OperationIdentifier) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *OperationIdentifier) GetNetworkIndex() int64 {
	if x != nil {
		return x.NetworkIndex
	}
	return 0
}

var File_coinbase_crypto_rosetta_types_operation_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_operation_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74,
	0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f,
	0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x04,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x14, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65,
	0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x13, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x61, 0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0b,
	0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0a, 0x63, 0x6f,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73,
	0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x50, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73,
	0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_operation_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_operation_proto_rawDescData = file_coinbase_crypto_rosetta_types_operation_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_operation_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_operation_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_operation_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_operation_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_crypto_rosetta_types_operation_proto_goTypes = []interface{}{
	(*Operation)(nil),           // 0: coinbase.crypto.rosetta.types.Operation
	(*OperationIdentifier)(nil), // 1: coinbase.crypto.rosetta.types.OperationIdentifier
	nil,                         // 2: coinbase.crypto.rosetta.types.Operation.MetadataEntry
	(*AccountIdentifier)(nil),   // 3: coinbase.crypto.rosetta.types.AccountIdentifier
	(*Amount)(nil),              // 4: coinbase.crypto.rosetta.types.Amount
	(*CoinChange)(nil),          // 5: coinbase.crypto.rosetta.types.CoinChange
	(*anypb.Any)(nil),           // 6: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_operation_proto_depIdxs = []int32{
	1, // 0: coinbase.crypto.rosetta.types.Operation.operation_identifier:type_name -> coinbase.crypto.rosetta.types.OperationIdentifier
	1, // 1: coinbase.crypto.rosetta.types.Operation.related_operations:type_name -> coinbase.crypto.rosetta.types.OperationIdentifier
	3, // 2: coinbase.crypto.rosetta.types.Operation.account:type_name -> coinbase.crypto.rosetta.types.AccountIdentifier
	4, // 3: coinbase.crypto.rosetta.types.Operation.amount:type_name -> coinbase.crypto.rosetta.types.Amount
	5, // 4: coinbase.crypto.rosetta.types.Operation.coin_change:type_name -> coinbase.crypto.rosetta.types.CoinChange
	2, // 5: coinbase.crypto.rosetta.types.Operation.metadata:type_name -> coinbase.crypto.rosetta.types.Operation.MetadataEntry
	6, // 6: coinbase.crypto.rosetta.types.Operation.MetadataEntry.value:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_operation_proto_init() }
func file_coinbase_crypto_rosetta_types_operation_proto_init() {
	if File_coinbase_crypto_rosetta_types_operation_proto != nil {
		return
	}
	file_coinbase_crypto_rosetta_types_account_identifer_proto_init()
	file_coinbase_crypto_rosetta_types_amount_proto_init()
	file_coinbase_crypto_rosetta_types_coin_change_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinbase_crypto_rosetta_types_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_crypto_rosetta_types_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_operation_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_operation_proto_depIdxs,
		MessageInfos:      file_coinbase_crypto_rosetta_types_operation_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_operation_proto = out.File
	file_coinbase_crypto_rosetta_types_operation_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_operation_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_operation_proto_depIdxs = nil
}
