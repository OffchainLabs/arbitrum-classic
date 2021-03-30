// +build ignore

package main

import (
	"io/ioutil"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
)

func main() {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		log.Fatal(err)
	}

	hash, err := mach.Hash()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("../../MACHINEHASH", []byte(hash.String()), 777)
	if err != nil {
		log.Fatal(err)
	}
}
