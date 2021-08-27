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

package utils

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var logger = log.With().Caller().Stack().Str("component", "rpc").Logger()

func setupPaths(r *mux.Router, path string) ([]*mux.Route, error) {
	if len(path) == 0 {
		return nil, errors.New("must have nonempty path")
	}
	basePath := r.Path(path)
	if path[len(path)-1] != '/' {
		path += "/"
	}
	prefixPath := r.PathPrefix(path)
	return []*mux.Route{basePath, prefixPath}, nil
}

func addRPCRoutes(r *mux.Router, handler http.Handler, path string) error {
	routes, err := setupPaths(r, path)
	if err != nil {
		return err
	}
	for _, route := range routes {
		route.Handler(handler).Methods("GET", "POST", "OPTIONS")
	}
	return nil
}

func addWSRoutes(r *mux.Router, handler http.Handler, path string) error {
	routes, err := setupPaths(r, path)
	if err != nil {
		return err
	}
	for _, route := range routes {
		route.Handler(handler).Methods("GET", "POST", "OPTIONS")
	}
	return nil
}

func LaunchRPC(ctx context.Context, handler http.Handler, addr, port, path string) error {
	r := mux.NewRouter()
	rpcRoutes, err := setupPaths(r, path)
	if err != nil {
		return err
	}
	for _, route := range rpcRoutes {
		route.Handler(handler).Methods("GET", "POST", "OPTIONS")
	}
	return launchServer(ctx, r, addr, port, "rpc")
}

func LaunchWS(ctx context.Context, server *rpc.Server, addr, port, path string) error {
	r := mux.NewRouter()
	wsRoutes, err := setupPaths(r, path)
	if err != nil {
		return err
	}
	wsHandler := server.WebsocketHandler([]string{"*"})
	for _, route := range wsRoutes {
		route.Handler(wsHandler)
	}
	return launchServer(ctx, r, addr, port, "websocket")
}

func LaunchRPCAndWS(ctx context.Context, server *rpc.Server, addr, port, rpcPath, wsPath string) error {
	r := mux.NewRouter()
	rpcRoutes, err := setupPaths(r, rpcPath)
	if err != nil {
		return err
	}
	wsRoutes, err := setupPaths(r, wsPath)
	if err != nil {
		return err
	}
	for _, route := range rpcRoutes {
		route.Handler(server).Methods("GET", "POST", "OPTIONS")
	}
	wsHandler := server.WebsocketHandler([]string{"*"})
	for _, route := range wsRoutes {
		route.Handler(wsHandler)
	}
	return launchServer(ctx, r, addr, port, "rpc and websocket")
}

func launchServer(ctx context.Context, handler http.Handler, addr string, port string, serverType string) error {
	headersOk := handlers.AllowedHeaders(
		[]string{"X-Requested-With", "Content-Type", "Authorization"},
	)
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods(
		[]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	)
	h := handlers.CORS(headersOk, originsOk, methodsOk)(handler)

	logger.Info().Str("port", port).Msgf("Launching %s server over http", serverType)
	server := &http.Server{Addr: addr + ":" + port, Handler: h}

	errChan := make(chan error, 1)
	defer close(errChan)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err.Error() == http.ErrServerClosed.Error() {
			errChan <- nil
		}
	}()

	select {
	case <-ctx.Done():
		return server.Shutdown(context.Background())
	case err := <-errChan:
		return err
	}
}
