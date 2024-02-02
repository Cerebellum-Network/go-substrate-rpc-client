package retriever

import libErr "github.com/Cerebellum-Network/go-substrate-rpc-client/v7/error"

const (
	ErrInternalStateUpdate   = libErr.Error("internal state update")
	ErrBlockRetrieval        = libErr.Error("block retrieval")
	ErrExtrinsicParsing      = libErr.Error("extrinsic parsing")
	ErrMetadataRetrieval     = libErr.Error("metadata retrieval")
	ErrCallRegistryCreation  = libErr.Error("call registry creation")
	ErrStorageEventRetrieval = libErr.Error("storage event retrieval")
	ErrEventParsing          = libErr.Error("event parsing")
	ErrEventRegistryCreation = libErr.Error("event registry creation")
)
