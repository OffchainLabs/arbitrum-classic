package arbos

import (
	"path/filepath"
	"runtime"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")
var ARB_ADDRESS_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000066")
var ARB_BLS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000067")
var ARB_FUNCTION_TABLE_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000068")
var ARB_TEST_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000069")
var ARB_OWNER_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006B")
var ARB_GAS_INFO_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006C")
var ARB_AGGREGATOR_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006D")
var ARB_RETRYABLE_ADDRESS = ethcommon.HexToAddress("0x000000000000000000000000000000000000006E")

var ARB_NODE_INTERFACE_ADDRESS = ethcommon.HexToAddress("0x00000000000000000000000000000000000000C8")

func Path(before bool) (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	if before {
		return filepath.Join(dir, "arbos_before.mexe"), nil
	} else {
		return filepath.Join(dir, "arbos.mexe"), nil
	}
}

func Dir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get arbos path")
	}

	return filepath.Join(filepath.Dir(filename), "../../arb-os/arb_os"), nil
}
