package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/pb"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx echo.Context, c pb.ProductServiceClient) error {
	b := CreateProductRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		//ctx.AbortWithError(http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  b.Name,
		Stock: b.Stock,
		Price: b.Price,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, &res)
}
