package arbos

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"runtime"
)

var logger = log.With().Str("component", "arbos").Logger()

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")

func Path() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logger.Fatal().Msg("Failed to get arbos path")
	}

	return filepath.Join(filepath.Dir(filename), "../../../arbos.mexe")
}
