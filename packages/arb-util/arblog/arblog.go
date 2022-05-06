package arblog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

var Logger zerolog.Logger

func init() {
	isPretty := false
	for _, arg := range os.Args {
		flag := strings.TrimLeft(arg, "-")

		if flag == "prettyprint" {
			isPretty = true
		}
	}
	if isPretty {
		Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		Logger = log.Logger.With().Caller().Stack().Logger()
	}
}
