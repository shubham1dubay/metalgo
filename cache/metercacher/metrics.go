// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metercacher

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/shubham1dubay/metalgo/utils/metric"
	"github.com/shubham1dubay/metalgo/utils/wrappers"
)

func newAveragerMetric(namespace, name string, reg prometheus.Registerer, errs *wrappers.Errs) metric.Averager {
	return metric.NewAveragerWithErrs(
		namespace,
		name,
		"time (in ns) of a "+name,
		reg,
		errs,
	)
}

func newCounterMetric(namespace, name string, reg prometheus.Registerer, errs *wrappers.Errs) prometheus.Counter {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      fmt.Sprintf("# of times a %s occurred", name),
	})
	errs.Add(reg.Register(c))
	return c
}

type metrics struct {
	get           metric.Averager
	put           metric.Averager
	len           prometheus.Gauge
	portionFilled prometheus.Gauge
	hit           prometheus.Counter
	miss          prometheus.Counter
}

func (m *metrics) Initialize(
	namespace string,
	reg prometheus.Registerer,
) error {
	errs := wrappers.Errs{}
	m.get = newAveragerMetric(namespace, "get", reg, &errs)
	m.put = newAveragerMetric(namespace, "put", reg, &errs)
	m.len = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "len",
			Help:      "number of entries",
		},
	)
	errs.Add(reg.Register(m.len))
	m.portionFilled = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "portion_filled",
			Help:      "fraction of cache filled",
		},
	)
	errs.Add(reg.Register(m.portionFilled))
	m.hit = newCounterMetric(namespace, "hit", reg, &errs)
	m.miss = newCounterMetric(namespace, "miss", reg, &errs)
	return errs.Err
}