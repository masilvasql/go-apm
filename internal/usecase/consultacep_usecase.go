package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.elastic.co/apm/v2"
)

type ConsultaCepInputDTO struct {
	Cep string `json:"cep"`
}

type ConsultaCepOutputDTO struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func ConsultarCep(ctx context.Context, input ConsultaCepInputDTO) (ConsultaCepOutputDTO, error) {

	spanUsecase, ctx := apm.StartSpan(ctx, "consultar_cep", "usecase")

	spanConultaCep, ctx := apm.StartSpan(ctx, "via-cep", "consulta-externa")
	res, err := http.Get("https://viacep.com.br/ws/" + input.Cep + "/json/")
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		spanUsecase.End()
		spanConultaCep.End()
		return ConsultaCepOutputDTO{}, err
	}
	defer res.Body.Close()
	time.Sleep(5 * time.Second)
	spanConultaCep.End()

	var output ConsultaCepOutputDTO

	spanDecoder, ctx := apm.StartSpan(ctx, "json-decoder", "decode")
	err = json.NewDecoder(res.Body).Decode(&output)
	spanDecoder.End()

	if err != nil {
		apm.CaptureError(ctx, err).Send()
		spanUsecase.End()
		spanConultaCep.End()
		return ConsultaCepOutputDTO{}, err
	}

	spanUsecase.End()

	return output, nil
}
