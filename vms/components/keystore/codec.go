// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"math"

	"github.com/shubham1dubay/metalgo/codec"
	"github.com/shubham1dubay/metalgo/codec/linearcodec"
	"github.com/shubham1dubay/metalgo/utils"
)

const CodecVersion = 0

var (
	Codec       codec.Manager
	LegacyCodec codec.Manager
)

func init() {
	c := linearcodec.NewDefault()
	Codec = codec.NewDefaultManager()
	lc := linearcodec.NewDefault()
	LegacyCodec = codec.NewManager(math.MaxInt32)

	err := utils.Err(
		Codec.RegisterCodec(CodecVersion, c),
		LegacyCodec.RegisterCodec(CodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}
