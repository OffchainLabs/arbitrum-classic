package web3

import "net/http"

type Net struct {
}

func (net *Net) Version(r *http.Request, args *VersionArgs, reply *string) error {
	*reply = "123456789"
	return nil
}
