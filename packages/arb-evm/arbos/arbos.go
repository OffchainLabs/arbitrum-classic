package arbos

import (
	"path/filepath"
	"runtime"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Stack().Str("component", "arbos").Logger()

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")
var ARB_ADDRESS_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000066")
var ARB_BLS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000067")
var ARB_FUNCTION_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000068")
var ARB_OWNER_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006B")
var ARB_GAS_INFO_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006C")
var ARB_AGGREGATOR_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006D")
var ARB_RETRYABLE_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006E")

var ARB_NODE_INTERFACE_ADDRESS = ethcommon.HexToAddress("0x00000000000000000000000000000000000000C8")

func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "arbos.mexe"), nil
}

func Dir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get arbos path")
	}

	return filepath.Join(filepath.Dir(filename), "../../arb-os/arb_os"), nil
}
