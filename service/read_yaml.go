package service

import (
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ReadYamlService struct {
}

func NewReadYamlService() domain.ReadYamlInterface {
	return new(ReadYamlService)
}

func (r *ReadYamlService) Read(defaultPropertiesGroup string) (map[string]string, error) {
	pathConfigs := "configs/default_configs.yaml"
	buf, _ := ioutil.ReadFile(pathConfigs)
	configsMap := make(map[string]map[string]string)
	yaml.Unmarshal(buf, configsMap)
	groupConfigs := configsMap[defaultPropertiesGroup]
	return groupConfigs, nil
}
