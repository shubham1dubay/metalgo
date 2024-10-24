// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/shubham1dubay/metalgo/snow"
	"github.com/shubham1dubay/metalgo/snow/uptime"
	"github.com/shubham1dubay/metalgo/utils"
	"github.com/shubham1dubay/metalgo/utils/timer/mockable"
	"github.com/shubham1dubay/metalgo/vms/platformvm/config"
	"github.com/shubham1dubay/metalgo/vms/platformvm/fx"
	"github.com/shubham1dubay/metalgo/vms/platformvm/reward"
	"github.com/shubham1dubay/metalgo/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Calculator
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}
