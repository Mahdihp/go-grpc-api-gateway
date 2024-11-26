package zipkin

import (
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinHttpReporter "github.com/openzipkin/zipkin-go/reporter/http"
)

func NewTracer(cfg config.Config) (reporter.Reporter, *zipkin.Tracer, *zipkinhttp.Client, error) {
	endpoint, err := zipkin.NewEndpoint("api-getway", "")
	if err != nil {
		fmt.Println("error creating zipkin endpoint: %s", err)
	}
	newReporter := zipkinHttpReporter.NewReporter(cfg.Zipin_Api)
	traceTags := make(map[string]string)
	traceTags["api-getway"] = "api-getway-1"

	tracer, err := zipkin.NewTracer(newReporter, zipkin.WithLocalEndpoint(endpoint), zipkin.WithTags(traceTags))
	client, err := zipkinhttp.NewClient(tracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		fmt.Println("tracing init failed: %s", err)
	}

	return newReporter, tracer, client, err
}
