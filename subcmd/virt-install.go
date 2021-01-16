package subcmd

import (
	"../process"
)

// Build is virt-install [options...]
func Build() {
	process.ReadYAML()
	return
}
