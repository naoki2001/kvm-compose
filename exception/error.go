package exception

import (
	"fmt"
)

// Error is return error status
func Error(code int, status string) {
	switch code {
	case 1:
		fmt.Print("error: Need subcommand \n")
	case 2:
		fmt.Print("error: Unknown subcomannd: " + status + "\n")
	case 3:
		fmt.Print("error: Too few options \n")
	case 4:
		fmt.Print("error: Too many options \n")
	case 5:
		fmt.Print("error: Unknown options \n")
	case 10:
		fmt.Print("error: Command execution failed \n")
	default:
		fmt.Print("error: Unexpected error \n")
	}
	return
}
