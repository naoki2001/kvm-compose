package process

import (
	"fmt"
	"io/ioutil"

	"../exception"
)

// ReadYAML is read yaml and encode
func ReadYAML() (yaml map[interface{}]interface{}) {
	buf, err := ioutil.ReadFile("kvm-compose.yml")
	if err != nil {
		exception.Error(30, "faild read file")
		return nil
	}

	yaml = make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &yaml)
	if err != nil {
		panic(err)
	}

	fmt.Print(buf)
	return yaml
}
