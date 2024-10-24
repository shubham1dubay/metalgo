// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/shubham1dubay/metalgo/codec"
	"github.com/shubham1dubay/metalgo/codec/linearcodec"
	"github.com/shubham1dubay/metalgo/codec/reflectcodec"
	"github.com/shubham1dubay/metalgo/utils"
	"github.com/shubham1dubay/metalgo/utils/units"
)

const (
	CodecVersion            uint16 = 0
	CodecVersionWithStopVtx uint16 = 1

	// maxSize is the maximum allowed vertex size. It is necessary to deter DoS
	maxSize = units.MiB
)

var Codec codec.Manager

func init() {
	lc0 := linearcodec.New([]string{reflectcodec.DefaultTagName + "V0"})
	lc1 := linearcodec.New([]string{reflectcodec.DefaultTagName + "V1"})

	Codec = codec.NewManager(maxSize)
	err := utils.Err(
		Codec.RegisterCodec(CodecVersion, lc0),
		Codec.RegisterCodec(CodecVersionWithStopVtx, lc1),
	)
	if err != nil {
		panic(err)
	}
}
