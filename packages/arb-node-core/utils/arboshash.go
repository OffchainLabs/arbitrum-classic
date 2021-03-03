// +build ignore

package main

import (
	"io/ioutil"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
)

func main() {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("../../MACHINEHASH", []byte(mach.Hash().String()), 777)
	if err != nil {
		log.Fatal(err)
	}
}
