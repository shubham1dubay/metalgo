// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"github.com/shubham1dubay/metalgo/api/info"
	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/snow"
	"github.com/shubham1dubay/metalgo/utils/constants"
	"github.com/shubham1dubay/metalgo/utils/logging"
	"github.com/shubham1dubay/metalgo/vms/avm"

	stdcontext "context"
)

const Alias = "C"

var _ Context = (*context)(nil)

type Context interface {
	NetworkID() uint32
	BlockchainID() ids.ID
	AVAXAssetID() ids.ID
}

type context struct {
	networkID    uint32
	blockchainID ids.ID
	avaxAssetID  ids.ID
}

func NewContextFromURI(ctx stdcontext.Context, uri string) (Context, error) {
	infoClient := info.NewClient(uri)
	xChainClient := avm.NewClient(uri, "X")
	return NewContextFromClients(ctx, infoClient, xChainClient)
}

func NewContextFromClients(
	ctx stdcontext.Context,
	infoClient info.Client,
	xChainClient avm.Client,
) (Context, error) {
	networkID, err := infoClient.GetNetworkID(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := infoClient.GetBlockchainID(ctx, Alias)
	if err != nil {
		return nil, err
	}

	asset, err := xChainClient.GetAssetDescription(ctx, "METAL")
	if err != nil {
		return nil, err
	}

	return NewContext(
		networkID,
		chainID,
		asset.AssetID,
	), nil
}

func NewContext(
	networkID uint32,
	blockchainID ids.ID,
	avaxAssetID ids.ID,
) Context {
	return &context{
		networkID:    networkID,
		blockchainID: blockchainID,
		avaxAssetID:  avaxAssetID,
	}
}

func (c *context) NetworkID() uint32 {
	return c.networkID
}

func (c *context) BlockchainID() ids.ID {
	return c.blockchainID
}

func (c *context) AVAXAssetID() ids.ID {
	return c.avaxAssetID
}

func newSnowContext(c Context) (*snow.Context, error) {
	chainID := c.BlockchainID()
	lookup := ids.NewAliaser()
	return &snow.Context{
		NetworkID:   c.NetworkID(),
		SubnetID:    constants.PrimaryNetworkID,
		ChainID:     chainID,
		CChainID:    chainID,
		AVAXAssetID: c.AVAXAssetID(),
		Log:         logging.NoLog{},
		BCLookup:    lookup,
	}, lookup.Alias(chainID, Alias)
}
