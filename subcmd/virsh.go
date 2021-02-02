package subcmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"../exception"
)

// VMList executes virsh list [option]
func VMList(option string) {
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

// VMOperation is doing VM operation: virsh [cont] [name]
func VMOperation(name string, cont string) {
	status := SearchVM(name)
	if status == "NotFound" {
		exception.Error(20, name)
		return
	} else if status == "active" && cont == "start" {
		exception.Error(21, name)
		return
	} else if status == "inactive" && ( cont == "shutdown" || cont =="destroy" ) {
		exception.Error(22, name)
		return
	} else {
		exception.Error(10, "none")
		return
	}
	
	fmt.Print(cont + "VM \n")
	
	arg := []string{
		cont,
		name,
	}
	
	exec.Command("virsh", arg...).Start()
	for {
		status := SearchVM(name)
		if (status == "active" && cont == "start") || (status == "inactive" && ( cont == "shutdown" || cont =="destroy" )) {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}

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
