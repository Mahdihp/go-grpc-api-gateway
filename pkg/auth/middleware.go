package auth

import (
	"context"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("authorization")
		if authorization == "" {
			//ctx.AbortWithStatus(http.StatusUnauthorized)
			//return
			//return ctx.JSON(http.StatusUnauthorized, "")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		token := strings.Split(authorization, "Bearer ")

		if len(token) < 2 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token[1],
		})

		if err != nil || res.Status != http.StatusOK {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		ctx.Set("userId", res.UserId)
		return next(ctx)
	}
}
