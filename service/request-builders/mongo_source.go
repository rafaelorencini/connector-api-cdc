package request_builders

import "github.com/rafaelorencini/connector-api-cdc/domain"

type MongoSourceBuilder struct {
}

func NewMongoSourceBuilder() domain.RequestBuilder {
	return new(MongoSourceBuilder)
}

func (b *MongoSourceBuilder) Build(connector *domain.Connector, mergedConfig map[string]string, tablesNames []string) map[string]string {
	map_ := map[string]string{"a": "a"}
	return map_
}
