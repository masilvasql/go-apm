package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/masilvasql/go-apm/config"
	"github.com/masilvasql/go-apm/internal/handler"
	"go.elastic.co/apm/module/apmchi/v2"
)

func main() {
	tracer := config.InitApm()
	defer tracer.Flush(nil)

	r := chi.NewRouter()
	r.Use(apmchi.Middleware())

	r.Post("/dividir", handler.DividrHandler)
	r.Post("/consultar-cep", handler.ConsultaCepHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
