package nodehealth

import (
	"fmt"
	"math/big"
	"net/http"
	"nodehealth"
	"time"

	"github.com/heptiolabs/healthcheck"
)

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

func nodeHealthTest() error {
	//Generate sample servers for testing
	go startTestingServerFail()
	go startTestingServerPass()

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

	healthChan := make(chan nodehealth.Log, largeBufferSize)
	go nodehealth.NodeHealthCheck(healthChan)

	//Test startup configuration delay
	fmt.Println("Startup delay")
	time.Sleep(startUpSleepTime)
	res, err := http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		fmt.Println(passMessage)
	} else {
		fmt.Println(failMessage)
		return nil
	}
	fmt.Println("")

	//Primary aSync Test
	fmt.Println("Test Removing Primary aSync")
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:8089"}
	time.Sleep(timeDelayTests)

	//Test server response
	res, err = http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	}
	fmt.Println(passMessage)
	fmt.Println("")

	//Test failing OpenEthereum Node
	fmt.Println("Failing OpenEthereum Node Test")
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:8088"}
	time.Sleep(timeDelayTests)

	//Test server response
	res, err = http.Get(nodehealthAddress + "/ready")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	} else {
		fmt.Println(passMessage)
	}

	fmt.Println("")

	//Test adding primary after start
	fmt.Println("Adding Primary Late Test")
	healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:8089"}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:8089"}
	time.Sleep(timeDelayTests)

	//Test server response
	res, err = http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		return err
	}
	if res.StatusCode != successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	}
	healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:8088"}
	time.Sleep(timeDelayTests)

	//Test server response
	res, err = http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	} else {
		fmt.Println(passMessage)
	}
	fmt.Println("")

	//Test InboxReader block status check
	fmt.Println("Test InboxReader blockStatus")
	healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:8089"}
	time.Sleep(timeDelayTests)
	const largeBigInt = 20
	testBigInt := big.NewInt(largeBigInt)
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "currentHeight", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "caughtUpTarget", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "arbCorePosition", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "getNextBlockToRead", ValBigInt: *testBigInt}

	//Test server response
	res, err = http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != successfulStatus {
		fmt.Println(failMessage)
		//The server is returning an unexpected status code
		return nil
	}

	const smallBigInt = 10
	blockTest := big.NewInt(smallBigInt)
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "currentHeight", ValBigInt: *blockTest}
	healthChan <- nodehealth.Log{Comp: inboxReaderName, Var: "caughtUpTarget", ValBigInt: *testBigInt}
	time.Sleep(timeDelayTests)

	//Test server response
	res, err = http.Get(nodehealthAddress + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode == successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	} else {
		fmt.Println(passMessage)
	}
	return nil
}
