package common

import "strings"

func IsFatalError(err error) bool {
	if strings.Contains(err.Error(), "arbcore thread aborted") ||
		strings.Contains(err.Error(), "aborting inbox reader thread") {
		return true
	}

	return false
}
