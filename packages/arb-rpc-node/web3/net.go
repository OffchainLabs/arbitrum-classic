/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package web3

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type Net struct {
	chainId uint64
	counter *prometheus.CounterVec
}

func (net *Net) Version() string {
	net.counter.WithLabelValues("net_version", "true").Inc()
	return strconv.FormatUint(net.chainId, 10)
}
