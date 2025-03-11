package usecases

import (
	"context"
	"labs-one/config"
	"labs-one/internal/entities"
	"labs-one/internal/infra/services"
	"regexp"
)

type GetTempoUseCaseInterface interface {
	GetTempo(cep string) (entities.GetTempoResponseDto, error)
}

type GetTempoUseCase struct {
	appConfid               *config.AppSettings
	ViaCepServiceInterface  services.ServiceCepInterface
	WeatherServiceInterface services.ServiceTempoInterface
}

func NewGetTempoUseCase(appConfig *config.AppSettings, viaCepService services.ServiceCepInterface, weatherService services.ServiceTempoInterface) *GetTempoUseCase {
	return &GetTempoUseCase{
		appConfid:               appConfig,
		ViaCepServiceInterface:  viaCepService,
		WeatherServiceInterface: weatherService,
	}
}

func (u *GetTempoUseCase) GetTempo(cep string) (entities.GetTempoResponseDto, error) {
	ctx := context.Background()

	isValidCep := ValidateCEP(cep)
	if !isValidCep {
		return entities.GetTempoResponseDto{}, &entities.CustomError{
			Code:    422,
			Message: "invalid zipcode",
		}
	}

	cepResponse, err := u.ViaCepServiceInterface.GetCep(ctx, cep)
	if err != nil {
		return entities.GetTempoResponseDto{}, &entities.CustomError{
			Code:    404,
			Message: "can not find zipcode",
		}
	}

	weather, err := u.WeatherServiceInterface.GetTempo(ctx, cepResponse.Localidade)
	if err != nil {
		return entities.GetTempoResponseDto{}, &entities.CustomError{
			Code:    404,
			Message: "can not find temperature",
		}
	}

	celcius := weather.Current.TempC
	Kelvin := celcius + 273
	Fahrenheit := celcius*1.8 + 32

	result := entities.GetTempoResponseDto{
		Kelvin:     Kelvin,
		Celsius:    celcius,
		Fahrenheit: Fahrenheit,
	}

	return result, nil
}

func ValidateCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-\d{3}$`)
	return re.MatchString(cep)
}
