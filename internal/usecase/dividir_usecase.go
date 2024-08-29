package usecase

import "errors"

type DividirUseCaseInput struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}

type DividirUseCaseOutput struct {
	Resultado float64 `json:"resultado"`
}

func Dividir(input DividirUseCaseInput) (DividirUseCaseOutput, error) {
	if input.Numero1 == 0 || input.Numero2 == 0 {
		return DividirUseCaseOutput{}, errors.New("Não é possível dividir por zero")
	}

	result := float64(input.Numero1) / float64(input.Numero2)

	return DividirUseCaseOutput{Resultado: result}, nil

}
