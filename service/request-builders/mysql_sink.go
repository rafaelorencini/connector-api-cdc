package request_builders

import "github.com/rafaelorencini/connector-api-cdc/domain"

type MysqlSinkBuilder struct {
}

func NewMysqlSinkBuilder() domain.RequestBuilder {
	return new(MysqlSinkBuilder)
}

func (b *MysqlSinkBuilder) Build(connector *domain.Connector, mergedConfig map[string]string, tablesNames []string) map[string]string {
	//topics_regex
	map_ := map[string]string{"a": "a"}
	return map_
}
