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

	successfulStatus := 200
	largeBufferSize := 200
	readinessEndpoint := "/ready"
	failMessage := "Failed"
	passMessage := "Passed"

	healthChan := make(chan nodehealth.Log, largeBufferSize)
	go nodehealth.NodeHealthCheck(healthChan)

	//Test startup configuration delay
	fmt.Println("Startup delay")
	time.Sleep(5 * time.Second)
	res, err := http.Get("http://127.0.0.1:8080" + readinessEndpoint)
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
	time.Sleep(11 * time.Second)

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	} else {
		fmt.Println(passMessage)
	}
	fmt.Println("")

	//Test failing OpenEthereum Node
	fmt.Println("Failing OpenEthereum Node Test")
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: "http://127.0.0.1:8088"}
	time.Sleep(11 * time.Second)

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + "/ready")
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
	time.Sleep(11 * time.Second)

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + readinessEndpoint)
	if err != nil {
		return err
	}
	if res.StatusCode != successfulStatus {
		//The server is returning an unexpected status code
		fmt.Println(failMessage)
		return nil
	}
	healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: "http://127.0.0.1:8088"}
	time.Sleep(11 * time.Second)

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + readinessEndpoint)
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
	time.Sleep(11 * time.Second)
	testBigInt := big.NewInt(20)
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "currentHeight", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "caughtUpTarget", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "arbCorePosition", ValBigInt: *testBigInt}
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "getNextBlockToRead", ValBigInt: *testBigInt}

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + readinessEndpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode != successfulStatus {
		fmt.Println(failMessage)
		//The server is returning an unexpected status code
		return nil
	}

	blockTest := big.NewInt(10)
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "currentHeight", ValBigInt: *blockTest}
	healthChan <- nodehealth.Log{Comp: "InboxReader", Var: "caughtUpTarget", ValBigInt: *testBigInt}
	time.Sleep(11 * time.Second)

	//Test server response
	res, err = http.Get("http://127.0.0.1:8080" + readinessEndpoint)
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

func main() {
	err := nodeHealthTest()
	fmt.Println(err)
}
