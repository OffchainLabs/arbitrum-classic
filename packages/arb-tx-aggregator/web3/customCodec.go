package web3

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gorilla/rpc/v2"
)

// UpCodec creates a CodecRequest to process each request.
type UpCodec struct {
}

// NewUpCodec returns a new UpCodec.
func NewUpCodec() *UpCodec {
	return &UpCodec{}
}

// NewRequest returns a new CodecRequest of type UpCodecRequest.
func (c *UpCodec) NewRequest(r *http.Request) rpc.CodecRequest {
	outerCR := &UpCodecRequest{}
	jsonC := NewCodec()
	innerCR := jsonC.NewRequest(r)

	outerCR.CodecRequest = innerCR.(*CodecRequest)
	return outerCR
}

type UpCodecRequest struct {
	*CodecRequest
}

func (c *UpCodecRequest) Method() (string, error) {
	m, err := c.CodecRequest.Method()
	if len(m) > 1 && err == nil {
		parts := strings.Split(m, "_")
		if len(parts) != 2 {
			return "", errors.New("malformed request")
		}
		service, method := parts[0], parts[1]

		r, n := utf8.DecodeRuneInString(service) // get the first rune, and it's length
		if unicode.IsLower(r) {
			service = string(unicode.ToUpper(r)) + service[n:]
		}
		r, n = utf8.DecodeRuneInString(method)
		if unicode.IsLower(r) {
			method = string(unicode.ToUpper(r)) + method[n:]
		}
		modifiedRequest := service + "." + method
		log.Println("Made request", modifiedRequest)
		return modifiedRequest, nil
	}
	return m, err
}
