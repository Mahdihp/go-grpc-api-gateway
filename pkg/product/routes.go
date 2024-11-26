package product

import (
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(r *echo.Echo, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	masterRoutes := r.Group("/product")
	masterRoutes.Use(a.AuthRequired)
	masterRoutes.POST("/", svc.CreateProduct)
	masterRoutes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx echo.Context) error {
	return routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx echo.Context) error {
	return routes.CreateProduct(ctx, svc.Client)
}
