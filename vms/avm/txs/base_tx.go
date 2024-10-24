// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/snow"
	"github.com/shubham1dubay/metalgo/utils/set"
	"github.com/shubham1dubay/metalgo/vms/components/avax"
	"github.com/shubham1dubay/metalgo/vms/secp256k1fx"
)

var (
	_ UnsignedTx             = (*BaseTx)(nil)
	_ secp256k1fx.UnsignedTx = (*BaseTx)(nil)
)

// BaseTx is the basis of all transactions.
type BaseTx struct {
	avax.BaseTx `serialize:"true"`

	bytes []byte
}

func (t *BaseTx) InitCtx(ctx *snow.Context) {
	for _, out := range t.Outs {
		out.InitCtx(ctx)
	}
}

func (t *BaseTx) SetBytes(bytes []byte) {
	t.bytes = bytes
}

func (t *BaseTx) Bytes() []byte {
	return t.bytes
}

func (t *BaseTx) InputIDs() set.Set[ids.ID] {
	inputIDs := set.NewSet[ids.ID](len(t.Ins))
	for _, in := range t.Ins {
		inputIDs.Add(in.InputID())
	}
	return inputIDs
}

func (t *BaseTx) Visit(v Visitor) error {
	return v.BaseTx(t)
}