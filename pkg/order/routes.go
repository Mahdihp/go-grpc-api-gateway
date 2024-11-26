package order

import (
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/order/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(r *echo.Echo, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	orderRoutes := r.Group("/order")
	orderRoutes.Use(a.AuthRequired)
	orderRoutes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx echo.Context) error {
	return routes.CreateOrder(ctx, svc.Client)
}
