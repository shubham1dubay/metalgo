// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package payload

import (
	"github.com/shubham1dubay/metalgo/codec"
	"github.com/shubham1dubay/metalgo/codec/linearcodec"
	"github.com/shubham1dubay/metalgo/utils"
	"github.com/shubham1dubay/metalgo/utils/units"
)

const (
	CodecVersion = 0

	MaxMessageSize = 24 * units.KiB
)

var Codec codec.Manager

func init() {
	Codec = codec.NewManager(MaxMessageSize)
	lc := linearcodec.NewDefault()

	err := utils.Err(
		lc.RegisterType(&Hash{}),
		lc.RegisterType(&AddressedCall{}),
		Codec.RegisterCodec(CodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}
