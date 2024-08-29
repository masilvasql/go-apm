package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/masilvasql/go-apm/internal/usecase"
	"go.elastic.co/apm/v2"
)

func DividrHandler(h http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	spanReadBody, ctx := apm.StartSpan(ctx, "read_body", "processamento do corpo da requisição")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(h, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	spanReadBody.End()

	defer r.Body.Close()

	var input usecase.DividirUseCaseInput
	spanUnmarshal, ctx := apm.StartSpan(ctx, "unmarshal", "processamento do corpo da requisição")
	err = json.Unmarshal(body, &input)

	spanUnmarshal.Context.SetLabel("input", input)
	spanUnmarshal.Context.SetLabel("numero1", input.Numero1)
	spanUnmarshal.Context.SetLabel("numero2", input.Numero2)

	if err != nil {
		http.Error(h, "Erro ao converter o corpo da requisição", http.StatusBadRequest)
		return
	}
	spanUnmarshal.End()

	spanUsecase, ctx := apm.StartSpan(ctx, "dividir_usecase", "processamento da divisão")
	resultado, err := usecase.Dividir(input)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		http.Error(h, err.Error(), http.StatusInternalServerError)
		spanUsecase.End()
		return
	}
	spanUsecase.End()

	spanResponse, _ := apm.StartSpan(ctx, "response", "processamento da resposta")
	response, err := json.Marshal(resultado)
	if err != nil {
		http.Error(h, "Erro ao converter o resultado", http.StatusBadRequest)
		return
	}
	spanResponse.End()

	h.Header().Set("Content-Type", "application/json")
	h.Write(response)

}
