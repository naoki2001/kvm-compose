package main

import (
	"os"

	"./exception"
	"./subcmd"
)

func main() {
	if len(os.Args) <= 1 {
		exception.Error(1, "need")
	} else {
		switch os.Args[1] {
		case "build":
			if len(os.Args) == 2 {
				subcmd.Build()
			} else {
				exception.Error(4, "meny")
			}
		case "list":
			option := "default"
			if len(os.Args) == 3 {
				if os.Args[2] == "--all" || os.Args[2] == "--inactive" {
					option = os.Args[2]
				} else {
					exception.Error(5, os.Args[2])
					return
				}
			} else if len(os.Args) >= 4 {
				exception.Error(4, "meny")
				return
			}
			subcmd.List(option)
		case "start":
			if len(os.Args) <= 2 {
				exception.Error(3, "few")
			} else if len(os.Args) == 3 {
				subcmd.VMOperation(os.Args[2], os.Args[1])
			} else if len(os.Args) >= 4 {
				exception.Error(4, "meny")
			}
		case "shutdown":
			if len(os.Args) <= 2 {
				exception.Error(3, "few")
			} else if len(os.Args) == 3 {
				subcmd.VMOperation(os.Args[2], os.Args[1])
			} else if len(os.Args) >= 4 {
				exception.Error(4, "meny")
			}
		case "destroy":
			if len(os.Args) <= 2 {
				exception.Error(3, "few")
			} else if len(os.Args) == 3 {
				subcmd.VMOperation(os.Args[2], os.Args[1])
			} else if len(os.Args) >= 4 {
				exception.Error(4, "meny")
			}
		case "version":
			if len(os.Args) == 2 {
				subcmd.Version()
			} else {
				exception.Error(4, "meny")
			}
		default:
			exception.Error(2, os.Args[1])
		}
	}
	return
}
