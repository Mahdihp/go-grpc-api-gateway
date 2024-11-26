package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx echo.Context, c pb.OrderServiceClient) error {
	b := CreateOrderRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	userId := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, &res)
}
