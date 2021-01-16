package subcmd

import (
	"fmt"

	"../process"
)

// Build is virt-install [options...]
func Build() {
	yml := process.ReadYAML()

	fmt.Print(yml)

	return
}
