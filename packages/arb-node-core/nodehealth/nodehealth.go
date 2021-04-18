package nodehealth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
)

//Configuration struct
type configStruct struct {
	//Mutex for config struct
	mu                   sync.Mutex
	prometheusHistograms map[string]*prometheus.HistogramVec
	prometheusRegistry   *prometheus.Registry

	//Aggregator Healthcheck Config
	//Rate to poll the remote APIs at
	pollingRate                    time.Duration
	loopDelayTimer                 time.Duration
	healthcheckRPC                 string
	openethereumHealthcheckRPC     string
	openethereumHealthcheckRPCPort string
	primaryHealthcheckRPC          string
	primaryHealthcheckRPCPort      string
	successCode                    int
	blockDifferenceTolerance       int64

	//OpenEthereum Healthcheck Config
	//Address to the OpenEthereum API
	openethereumAPI string
	//Maximum time to wait for http requests
	requestTimeout time.Duration
	//Number of blocks OpenEthereum's current block can be away from the
	//estimated current block before a healthcheck error is triggered
	blockSyncDifference int64
	//Minimum number of peers that can be connected to OpenEthereum
	//without a healthcheck error being triggered
	peerMinimum int
	//Maximum time OpenEthereum can take to change its current block
	//before a healthcheck error is triggered
	blockUpdateTimeout time.Duration
	//Debug variable to print the responses from the OpenEthereum API
	printRequests bool
	//Debug variable to print the status of the configuration load
	printConfigMsg bool
}

// Default configuration values for the healthcheck server
func newConfig() *configStruct {
	config := configStruct{}
	//Global configuration
	const healthcheckRPC = ""

	//Node health configuration
	const defaultSuccessCode = 200
	const defaultBlockDifferenceTolerance = 2
	const defaultPollingRate = 10 * time.Second
	const loopDelayTimer = 1 * time.Second
	const defaultHealthCheckPort = "8080"

	//OpenEthereum health configuration
	const requestTimeout = 10 * time.Second
	const blockSyncDifference = 10
	const peerMinimum = 1
	const blockUpdateTimeout = 45 * time.Second
	const printRequests = false
	const printConfigMsg = false

	config.prometheusHistograms = make(map[string]*prometheus.HistogramVec)
	config.prometheusRegistry = prometheus.NewRegistry()

	config.healthcheckRPC = healthcheckRPC
	config.openethereumHealthcheckRPCPort = defaultHealthCheckPort
	config.primaryHealthcheckRPCPort = defaultHealthCheckPort
	config.pollingRate = defaultPollingRate
	config.loopDelayTimer = loopDelayTimer
	config.openethereumHealthcheckRPC = ""
	config.primaryHealthcheckRPC = ""
	config.successCode = defaultSuccessCode
	config.blockDifferenceTolerance = defaultBlockDifferenceTolerance

	config.openethereumAPI = ""
	config.requestTimeout = requestTimeout
	config.blockDifferenceTolerance = blockSyncDifference
	config.peerMinimum = peerMinimum
	config.blockUpdateTimeout = blockUpdateTimeout
	config.printRequests = printRequests
	config.printConfigMsg = printConfigMsg

	return &config
}

type healthState struct {
	mu          sync.Mutex
	inboxReader inboxReaderState
}

type inboxReaderState struct {
	loadingDatabase    bool
	getNextBlockToRead *big.Int
	currentHeight      *big.Int
	arbCorePosition    *big.Int
	caughtUpTarget     *big.Int
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

type asyncDataStruct struct {
	//Healthchecks to allow the status to be shared between handlers
	checkOpenethereum healthcheck.Check
	checkPrimary      healthcheck.Check
	inboxReaderStatus healthcheck.Check

	//Response structs to process the OpenEthereum responses
	ethSyncResp        OpenEthereumResponse
	parityNetPeersResp OpenEthereumResponse

	//Healthchecks to allow the status to be shared between handlers
	ethSyncCheck        healthcheck.Check
	parityNetPeersCheck healthcheck.Check
	tcpDialCheck        healthcheck.Check
	blockRefreshCheck   healthcheck.Check
	minimumPeersCheck   healthcheck.Check
	blockSyncCheck      healthcheck.Check
}

//Perform all upstream checks at a set time interval in an asynchronous manner
func newAsyncUpstream(state *healthState, config *configStruct) *asyncDataStruct {
	asyncData := asyncDataStruct{}

	if config.openethereumAPI != "" {
		//Request eth_syncing status from OpenEthereum
		asyncData.ethSyncCheck = ethSyncCheck(config, &asyncData)

		//Request eth_syncing status from OpenEthereum
		asyncData.parityNetPeersCheck = netPeersCheck(config, &asyncData)

		//Check OpenEthereum has more than peerMinimum peers currently connected to it
		asyncData.minimumPeersCheck = openEthereumPeerCount(config, &asyncData)

		//Check OpenEthereum is refreshing its currentBlock quicker than the blockUpdateTimeout
		asyncData.blockRefreshCheck = openEthereumBlockUpdateCheck(config, &asyncData)

		//Check if OpenEthereum is within blockSyncDifference from the estimated block
		asyncData.blockSyncCheck = openEthereumBlockSyncCheck(config, &asyncData)

		//Check if the OpenEthereum API is accepting pings
		asyncData.tcpDialCheck = healthcheck.Async(
			healthcheck.TCPDialCheck(config.openethereumAPI,
				config.requestTimeout), config.pollingRate)

	} else {
		//Check the healthcheck endpoint for OpenEthereum
		asyncData.checkOpenethereum = checkEndpoint(config, &config.openethereumHealthcheckRPC, &config.openethereumHealthcheckRPCPort)
	}

	//Check the primary endpoint
	asyncData.checkPrimary = checkEndpoint(config, &config.primaryHealthcheckRPC, &config.primaryHealthcheckRPCPort)

	//Check how many blocks the inboxReader is behind
	asyncData.inboxReaderStatus = checkInboxReader(config, state)

	return &asyncData
}

//OpenEthereum response struct for json parsing
type OpenEthereumResponse struct {
	//Raw json response from the OpenEthereum API
	respBody string
	//Request ID
	Id int `json:"id"`
	//JSON RPC version
	Jsonrpc string `json:"jsonrpc"`
	//Result struct to process the json
	Result OpenEthereumResult `json:"result"`
}

//OpenEthereum result field struct for json parsing
type OpenEthereumResult struct {
	//Result field for parityNetPeers
	Active    int                `json:"active"`
	Connected int                `json:"connected"`
	Max       int                `json:"max"`
	Peers     []OpenEthereumPeer `json:"peers"`

	//Result field for ethSyncing
	StartingBlock       string `json:"startingBlock"`
	CurrentBlock        string `json:"currentBlock"`
	HighestBlock        string `json:"highestBlock"`
	WarpChunksAmount    string `json:"warpChunksAmount"`
	WarpChunksProcessed string `json:"warpChunksProcessed"`
}

//OpenEthereum result peer field struct for json parsing
type OpenEthereumPeer struct {
	//Peer client ID
	Id string `json:"id"`
}

//OpenEthereum result peer network field struct for json parsing
type OpenEthereumPeerNetwork struct {
	//Peer local IP address
	localAddress string
	//Peer remote IP address
	remoteAddress string
}

//Log structure for passing messages on healthChan to logger
type Log struct {
	Err     error
	Sev     string
	Var     string
	Comp    string
	Debug   bool
	Config  bool
	Metrics bool

	ValStr    string
	ValInt    int64
	ValBigInt *big.Int
	ValBool   bool
	ValTime   time.Duration
}

func logger(config *configStruct, state *healthState, logMsgChan <-chan Log) {
	for {
		//Read log structure from channel
		logMessage := <-logMsgChan
		//Check messsage type
		if logMessage.Metrics {
			metricsHandler(config, logMessage)
		} else if logMessage.Config {
			updateConfig(config, logMessage)
		} else {
			//Check if the InboxReader is sending logs
			if logMessage.Comp == "InboxReader" {
				updateInboxReader(state, logMessage)
			}
		}
	}
}

func LogTime(comp string, function string, start time.Time, healthChan chan Log) {
	healthChan <- Log{Metrics: true, Comp: comp, Var: function, ValTime: time.Since(start)}
}

func metricsHandler(config *configStruct, logMessage Log) {
	_, ok := config.prometheusHistograms[logMessage.Comp]
	if !ok {
		config.prometheusHistograms[logMessage.Comp] = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "kovan4-0",
			Name:      logMessage.Comp,
			Help:      logMessage.Comp + " latency distributions.",
			Buckets:   prometheus.ExponentialBuckets(1, 10, 4),
		}, []string{logMessage.Comp})

		config.prometheusRegistry.MustRegister(config.prometheusHistograms[logMessage.Comp])
	}
	config.prometheusHistograms[logMessage.Comp].WithLabelValues(logMessage.Var).Observe(float64(logMessage.ValTime.Milliseconds()))
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
		fmt.Println(config.openethereumHealthcheckRPC)
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
	if logMessage.Var == "healthcheckRPC" {
		config.healthcheckRPC = logMessage.ValStr
	}
	if logMessage.Var == "primaryHealthcheckRPCPort" {
		config.primaryHealthcheckRPCPort = logMessage.ValStr
	}
	if logMessage.Var == "blockDifferenceTolerance" {
		config.blockDifferenceTolerance = logMessage.ValInt
	}
	if logMessage.Var == "peerMinimum" {
		config.peerMinimum = int(logMessage.ValInt)
	}
	if logMessage.Var == "blockUpdateTimeout" {
		config.blockUpdateTimeout = logMessage.ValTime
	}
	if logMessage.Var == "printRequests" {
		config.printRequests = logMessage.ValBool
	}
	if logMessage.Var == "printConfigMsg" {
		config.printConfigMsg = logMessage.ValBool
	}
	if logMessage.Var == "openEthereumAPI" {
		config.openethereumAPI = logMessage.ValStr
	}
}

func ethSyncCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//OpenEthereum API call to send
		var jsonRequest = []byte(`{"method":"eth_syncing","params":[],"id":1,"jsonrpc":"2.0"}`)

		//Generate POST request to OpenEthereum
		req, err := http.NewRequest("POST", config.openethereumAPI,
			bytes.NewBuffer(jsonRequest))
		if err != nil {
			panic(err)
		}
		req.Header.Set("X-Custom-Header", "openethereum-healthcheck-client")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		//Perform POST request
		resp, err := client.Do(req)
		if err != nil {
			aSyncData.ethSyncResp.respBody = "failed"
			return (err)
		}
		defer resp.Body.Close()

		//Decode reponse into a string for ease of use
		body, err := ioutil.ReadAll(resp.Body)
		aSyncData.ethSyncResp.respBody = string(body)

		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result":false`) {
			return nil
		}
		if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
			err := fmt.Errorf("GET request failed")
			return err
		}

		//If OpenEthereum is currently syncing, parse the response
		err = json.Unmarshal(body, &aSyncData.ethSyncResp)
		if err != nil {
			return (err)
		}

		//Debug statement to print the response status, header, and body
		if config.printRequests == true {
			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			fmt.Println("response Body:", aSyncData.ethSyncResp.respBody)
		}

		return err
	}, config.pollingRate)
	return check
}

func netPeersCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//OpenEthereum API call to send
		var jsonRequest = []byte(`{"method":"parity_netPeers","params":[],"id":1,"jsonrpc":"2.0"}`)

		//Generate POST request to OpenEthereum
		req, err := http.NewRequest("POST", config.openethereumAPI,
			bytes.NewBuffer(jsonRequest))
		if err != nil {
			panic(err)
		}
		req.Header.Set("X-Custom-Header",
			"openethereum-healthcheck-client")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		//Perform POST request
		resp, err := client.Do(req)
		if err != nil {
			aSyncData.parityNetPeersResp.respBody = "failed"
			return (err)
		}
		defer resp.Body.Close()

		//Decode reponse into a string for ease of use
		body, err := ioutil.ReadAll(resp.Body)
		aSyncData.parityNetPeersResp.respBody = string(
			body)

		//Parse the response into a struct
		err = json.Unmarshal(body, &aSyncData.parityNetPeersResp)
		if err != nil {
			return (err)
		}

		//Debug statement to print the response status, header, and body
		if config.printRequests == true {
			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			fmt.Println("response Body:", aSyncData.parityNetPeersResp.respBody)
		}

		return err
	}, config.pollingRate)
	return check
}

func openEthereumPeerCount(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Check if GET request to OpenEthereum failed
		if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
			err := fmt.Errorf("GET request failed")
			return err
		}
		//Compare currently connected peers to config struct
		if aSyncData.parityNetPeersResp.Result.Connected < config.peerMinimum {
			err := fmt.Errorf("minimumPeers :%d",
				aSyncData.parityNetPeersResp.Result.Connected)
			return err
		}
		return nil
	}, config.pollingRate)
	return check
}

func openEthereumBlockUpdateCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Pause to allow request to be captured
		time.Sleep(2 * time.Second)
		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result":false`) {
			return nil
		}
		if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
			err := fmt.Errorf("GET request failed")
			return err
		}

		//Retrieve current block from response struct
		currentBlock := aSyncData.ethSyncResp.Result.CurrentBlock

		//Retrieve snapshot status
		warpBlock := aSyncData.ethSyncResp.Result.WarpChunksProcessed

		//Wait the blockUpdateTimeout
		time.Sleep(config.blockUpdateTimeout)

		//Check if the new block is equal to the old block
		if currentBlock == aSyncData.ethSyncResp.Result.CurrentBlock {
			if warpBlock == aSyncData.ethSyncResp.Result.WarpChunksProcessed {
				err := fmt.Errorf("currentBlock/warpBlock have not refreshed within timeout")
				return err
			}
		}
		return nil
	}, 2*config.blockUpdateTimeout)
	return check
}

func openEthereumBlockDifference(aSyncData *asyncDataStruct) (int64, error) {
	//Check if OpenEthereum is not currently syncing
	if strings.Contains(aSyncData.ethSyncResp.respBody, `"result":false`) {
		return 0, nil
	}
	if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
		err := fmt.Errorf("GET request failed")
		return 0, err
	}

	//Convert the hex string to an integer
	currentBlockInt, err := strconv.ParseInt(strings.Replace(
		aSyncData.ethSyncResp.Result.CurrentBlock, "0x", "", -1), 16, 64)
	if err != nil {
		return -1, err
	}

	//Convert the hex string to an integer
	highestBlockInt, err := strconv.ParseInt(strings.Replace(
		aSyncData.ethSyncResp.Result.HighestBlock, "0x", "", -1), 16, 64)
	if err != nil {
		return -1, err
	}

	//Calculate how many blocks behind OpenEthereum is
	blockDifference := highestBlockInt - currentBlockInt

	return blockDifference, nil
}

func openEthereumBlockSyncCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result":false`) {
			return nil
		}
		if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
			err := fmt.Errorf("GET request failed")
			return err
		}

		//Retrieve the current block difference between the estimated block and the current block
		currentBlockDifference, err := openEthereumBlockDifference(aSyncData)
		if err != nil {
			return err
		}

		//Compare to blockSyncDifference
		if currentBlockDifference > config.blockSyncDifference {
			err := fmt.Errorf("blockDifference :%d",
				currentBlockDifference)
			return err
		}

		return nil
	}, config.pollingRate)
	return check
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
func nodeReadinessChecks(health healthcheck.Handler, config *configStruct, httpMux *http.ServeMux, aSyncData *asyncDataStruct) {
	//Add healthchecks to the readiness check

	health.AddReadinessCheck(
		"primary-status",
		aSyncData.checkPrimary)

	health.AddReadinessCheck(
		"inbox-reader-status",
		aSyncData.inboxReaderStatus)

	//OpenEthereum healthchecks
	//Add healthchecks to the readiness check
	if config.openethereumAPI != "" {
		health.AddReadinessCheck(
			"openethereum-api-status",
			aSyncData.tcpDialCheck)
		health.AddReadinessCheck(
			"openethereum-sync-response-status",
			aSyncData.ethSyncCheck)

		health.AddReadinessCheck(
			"openethereum-netpeers-response-status",
			aSyncData.parityNetPeersCheck)

		health.AddReadinessCheck(
			"openethereum-sync-status",
			aSyncData.blockSyncCheck)

		health.AddReadinessCheck(
			"openethereum-peer-status",
			aSyncData.minimumPeersCheck)

		health.AddReadinessCheck(
			"openethereum-block-refresh-status",
			aSyncData.blockRefreshCheck)
	} else {
		health.AddReadinessCheck(
			"openethereum-status",
			aSyncData.checkOpenethereum)
	}

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)
}

func waitConfig(config *configStruct) {
	config.mu.Lock()
	defer config.mu.Unlock()
	for {
		fmt.Println("waiting" + config.openethereumHealthcheckRPC + config.healthcheckRPC + config.openethereumAPI)
		if config.healthcheckRPC != "" {
			if config.openethereumAPI != "" || config.openethereumHealthcheckRPC != "" {
				return
			}
		}
		config.mu.Unlock()

		time.Sleep(config.loopDelayTimer)
		config.mu.Lock()
	}
}

//Start the healthcheck for OpenEthereum
func startHealthCheck(config *configStruct, state *healthState) error {
	//Allocate storage for the aSync calls

	//Create the main healthcheck handler
	health := healthcheck.NewMetricsHandler(config.prometheusRegistry, "healthcheck")
	//Create an HTTP server mux to serve the endpoints
	httpMux := http.NewServeMux()

	waitConfig(config)

	//Schedule the async calls
	asyncUpstream := newAsyncUpstream(state, config)

	//Create an endpoint to serve the liveness check
	httpMux.HandleFunc("/live", health.LiveEndpoint)

	//Define which healthchecks to use for the readiness API and expose the readiness API
	nodeReadinessChecks(health, config, httpMux, asyncUpstream)

	//Create an endpoint to serve the prometheus endpoint
	httpMux.Handle("/metrics", promhttp.HandlerFor(
		config.prometheusRegistry,
		promhttp.HandlerOpts{},
	))

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
