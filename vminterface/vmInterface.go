package vminterface

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
	"github.com/offchainlabs/arb-avm/loader"
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-util/protocol"
	"log"
	"runtime"
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
	c unsafe.Pointer
	g *vm.Machine
}

func CCreateVM(codeFile string, inboxFile string) *VM {
	fmt.Println("CCreartVM codeFile=", codeFile)
	var ret VM
	retptr := &ret
	//****************
	// C stuff
	cFilename := C.CString(codeFile)
	cInboxFilename := C.CString(inboxFile)

	cMachine := C.machine_create(cFilename, cInboxFilename)
	retptr.c = cMachine
	runtime.SetFinalizer(retptr, cdestroyVM)
	C.free(unsafe.Pointer(cFilename))
	C.free(unsafe.Pointer(cInboxFilename))
	return retptr
}

func cdestroyVM(cMachine *VM) {
	fmt.Println("Calling C.machine_destroy")
	C.machine_destroy(cMachine.c)
}

func CCloneVM(cMachine *VM) *VM {
	var ret VM
	retptr := &ret
	cMach := C.machine_clone(cMachine.c)
	retptr.c = cMach
	return retptr
}

//func RunVM(cMachine unsafe.Pointer, steps int, timebounds protocol.TimeBounds) int {
func CRunVM(cMachine *VM, steps uint64) uint64 {
	fmt.Println("Starting cMachine")
	//cStart := time.Now()
	//            machine_run(void *m, uint64_t maxSteps);
	cSteps := C.machine_run(cMachine.c, C.ulonglong(steps))
	//cEnd := time.Now()
	//cSteps := 0
	fmt.Println("cMachine ended ", cSteps, " steps run.")
	// C stuff
	//*************
	return uint64(cSteps)
}

func CreateVM(codeFile string) *VM {
	var ret VM
	retptr := &ret
	machine, err := loader.LoadMachineFromFile(codeFile, true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}
	retptr.g = machine

	return retptr
}

func CreateVMwithMessages(codeFile string, inboxFile string) *VM {
	var ret VM
	retptr := &ret
	machine, err := loader.LoadMachineFromFile(codeFile, true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}
	retptr.g = machine

	return retptr
}

func destroyVM(machine *VM) {
	//C.machine_destroy(machine.g)
}

//func RunVM(cMachine unsafe.Pointer, steps int, timebounds protocol.TimeBounds) int {
func RunVM(machine *VM, steps uint64) uint64 {
	fmt.Println("Starting machine")
	//cStart := time.Now()
	//            machine_run(void *m, uint64_t maxSteps);
	//cSteps := C.machine_run(cMachine.m, C.ulonglong(steps))
	//cEnd := time.Now()
	//cSteps := 0
	//fmt.Println("cMachine ended ", cSteps, " steps run.")
	// C stuff
	//*************
	//return uint64(cSteps)
	return 0
}

func SendMessageToVM(machine *VM, msg protocol.Message) {
	//machine.g.
}
