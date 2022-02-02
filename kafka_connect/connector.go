package kafka_connect

import (
	"connector-api-cdc/domain"
	"net/http"
)

type ConnectorKafkaConnect struct {
	Connect *KafkaConnect
}

func NewKafkaConnectRequests(connect *KafkaConnect) domain.KafkaConnectRequests {
	return &ConnectorKafkaConnect{
		Connect: connect,
	}
}

func (r *ConnectorKafkaConnect) CreateConnectorKafkaConnect(payload *string) (err error) {
	client := r.Connect.httpClient()
	*payload = `
{
  "name": "inventory-connector",  
  "config": {  
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "tasks.max": "1",  
    "database.hostname": "mysql",  
    "database.port": "3306",
    "database.user": "debezium",
    "database.password": "dbz",
    "database.server.id": "184054",  
    "database.server.name": "dbserver1",  
    "database.include.list": "inventory",  
    "database.history.kafka.bootstrap.servers": "kafka:9092",  
    "database.history.kafka.topic": "schema-changes.inventory"  
  }
}`
	r.Connect.sendRequest(client, http.MethodPost, *payload)
	return
}
