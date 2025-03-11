package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"labs-one/config"
	"labs-one/internal/entities"

	"net/http"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ServiceCepInterface interface {
	GetCep(ctx context.Context, cep string) (entities.ViaCepDto, error)
}

type ServiceCep struct {
	HttpClient HttpClient
	appConfig  *config.AppSettings
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func NewServiceCep(httpClient HttpClient, appConfig *config.AppSettings) *ServiceCep {
	return &ServiceCep{
		HttpClient: httpClient,
		appConfig:  appConfig,
	}
}

func (s *ServiceCep) GetCep(ctx context.Context, cep string) (entities.ViaCepDto, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := s.appConfig.UrlCep + "/" + cep + "/json"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entities.ViaCepDto{}, err
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return entities.ViaCepDto{}, fmt.Errorf("timeout de 5s excedido ao consultar o serviço ViaCep: %v", err)
		}
	}

	if res == nil {
		return entities.ViaCepDto{}, errors.New("resposta nula ao consultar o viacep")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("erro ao fechar o corpo da resposta ViaCep: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return entities.ViaCepDto{}, fmt.Errorf("erro ao consultar o serviço ViaCep: %d", res.StatusCode)
	}

	var data entities.ViaCepDto
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return entities.ViaCepDto{}, fmt.Errorf("erro ao decodificar resposta JSON do ViaCep: %w", err)
	}

	return data, nil
}
