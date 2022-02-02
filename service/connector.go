package service

import (
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"net/http"
)

type ConnectorService struct {
	connectorRepository domain.ConnectorRepository
	buildRequest        domain.BuildRequestInterface
	kafkaConnect        domain.KafkaConnectInterface
}

func NewConnectorService(r domain.ConnectorRepository) domain.ConnectorServices {
	return &ConnectorService{
		connectorRepository: r,
	}
}

//create connector
func (c *ConnectorService) CreateConnector(connector *domain.Connector) error {
	err := c.connectorRepository.CreateConnector(connector)
	if err != nil {
		return err
	}

	c.registrateConnector(connector, err)

	return nil
}

func (c *ConnectorService) registrateConnector(connector *domain.Connector, err error) {
	payload, err := c.buildRequest.Build(connector)
	c.kafkaConnect.SendRequest(http.MethodPost, payload)
}

//get connectors
func (s *ConnectorService) GetConnectors(connector *[]domain.Connector) error {
	err := s.connectorRepository.GetConnectors(connector)
	if err != nil {
		return err
	}
	return nil
}

//get connector by id
func (s *ConnectorService) GetConnector(connector *domain.Connector, id string) error {
	err := s.connectorRepository.GetConnector(connector, id)
	if err != nil {
		return err
	}
	return nil
}

// update connector
func (s *ConnectorService) UpdateConnector(connector *domain.Connector) error {
	err := s.connectorRepository.UpdateConnector(connector)
	if err != nil {
		return err
	}
	return nil
}

// delete connector
func (s *ConnectorService) DeleteConnector(connector *domain.Connector, id string) error {
	err := s.connectorRepository.DeleteConnector(connector, id)
	if err != nil {
		return err
	}
	return nil
}
