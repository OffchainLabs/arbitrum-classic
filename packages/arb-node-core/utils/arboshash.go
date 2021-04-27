// +build ignore

package main

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("error generating MACHINEHASH file")
	}
}

func run() error {
	arbosPath, err := arbos.Path()
	if err != nil {
		return err
	}

	mach, err := cmachine.New(arbosPath)
	if err != nil {
		return err
	}

	hash, err := mach.Hash()
	if err != nil {
		return err
	}

	return ioutil.WriteFile("../../MACHINEHASH", []byte(hash.String()), 777)
}
