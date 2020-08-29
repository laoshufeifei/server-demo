package utils

import (
	"log"
	"syscall"
)

// IncreaseResourcesLimit set ulimit -n (1024 * 1024)
func IncreaseResourcesLimit() {
	log.Println("from dal_linux.go")

	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("rLimit.Cur is %v, rLimit.Max is %v\n", rLimit.Cur, rLimit.Max)
	if rLimit.Cur != rLimit.Max {
		rLimit.Cur = rLimit.Max
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}
	}
}
