package common

import "strings"

func IsFatalError(err error) bool {
	return strings.Contains(err.Error(), "arbcore thread aborted")
}
