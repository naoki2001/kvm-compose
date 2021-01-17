package process

import (
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"

	"../exception"
)

// ReadYAML is read yaml and encode
func ReadYAML() (yml map[interface{}]interface{}, error int) {
	buf, err := ioutil.ReadFile("kvm-compose.yml")
	if err != nil {
		exception.Error(30, "faild read file")
		error = 30
		return
	}

	yml = make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &yml)
	if err != nil {
		panic(err)
	}
	return
}

// CountVM is count vm? in kvm-compose.yml
func CountVM(yml map[interface{}]interface{}) (count int) {
	count = 1
	for {
		switch yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(count)].(type) {
		case nil:
			return count - 1
		default:
			count++
		}
	}
}

// CheckYAML check if the required options exist
func CheckYAML(yml map[interface{}]interface{}, count int) (error int) {
	i := 1
	for i <= count {
		if yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["name"] == nil {
			exception.Error(35, "name")
			error = error + 1
		}
		if yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["vcpus"] == nil {
			exception.Error(35, "vpcus")
			error = error + 1
		}
		if yml["vm"].(map[interface{}]interface{})["vm"+strconv.Itoa(i)].(map[interface{}]interface{})["memory"] == nil {
			exception.Error(35, "memory")
			error = error + 1
		}
		i++
	}

	return
}
