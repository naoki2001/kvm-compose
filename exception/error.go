package exception

import (
	"fmt"
)

// Error is view error message
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
		fmt.Print("error: Unknown options: " + status + "\n")
	case 10:
		fmt.Print("error: Command execution failed \n")
	case 20:
		fmt.Print("error: " + status + " not found\n")
	case 21:
		fmt.Print("error: " + status + " is already running\n")
	case 22:
		fmt.Print("error: " + status + " is already stopping\n")
	case 30:
		fmt.Print("error: Failed to read kvm-compose.yml\n")
	default:
		fmt.Print("error: Unexpected error \n")
	}
	return
}
