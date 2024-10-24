// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package upgrade

import (
	"flag"
	"fmt"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"

	"github.com/shubham1dubay/metalgo/tests/fixture/e2e"
	"github.com/shubham1dubay/metalgo/tests/fixture/tmpnet"
)

func TestUpgrade(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "upgrade test suites")
}

var (
	avalancheGoExecPath            string
	avalancheGoExecPathToUpgradeTo string
)

func init() {
	flag.StringVar(
		&avalancheGoExecPath,
		"metalgo-path",
		"",
		"metalgo executable path",
	)
	flag.StringVar(
		&avalancheGoExecPathToUpgradeTo,
		"metalgo-path-to-upgrade-to",
		"",
		"metalgo executable path to upgrade to",
	)
}

var _ = ginkgo.Describe("[Upgrade]", func() {
	require := require.New(ginkgo.GinkgoT())

	ginkgo.It("can upgrade versions", func() {
		network := &tmpnet.Network{
			Owner: "avalanchego-upgrade",
		}
		e2e.StartNetwork(network, avalancheGoExecPath, "" /* pluginDir */, 0 /* shutdownDelay */)

		ginkgo.By(fmt.Sprintf("restarting all nodes with %q binary", avalancheGoExecPathToUpgradeTo))
		for _, node := range network.Nodes {
			ginkgo.By(fmt.Sprintf("restarting node %q with %q binary", node.NodeID, avalancheGoExecPathToUpgradeTo))
			require.NoError(node.Stop(e2e.DefaultContext()))

			node.RuntimeConfig.AvalancheGoPath = avalancheGoExecPathToUpgradeTo

			require.NoError(network.StartNode(e2e.DefaultContext(), ginkgo.GinkgoWriter, node))

			ginkgo.By(fmt.Sprintf("waiting for node %q to report healthy after restart", node.NodeID))
			e2e.WaitForHealthy(node)
		}

		e2e.CheckBootstrapIsPossible(network)
	})
})
