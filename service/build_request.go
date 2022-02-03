package service

import (
	"encoding/json"
	"github.com/rafaelorencini/connector-api-cdc/domain"
)

type BuildRequestService struct {
	ReadYaml     domain.ReadYamlInterface
	GetSecrets   domain.GetSecretsInterface
	MapFunctions map[string]domain.RequestBuilder
}

func NewBuildRequestService(readYaml domain.ReadYamlInterface, getSecrets domain.GetSecretsInterface, mapFunctions map[string]domain.RequestBuilder) domain.BuildRequestInterface {
	return &BuildRequestService{
		ReadYaml:     readYaml,
		GetSecrets:   getSecrets,
		MapFunctions: mapFunctions,
	}
}

func (br *BuildRequestService) Build(connector *domain.Connector) (payload string, err error) {
	mergedConfig := br.mountConfigs(connector)
	payloadMap := br.mountPayloadMap(connector, mergedConfig)
	payload = br.generatePayloadString(payloadMap)
	return payload, nil
}

func (br *BuildRequestService) mountPayloadMap(connector *domain.Connector, mergedConfig map[string]string) map[string]interface{} {
	payloadMap := make(map[string]interface{})
	payloadMap["name"] = connector.ConnectorName
	payloadMap["config"] = mergedConfig
	return payloadMap
}

func (br *BuildRequestService) mountConfigs(connector *domain.Connector) map[string]string {
	secrets, _ := br.GetSecrets.Get()
	defaultGroup := connector.DatabaseType + "." + connector.ConnectorType
	defaultConfig, _ := br.ReadYaml.Read(defaultGroup)

	mergedConfig := br.mergeMaps(secrets, defaultConfig)
	tablesNames := br.getTablesNames(connector)
	key := connector.DatabaseType + "_" + connector.ConnectorType
	payloadMap := br.MapFunctions[key].Build(connector, mergedConfig, tablesNames)
	return payloadMap
}

func (br *BuildRequestService) generatePayloadString(payloadMap map[string]interface{}) string {
	payloadByte, _ := json.Marshal(payloadMap)
	return string(payloadByte)
}

//TODO Usar um map function (ou algo assim) para extrair os nomes
func (br *BuildRequestService) getTablesNames(connector *domain.Connector) []string {
	strList := make([]string, 0)
	for _, table := range connector.Tables {
		tableName := connector.DatabaseName + "." + table.Table
		strList = append(strList, tableName)
	}
	return strList
}

func (br *BuildRequestService) mergeMaps(secrets map[string]string, defaultConfig map[string]string) map[string]string {
	for k, v := range secrets {
		defaultConfig[k] = v
	}

	return defaultConfig
}
