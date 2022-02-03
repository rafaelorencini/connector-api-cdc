package request_builders

import (
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"strings"
)

type MysqlSourceBuilder struct {
}

func NewMysqlSourceBuilder() domain.RequestBuilder {
	return new(MysqlSourceBuilder)
}

func (b *MysqlSourceBuilder) Build(connector *domain.Connector, mergedConfig map[string]string, tablesNames []string) map[string]string {
	tablesNamesJoined := strings.Join(tablesNames[:], ",")

	// TODO Extrair para constantes
	mergedConfig["table.include.list"] = tablesNamesJoined
	mergedConfig["database.server.name"] = connector.DatabaseName

	return mergedConfig
}
