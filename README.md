# Erigon Helios

Fork of erigon, with custom rpc methods.

### Notes on RawDB

```
# accessors_chain.go

ReadCanonicalHash()    : blockNumber -> common.Hash
ReadHeaderNumber()     : common.Hash -> uint64
ReadHeadHeaderHash()   : current canonical header (common.Hash)
ReadHeadBlockHash()    : current canonical head block (common.Hash)
ReadTd()               : total difficults of a block
ReadReceipts()         : get all txn receipts in a block (with metadata)
ReadBlock()            : assemble entire block -> types.Block
ReadBlockWithSenders() : returns types.Block, []common.Address

# accessors_account.go

ReadAccount()

# accessors_indexes.go

ReadTransaction()      : Get specific txn hash from database with metadata
ReadReceipt()
```
