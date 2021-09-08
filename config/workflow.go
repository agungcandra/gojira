package config

import (
	"io/ioutil"

	"github.com/agungcandra/gojira/entity"
	"gopkg.in/yaml.v3"
)

func LoadWorklow() entity.Workflow {
	workflow := make(entity.Workflow)
	workflowByte, err := ioutil.ReadFile("workflow.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(workflowByte, &workflow)
	if err != nil {
		panic(err)
	}

	return workflow
}
