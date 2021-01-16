package subcmd

import (
	"fmt"
)

// Version view version of kvm-compose
func Version() {
	fmt.Print("kvm-compose version 1.0.0alpha\n")
	return
}
