// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"labs-one/config"
	"labs-one/internal/infra/services"
	"labs-one/internal/infra/web"
	"labs-one/internal/usecases"
)

import (
	_ "labs-one/docs"
)

// Injectors from wire.go:

func NewConfig() *config.AppSettings {
	appSettings := config.ProvideConfig()
	return appSettings
}

func NewGetTempUseCase() *usecases.GetTempoUseCase {
	appSettings := config.ProvideConfig()
	httpClient := services.NewHttpClient()
	serviceCep := services.NewServiceCep(httpClient, appSettings)
	serviceTempo := services.NewServiceTempo(httpClient, appSettings)
	getTempoUseCase := usecases.NewGetTempoUseCase(appSettings, serviceCep, serviceTempo)
	return getTempoUseCase
}

func NewGetTempoHandler() *web.GetTempoHandler {
	appSettings := config.ProvideConfig()
	httpClient := services.NewHttpClient()
	serviceCep := services.NewServiceCep(httpClient, appSettings)
	serviceTempo := services.NewServiceTempo(httpClient, appSettings)
	getTempoUseCase := usecases.NewGetTempoUseCase(appSettings, serviceCep, serviceTempo)
	getTempoHandler := web.NewGetTempoHandler(appSettings, getTempoUseCase, serviceCep, serviceTempo)
	return getTempoHandler
}

// wire.go:

var ProviderConfig = wire.NewSet(config.ProvideConfig)

var ProviderHttpClient = wire.NewSet(services.NewHttpClient)

var ProviderCep = wire.NewSet(services.NewServiceCep, wire.Bind(new(services.ServiceCepInterface), new(*services.ServiceCep)))

var ProviderTempo = wire.NewSet(services.NewServiceTempo, wire.Bind(new(services.ServiceTempoInterface), new(*services.ServiceTempo)))

var ProviderGlobal = wire.NewSet(
	ProviderHttpClient,
	ProviderConfig,
	ProviderCep,
	ProviderTempo,
)

var ProviderUseCase = wire.NewSet(usecases.NewGetTempoUseCase, wire.Bind(new(usecases.GetTempoUseCaseInterface), new(*usecases.GetTempoUseCase)))

var ProviderHandler = wire.NewSet(web.NewGetTempoHandler)
