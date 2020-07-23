package web3

import (
	"net/http"
	"strconv"
)

type Net struct {
	chainId uint64
}

func (net *Net) Version(r *http.Request, args *EmptyArgs, reply *string) error {
	*reply = strconv.FormatUint(net.chainId, 10)
	return nil
}
