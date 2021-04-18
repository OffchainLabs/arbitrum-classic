package nodehealth

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/heptiolabs/healthcheck"
)

type configTestStruct struct {
	successfulStatus  int
	largeBufferSize   int
	readinessEndpoint string
	failMessage       string
	passMessage       string
	startUpSleepTime  time.Duration
	timeDelayTests    time.Duration
	nodehealthAddress string
	inboxReaderName   string
}

func (config *configTestStruct) newTestConfig() {
	//Configuration constants
	const successfulStatus = 200
	const largeBufferSize = 200
	const readinessEndpoint = "/ready"
	const failMessage = "Failed"
	const passMessage = "Passed"
	const startUpSleepTime = 5 * time.Second
	const timeDelayTests = 11 * time.Second
	const nodehealthAddress = "http://127.0.0.1:8080"
	const inboxReaderName = "InboxReader"

	config.successfulStatus = successfulStatus
	config.largeBufferSize = largeBufferSize
	config.readinessEndpoint = readinessEndpoint
	config.failMessage = failMessage
	config.passMessage = passMessage
	config.startUpSleepTime = startUpSleepTime
	config.timeDelayTests = timeDelayTests
	config.nodehealthAddress = nodehealthAddress
	config.inboxReaderName = inboxReaderName
}

func startTestingServerFail() {
	health := healthcheck.NewHandler()
	httpMux := http.NewServeMux()

	//Readiness check that always fails
	health.AddReadinessCheck("failing-check", func() error {
		return fmt.Errorf("example failure")
	})

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)

	http.ListenAndServe("0.0.0.0:8088", httpMux)
}

func startTestingServerPass() {
	health := healthcheck.NewHandler()
	httpMux := http.NewServeMux()

	//Readiness check that always fails
	health.AddReadinessCheck("pass-check", func() error {
		return nil
	})

	//Create an endpoint to serve the readiness check
	httpMux.HandleFunc("/ready", health.ReadyEndpoint)

	http.ListenAndServe("0.0.0.0:8089", httpMux)
}

func startUpTest(config *configTestStruct) error {
	fmt.Println("Startup delay")
	time.Sleep(config.startUpSleepTime)
	_, err := http.Get(config.nodehealthAddress + config.readinessEndpoint)
	if err != nil {
		fmt.Println(config.passMessage)
	} else {
		fmt.Println(config.failMessage)
		return errors.New("Failed startup delay test - exiting")
	}
	fmt.Println("")
	return nil
}

func aSyncTest(healthChan chan Log, config *configTestStruct) error {
	fmt.Println("Test Removing Primary aSync")
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:8092"}
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPCPort", ValStr: "8089"}
	healthChan <- Log{Comp: "InboxReader", Var: "loadingDatabase", ValBool: false}
	const smallBigInt = 10
	blockTest := big.NewInt(10)
	healthChan <- Log{Comp: config.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: config.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(blockTest)}
	time.Sleep(config.timeDelayTests)

	//Test server response
	res, err := http.Get(config.nodehealthAddress + config.readinessEndpoint + "?full=1")

	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != config.successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(config.failMessage)
		return errors.New("Failed test without primary - exiting")
	}
	fmt.Println(config.passMessage)
	fmt.Println("")
	return nil
}

func openEthereumFailure(healthChan chan Log, config *configTestStruct) error {
	fmt.Println("Failing OpenEthereum Node Test")
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPCPort", ValStr: "8088"}

	time.Sleep(config.timeDelayTests)
	//Test server response
	res, err := http.Get(config.nodehealthAddress + "/ready")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == config.successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(config.failMessage)
		return errors.New("Failed test of OpenEthereum failure - exiting")
	}
	fmt.Println(config.passMessage)
	fmt.Println("")
	return nil
}

func addPrimaryWhileRunning(healthChan chan Log, config *configTestStruct) error {
	fmt.Println("Adding Primary Late Test")
	healthChan <- Log{Config: true, Var: "openethereumHealthcheckRPCPort", ValStr: "8089"}
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:9010"}
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPCPort", ValStr: "8089"}

	time.Sleep(config.timeDelayTests)

	//Test server response
	res, err := http.Get(config.nodehealthAddress + config.readinessEndpoint)
	if err != nil {
		return err
	}
	if res.StatusCode != config.successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(config.failMessage)
		return errors.New("Failed adding primary while running test - exiting")
	}
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPCPort", ValStr: "8088"}

	time.Sleep(config.timeDelayTests)

	//Test server response
	res, err = http.Get(config.nodehealthAddress + config.readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == config.successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(config.failMessage)
		return errors.New("Failed adding primary while running test - exiting")
	}
	fmt.Println(config.passMessage)

	fmt.Println("")
	return nil
}

func inboxReaderTest(healthChan chan Log, config *configTestStruct) error {
	fmt.Println("Test InboxReader blockStatus")
	healthChan <- Log{Config: true, Var: "primaryHealthcheckRPCPort", ValStr: "8089"}
	time.Sleep(config.timeDelayTests)
	const largeBigInt = 20
	testBigInt := big.NewInt(largeBigInt)
	healthChan <- Log{Comp: config.inboxReaderName, Var: "currentHeight", ValBigInt: new(big.Int).Set(testBigInt)}
	healthChan <- Log{Comp: config.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(testBigInt)}
	healthChan <- Log{Comp: config.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(testBigInt)}
	healthChan <- Log{Comp: config.inboxReaderName, Var: "getNextBlockToRead", ValBigInt: new(big.Int).Set(testBigInt)}

	//Test server response
	res, err := http.Get(config.nodehealthAddress + config.readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != config.successfulStatus {
		fmt.Println(config.failMessage)
		//The server is returning an unexpected status code
		return errors.New("Failed inbox reader test - exiting")
	}

	const smallBigInt = 10
	blockTest := big.NewInt(smallBigInt)
	healthChan <- Log{Comp: config.inboxReaderName, Var: "arbCorePosition", ValBigInt: new(big.Int).Set(blockTest)}
	healthChan <- Log{Comp: config.inboxReaderName, Var: "caughtUpTarget", ValBigInt: new(big.Int).Set(testBigInt)}
	time.Sleep(config.timeDelayTests)

	//Test server response
	res, err = http.Get(config.nodehealthAddress + config.readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == config.successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(config.failMessage)
		return errors.New("Failed inbox reader test - exiting")
	}
	fmt.Println(config.passMessage)
	return nil
}

func TestNodeHealth(t *testing.T) {
	config := configTestStruct{}
	config.newTestConfig()

	//Generate sample servers for testing
	go startTestingServerFail()
	go startTestingServerPass()

	healthChan := make(chan Log, config.largeBufferSize)
	go NodeHealthCheck(healthChan)
	healthChan <- Log{Config: true, Var: "healcheckRPC", ValStr: "0.0.0.0:8080"}
	healthChan <- Log{Config: true, Var: "openEthereumInternalCheckEnable", ValStr: "https://eth-kovan.alchemyapi.io/v2/yvzMZUhX0jmdpRfqrUEGwh--U59mJNhf"}
	//Test startup configuration delay
	err := startUpTest(&config)
	if err != nil {
		t.Fatal(err)
	}

	//Primary aSync Test
	err = aSyncTest(healthChan, &config)
	if err != nil {
		t.Fatal(err)
	}

	//Test failing OpenEthereum Node
	err = openEthereumFailure(healthChan, &config)
	if err != nil {
		t.Fatal(err)
	}

	//Test adding primary after start
	err = addPrimaryWhileRunning(healthChan, &config)
	if err != nil {
		t.Fatal(err)
	}

	//Test InboxReader block status check
	err = inboxReaderTest(healthChan, &config)
	if err != nil {
		t.Fatal(err)
	}
}
