package auth

import (
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	kafka "github.com/segmentio/kafka-go"
)

func SetRegisterRoutes(r *echo.Echo, c *config.Config, kafka *kafka.Writer) *ServiceClient {
	svc := &ServiceClient{
		Client:      InitServiceClient(c),
		KafkaWriter: kafka,
	}

	masterRoutes := r.Group("/auth")
	masterRoutes.POST("/register", svc.Register)
	masterRoutes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx echo.Context) error {
	return routes.Register(ctx, svc.Client, svc.KafkaWriter)
}

func (svc *ServiceClient) Login(ctx echo.Context) error {
	return routes.Login(ctx, svc.Client)
}
