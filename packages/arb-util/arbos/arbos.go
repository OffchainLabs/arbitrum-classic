package arbos

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"
)

func Path() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal(errors.New("failed to get arbos path"))
	}

	return filepath.Join(filepath.Dir(filename), "../../../arbos.mexe")
}
