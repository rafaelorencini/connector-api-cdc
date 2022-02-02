package service

import "github.com/rafaelorencini/connector-api-cdc/domain"

type GetSecretsService struct {
}

func NewGetSecretsServiceService() domain.GetSecretsInterface {
	return new(GetSecretsService)
}

func (r *GetSecretsService) Get() (map[string]string, error) {
	m := map[string]string{"mock": "mock"}
	return m, nil
}
