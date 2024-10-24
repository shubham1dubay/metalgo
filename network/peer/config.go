// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/message"
	"github.com/shubham1dubay/metalgo/network/throttling"
	"github.com/shubham1dubay/metalgo/snow/networking/router"
	"github.com/shubham1dubay/metalgo/snow/networking/tracker"
	"github.com/shubham1dubay/metalgo/snow/uptime"
	"github.com/shubham1dubay/metalgo/snow/validators"
	"github.com/shubham1dubay/metalgo/utils/logging"
	"github.com/shubham1dubay/metalgo/utils/set"
	"github.com/shubham1dubay/metalgo/utils/timer/mockable"
	"github.com/shubham1dubay/metalgo/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize int
	Clock           mockable.Clock
	Metrics         *Metrics
	MessageCreator  message.Creator

	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	MySubnets            set.Set[ids.ID]
	Beacons              validators.Manager
	Validators           validators.Manager
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	SupportedACPs []uint32
	ObjectedACPs  []uint32

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker tracker.ResourceTracker

	// Calculates uptime of peers
	UptimeCalculator uptime.Calculator

	// Signs my IP so I can send my signed IP address in the Handshake message
	IPSigner *IPSigner
}
