/*
 * Copyright 2019, Offchain Labs, Inc.
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

package main

func main() {
	//
	// cMac := vm.CreateVM(os.Args[1], os.Args[2])
	// cSteps := vm.RunVM(cMac, 1000000)
	// fmt.Println("cMachine ended ", cSteps, " steps run.")
	//
	// candidatesCountBytes, _ := hexutil.Decode("0x2d35a8a2")
	// candidatesCount, _ := evm.BytesToSizedByteArray(candidatesCountBytes)
	//
	// candidates1Bytes, _ := hexutil.Decode("0x3477ee2e0000000000000000000000000000000000000000000000000000000000000001")
	// candidates1, _ := evm.BytesToSizedByteArray(candidates1Bytes)
	//
	// candidates2Bytes, _ := hexutil.Decode("0x3477ee2e0000000000000000000000000000000000000000000000000000000000000002")
	// candidates2, _ := evm.BytesToSizedByteArray(candidates2Bytes)
	//
	////var machine *vm.Machine
	// machine, err := loader.LoadMachineFromFile(os.Args[1], true)
	// if err != nil {
	//	log.Fatal("Loader Error: ", err)
	//}
	//
	// balanceTracker := protocol.NewBalanceTracker()
	// inbox := protocol.NewEmptyInbox()
	// rawAddress, _ := new(big.Int).SetString("784030224795475933405737832577560929931042096197", 10)
	// addressVal := value.NewIntValue(rawAddress)
	// inbox.SendMessage(protocol.NewMessage(value.NewTuple2(addressVal, candidatesCount), [21]byte{}, big.NewInt(0), protocol.Identity{}))
	// inbox.SendMessage(protocol.NewMessage(value.NewTuple2(addressVal, candidates1), [21]byte{}, big.NewInt(0), protocol.Identity{}))
	// inbox.SendMessage(protocol.NewMessage(value.NewTuple2(addressVal, candidates2), [21]byte{}, big.NewInt(0), protocol.Identity{}))
	// inbox.DeliverMessages()
	//
	// ctx := protocol.NewMachineAssertionContext(machine, balanceTracker, [2]uint64{0, 10000}, inbox.Receive())
	// steps := machine.Run(800000)
	// fmt.Println(steps)
	// ad := ctx.Finalize(machine)
	// assertion := ad.GetAssertion()
	// for _, val := range assertion.Logs {
	//	log, err := evm.ProcessLog(val)
	//	if err != nil {
	//		fmt.Printf("Logged1 invalid %v: %v\n", err, val)
	//	} else {
	//		fmt.Printf("Logged1 %v\n", log)
	//	}
	//
	//}
	// precondition := ad.GetAssertion().Stub().GeneratePostcondition(ad.GetPrecondition())
	//
	// ctx = protocol.NewMachineAssertionContext(machine, precondition.BeforeBalance, precondition.TimeBounds, inbox.Receive())
	// steps = machine.Run(110000)
	// fmt.Println(steps)
	// ad = ctx.Finalize(machine)
	// assertion = ad.GetAssertion()
	// for _, val := range assertion.Logs {
	//	log, err := evm.ProcessLog(val)
	//	if err != nil {
	//		fmt.Printf("Logged1 invalid %v: %v\n", err, val)
	//	} else {
	//		fmt.Printf("Logged1 %v\n", log)
	//	}
	//}
}
