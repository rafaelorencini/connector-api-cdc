package service

import "github.com/rafaelorencini/connector-api-cdc/domain"

type GetSecretsService struct {
}

func NewGetSecretsServiceService() domain.GetSecretsInterface {
	return new(GetSecretsService)
}

func (r *GetSecretsService) Get() (map[string]string, error) {

	m := map[string]string{
		"database.user":     "root",
		"database.port":     "3306",
		"database.hostname": "db_mysql",
		"database.password": "root",
	}

	return m, nil
}
