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
	status := SearchVM(name)
	if status == "NotFound" {
		exception.Error(20, name)
		return
	} else if status == "active" {
		exception.Error(21, name)
		return
	} else if status == "inactive" {
	} else {
		exception.Error(-1, "none")
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

//internal fuction

// SearchVM is serch VM by virsh command
func SearchVM(name string) (status string) {
	status := ""
	cmd := "virsh list --all | grep " + name
	exist, err := exec.Command("sh", "-c", cmd).Output()

	if exist == "" {
		status = "NotFound"
	} else {
		cmd = "virsh list | grep " + name
		exist, err = exec.Command("sh", "-c", cmd).Output()
		if exist == "" {
			status = "inactive"
		} else {
			status = "active"
		}
	}
	return status
}
