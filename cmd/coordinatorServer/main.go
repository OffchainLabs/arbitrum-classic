/*
 * Copyright 2019, Offchain Labs, Inc.
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

package main

import (
	jsonenc "encoding/json"
	"github.com/offchainlabs/arb-validator/coordinator"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/offchainlabs/arb-avm/loader"

	"github.com/offchainlabs/arb-validator/ethvalidator"
)



func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	router.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
}

// Launches the Coordinator validator with the following command line arguments:
// 1) Compiled Arbitrum bytecode file
// 2) private key file
// 3) public addresses file (newline separated)
// 4) Global EthBridge addresses json file
// 5) ethURL
func main() {
	// Check number of args
	if len(os.Args)-1 != 5 {
		log.Fatalln("Expected five arguments")
	}

	// 1) Compiled Arbitrum bytecode
	machine, err := loader.LoadMachineFromFile(os.Args[1], true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	// 2) Private key
	keyFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err := ioutil.ReadAll(keyFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := keyFile.Close(); err != nil {
		log.Fatalln(err)
	}
	rawKey := strings.TrimSpace(string(byteValue))
	key, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		log.Fatal("HexToECDSA private key error: ", err)
	}

	// 3) All public key addresses
	addrFile, err := os.Open(os.Args[3])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err = ioutil.ReadAll(addrFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := addrFile.Close(); err != nil {
		log.Fatalln(err)
	}
	validatorHexAddrs := strings.Split(strings.TrimSpace(string(byteValue)), "\n")
	validators := make([]common.Address, len(validatorHexAddrs))
	for i, v := range validatorHexAddrs {
		validators[i] = common.HexToAddress(v)
	}

	// 4) Global EthBridge addresses json
	jsonFile, err := os.Open(os.Args[4])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		log.Fatalln(err)
	}

	var connectionInfo ethvalidator.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		log.Fatalln(err)
	}

	// 5) URL
	ethURL := os.Args[5]

	// Validator creation
	server := coordinator.NewCoordinatorServer(machine, key, validators, connectionInfo, ethURL)

	// Run server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(server, "Validator"); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/", s).Methods("GET", "POST", "OPTIONS")
	AttachProfiler(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(":1235", handlers.CORS(headersOk, originsOk, methodsOk)(r))
	if err != nil {
		panic(err)
	}
}
