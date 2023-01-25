package tests

import (
	"fmt"
	"os"
)

func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
		os.Exit(0)
	}
}
