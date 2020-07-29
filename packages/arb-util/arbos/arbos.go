package arbos

import (
	"errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"log"
	"path/filepath"
	"runtime"
)

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")

func Path() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal(errors.New("failed to get arbos path"))
	}

	return filepath.Join(filepath.Dir(filename), "../../../arbos.mexe")
}
