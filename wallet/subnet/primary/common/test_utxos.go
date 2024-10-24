// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"slices"

	"github.com/stretchr/testify/require"

	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/utils/constants"
	"github.com/shubham1dubay/metalgo/vms/components/avax"
)

func NewDeterministicChainUTXOs(require *require.Assertions, utxoSets map[ids.ID][]*avax.UTXO) *DeterministicChainUTXOs {
	globalUTXOs := NewUTXOs()
	for subnetID, utxos := range utxoSets {
		for _, utxo := range utxos {
			require.NoError(
				globalUTXOs.AddUTXO(context.Background(), subnetID, constants.PlatformChainID, utxo),
			)
		}
	}
	return &DeterministicChainUTXOs{
		ChainUTXOs: NewChainUTXOs(constants.PlatformChainID, globalUTXOs),
	}
}

type DeterministicChainUTXOs struct {
	ChainUTXOs
}

func (c *DeterministicChainUTXOs) UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error) {
	utxos, err := c.ChainUTXOs.UTXOs(ctx, sourceChainID)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(utxos, func(a, b *avax.UTXO) int {
		return a.Compare(&b.UTXOID)
	})
	return utxos, nil
}
