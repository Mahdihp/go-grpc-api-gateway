package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/pb"
)

func FineOne(ctx echo.Context, c pb.ProductServiceClient) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return err
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return err
	}

	return ctx.JSON(http.StatusCreated, &res)
}
