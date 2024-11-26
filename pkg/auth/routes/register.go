package routes

import (
	"context"
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/data_util"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/labstack/echo/v4"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx echo.Context, c pb.AuthServiceClient, kafkaWriter *kafka.Writer) error {
	b := RegisterRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		//ctx.AbortWithError(http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	//zipkintracing.DoHTTP(ctx, ctx.Request(), zipkinClient)
	//defer zipkintracing.TraceFunc(ctx, "RegisterAccount", zipkintracing.DefaultSpanTags, tracer)()
	//zipkinClient.DoWithAppSpan(ctx.Request(), "DoWithAppSpan")

	//kafka producer...
	message, _ := data_util.SerializeMessage(b)
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("address-%s", ctx.Request().RemoteAddr)),
		Value: message,
	}
	err2 := kafkaWriter.WriteMessages(ctx.Request().Context(), msg)

	if err2 != nil {
		log.Fatalln(err2)
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		//return ctx.JSON(http.StatusBadGateway, "هنوز grpc راه ننداختیم")
		return echo.NewHTTPError(http.StatusBadGateway, "هنوز grpc راه ننداختیم")
		//return ctx.HTML(http.StatusBadGateway, "<strong>Hello, World!</strong>")
	}

	return ctx.JSON(http.StatusCreated, &res)
}
