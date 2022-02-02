package service

import (
	"fmt"
	"github.com/rafaelorencini/connector-api-cdc/domain"
)

type BuildRequestService struct {
	ReadYaml   domain.ReadYamlInterface
	GetSecrets domain.GetSecretsInterface
}

func NewBuildRequestService(readYaml domain.ReadYamlInterface, getSecrets domain.GetSecretsInterface) domain.BuildRequestInterface {
	return &BuildRequestService{
		ReadYaml:   readYaml,
		GetSecrets: getSecrets,
	}
}

func (br *BuildRequestService) Build(connector *domain.Connector) (payload string, err error) {
	var paylod string
	secrets, err := br.GetSecrets.Get()
	if err != nil {
		return
	}

	default_config, err := br.ReadYaml.Read()
	if err != nil {
		return
	}

	fmt.Printf("", secrets)
	fmt.Printf("", default_config)

	paylod = "mock"

	return paylod, nil

}
