package routes

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx echo.Context, c pb.AuthServiceClient) error {
	b := LoginRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		//ctx.AbortWithError(http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return echo.NewHTTPError(http.StatusBadGateway, "هنوز grpc راه ننداختیم")
	}

	return ctx.JSON(http.StatusCreated, &res)
}
