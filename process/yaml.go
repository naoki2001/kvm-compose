package process

import (
	"io/ioutil"

	"../exception"
	"gopkg.in/yaml.v2"
)

// ReadYAML is read yaml and encode
func ReadYAML() (yml map[interface{}]interface{}) {
	buf, err := ioutil.ReadFile("kvm-compose.yml")
	if err != nil {
		exception.Error(30, "faild read file")
		return nil
	}

	yml = make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &yml)
	if err != nil {
		panic(err)
	}
	return yml
}
