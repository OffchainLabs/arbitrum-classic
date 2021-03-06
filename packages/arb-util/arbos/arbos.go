package arbos

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"runtime"
)

var logger = log.With().Caller().Str("component", "arbos").Logger()

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")
var ARB_ADDRESS_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000066")
var ARB_BLS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000067")
var ARB_FUNCTION_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000068")
var ARB_OWNER_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006B")

var ARB_NODE_INTERFACE_ADDRESS = ethcommon.HexToAddress("0x00000000000000000000000000000000000000C8")

func Path() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logger.Fatal().Msg("Failed to get arbos path")
	}

	return filepath.Join(filepath.Dir(filename), "../../arb-os/arb_os/arbos.mexe")
}
