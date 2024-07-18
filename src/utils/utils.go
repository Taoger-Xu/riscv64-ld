package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

const (
	Red   = "\033[0;31m"
	Reset = "\033[0m"
)

func Fatal(v any) {
	fmt.Printf("%srvld: fatal:%s %v\n", Red, Reset, v)
	debug.PrintStack()
	os.Exit(1)
}

func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}
