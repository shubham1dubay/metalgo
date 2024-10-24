// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/vms/components/avax"
	"github.com/shubham1dubay/metalgo/vms/components/verify"
	"github.com/shubham1dubay/metalgo/vms/platformvm/txs"
)

func TestNewApricotAtomicBlock(t *testing.T) {
	require := require.New(t)

	parentID := ids.GenerateTestID()
	height := uint64(1337)
	tx := &txs.Tx{
		Unsigned: &txs.ImportTx{
			BaseTx: txs.BaseTx{
				BaseTx: avax.BaseTx{
					Ins:  []*avax.TransferableInput{},
					Outs: []*avax.TransferableOutput{},
				},
			},
			ImportedInputs: []*avax.TransferableInput{},
		},
		Creds: []verify.Verifiable{},
	}
	require.NoError(tx.Initialize(txs.Codec))

	blk, err := NewApricotAtomicBlock(
		parentID,
		height,
		tx,
	)
	require.NoError(err)

	// Make sure the block and tx are initialized
	require.NotEmpty(blk.Bytes())
	require.NotEmpty(blk.Tx.Bytes())
	require.NotEqual(ids.Empty, blk.Tx.ID())
	require.Equal(tx.Bytes(), blk.Tx.Bytes())
	require.Equal(parentID, blk.Parent())
	require.Equal(height, blk.Height())
}