package utils

import (
	"flag"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type RollupArgs struct {
	ValidatorFolder string
	EthURL          string
	Address         common.Address
}

func ParseRollupCommand(fs *flag.FlagSet, startIndex int) RollupArgs {
	validatorFolder := fs.Arg(startIndex)
	ethURL := fs.Arg(startIndex + 1)
	addressString := fs.Arg(startIndex + 2)

	address := common.HexToAddress(addressString)

	return RollupArgs{
		ValidatorFolder: validatorFolder,
		EthURL:          ethURL,
		Address:         address,
	}
}

const RollupArgsString = "<validator_folder> <ethURL> <rollup_address>"
