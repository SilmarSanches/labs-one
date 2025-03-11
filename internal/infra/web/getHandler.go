package web

import (
	"encoding/json"
	"labs-one/config"
	"labs-one/internal/entities"
	"labs-one/internal/infra/services"
	"labs-one/internal/usecases"
	"net/http"
)

type GetTempoHandler struct {
	config *config.AppSettings
	GetTempoUseCase usecases.GetTempoUseCaseInterface
	ServiceCepInterface services.ServiceCepInterface
	ServiceTempoInterface services.ServiceTempoInterface
}

func NewGetTempoHandler(appConfig *config.AppSettings, getTempoUseCase usecases.GetTempoUseCaseInterface, serviceCep services.ServiceCepInterface, serviceTempo services.ServiceTempoInterface) *GetTempoHandler {
	return &GetTempoHandler{
		config: appConfig,
		GetTempoUseCase: getTempoUseCase,
		ServiceCepInterface: serviceCep,
		ServiceTempoInterface: serviceTempo,
	}
}

// HandleLabsOne godoc
// @Summary Consulta temperatura baseado no CEP
// @Description Consulta a temperatura atual baseada no CEP fornecido
// @Tags Labs-One
// @Accept json
// @Produce json
// @Param cep query string true "CEP"
// @Success 200 {object} entities.GetTempoResponseDto "OK"
// @Failure 404 {object} entities.CustomError "Not Found"
// @Failure 422 {object} entities.CustomError "Invalid Zipcode"
// @Router /get-temp [get]
func (h *GetTempoHandler) HandleLabsOne(w http.ResponseWriter, r *http.Request) {
    cep := r.URL.Query().Get("cep")
    if cep == "" {
        http.Error(w, "cep is required", http.StatusBadRequest)
        return
    }

    response, err := h.GetTempoUseCase.GetTempo(cep)
    if err != nil {
        customErr, ok := err.(*entities.CustomError)
        if ok {
            w.WriteHeader(customErr.Code)
            json.NewEncoder(w).Encode(customErr)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

