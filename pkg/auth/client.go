package auth

import (
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	kafka "github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client       pb.AuthServiceClient
	ZipkinClient *zipkinhttp.Client
	Tracer       *zipkin.Tracer
	KafkaWriter  *kafka.Writer
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
