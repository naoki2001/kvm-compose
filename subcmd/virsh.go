package subcmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"../exception"
)

// List executes virsh list [option]
func List(option string) {
	arg := []string{
		"list",
		option,
	}
	if option == "default" {
		arg = []string{
			"list",
		}
	}

	cmd := exec.Command("virsh", arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		exception.Error(10, "execution")
	}
	return
}

// Start executes virsh start [name]
func Start(name string) {
	cmd := "virsh list | grep " + name
	status, _ := exec.Command("sh", "-c", cmd).Output()
	if string(status) != "" {
		fmt.Print("error: " + name + " is allready running \n")
		return
	}

	arg := []string{
		"start",
		name,
	}

	exec.Command("virsh", arg...).Start()
	for {
		status, _ = exec.Command("sh", "-c", cmd).Output()
		if string(status) != "" {
			fmt.Print("start \n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}

	return
}
