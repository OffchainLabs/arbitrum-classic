package vm

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lavm -lstdc++
#include <cmachine.h>
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/offchainlabs/arb-util/evm"
	//"github.com/offchainlabs/arb-avm/loader"
	//"github.com/offchainlabs/arb-util/protocol"
	//"github.com/offchainlabs/arb-util/value"
	//"log"
	//"math/big"
	//"os"
)

type VM struct {
	m unsafe.Pointer
}

func CreateVM(codeFile string, inboxFile string) unsafe.Pointer {

	//****************
	// C stuff
	cFilename := C.CString(codeFile)
	cInboxFilename := C.CString(inboxFile)

	cMachine := C.machine_create(cFilename, cInboxFilename)

	C.free(unsafe.Pointer(cFilename))
	C.free(unsafe.Pointer(cInboxFilename))
	return cMachine
}

//func RunVM(cMachine unsafe.Pointer, steps int, timebounds protocol.TimeBounds) int {
func RunVM(cMachine unsafe.Pointer, steps uint64) uint64 {
	fmt.Println("Starting cMachine")
	//cStart := time.Now()
	//            machine_run(void *m, uint64_t maxSteps);
	cSteps := C.machine_run(cMachine, C.ulonglong(steps))
	//cEnd := time.Now()
	//cSteps := 0
	fmt.Println("cMachine ended ", cSteps, " steps run.")
	// C stuff
	//*************
	return uint64(cSteps)
}
