package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/masilvasql/go-apm/internal/usecase"
)

func ConsultaCepHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var input usecase.ConsultaCepInputDTO

	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Erro ao fazer parse do corpo da requisição", http.StatusBadRequest)
		return
	}

	output, err := usecase.ConsultarCep(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao consultar o CEP", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Erro ao serializar a resposta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
