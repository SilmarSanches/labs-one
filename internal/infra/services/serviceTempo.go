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
	"net/url"
	"time"
)

type ServiceTempoInterface interface {
	GetTempo(ctx context.Context, cidade string) (entities.TempoDto, error)
}

type ServiceTempo struct {
	HttpClient HttpClient
	appConfig  *config.AppSettings
}

func NewServiceTempo(httpClient HttpClient, appConfig *config.AppSettings) *ServiceTempo {
	return &ServiceTempo{
		HttpClient: httpClient,
		appConfig:  appConfig,
	}
}

func (s *ServiceTempo) GetTempo(ctx context.Context, cidade string) (entities.TempoDto, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cidadeEncoded := url.QueryEscape(cidade)
	url := s.appConfig.UrlTempo + "/current.json?q=" + cidadeEncoded + "&key=" + s.appConfig.TempoApiKey

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entities.TempoDto{}, err
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return entities.TempoDto{}, fmt.Errorf("timeout de 5s excedido ao consultar o serviço ViaCep: %v", err)
		}
	}

	if res == nil {
		return entities.TempoDto{}, errors.New("resposta nula ao consultar o weatherapi")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("erro ao fechar o corpo da resposta weatherapi: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return entities.TempoDto{}, fmt.Errorf("erro ao consultar o serviço weatherapi: %v", res.StatusCode)
	}

	var data entities.TempoDto
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return entities.TempoDto{}, err
	}

	return data, nil
}
