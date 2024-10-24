// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"reflect"

	"github.com/shubham1dubay/metalgo/codec"
	"github.com/shubham1dubay/metalgo/utils"
	"github.com/shubham1dubay/metalgo/utils/logging"
	"github.com/shubham1dubay/metalgo/utils/timer/mockable"
	"github.com/shubham1dubay/metalgo/vms/avm/fxs"
	"github.com/shubham1dubay/metalgo/vms/avm/txs"
)

// CodecVersion is the current default codec version
const CodecVersion = txs.CodecVersion

var _ Parser = (*parser)(nil)

type Parser interface {
	txs.Parser

	ParseBlock(bytes []byte) (Block, error)
	ParseGenesisBlock(bytes []byte) (Block, error)
}

type parser struct {
	txs.Parser
}

func NewParser(fxs []fxs.Fx) (Parser, error) {
	p, err := txs.NewParser(fxs)
	if err != nil {
		return nil, err
	}
	c := p.CodecRegistry()
	gc := p.GenesisCodecRegistry()

	err = utils.Err(
		c.RegisterType(&StandardBlock{}),
		gc.RegisterType(&StandardBlock{}),
	)
	return &parser{
		Parser: p,
	}, err
}

func NewCustomParser(
	typeToFxIndex map[reflect.Type]int,
	clock *mockable.Clock,
	log logging.Logger,
	fxs []fxs.Fx,
) (Parser, error) {
	p, err := txs.NewCustomParser(typeToFxIndex, clock, log, fxs)
	if err != nil {
		return nil, err
	}
	c := p.CodecRegistry()
	gc := p.GenesisCodecRegistry()

	err = utils.Err(
		c.RegisterType(&StandardBlock{}),
		gc.RegisterType(&StandardBlock{}),
	)
	return &parser{
		Parser: p,
	}, err
}

func (p *parser) ParseBlock(bytes []byte) (Block, error) {
	return parse(p.Codec(), bytes)
}

func (p *parser) ParseGenesisBlock(bytes []byte) (Block, error) {
	return parse(p.GenesisCodec(), bytes)
}

func parse(cm codec.Manager, bytes []byte) (Block, error) {
	var blk Block
	if _, err := cm.Unmarshal(bytes, &blk); err != nil {
		return nil, err
	}
	return blk, blk.initialize(bytes, cm)
}
