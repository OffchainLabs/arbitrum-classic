package nodehealth

import (
	"errors"
	"math/big"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Configuration struct
type configStruct struct {
	mu                             sync.Mutex
	pollingRate                    time.Duration
	loopDelayTimer                 time.Duration
	healthcheckRPC                 string
	openethereumHealthcheckRPC     string
	openethereumHealthcheckRPCPort string
	primaryHealthcheckRPC          string
	primaryHealthcheckRPCPort      string
	successCode                    int
	blockDifferenceTolerance       int64
}

//Log structure for passing messages on healthChan to logger
type Log struct {
	Err    error
	Sev    string
	Var    string
	Comp   string
	Debug  bool
	Config bool

	ValStr    string
	ValBigInt *big.Int
	ValBool   bool
	ValTime   time.Duration
}

type inboxReaderState struct {
	loadingDatabase    bool
	getNextBlockToRead *big.Int
	currentHeight      *big.Int
	arbCorePosition    *big.Int
	caughtUpTarget     *big.Int
}

type healthState struct {
	mu          sync.Mutex
	inboxReader inboxReaderState
}

type asyncDataStruct struct {
	//Healthchecks to allow the status to be shared between handlers
	checkOpenethereum healthcheck.Check
	checkPrimary      healthcheck.Check
	inboxReaderStatus healthcheck.Check
}

//Perform all upstream checks at a set time interval in an asynchronous manner
func newAsyncUpstream(state *healthState, config *configStruct) *asyncDataStruct {
	asyncData := asyncDataStruct{}

	//Check the healthcheck endpoint for OpenEthereum
	asyncData.checkOpenethereum = checkEndpoint(config, &config.openethereumHealthcheckRPC, &config.openethereumHealthcheckRPCPort)

	//Check the primary endpoint
	asyncData.checkPrimary = checkEndpoint(config, &config.primaryHealthcheckRPC, &config.primaryHealthcheckRPCPort)

	//Check how many blocks the inboxReader is behind
	asyncData.inboxReaderStatus = checkInboxReader(config, state)

	return &asyncData
}

func updateInboxReader(state *healthState, logMessage Log) {
	state.mu.Lock()
	defer state.mu.Unlock()

	//Load the log into the correct struct field inside the state array
	if logMessage.Var == "loadingDatabase" {
		state.inboxReader.loadingDatabase = logMessage.ValBool
	}
	if logMessage.Var == "getNextBlockToRead" {
		state.inboxReader.getNextBlockToRead.Set(logMessage.ValBigInt)
	}
	if logMessage.Var == "currentHeight" {
		state.inboxReader.currentHeight.Set(logMessage.ValBigInt)
	}
	if logMessage.Var == "arbCorePosition" {
		state.inboxReader.arbCorePosition.Set(logMessage.ValBigInt)
	}
	if logMessage.Var == "caughtUpTarget" {
		state.inboxReader.caughtUpTarget.Set(logMessage.ValBigInt)
	}
}

func updateConfig(config *configStruct, logMessage Log) {
	config.mu.Lock()
	defer config.mu.Unlock()
	//Load the configuration message into the config struct
	if logMessage.Var == "openethereumHealthcheckRPC" {
		u, err := url.Parse(logMessage.ValStr)
		if err != nil {
			return
		}
		config.openethereumHealthcheckRPC = u.Hostname()
	}
	if logMessage.Var == "openethereumHealthcheckRPCPort" {
		config.openethereumHealthcheckRPCPort = logMessage.ValStr
	}
	if logMessage.Var == "primaryHealthcheckRPC" {
		u, err := url.Parse(logMessage.ValStr)
		if err != nil {
			return
		}
		config.primaryHealthcheckRPC = u.Hostname()
	}
	if logMessage.Var == "primaryHealthcheckRPCPort" {
		config.primaryHealthcheckRPCPort = logMessage.ValStr
	}
}

func logger(config *configStruct, state *healthState, logMsgChan <-chan Log) {
	for {
		//Read log structure from channel
		logMessage := <-logMsgChan
		//Check if a configuration message has been sent
		if logMessage.Config {
			updateConfig(config, logMessage)
		} else {
			//Check if the InboxReader is sending logs
			if logMessage.Comp == "InboxReader" {
				updateInboxReader(state, logMessage)
			}
		}
	}
}

// Default configuration values for the healthcheck server
func newConfig() *configStruct {
	config := configStruct{}

	const defaultSuccessCode = 200
	const defaultBlockDifferenceTolerance = 2
	const defaultPollingRate = 10 * time.Second
	const loopDelayTimer = 1 * time.Second
	const defaultHealthCheckPort = "8080"

	config.healthcheckRPC = "0.0.0.0:8080"
	config.openethereumHealthcheckRPCPort = defaultHealthCheckPort
	config.primaryHealthcheckRPCPort = defaultHealthCheckPort
	config.pollingRate = defaultPollingRate
	config.loopDelayTimer = loopDelayTimer
	config.openethereumHealthcheckRPC = ""
	config.primaryHealthcheckRPC = ""
	config.successCode = defaultSuccessCode
	config.blockDifferenceTolerance = defaultBlockDifferenceTolerance

	return &config
}

func newHealthState() *healthState {
	state := healthState{}

	state.inboxReader.loadingDatabase = true
	state.inboxReader.currentHeight = new(big.Int)
	state.inboxReader.caughtUpTarget = new(big.Int)
	state.inboxReader.arbCorePosition = new(big.Int)
	state.inboxReader.getNextBlockToRead = new(big.Int)

	return &state
}

func checkEndpoint(config *configStruct, endpoint *string, port *string) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Lock config mutex for read operation
		config.mu.Lock()
		endpointStr := *endpoint
		portStr := *port
		config.mu.Unlock()
		if endpointStr == "" {
			return nil
		}

		//Retrieve status code from healthcheck endpoint
		res, err := http.Get("http://" + endpointStr + ":" + portStr + "/ready")

		//Check the response code to determine if OpenEthereum is reeady
		if err != nil {
			return err
		}
		if res.StatusCode != config.successCode {
			//The server is returning an unexpected status code
			return errors.New("OpenEthereum not ready")
		}
		return nil
	}, config.pollingRate)
	return check
}

func checkInboxReader(config *configStruct, state *healthState) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Lock config mutex for read operation
		state.mu.Lock()
		defer state.mu.Unlock()

		if state.inboxReader.loadingDatabase == true {
			return errors.New("Loading database snapshot")
		}

		//Calculate out the block difference
		blockDifference := new(big.Int).Sub(state.inboxReader.caughtUpTarget, state.inboxReader.arbCorePosition)
		//Set the tolerance we are willing to accept
		tolerance := big.NewInt(config.blockDifferenceTolerance)
		//Compare the tolerance using CmpAbs, fail if > then tolerance
		if blockDifference.CmpAbs(tolerance) > 0 {
			return errors.New("InboxReader catching up block " + state.inboxReader.arbCorePosition.String() + " of " + state.inboxReader.caughtUpTarget.String())
		}
		return nil
	}, config.pollingRate)
	return check
}

//Define which healthchecks to use for the readiness API and expose the readiness API
func nodeReadinessChecks(health healthcheck.Handler, httpMux *http.ServeMux, aSyncData *asyncDataStruct) {
	//Add healthchecks to the readiness check
	health.AddReadinessCheck(
		"openethereum-status",
		aSyncData.checkOpenethereum)

	health.AddReadinessCheck(
		"primary-status",
		aSyncData.checkPrimary)

	health.AddReadinessCheck(
		"inbox-reader-status",
		aSyncData.inboxReaderStatus)

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)
}

func waitConfig(config *configStruct) {
	config.mu.Lock()
	defer config.mu.Unlock()
	for config.openethereumHealthcheckRPC == "" {
		config.mu.Unlock()
		time.Sleep(config.loopDelayTimer)
		config.mu.Lock()
	}
}

//Start the healthcheck for OpenEthereum
func startHealthCheck(config *configStruct, state *healthState) error {
	//Allocate storage for the aSync calls

	//Create the main healthcheck handler
	prometheusRegistry := prometheus.NewRegistry()
	health := healthcheck.NewMetricsHandler(prometheusRegistry, "healthcheck")
	//Create an HTTP server mux to serve the endpoints
	httpMux := http.NewServeMux()

	waitConfig(config)

	//Schedule the async calls
	asyncUpstream := newAsyncUpstream(state, config)

	//Create an endpoint to serve the liveness check
	httpMux.HandleFunc("/live", health.LiveEndpoint)

	//Define which healthchecks to use for the readiness API and expose the readiness API
	nodeReadinessChecks(health, httpMux, asyncUpstream)

	//Create an endpoint to serve the prometheus endpoint
	httpMux.Handle("/metrics", promhttp.HandlerFor(
		prometheusRegistry, promhttp.HandlerOpts{}))

	//Create the HTTP server and start a watchdog to monitor its return codes
	err := http.ListenAndServe(config.healthcheckRPC, httpMux)
	return err
}

// NodeHealthCheck Create a node healthcheck that listens on the given channel
func NodeHealthCheck(logMsgChan <-chan Log) error {
	//Create the configuration struct
	state := newHealthState()

	//Load the default configuration
	config := newConfig()

	go logger(config, state, logMsgChan)

	err := startHealthCheck(config, state)

	return err
}
