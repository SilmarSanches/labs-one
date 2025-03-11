//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"labs-one/config"
	"labs-one/internal/infra/services"
	"labs-one/internal/infra/web"
	"labs-one/internal/usecases"

	"github.com/google/wire"
)

var ProviderConfig = wire.NewSet(config.ProvideConfig)

var ProviderCep = wire.NewSet(
	services.NewHttpClient,
	services.NewServiceCep,
	wire.Bind(new(services.HttpClient), new(*http.Client)),
	wire.Bind(new(services.ServiceCepInterface), new(*services.ServiceCep)),
)

var ProviderTempo = wire.NewSet(
	services.NewServiceTempo,
	wire.Bind(new(services.ServiceTempoInterface), new(*services.ServiceTempo)),
)

var ProviderUseCase = wire.NewSet(
	usecases.NewGetTempoUseCase,
	wire.Bind(new(usecases.GetTempoUseCaseInterface), new(*usecases.GetTempoUseCase)),
)

var ProviderHandler = wire.NewSet(web.NewGetTempoHandler)

func NewConfig() *config.AppSettings {
	wire.Build(ProviderConfig)
	return &config.AppSettings{}
}

func NewGetTempUseCase() *usecases.GetTempoUseCase {
	wire.Build(ProviderConfig, ProviderCep, ProviderTempo, ProviderUseCase)
	return &usecases.GetTempoUseCase{}
}

func NewGetTempoHandler() *web.GetTempoHandler {
	wire.Build(ProviderConfig, ProviderCep, ProviderTempo, ProviderUseCase, ProviderHandler)
	return &web.GetTempoHandler{}
}
