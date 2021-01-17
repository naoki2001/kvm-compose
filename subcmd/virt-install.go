package subcmd

import (
	"fmt"
	"strconv"

	"../process"
)

// Build is virt-install [options...]
func Build() {
	yml, err := process.ReadYAML()
	if err == 30 {
		return
	}
	count := process.CountVM(yml)
	error := process.CheckYAML(yml, count)
	if error != 0 {
		return
	}

	i := 1
	var arg []string
	for i <= count {
		arg = []string{
			"--name", yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["name"].(string),
			"--vcpus", strconv.Itoa(yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["vcpus"].(int)),
			"--memory", strconv.Itoa(yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["memory"].(int)),
		}
		fmt.Print(arg)
		i++
	}

	fmt.Print(count)

	return
}
