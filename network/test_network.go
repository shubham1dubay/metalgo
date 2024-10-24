// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"crypto"
	"errors"
	"math"
	"net"
	"runtime"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/shubham1dubay/metalgo/ids"
	"github.com/shubham1dubay/metalgo/message"
	"github.com/shubham1dubay/metalgo/network/dialer"
	"github.com/shubham1dubay/metalgo/network/peer"
	"github.com/shubham1dubay/metalgo/network/throttling"
	"github.com/shubham1dubay/metalgo/snow/networking/router"
	"github.com/shubham1dubay/metalgo/snow/networking/tracker"
	"github.com/shubham1dubay/metalgo/snow/uptime"
	"github.com/shubham1dubay/metalgo/snow/validators"
	"github.com/shubham1dubay/metalgo/staking"
	"github.com/shubham1dubay/metalgo/subnets"
	"github.com/shubham1dubay/metalgo/utils/constants"
	"github.com/shubham1dubay/metalgo/utils/crypto/bls"
	"github.com/shubham1dubay/metalgo/utils/ips"
	"github.com/shubham1dubay/metalgo/utils/logging"
	"github.com/shubham1dubay/metalgo/utils/math/meter"
	"github.com/shubham1dubay/metalgo/utils/resource"
	"github.com/shubham1dubay/metalgo/utils/set"
	"github.com/shubham1dubay/metalgo/utils/units"
)

var (
	errClosed = errors.New("closed")

	_ net.Listener    = (*noopListener)(nil)
	_ subnets.Allower = (*nodeIDConnector)(nil)
)

type noopListener struct {
	once   sync.Once
	closed chan struct{}
}

func newNoopListener() net.Listener {
	return &noopListener{
		closed: make(chan struct{}),
	}
}

func (l *noopListener) Accept() (net.Conn, error) {
	<-l.closed
	return nil, errClosed
}

func (l *noopListener) Close() error {
	l.once.Do(func() {
		close(l.closed)
	})
	return nil
}

func (*noopListener) Addr() net.Addr {
	return &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 1,
	}
}

func NewTestNetwork(
	log logging.Logger,
	networkID uint32,
	currentValidators validators.Manager,
	trackedSubnets set.Set[ids.ID],
	router router.ExternalHandler,
) (Network, error) {
	metrics := prometheus.NewRegistry()
	msgCreator, err := message.NewCreator(
		logging.NoLog{},
		metrics,
		"",
		constants.DefaultNetworkCompressionType,
		constants.DefaultNetworkMaximumInboundTimeout,
	)
	if err != nil {
		return nil, err
	}

	networkConfig := Config{
		ThrottlerConfig: ThrottlerConfig{
			InboundConnUpgradeThrottlerConfig: throttling.InboundConnUpgradeThrottlerConfig{
				UpgradeCooldown:        constants.DefaultInboundConnUpgradeThrottlerCooldown,
				MaxRecentConnsUpgraded: int(math.Ceil(constants.DefaultInboundThrottlerMaxConnsPerSec * constants.DefaultInboundConnUpgradeThrottlerCooldown.Seconds())),
			},

			InboundMsgThrottlerConfig: throttling.InboundMsgThrottlerConfig{
				MsgByteThrottlerConfig: throttling.MsgByteThrottlerConfig{
					VdrAllocSize:        constants.DefaultInboundThrottlerVdrAllocSize,
					AtLargeAllocSize:    constants.DefaultInboundThrottlerAtLargeAllocSize,
					NodeMaxAtLargeBytes: constants.DefaultInboundThrottlerNodeMaxAtLargeBytes,
				},

				BandwidthThrottlerConfig: throttling.BandwidthThrottlerConfig{
					RefillRate:   constants.DefaultInboundThrottlerBandwidthRefillRate,
					MaxBurstSize: constants.DefaultInboundThrottlerBandwidthMaxBurstSize,
				},

				CPUThrottlerConfig: throttling.SystemThrottlerConfig{
					MaxRecheckDelay: constants.DefaultInboundThrottlerCPUMaxRecheckDelay,
				},

				DiskThrottlerConfig: throttling.SystemThrottlerConfig{
					MaxRecheckDelay: constants.DefaultInboundThrottlerDiskMaxRecheckDelay,
				},

				MaxProcessingMsgsPerNode: constants.DefaultInboundThrottlerMaxProcessingMsgsPerNode,
			},
			OutboundMsgThrottlerConfig: throttling.MsgByteThrottlerConfig{
				VdrAllocSize:        constants.DefaultOutboundThrottlerVdrAllocSize,
				AtLargeAllocSize:    constants.DefaultOutboundThrottlerAtLargeAllocSize,
				NodeMaxAtLargeBytes: constants.DefaultOutboundThrottlerNodeMaxAtLargeBytes,
			},

			MaxInboundConnsPerSec: constants.DefaultInboundThrottlerMaxConnsPerSec,
		},

		HealthConfig: HealthConfig{
			Enabled:                      true,
			MinConnectedPeers:            constants.DefaultNetworkHealthMinPeers,
			MaxTimeSinceMsgReceived:      constants.DefaultNetworkHealthMaxTimeSinceMsgReceived,
			MaxTimeSinceMsgSent:          constants.DefaultNetworkHealthMaxTimeSinceMsgSent,
			MaxPortionSendQueueBytesFull: constants.DefaultNetworkHealthMaxPortionSendQueueFill,
			MaxSendFailRate:              constants.DefaultNetworkHealthMaxSendFailRate,
			SendFailRateHalflife:         constants.DefaultHealthCheckAveragerHalflife,
		},

		ProxyEnabled:           constants.DefaultNetworkTCPProxyEnabled,
		ProxyReadHeaderTimeout: constants.DefaultNetworkTCPProxyReadTimeout,

		DialerConfig: dialer.Config{
			ThrottleRps:       constants.DefaultOutboundConnectionThrottlingRps,
			ConnectionTimeout: constants.DefaultOutboundConnectionTimeout,
		},

		TimeoutConfig: TimeoutConfig{
			PingPongTimeout:      constants.DefaultPingPongTimeout,
			ReadHandshakeTimeout: constants.DefaultNetworkReadHandshakeTimeout,
		},

		PeerListGossipConfig: PeerListGossipConfig{
			PeerListNumValidatorIPs: constants.DefaultNetworkPeerListNumValidatorIPs,
			PeerListPullGossipFreq:  constants.DefaultNetworkPeerListPullGossipFreq,
			PeerListBloomResetFreq:  constants.DefaultNetworkPeerListBloomResetFreq,
		},

		DelayConfig: DelayConfig{
			InitialReconnectDelay: constants.DefaultNetworkInitialReconnectDelay,
			MaxReconnectDelay:     constants.DefaultNetworkMaxReconnectDelay,
		},

		MaxClockDifference:           constants.DefaultNetworkMaxClockDifference,
		CompressionType:              constants.DefaultNetworkCompressionType,
		PingFrequency:                constants.DefaultPingFrequency,
		AllowPrivateIPs:              !constants.ProductionNetworkIDs.Contains(networkID),
		UptimeMetricFreq:             constants.DefaultUptimeMetricFreq,
		MaximumInboundMessageTimeout: constants.DefaultNetworkMaximumInboundTimeout,

		RequireValidatorToConnect: constants.DefaultNetworkRequireValidatorToConnect,
		PeerReadBufferSize:        constants.DefaultNetworkPeerReadBufferSize,
		PeerWriteBufferSize:       constants.DefaultNetworkPeerWriteBufferSize,
	}

	networkConfig.NetworkID = networkID
	networkConfig.TrackedSubnets = trackedSubnets

	tlsCert, err := staking.NewTLSCert()
	if err != nil {
		return nil, err
	}
	tlsConfig := peer.TLSConfig(*tlsCert, nil)
	networkConfig.TLSConfig = tlsConfig
	networkConfig.TLSKey = tlsCert.PrivateKey.(crypto.Signer)
	networkConfig.BLSKey, err = bls.NewSecretKey()
	if err != nil {
		return nil, err
	}

	networkConfig.Validators = currentValidators
	networkConfig.Beacons = validators.NewManager()
	// This never actually does anything because we never initialize the P-chain
	networkConfig.UptimeCalculator = uptime.NoOpCalculator

	// TODO actually monitor usage
	// TestNetwork doesn't use disk so we don't need to track it, but we should
	// still have guardrails around cpu/memory usage.
	networkConfig.ResourceTracker, err = tracker.NewResourceTracker(
		metrics,
		resource.NoUsage,
		&meter.ContinuousFactory{},
		constants.DefaultHealthCheckAveragerHalflife,
	)
	if err != nil {
		return nil, err
	}
	networkConfig.CPUTargeter = tracker.NewTargeter(
		logging.NoLog{},
		&tracker.TargeterConfig{
			VdrAlloc:           float64(runtime.NumCPU()),
			MaxNonVdrUsage:     .8 * float64(runtime.NumCPU()),
			MaxNonVdrNodeUsage: float64(runtime.NumCPU()) / 8,
		},
		currentValidators,
		networkConfig.ResourceTracker.CPUTracker(),
	)
	networkConfig.DiskTargeter = tracker.NewTargeter(
		logging.NoLog{},
		&tracker.TargeterConfig{
			VdrAlloc:           1000 * units.GiB,
			MaxNonVdrUsage:     1000 * units.GiB,
			MaxNonVdrNodeUsage: 1000 * units.GiB,
		},
		currentValidators,
		networkConfig.ResourceTracker.DiskTracker(),
	)

	networkConfig.MyIPPort = ips.NewDynamicIPPort(net.IPv4zero, 1)

	return NewNetwork(
		&networkConfig,
		msgCreator,
		metrics,
		log,
		newNoopListener(),
		dialer.NewDialer(
			constants.NetworkType,
			dialer.Config{
				ThrottleRps:       constants.DefaultOutboundConnectionThrottlingRps,
				ConnectionTimeout: constants.DefaultOutboundConnectionTimeout,
			},
			log,
		),
		router,
	)
}

type nodeIDConnector struct {
	nodeID ids.NodeID
}

func newNodeIDConnector(nodeID ids.NodeID) *nodeIDConnector {
	return &nodeIDConnector{nodeID: nodeID}
}

func (f *nodeIDConnector) IsAllowed(nodeID ids.NodeID, _ bool) bool {
	return nodeID == f.nodeID
}
