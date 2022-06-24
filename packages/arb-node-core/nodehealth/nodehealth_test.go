/*
 * Copyright 2021, Offchain Labs, Inc.
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

package nodehealth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/heptiolabs/healthcheck"
)

type testConfigStruct struct {
	successfulStatus  int
	verbose           bool
	bufferSize        int
	readinessEndpoint string
	failMessage       string
	passMessage       string
	startUpSleepTime  time.Duration
	timeDelayTests    time.Duration
	nodehealthAddress string
	inboxReaderName   string
	failServerPort    string
	passServerPort    string
}

func newTestConfig() *testConfigStruct {
	testConfig := testConfigStruct{}

	//Configuration constants
	const successfulStatus = 200
	const verbose = true
	const bufferSize = 10
	const readinessEndpoint = "/ready"
	const failMessage = "Failed"
	const passMessage = "Passed"
	const startUpSleepTime = 5 * time.Second
	const timeDelayTests = 11 * time.Second
	const nodehealthAddress = "http://127.0.0.1:8087"
	const inboxReaderName = "InboxReader"
	const failServerPort = "8088"
	const passServerPort = "8089"

	testConfig.successfulStatus = successfulStatus
	testConfig.verbose = verbose
	testConfig.bufferSize = bufferSize
	testConfig.readinessEndpoint = readinessEndpoint
	testConfig.failMessage = failMessage
	testConfig.passMessage = passMessage
	testConfig.startUpSleepTime = startUpSleepTime
	testConfig.timeDelayTests = timeDelayTests
	testConfig.nodehealthAddress = nodehealthAddress
	testConfig.inboxReaderName = inboxReaderName
	testConfig.failServerPort = failServerPort
	testConfig.passServerPort = passServerPort

	return &testConfig
}

func startTestingServerFail(testConfig *testConfigStruct) {
	health := healthcheck.NewHandler()
	httpMux := http.NewServeMux()

	//Readiness check that always fails
	health.AddReadinessCheck("failing_check", func() error {
		return fmt.Errorf("example failure")
	})

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)

	http.ListenAndServe("127.0.0.1:"+testConfig.failServerPort, httpMux)
}

func startTestingServerPass(testConfig *testConfigStruct) {
	health := healthcheck.NewHandler()
	httpMux := http.NewServeMux()

	//Readiness check that always fails
	health.AddReadinessCheck("pass_check", func() error {
		return nil
	})

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)

	http.ListenAndServe("127.0.0.1:"+testConfig.passServerPort, httpMux)
}

func setOpenEthereumEndpoint(healthChan chan Log) {
	healthChan <- Log{Config: true, Var: "openEthereumAPI", ValStr: "https://eth-kovan.alchemyapi.io/v2/yvzMZUhX0jmdpRfqrUEGwh--U59mJNhf"}
}

func setNodeHealthBaseConfig(healthChan chan Log) {
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}
}

func healthEndpointStatus(testConfig *testConfigStruct, mode string, healthChan chan Log) error {
	function := "healthEndpointStatus: "
	time.Sleep(testConfig.startUpSleepTime)
	resp, err := http.Get(testConfig.nodehealthAddress + testConfig.readinessEndpoint)
	if err == nil {
		defer resp.Body.Close()
	}
	if mode == "unavailable" {
		if err != nil {
			if testConfig.verbose {
				fmt.Println(function + testConfig.passMessage)
			}
		} else {
			if testConfig.verbose {
				fmt.Println(function + testConfig.failMessage)
			}
			return errors.New(function + testConfig.failMessage)
		}
	} else if mode == "available" {
		if err == nil {
			if testConfig.verbose {
				fmt.Println(function + testConfig.passMessage)
			}
		} else {
			if testConfig.verbose {
				fmt.Println(function + testConfig.failMessage)
			}
			return errors.New(function + testConfig.failMessage)
		}
	}
	fmt.Println("")
	return nil
}

//Test that nodehealth properly waits for its configuration
func startUpDelayTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("startUpDelayTest")

	if testConfig.verbose {
		fmt.Println("Configure the basic required nodehealth setting")
	}
	setNodeHealthBaseConfig(healthChan)

	if testConfig.verbose {
		fmt.Println("The healthcheck server should not be running unit the OE endpoint is set and init is called")
	}
	err := healthEndpointStatus(testConfig, "unavailable", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Send init to healthcheck")
	}
	Init(healthChan)

	if testConfig.verbose {
		fmt.Println("The healthcheck server should still be waiting for the OE endpoint")
	}
	err = healthEndpointStatus(testConfig, "unavailable", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("The healthcheck should be running after the OE endpoint is set")
	}
	setOpenEthereumEndpoint(healthChan)
	err = healthEndpointStatus(testConfig, "available", healthChan)
	if err != nil {
		return err
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

//Test the healthcheck server is gracefully shutdown when the CancelFunc is called
func gracefulShutdownTest(cancel context.CancelFunc, testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("gracefulShutdownTest")

	if testConfig.verbose {
		fmt.Println("The healthcheck endpoint should be running")
	}
	err := healthEndpointStatus(testConfig, "available", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Cancel the context")
	}
	cancel()

	if testConfig.verbose {
		fmt.Println("The healthcheck endpoint should be stopped")
	}
	time.Sleep(5 * time.Second) //TCP timeout
	err = healthEndpointStatus(testConfig, "unavailable", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Check the channel logger no longer running")
	}
	for i := 0; i <= testConfig.bufferSize; i++ {
		healthChan <- Log{Config: true, Var: "TESTVARIABLE", ValBool: false}
	}
	select {
	case healthChan <- Log{Config: true, Var: "TESTVARIABLE", ValBool: false}:
		return errors.New("Channel logger not stopped")
	default:
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

//Check the healthcheck disables itself when healthcheckRPC is not set
func disableHealthcheckTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("disableHealthcheckTest")

	if testConfig.verbose {
		fmt.Println("Disable the healthcheck by not setting healthcheckRPC")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}

	if testConfig.verbose {
		fmt.Println("The healthcheck endpoint should be stopped")
	}
	time.Sleep(5 * time.Second) //TCP timeout
	err := healthEndpointStatus(testConfig, "unavailable", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("The healthcheck server should not start")
	}
	Init(healthChan)

	if testConfig.verbose {
		fmt.Println("The healthcheck endpoint should be stopped")
	}
	time.Sleep(5 * time.Second) //TCP timeout
	err = healthEndpointStatus(testConfig, "unavailable", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Check the channel logger is still running to prevent blocking")
	}
	for i := 0; i <= testConfig.bufferSize; i++ {
		healthChan <- Log{Config: true, Var: "TESTVARIABLE", ValBool: false}
	}
	time.Sleep(testConfig.startUpSleepTime)
	select {
	case healthChan <- Log{Config: true, Var: "TESTVARIABLE", ValBool: false}:
	default:
		return errors.New("Channel logger improperly stopped causing a blocking condition")
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func getHealthcheckStatus(testConfig *testConfigStruct, healthChan chan Log) (map[string]string, error) {
	function := "getHealthcheckStatus"
	var parsedResp map[string]string

	time.Sleep(testConfig.startUpSleepTime)

	resp, err := http.Get(testConfig.nodehealthAddress + testConfig.readinessEndpoint + "?full=1")
	if err == nil {
		defer resp.Body.Close()
		if testConfig.verbose {
			fmt.Println(function + testConfig.passMessage)
		}
	} else {
		if testConfig.verbose {
			fmt.Println(function + testConfig.failMessage)
		}
		return nil, errors.New(function + testConfig.failMessage)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &parsedResp)
	if err != nil {
		return nil, err
	}

	return parsedResp, err
}

//Check the healthcheck disables the primary check
func disablePrimaryCheckTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("disablePrimaryCheck")

	if testConfig.verbose {
		fmt.Println("Disable the primary by setting disablePrimaryCheck to true")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}
	setOpenEthereumEndpoint(healthChan)
	Init(healthChan)

	if testConfig.verbose {
		fmt.Println("Retrieve and parse the full healthcheck server response")
	}
	respMap, err := getHealthcheckStatus(testConfig, healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Check if the response contains the primary healthcheck")
	}
	_, ok := respMap["primary_status"]
	if ok {
		return errors.New("Primary healthcheck still present after being disabled")
	}
	_, ok = respMap["openethereum_api_status"]
	if !ok {
		return errors.New("OpenEthereum healthcheck improperly disabled")
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func retrieveVerifyOpenEthereumDisabled(testConfig *testConfigStruct, healthChan chan Log) error {
	if testConfig.verbose {
		fmt.Println("Retrieve and parse the full healthcheck server response")
	}
	respMap, err := getHealthcheckStatus(testConfig, healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Check if the response contains the OpenEthereum healthcheck")
	}
	_, ok := respMap["openethereum_api_status"]
	if ok {
		return errors.New("OpenEthereum healthcheck still present after being disabled")
	}
	_, ok = respMap["primary_status"]
	if !ok {
		return errors.New("Primary healthcheck improperly disabled")
	}

	return nil
}

//Check the healthcheck disables the OpenEthereumCheck
func disableOpenEthereumCheckTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("disableOpenEthereumCheck")

	if testConfig.verbose {
		fmt.Println("Disable the OpenEthereum check by setting disableOpenEthereumCheck to true")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}
	Init(healthChan)

	err := retrieveVerifyOpenEthereumDisabled(testConfig, healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Verify setting the OpenEthereum endpoint does not affect the check being disabled")
	}
	setOpenEthereumEndpoint(healthChan)

	err = retrieveVerifyOpenEthereumDisabled(testConfig, healthChan)
	if err != nil {
		return err
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

//Check the healthcheck can disable both OpenEthereum and primary check
func disableOpenEthereumPrimaryCheckTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("disableOpenEthereumPrimaryCheck")

	if testConfig.verbose {
		fmt.Println("Disable the both OpenEthereum and primary healthcheck")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}
	Init(healthChan)

	respMap, err := getHealthcheckStatus(testConfig, healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Check if the response contains the primary healthcheck")
	}
	_, ok := respMap["primary_status"]
	if ok {
		return errors.New("Primary healthcheck still present after being disabled")
	}
	_, ok = respMap["openethereum_api_status"]
	if ok {
		return errors.New("OpenEthereum healthcheck improperly disabled")
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

//Check the healthcheck successfully disables the prometheus endpoint
func disableMetricsTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("disableOpenEthereumPrimaryCheck")

	if testConfig.verbose {
		fmt.Println("Disable the both OpenEthereum and primary healthcheck")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}
	Init(healthChan)

	if testConfig.verbose {
		fmt.Println("Retrieve metrics endpoint response code")
	}
	res, err := http.Get(testConfig.nodehealthAddress + "/metrics")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 404 {
		return errors.New("Prometheus metrics not properly disabled")
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func testServerResponse(testConfig *testConfigStruct, mode string, healthChan chan Log) error {
	function := "testServerResponse: "
	time.Sleep(testConfig.startUpSleepTime)
	//Test server response
	res, err := http.Get(testConfig.nodehealthAddress + testConfig.readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	if mode == "ready" {
		if res.StatusCode != testConfig.successfulStatus {
			return errors.New(function + "Error - node not ready")
		}
	} else if mode == "notReady" {
		if res.StatusCode == testConfig.successfulStatus {
			return errors.New(function + "Error - node ready")
		}
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func openethereumFailureTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("openethereumFailureTest")

	if testConfig.verbose {
		fmt.Println("Prepare healthcheck server for failure testing ")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}

	Init(healthChan)

	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:0000"}
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPCPort", ValStr: testConfig.passServerPort}
	healthChan <- Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: false}

	blockTest := big.NewInt(10)
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(blockTest)}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is ready before OpenEthereum fails")
	}
	err := testServerResponse(testConfig, "ready", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Simulate OpenEthereum failure")
	}
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPCPort", ValStr: testConfig.failServerPort}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is not ready after OpenEthereum fails")
	}
	err = testServerResponse(testConfig, "notReady", healthChan)
	if err != nil {
		return err
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func primaryFailureTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("primaryFailureTest")

	if testConfig.verbose {
		fmt.Println("Prepare healthcheck server for failure testing ")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}

	Init(healthChan)

	healthChan <- Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: false}
	blockTest := big.NewInt(10)
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(blockTest)}

	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:8087"}
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPCPort", ValStr: testConfig.passServerPort}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is ready before the primary fails")
	}
	err := testServerResponse(testConfig, "ready", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Simulate primary failure")
	}
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPCPort", ValStr: testConfig.failServerPort}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is not ready after primary fails")
	}
	err = testServerResponse(testConfig, "notReady", healthChan)
	if err != nil {
		return err
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func inboxReaderCatchUpTest(testConfig *testConfigStruct, healthChan chan Log) error {
	fmt.Println("inboxReaderCatchUpTest")

	if testConfig.verbose {
		fmt.Println("Prepare healthcheck server for InboxReader testing ")
	}
	healthChan <- Log{Config: true, Var: "disablePrimaryCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: true}
	healthChan <- Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- Log{Config: true, Var: "healthcheckRPC", ValStr: "127.0.0.1:8087"}

	Init(healthChan)

	healthChan <- Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: false}
	blockTest := big.NewInt(10)
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(blockTest)}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is ready when arbCorePosition == caughtUpTarget")
	}
	err := testServerResponse(testConfig, "ready", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Set caughtUpTarget > tolerance blocks away from arbCorePosition")
	}
	caughtUpTarget := big.NewInt(20)
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(caughtUpTarget)}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is not ready when arbCorePosition is greater than the tolerance from caughtUpTarget")
	}
	err = testServerResponse(testConfig, "notReady", healthChan)
	if err != nil {
		return err
	}

	if testConfig.verbose {
		fmt.Println("Set arbCorePosition > caughtUpTarget (specific issue that occurs when primary fails during refresh)")
	}
	arbCorePosition := big.NewInt(40)
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(arbCorePosition)}
	healthChan <- Log{Comp: testConfig.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(caughtUpTarget)}
	time.Sleep(testConfig.timeDelayTests)

	if testConfig.verbose {
		fmt.Println("Check the server is ready when arbCorePosition is greater than caughtUpTarget")
	}
	err = testServerResponse(testConfig, "ready", healthChan)
	if err != nil {
		return err
	}

	fmt.Println(testConfig.passMessage)
	return nil
}

func TestNodeHealth(t *testing.T) {
	//Load the unit test configuration variables
	testConfig := newTestConfig()
	healthChan := make(chan Log, testConfig.bufferSize)

	//Generate sample servers for testing
	go startTestingServerFail(testConfig)
	go startTestingServerPass(testConfig)

	//Start the healthcheck server with a background context for testing
	ctx, cancel := context.WithCancel(context.Background())

	registry := metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test the startup delay works properly
	err := startUpDelayTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Test the server performs a graceful shutdown when the CloseFunc is called
	err = gracefulShutdownTest(cancel, testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test the server doesn't bind when the healthcheck is disabled
	err = disableHealthcheckTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test the server removes the primary healthcheck when it is disabled
	err = disablePrimaryCheckTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test the server removes the OpenEthereum healthcheck when it is disabled
	err = disableOpenEthereumCheckTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test the healthcheck can disable both OpenEthereum and primary check
	err = disableOpenEthereumPrimaryCheckTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Test whether metrics are disabled on the healthcheck
	err = disableMetricsTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//OpenEthereum failure test
	err = openethereumFailureTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx, cancel = context.WithCancel(context.Background())
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Primary failure test
	err = primaryFailureTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}

	//Restart the healthcheck server with a fresh context and healthchan
	cancel()
	time.Sleep(5 * time.Second) //TCP Timeout
	ctx = context.Background()
	healthChan = make(chan Log, testConfig.bufferSize)
	registry = metrics.NewRegistry()
	go StartNodeHealthCheck(ctx, healthChan, registry)

	//Test InboxReader catch up status
	err = inboxReaderCatchUpTest(testConfig, healthChan)
	if err != nil {
		t.Fatal(err)
	}
	cancel()
}
