package service

import (
	"connector-api-cdc/domain"
)

type ConnectorService struct {
	connectorRepository domain.ConnectorRepository
}

func NewConnectorService(r domain.ConnectorRepository) domain.ConnectorServices {
	return &ConnectorService{
		connectorRepository: r,
	}
}

//create connector
func (s *ConnectorService) CreateConnector(connector *domain.Connector) error {
	err := s.connectorRepository.CreateConnector(connector)
	if err != nil {
		return err
	}
	return nil
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
