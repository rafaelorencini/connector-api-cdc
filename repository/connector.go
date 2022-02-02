package repository

import (
	"gorm.io/gorm"
)

type ConnectorRepository struct {
	DB *gorm.DB
}

func NewConnectorRepository(db *gorm.DB) domain.ConnectorRepository {
	return &ConnectorRepository{
		DB: db,
	}
}

//create a connector
func (r *ConnectorRepository) CreateConnector(Connector *domain.Connector) (err error) {
	err = r.DB.Create(Connector).Error
	if err != nil {
		return err
	}
	return
}

//get connectors
func (r *ConnectorRepository) GetConnectors(Connector *[]domain.Connector) (err error) {
	err = r.DB.Preload("Tables").Find(Connector).Error
	if err != nil {
		return err
	}
	return
}

//get connector by id
func (r *ConnectorRepository) GetConnector(Connector *domain.Connector, id string) (err error) {
	err = r.DB.Preload("Tables").Where("ID = ?", id).First(Connector).Error
	if err != nil {
		return err
	}
	return
}

//update connector
func (r *ConnectorRepository) UpdateConnector(Connector *domain.Connector) (err error) {
	err = r.DB.Save(Connector).Error
	if err != nil {
		return err
	}
	return
}

//delete user
func (r *ConnectorRepository) DeleteConnector(Connector *domain.Connector, id string) (err error) {
	err = r.DB.Where("id = ?", id).Delete(Connector).Error
	if err != nil {
		return err
	}
	return
}
