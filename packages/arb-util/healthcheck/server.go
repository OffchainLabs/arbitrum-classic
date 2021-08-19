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

package healthcheck

import (
	"context"
	"net/http"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type NodeHealth struct {
	Ready  gosundheit.Health
	Synced gosundheit.Health
	config configuration.Healthcheck
}

func NewNodeHealth(config configuration.Healthcheck, registry metrics.Registry) (*NodeHealth, error) {
	healthCheck := New(metrics.NewPrefixedChildRegistry(registry, "health/"))
	syncedCheck := New(metrics.NewPrefixedChildRegistry(registry, "sync_health/"))
	if err := syncedCheck.RegisterCheck(&checks.CustomCheck{
		CheckName: "ready",
		CheckFunc: func(ctx context.Context) (details interface{}, err error) {
			if !healthCheck.IsHealthy() {
				return nil, errors.New("core system not ready")
			}
			return nil, nil
		},
	}); err != nil {
		return nil, err
	}
	return &NodeHealth{
		Ready:  New(registry),
		Synced: New(registry),
		config: config,
	}, nil
}

func (n *NodeHealth) Launch() {
	if n.config.Enable {
		go func() {
			mux := http.NewServeMux()
			mux.Handle("/health/ready", healthhttp.HandleHealthJSON(n.Ready))
			mux.Handle("/synced/ready", healthhttp.HandleHealthJSON(n.Synced))
			if err := http.ListenAndServe(n.config.Endpoint.Addr+":"+n.config.Endpoint.Port, mux); err != nil {
				log.Error().Err(err).Msg("healthcheck server failed")
			}
		}()
	}
}
