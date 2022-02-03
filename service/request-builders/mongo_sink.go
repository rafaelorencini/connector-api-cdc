package request_builders

import "github.com/rafaelorencini/connector-api-cdc/domain"

type MongoSinkBuilder struct {
}

func NewMongoSinkBuilder() domain.RequestBuilder {
	return new(MongoSinkBuilder)
}

func (b *MongoSinkBuilder) Build(connector *domain.Connector, mergedConfig map[string]string, tablesNames []string) map[string]string {
	map_ := map[string]string{"a": "a"}
	return map_
}
