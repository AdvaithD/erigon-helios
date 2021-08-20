package commands

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
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

func(api *HeliosAPIImpl) GetInternalOperations(ctx context.Context, hash common.Hash) ([]string, error) {
    fmt.Println("Not implemented!)
}
