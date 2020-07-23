package web3

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"net/http"
)

type Net struct {
	chainId uint64
}

func (net *Net) Version(r *http.Request, args *EmptyArgs, reply *string) error {
	*reply = hexutil.EncodeUint64(net.chainId)
	return nil
}
