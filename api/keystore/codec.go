// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/shubham1dubay/metalgo/codec"
	"github.com/shubham1dubay/metalgo/codec/linearcodec"
	"github.com/shubham1dubay/metalgo/utils/units"
)

const (
	CodecVersion = 0

	maxPackerSize = 1 * units.GiB // max size, in bytes, of something being marshalled by Marshal()
)

var Codec codec.Manager

func init() {
	lc := linearcodec.NewDefault()
	Codec = codec.NewManager(maxPackerSize)
	if err := Codec.RegisterCodec(CodecVersion, lc); err != nil {
		panic(err)
	}
}
