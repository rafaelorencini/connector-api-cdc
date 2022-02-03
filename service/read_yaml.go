package service

import "github.com/rafaelorencini/connector-api-cdc/domain"

type ReadYamlService struct {
}

func NewReadYamlService() domain.ReadYamlInterface {
	return new(ReadYamlService)
}

func (r *ReadYamlService) Read() (map[string]string, error) {
	m := map[string]string{
		"connector.class":                          "io.debezium.connector.mysql.MySqlConnector",
		"tasks.max":                                "1",
		"database.server.id":                       "184054",
		"database.server.name":                     "dbserver1",
		"database.history.kafka.bootstrap.servers": "kafka:9092",
		"database.history.kafka.topic":             "schema-changes.my_db",
	}

	return m, nil
}
