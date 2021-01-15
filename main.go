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
			}
		case "list":
			option := "default"
			if len(os.Args) == 3 {
				option = os.Args[2]
			} else if len(os.Args) >= 4 {
				exception.Error(4, "meny")
				return
			}
			subcmd.List(option)
		case "start":
		case "shutdown":

		case "destroy":

		case "version":

		default:
			exception.Error(2, os.Args[2])
		}
	}
	return
}
