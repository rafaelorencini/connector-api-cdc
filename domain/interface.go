package domain

import "net/http"

type ConnectorServices interface {
	CreateConnector(connector *Connector) error
	GetConnectors(connector *[]Connector) error
	GetConnector(connector *Connector, id string) error
	UpdateConnector(connector *Connector) error
	DeleteConnector(connector *Connector, id string) error
}

type ConnectorRepository interface {
	Reader
	Writer
}

type Reader interface {
	GetConnectors(Connector *[]Connector) error
	GetConnector(Connector *Connector, id string) error
}

type Writer interface {
	//Create(c *fiber.Ctx) error
	CreateConnector(Connector *Connector) error
	UpdateConnector(Connector *Connector) error
	DeleteConnector(Connector *Connector, id string) error
}

type ReadYamlInterface interface {
	Read() (map[string]string, error)
}

type BuildRequestInterface interface {
	Build(connector *Connector) (string, error)
}

type GetSecretsInterface interface {
	Get() (map[string]string, error)
}

type KafkaConnectInterface interface {
	SendRequest(method string, payload string) *http.Response
}
