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
		exception.Error(10, "none")
		return
	}

	arg := []string{
		"start",
		name,
	}

	exec.Command("virsh", arg...).Start()
	for {
		status := SearchVM(name)
		if status == "active" {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}

// Shutdown executes virsh shutdown [name]
func Shutdown(name string) {
	status := SearchVM(name)
	if status == "NotFound" {
		exception.Error(20, name)
		return
	} else if status == "active" {
	} else if status == "inactive" {
		exception.Error(22, name)
		return
	} else {
		exception.Error(10, "none")
		return
	}

	arg := []string{
		"shutdown",
		name,
	}

	exec.Command("virsh", arg...).Start()
	for {
		status := SearchVM(name)
		if status == "inactive" {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}

// Destroy executes virsh destroy [name]
func Destroy(name string) {
	status := SearchVM(name)
	if status == "NotFound" {
		exception.Error(20, name)
		return
	} else if status == "active" {
	} else if status == "inactive" {
		exception.Error(22, name)
		return
	} else {
		exception.Error(10, "none")
		return
	}

	arg := []string{
		"Destroy",
		name,
	}

	exec.Command("virsh", arg...).Start()
	for {
		status := SearchVM(name)
		if status == "inactive" {
			fmt.Print("success\n")
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
	cmd := "virsh list --all | grep " + name
	err := exec.Command("sh", "-c", cmd).Run()

	if err != nil {
		status = "NotFound"
	} else {
		cmd = "virsh list | grep " + name
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			status = "inactive"
		} else {
			status = "active"
		}
	}
	return status
}
