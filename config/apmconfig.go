package config

import (
	"fmt"
	"log"
	"net/url"

	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/transport"
)

func InitApm() *apm.Tracer {

	serverURL := "http://localhost:8200"
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		log.Fatalf("Error parsing APM server URL: %s", err)
	}

	tp, err := transport.NewHTTPTransport(transport.HTTPTransportOptions{
		ServerURLs: []*url.URL{parsedURL},
	})
	if err != nil {
		log.Fatalf("Error initializing APM transport: %s", err)
	}

	tracer, err := apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:    "go-apm",
		ServiceVersion: "1.0.0",
		Transport:      tp,
	})

	if err != nil {
		log.Fatalf("Error initializing APM tracer: %s", err)
	}

	// Configurar o tracer como o tracer padrão

	apm.SetDefaultTracer(tracer)

	fmt.Println("APM initialized")

	return tracer
}
