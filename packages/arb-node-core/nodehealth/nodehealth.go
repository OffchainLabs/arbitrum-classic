package nodehealth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
)

//Nodehealth configuration struct
type configStruct struct {
	//Mutex for config struct
	mu sync.Mutex

	//Initialize healthcheck execution when true
	init bool
	//Address to bind healthcheck to
	healthcheckRPC string
	//Disable the healthcheck metrics
	healthcheckMetrics bool
	//Disable checking the primary aggregator
	disablePrimaryCheck bool
	//Disable checking the OpenEthereum node
	disableOpenEthereumCheck bool

	//Map to dynamically allocate Prometheus Histograms
	prometheusHistograms map[string]*prometheus.HistogramVec
	//Prometheus registry to register histograms on
	prometheusRegistry *prometheus.Registry

	//Aggregator Healthcheck Config
	//Rate to poll the remote APIs at
	pollingRate time.Duration
	//Rate to recheck whether configuration variables are set
	loopDelayTimer time.Duration
	//Address the OpenEthereum healthcheck server is running on if internal OE check isn't used
	openethereumHealthcheckRPC string
	//Port the OpenEthereum healthcheck server is running on
	openethereumHealthcheckRPCPort string
	//Address the primary node healthcheck is running at if used
	primaryHealthcheckRPC string
	//Port the primary node healthcheck server is running on
	primaryHealthcheckRPCPort string
	//HTTP code the healthcheck is set to return to indicate success
	successCode int
	//Blocks between arbCorePosition and caughtUpTarget to consider acceptable
	blockDifferenceTolerance int64

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

//Struct for storing the state of a node's different components
type healthState struct {
	mu sync.Mutex
	//InboxReader state struct
	inboxReader inboxReaderState
}

//Struct for storing inboxReader's current state
type inboxReaderState struct {
	//Boolean to indicate whether we are waiting for the database to load
	loadingDatabase bool

	//InboxReader variables used to determine where we are in the sync process
	getNextBlockToRead *big.Int
	currentHeight      *big.Int
	arbCorePosition    *big.Int
	caughtUpTarget     *big.Int
}

//Struct for storing the asynchronous healthcheck calls
type asyncDataStruct struct {
	mu sync.Mutex

	//Map to dynamically allocate new healthchecks
	healthchecks map[string]healthcheck.Check

	//Response structs to process the OpenEthereum responses
	ethSyncResp        OpenEthereumResponse
	parityNetPeersResp OpenEthereumResponse
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
	//Different message types we could be sent
	Err     error
	Sev     string
	Var     string
	Comp    string
	Debug   bool
	Config  bool
	Metrics bool

	//Potential variable types a client could want to log to reduce casting
	ValStr    string
	ValInt    int64
	ValBigInt *big.Int
	ValBool   bool
	ValTime   time.Duration
}

// Default configuration values for the healthcheck server
func newConfig() *configStruct {
	config := configStruct{}
	const init = false

	//Global configuration
	const healthcheckRPC = ""
	const healthcheckMetrics = false
	const disablePrimaryCheck = false
	const disableOpenEthereumCheck = false

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

	//Load configuration into struct
	config.prometheusHistograms = make(map[string]*prometheus.HistogramVec)
	config.prometheusRegistry = prometheus.NewRegistry()

	config.init = init
	config.healthcheckRPC = healthcheckRPC
	config.healthcheckMetrics = healthcheckMetrics
	config.disablePrimaryCheck = disablePrimaryCheck
	config.disableOpenEthereumCheck = disableOpenEthereumCheck

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
	config.blockSyncDifference = blockSyncDifference
	config.peerMinimum = peerMinimum
	config.blockUpdateTimeout = blockUpdateTimeout
	config.printRequests = printRequests
	config.printConfigMsg = printConfigMsg

	return &config
}

//Initialize all upstream checks to run at a set time interval in an asynchronous manner
func newAsyncUpstream(state *healthState, config *configStruct) *asyncDataStruct {
	asyncData := asyncDataStruct{}
	//Allocate memory for healthcheck map
	asyncData.healthchecks = make(map[string]healthcheck.Check)
	if config.openethereumAPI != "" {
		//Request eth_syncing status from OpenEthereum
		asyncData.healthchecks["ethSyncCheck"] = ethSyncCheck(config, &asyncData)

		//Request eth_syncing status from OpenEthereum
		asyncData.healthchecks["parityNetPeersCheck"] = netPeersCheck(config, &asyncData)

		//Check OpenEthereum has more than peerMinimum peers currently connected to it
		asyncData.healthchecks["minimumPeersCheck"] = openEthereumPeerCount(config, &asyncData)

		//Check OpenEthereum is refreshing its currentBlock quicker than the blockUpdateTimeout
		asyncData.healthchecks["blockRefreshCheck"] = openEthereumBlockUpdateCheck(config, &asyncData)

		//Check if OpenEthereum is within blockSyncDifference from the estimated block
		asyncData.healthchecks["blockSyncCheck"] = openEthereumBlockSyncCheck(config, &asyncData)

		//Check if the OpenEthereum API is accepting pings
		asyncData.healthchecks["tcpDialCheck"] = openEthereumTCPDialCheck(config, &asyncData)

	} else {
		//Check the healthcheck endpoint for OpenEthereum
		asyncData.healthchecks["checkOpenethereum"] = checkEndpoint(config, &config.openethereumHealthcheckRPC, &config.openethereumHealthcheckRPCPort)
	}

	//Check the primary endpoint
	asyncData.healthchecks["checkPrimary"] = checkEndpoint(config, &config.primaryHealthcheckRPC, &config.primaryHealthcheckRPCPort)

	//Check how many blocks the inboxReader is behind
	asyncData.healthchecks["inboxReaderStatus"] = checkInboxReader(config, state)

	return &asyncData
}

//Initialize health state storage
func newHealthState() *healthState {
	state := healthState{}

	state.inboxReader.loadingDatabase = true
	state.inboxReader.currentHeight = new(big.Int)
	state.inboxReader.caughtUpTarget = new(big.Int)
	state.inboxReader.arbCorePosition = new(big.Int)
	state.inboxReader.getNextBlockToRead = new(big.Int)

	return &state
}

//Async logger to dequeue messages from channel buffer and load them into the state structs
func logger(ctx context.Context, config *configStruct, state *healthState, logMsgChan <-chan Log) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

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

func Init(healthChan chan Log) {
	healthChan <- Log{Config: true, Var: "init"}
}

//Send a function's runtime over the health channel
func LogTime(comp string, function string, start time.Time, healthChan chan Log) {
	healthChan <- Log{Metrics: true, Comp: comp, Var: function, ValTime: time.Since(start)}
}

//Create a histogram for a functions runtime and add it to the registry
func metricsHandler(config *configStruct, logMessage Log) {
	//Check if the prometheus histogram already exists
	_, ok := config.prometheusHistograms[logMessage.Comp]
	if !ok {
		//Initialize a histogram for the function if it doesn't exist
		config.prometheusHistograms[logMessage.Comp] = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "kovan4-0",
			Name:      logMessage.Comp,
			Help:      logMessage.Comp + " latency distributions.",
			Buckets:   prometheus.ExponentialBuckets(1, 10, 4),
		}, []string{logMessage.Comp})

		//Register the histogram with the prometheus registry
		config.prometheusRegistry.MustRegister(config.prometheusHistograms[logMessage.Comp])
	}

	//Observe the function's runtime using the function name as the label
	config.prometheusHistograms[logMessage.Comp].WithLabelValues(logMessage.Var).Observe(float64(logMessage.ValTime.Milliseconds()))
}

//Update the inboxReader state struct using a value from the health channel
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

//Update the configurations truct using a value from the health channel
func updateConfig(config *configStruct, logMessage Log) {
	config.mu.Lock()
	defer config.mu.Unlock()
	if logMessage.Var == "init" {
		config.init = true
	}
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
	if logMessage.Var == "healthcheckRPC" {
		config.healthcheckRPC = logMessage.ValStr
	}
	if logMessage.Var == "primaryHealthcheckRPCPort" {
		config.primaryHealthcheckRPCPort = logMessage.ValStr
	}
	if logMessage.Var == "blockSyncDifference" {
		config.blockSyncDifference = logMessage.ValInt
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
	if logMessage.Var == "healthcheckMetrics" {
		config.healthcheckMetrics = logMessage.ValBool
	}
	if logMessage.Var == "disablePrimaryCheck" {
		config.disablePrimaryCheck = logMessage.ValBool
	}
	if logMessage.Var == "disableOpenEthereumCheck" {
		config.disableOpenEthereumCheck = logMessage.ValBool
	}
}

//Resolve the IP of the OpenEthereum node and check if it can be dialed
func openEthereumTCPDialCheck(config *configStruct, asyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Parse the URL to extract the hostname and port
		u, err := url.Parse(config.openethereumAPI)
		if err != nil {
			return err
		}

		//Lookup the IP address of the hostname
		ipAddr, err := net.LookupIP(u.Hostname())
		if err != nil {
			return err
		}

		//Extract the port from the URL
		port := u.Port()

		//Default to port 80 if no port is provided
		if port == "" {
			port = "80"
		}

		//Set a timeout on the TCP dialer
		d := net.Dialer{Timeout: config.requestTimeout}

		//Dial the IP address
		conn, err := d.Dial("tcp", ipAddr[0].String()+":"+port)
		if err != nil {
			return err
		}

		//Close the connection before returning
		conn.Close()
		return nil
	}, config.pollingRate)
	return check
}

//Send the call defined in jsonRequest to the OpenEthereum server
func openEthereumCall(config *configStruct, jsonRequest []byte) ([]byte, error) {
	//Generate POST request to OpenEthereum
	req, err := http.NewRequest("POST", config.openethereumAPI,
		bytes.NewBuffer(jsonRequest))
	if err != nil {
		panic(err)
	}

	//Set request headers to identify the healthcheck
	req.Header.Set("X-Custom-Header", "openethereum-healthcheck-client")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	//Perform POST request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//Close the connection after the response
	defer resp.Body.Close()

	//Decode reponse into a string for ease of use
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

//Request the eth_syncing status from the OpenEthereum server and parse it into a state array
func ethSyncCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//OpenEthereum API call to send
		var jsonRequest = []byte(`{"method":"eth_syncing","params":[],"id":1,"jsonrpc":"2.0"}`)

		//Send the call to OpenEthereum
		respBody, err := openEthereumCall(config, jsonRequest)

		aSyncData.mu.Lock()
		defer aSyncData.mu.Unlock()

		if err != nil {
			aSyncData.ethSyncResp.respBody = "failed"
			return (err)
		}

		//Convert the response from a byte slice to a string
		aSyncData.ethSyncResp.respBody = string(respBody)

		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result": false`) {
			return nil
		}
		if strings.Contains(aSyncData.ethSyncResp.respBody, `failed`) {
			err := fmt.Errorf("GET request failed")
			return err
		}

		//If OpenEthereum is currently syncing, parse the response
		err = json.Unmarshal(respBody, &aSyncData.ethSyncResp)
		if err != nil {
			return (err)
		}

		return err
	}, config.pollingRate)
	return check
}

//Request the netPeers status from the OpenEthereum server and parse it into a state array
func netPeersCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//OpenEthereum API call to send
		var jsonRequest = []byte(`{"method":"parity_netPeers","params":[],"id":1,"jsonrpc":"2.0"}`)

		//Send the call to OpenEthereum
		respBody, err := openEthereumCall(config, jsonRequest)

		aSyncData.mu.Lock()
		defer aSyncData.mu.Unlock()

		if err != nil {
			aSyncData.ethSyncResp.respBody = "failed"
			return (err)
		}

		//Convert the response from a byte slice to a string
		aSyncData.parityNetPeersResp.respBody = string(respBody)

		//Check if the netPeers response is unsupported by the OpenEthereum node
		if strings.Contains(aSyncData.parityNetPeersResp.respBody, "Unsupported method") {
			aSyncData.ethSyncResp.respBody = "failed"
			return nil
		}

		//Parse the response into a struct
		err = json.Unmarshal(respBody, &aSyncData.parityNetPeersResp)
		if err != nil {
			return (err)
		}

		return err
	}, config.pollingRate)
	return check
}

//Check that OpenEthereum has more than minimumPeers currently connected to it
func openEthereumPeerCount(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		aSyncData.mu.Lock()
		defer aSyncData.mu.Unlock()
		//Check if the netPeers response is unsupported by the OpenEthereum node
		if strings.Contains(aSyncData.parityNetPeersResp.respBody, "Unsupported method") {
			return nil
		}

		//Check if GET request to OpenEthereum failed
		if strings.Contains(aSyncData.parityNetPeersResp.respBody, `failed`) {
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

//Check that the block OpenEthereum is on is updating at a rate faster then config.blockUpdateTimeout
func openEthereumBlockUpdateCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		aSyncData.mu.Lock()
		defer aSyncData.mu.Unlock()
		//Pause to allow request to be captured
		time.Sleep(2 * time.Second)

		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result": false`) {
			return nil
		}

		//Check if the call to OpenEthereum failed
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

//Helper function for openEthereumBlockUpdateCheck to calculate the block difference in OpenEthereum's response
func openEthereumBlockDifference(aSyncData *asyncDataStruct) (int64, error) {
	aSyncData.mu.Lock()
	defer aSyncData.mu.Unlock()
	//Check if OpenEthereum is not currently syncing
	if strings.Contains(aSyncData.ethSyncResp.respBody, `"result":false`) {
		return 0, nil
	}

	//Check if the GET request to OpenEthereum failed
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

//Check the current OpenEthereum block versus the expected block it should be at
func openEthereumBlockSyncCheck(config *configStruct, aSyncData *asyncDataStruct) healthcheck.Check {
	check := healthcheck.Async(func() error {
		aSyncData.mu.Lock()
		defer aSyncData.mu.Unlock()
		//Check if OpenEthereum is not currently syncing
		if strings.Contains(aSyncData.ethSyncResp.respBody, `"result": false`) {
			return nil
		}

		//Check if the GET request failed
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

//Asynchronously check a healthcheck endpoint to determine its status
func checkEndpoint(config *configStruct, endpoint *string, port *string) healthcheck.Check {
	check := healthcheck.Async(func() error {
		//Lock config mutex for read operation
		config.mu.Lock()

		//Copy the endpoint and port the mutex can be released
		endpointStr := *endpoint
		portStr := *port

		//Release the config mutex
		config.mu.Unlock()

		//Check if the health endpoint has been configured
		if endpointStr == "" {
			return nil
		}

		//Retrieve status code from healthcheck endpoint
		res, err := http.Get("http://" + endpointStr + ":" + portStr + "/ready")

		//Check the response code to determine if OpenEthereum is ready
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

//Check whether the InboxReader's arbCorePosition is caught up to the target within a tolerance
func checkInboxReader(config *configStruct, state *healthState) healthcheck.Check {
	check := healthcheck.Async(func() error {
		state.mu.Lock()
		defer state.mu.Unlock()

		//Check if the database is still loading
		if state.inboxReader.loadingDatabase == true {
			return errors.New("Loading database snapshot")
		}

		//Calculate out the block difference
		blockDifference := new(big.Int).Sub(state.inboxReader.caughtUpTarget, state.inboxReader.arbCorePosition)

		//Patch bug where caughtUpTarget exceeds arbCorePosition due to a primary failure
		if blockDifference.Sign() == -1 {
			return nil
		}

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
func nodeReadinessChecks(health healthcheck.Handler, config *configStruct, httpMux *http.ServeMux, asyncData *asyncDataStruct) {
	//Add healthchecks to the readiness check

	//Add primary healthcheck if it is not disabled
	if !config.disablePrimaryCheck {
		health.AddReadinessCheck(
			"primary-status",
			asyncData.healthchecks["checkPrimary"])
	}

	health.AddReadinessCheck(
		"inbox-reader-status",
		asyncData.healthchecks["inboxReaderStatus"])

	//OpenEthereum healthchecks
	//Add healthchecks to the readiness check if they are not disabled
	if !config.disableOpenEthereumCheck {
		if config.openethereumAPI != "" {
			health.AddReadinessCheck(
				"openethereum-api-status",
				asyncData.healthchecks["tcpDialCheck"])
			health.AddReadinessCheck(
				"openethereum-sync-response-status",
				asyncData.healthchecks["ethSyncCheck"])

			health.AddReadinessCheck(
				"openethereum-netpeers-response-status",
				asyncData.healthchecks["parityNetPeersCheck"])

			health.AddReadinessCheck(
				"openethereum-sync-status",
				asyncData.healthchecks["blockSyncCheck"])

			health.AddReadinessCheck(
				"openethereum-peer-status",
				asyncData.healthchecks["minimumPeersCheck"])

			health.AddReadinessCheck(
				"openethereum-block-refresh-status",
				asyncData.healthchecks["blockRefreshCheck"])
		} else {
			health.AddReadinessCheck(
				"openethereum-status",
				asyncData.healthchecks["checkOpenethereum"])
		}
	}

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)
}

//Wait for critical configuration variables to be loaded before continuing
func waitConfig(config *configStruct) {
	config.mu.Lock()
	defer config.mu.Unlock()

	//Loop while the configuration variables are not set
	for {
		if config.init {
			if config.disableOpenEthereumCheck || config.healthcheckRPC == "" ||
				config.openethereumAPI != "" || config.openethereumHealthcheckRPC != "" {
				return
			}
		}
		//Prevent the lock from being held over the sleep
		config.mu.Unlock()

		//Sleep loopDelayTimer duration to reduce load
		time.Sleep(config.loopDelayTimer)

		//Lock the configuration for the read operation
		config.mu.Lock()
	}
}

//Start the healthcheck
func startHealthCheck(ctx context.Context, config *configStruct, state *healthState) error {
	//Create the main healthcheck handler
	health := healthcheck.NewMetricsHandler(config.prometheusRegistry, "healthcheck")

	//Create an HTTP server mux to serve the endpoints
	httpMux := http.NewServeMux()

	//Wait for the configuration to be loaded
	waitConfig(config)

	//Exit if the healthcheck is disabled while leaving the logger running to prevent blocking calls
	if config.healthcheckRPC == "" {
		<-ctx.Done()
		return nil
	}

	//Schedule the async calls
	asyncUpstream := newAsyncUpstream(state, config)

	//Create an endpoint to serve the liveness check
	httpMux.HandleFunc("/live", health.LiveEndpoint)

	//Define which healthchecks to use for the readiness API and expose the readiness API
	nodeReadinessChecks(health, config, httpMux, asyncUpstream)

	//Create an endpoint to serve the prometheus endpoint
	if config.healthcheckMetrics {
		httpMux.Handle("/metrics", promhttp.HandlerFor(
			config.prometheusRegistry,
			promhttp.HandlerOpts{},
		))
	}

	//Create the HTTP server
	httpServer := &http.Server{
		Addr:        config.healthcheckRPC,
		Handler:     httpMux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	//Start serving requests
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return
		}
	}()

	//Gracefully shut down the server if the context is Done
	<-ctx.Done()

	gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefulCtx); err != nil {
		return err
	}

	return nil
}

// NodeHealthCheck Create a node healthcheck that listens on the given channel
func StartNodeHealthCheck(ctx context.Context, logMsgChan <-chan Log) error {
	//Create the configuration struct
	state := newHealthState()

	//Load the default configuration
	config := newConfig()

	//Start the channel logger
	go logger(ctx, config, state, logMsgChan)

	//Start the node healthcheck
	err := startHealthCheck(ctx, config, state)

	return err
}
