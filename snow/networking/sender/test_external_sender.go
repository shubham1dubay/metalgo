// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"errors"
	"testing"

	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/message"
	"github.com/shubham1dubay/metalgo/snow/engine/common"
	"github.com/shubham1dubay/metalgo/subnets"
	"github.com/shubham1dubay/metalgo/utils/set"
)

var (
	_ ExternalSender = (*ExternalSenderTest)(nil)

	errSend = errors.New("unexpectedly called Send")
)

// ExternalSenderTest is a test sender
type ExternalSenderTest struct {
	TB testing.TB

	CantSend bool

	SendF func(msg message.OutboundMessage, config common.SendConfig, subnetID ids.ID, allower subnets.Allower) set.Set[ids.NodeID]
}

// Default set the default callable value to [cant]
func (s *ExternalSenderTest) Default(cant bool) {
	s.CantSend = cant
}

func (s *ExternalSenderTest) Send(
	msg message.OutboundMessage,
	config common.SendConfig,
	subnetID ids.ID,
	allower subnets.Allower,
) set.Set[ids.NodeID] {
	if s.SendF != nil {
		return s.SendF(msg, config, subnetID, allower)
	}
	if s.CantSend {
		if s.TB != nil {
			s.TB.Helper()
			s.TB.Fatal(errSend)
		}
	}
	return nil
}
