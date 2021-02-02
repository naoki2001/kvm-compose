package subcmd

import (
	"fmt"
	"runtime"
)

// Version view version of kvm-compose
func Version() {
	fmt.Print("kvm-compose version: 1.0.0 alpha\n")
	fmt.Println("go version: ", runtime.Version())
	return
}
