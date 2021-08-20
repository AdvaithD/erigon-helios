package commands

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"

    "github.com/ledgerwatch/erigon/core",
    "github.com/ledgerwatch/erigon/core/types",
    "github.com/ledgerwatch/erigon/core/rawdb",
    "github.com/ledgerwatch/erigon/ethdb",
    "github.com/ledgerwatch/erigon/internal/ethapi",
    "github.com/ledgerwatch/erigon-lib/kv",
    "github.com/ledgerwatch/erigon/common",
    "github.com/ledgerwatch/erigon",
    "github.com/holiman/uint256"
)

type HeliosAPI interface {
	GetApiLevel() uint8
	GetInternalOperations(ctx context.Context, hash common.Hash) err
}

type HeliosAPIImpl struct {
	*BaseAPI
	db kv.RoDB // read only database
}

func NewHeliosAPI(base *BaseAPI, db kv.RoDB) *HeliosAPIImpl {
	return &HeliosAPIImpl{
		BaseAPI: base,
		db:      db,
	}
}

func(api *HeliosAPIImpl) GetInternalOperations(ctx context.Context, hash common.Hash) (string res, error) {

    // begin read only transaction
    tx, err := api.db.BeginRo(ctx)

    if err != nil {
        return nil, err
    }

    defer tx.Rollback()

    txn, blockHash, _, txIndex, err := rawdb.ReadTransaction(tx, hash)

    if err != nil {
        return err
    }
    if txn == nil {
        return nil, fmt.Errorf("txn: %x not found, txn)
    }

    block, err := rawdb.ReadBlockByHash(txn, blockHash)

    if err != bil {
        return nil, err
    }

    fmt.Println("Found txn: %x", txn)
}

