package nodehealth

import (
	"errors"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Configuration struct
type configStruct struct {
	mu                         sync.Mutex
	pollingRate                time.Duration
	loopDelayTimer             time.Duration
	healthcheckRPC             string
	openethereumHealthcheckRPC string
	primaryHealthcheckRPC      string
	successCode                int
	blockDifferenceTolerance   int64
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
	ValBigInt big.Int
	ValTime   time.Duration
}

type inboxReaderState struct {
	getNextBlockToRead big.Int
	currentHeight      big.Int
	arbCorePosition    big.Int
	caughtUpTarget     big.Int
}

type healthState struct {
	mu          sync.Mutex
	inboxReader inboxReaderState
}

type aSyncDataStruct struct {
	//Healthchecks to allow the status to be shared between handlers
	checkOpenethereum healthcheck.Check
	checkPrimary      healthcheck.Check
	inboxReaderStatus healthcheck.Check
}

func logger(config *configStruct, state *healthState, logMsgChan <-chan Log) {
	for {
		//Read log structure from channel
		logMessage := <-logMsgChan

		//Check if a configuration message has been sent
		if logMessage.Config {
			//Load the configuration message into the config struct
			if logMessage.Var == "openethereumHealthcheckRPC" {
				config.mu.Lock()
				config.openethereumHealthcheckRPC = logMessage.ValStr
				config.mu.Unlock()
			}
			if logMessage.Var == "primaryHealthcheckRPC" {
				config.mu.Lock()
				config.primaryHealthcheckRPC = logMessage.ValStr
				config.mu.Unlock()
			}
		} else {
			//Check if the InboxReader is sending logs
			if logMessage.Comp == "InboxReader" {
				//Load the log into the correct struct field inside the state array
				if logMessage.Var == "getNextBlockToRead" {
					state.mu.Lock()
					state.inboxReader.getNextBlockToRead = logMessage.ValBigInt
					state.mu.Unlock()
				}
				if logMessage.Var == "currentHeight" {
					state.mu.Lock()
					state.inboxReader.currentHeight = logMessage.ValBigInt
					state.mu.Unlock()
				}
				if logMessage.Var == "arbCorePosition" {
					state.mu.Lock()
					state.inboxReader.arbCorePosition = logMessage.ValBigInt
					state.mu.Unlock()
				}
				if logMessage.Var == "caughtUpTarget" {
					state.mu.Lock()
					state.inboxReader.caughtUpTarget = logMessage.ValBigInt
					state.mu.Unlock()
				}
			}
		}
	}
}

// Default configuration values for the healthcheck server
func (config *configStruct) loadConfig() {
	const defaultSuccessCode = 200
	const defaultBlockDifferenceTolerance = 2
	const defaultPollingRate = 10 * time.Second
	const loopDelayTimer = 1 * time.Second
	config.healthcheckRPC = "0.0.0.0:8080"
	config.pollingRate = defaultPollingRate
	config.loopDelayTimer = loopDelayTimer
	config.openethereumHealthcheckRPC = ""
	config.primaryHealthcheckRPC = ""
	config.successCode = defaultSuccessCode
	config.blockDifferenceTolerance = defaultBlockDifferenceTolerance
}

//Perform all upstream checks at a set time interval in an asynchronous manner
func aSyncUpstream(aSyncData *aSyncDataStruct, state *healthState, config *configStruct) {
	//Check the healthcheck endpoint for OpenEthereum
	aSyncData.checkOpenethereum = healthcheck.Async(func() error {
		//Lock config mutex for read operation
		config.mu.Lock()
		//Retrieve status code from healthcheck endpoint
		res, err := http.Get(config.openethereumHealthcheckRPC + "/ready")
		//Unlock config mutex
		config.mu.Unlock()
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

	//Check the primary endpoint
	aSyncData.checkPrimary = healthcheck.Async(func() error {
		//Lock config mutex for read operation
		config.mu.Lock()
		if config.primaryHealthcheckRPC != "" {
			//If the primary is being used, retrieve endpoint response code
			res, err := http.Get(config.primaryHealthcheckRPC + "/ready")
			//Unlock the config mutex
			config.mu.Unlock()
			//Check the response code to determine if the primary is ready
			if err != nil {
				return err
			}
			if res.StatusCode != config.successCode {
				//The server is returning an unexpected status code
				return errors.New("Primary not ready")
			}
		} else {
			//Make sure the config mutex is unlocked before exiting
			config.mu.Unlock()
		}
		return nil
	}, config.pollingRate)

	//Check how many blocks the inboxReader is behind
	aSyncData.inboxReaderStatus = healthcheck.Async(func() error {
		//Lock config mutex for read operation
		state.mu.Lock()
		//Calculate out the block difference
		const subtractionBigInt = 0
		blockDifference := big.NewInt(subtractionBigInt).Sub(&state.inboxReader.caughtUpTarget, &state.inboxReader.currentHeight)
		//Set the tolerance we are willing to accept
		tolerance := big.NewInt(config.blockDifferenceTolerance)
		//Compare the tolerance using CmpAbs, fail if > then tolerance
		const greaterThanReturn = 1
		if blockDifference.CmpAbs(tolerance) == greaterThanReturn {
			state.mu.Unlock()
			return errors.New("InboxReader catching up")
		}
		//Unlock config mutex
		state.mu.Unlock()
		return nil
	}, config.pollingRate)
}

//Define which healthchecks to use for the readiness API and expose the readiness API
func nodeReadinessChecks(health healthcheck.Handler, httpMux *http.ServeMux, aSyncData *aSyncDataStruct) {
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

//Define which healthchecks to use for the liveness API and expose the liveness API
func nodeLivenessChecks(health healthcheck.Handler, httpMux *http.ServeMux,
	aSyncData *aSyncDataStruct, config configStruct) {
	//Create an endpoint to serve the liveness check
	httpMux.HandleFunc("/live", health.LiveEndpoint)
}

//Expose the prometheus metrics API along with the raw responses from OpenEthereum
func nodeMetrics(httpMux *http.ServeMux, prometheusRegistry *prometheus.Registry) {
	//Create an endpoint to serve the prometheus endpoint
	httpMux.Handle("/metrics", promhttp.HandlerFor(
		prometheusRegistry, promhttp.HandlerOpts{}))
}

//Create the HTTP server and start a watchdog to monitor its return codes
func healthcheckServer(httpMux *http.ServeMux, config configStruct) {
	http.ListenAndServe(config.healthcheckRPC, httpMux)
}

func waitConfig(config *configStruct) {
	config.mu.Lock()
	for config.openethereumHealthcheckRPC == "" {
		config.mu.Unlock()
		time.Sleep(config.loopDelayTimer)
		config.mu.Lock()
	}
	config.mu.Unlock()
}

//Start the healthcheck for OpenEthereum
func startHealthCheck(config *configStruct, state *healthState) {
	//Allocate storage for the aSync calls
	aSyncData := aSyncDataStruct{}
	//Create the main healthcheck handler
	prometheusRegistry := prometheus.NewRegistry()
	health := healthcheck.NewMetricsHandler(prometheusRegistry, "healthcheck")
	//Create an HTTP server mux to serve the endpoints
	httpMux := http.NewServeMux()

	waitConfig(config)

	//Schedule the async calls
	aSyncUpstream(&aSyncData, state, config)

	//Define which healthchecks to use for the liveness API and expose the liveness API
	nodeLivenessChecks(health, httpMux, &aSyncData, *config)

	//Define which healthchecks to use for the readiness API and expose the readiness API
	nodeReadinessChecks(health, httpMux, &aSyncData)

	nodeMetrics(httpMux, prometheusRegistry)
	//Create the HTTP server and start a watchdog to monitor its return codes
	healthcheckServer(httpMux, *config)
}

//NodeHealthCheck Create a node healthcheck that listens on the given channel
//Pass configuration elements on the channel to configure endpoints
func NodeHealthCheck(logMsgChan <-chan Log) {
	//Create the configuration struct
	config := configStruct{}
	state := healthState{}

	//Load the default configuration
	config.loadConfig()

	go logger(&config, &state, logMsgChan)

	startHealthCheck(&config, &state)
}
