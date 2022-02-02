package service

import "github.com/rafaelorencini/connector-api-cdc/domain"

type ReadYamlService struct {
}

func NewReadYamlService() domain.ReadYamlInterface {
	return new(ReadYamlService)
}

func (r *ReadYamlService) Read() (map[string]string, error) {
	m := map[string]string{"mock": "mock"}
	return m, nil

	//{
	//	"name": "inventory-connector",
	//	"config": {
	//	"connector.class": "io.debezium.connector.mysql.MySqlConnector",
	//		"tasks.max": "1",
	//		"database.hostname": "mysql",
	//		"database.port": "3306",
	//		"database.user": "debezium",
	//		"database.password": "dbz",
	//		"database.server.id": "184054",
	//		"database.server.name": "dbserver1",
	//		"database.include.list": "inventory",
	//		"database.history.kafka.bootstrap.servers": "kafka:9092",
	//		"database.history.kafka.topic": "schema-changes.inventory"
	//}
	//}`
}
