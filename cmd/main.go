package main

import (
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/adapter/kafka"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/order"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	echoServ := echo.New()
	echoServ.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost, http.MethodGet},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	echoServ.Use(middleware.Logger())
	echoServ.Use(middleware.Recover())

	//reporter, tracer, client, err := zipkin.NewTracer(cfg)
	//if err != nil {
	//	log.Fatalln("Error NewTracer", err)
	//}
	echoServ.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println(string(reqBody))
	}))

	//echoServ.Use(zipkintracing.TraceServer(tracer))

	kafkaWriter := kafka.KafkaProducer(cfg, "test-topic")

	//authSvc := *auth.SetRegisterRoutes(echoServ, &cfg, client, tracer, kafkaWriter)
	authSvc := *auth.SetRegisterRoutes(echoServ, &cfg, kafkaWriter)

	product.RegisterRoutes(echoServ, &cfg, &authSvc)
	order.RegisterRoutes(echoServ, &cfg, &authSvc)

	//defer reporter.Close()
	defer kafkaWriter.Close()

	echoServ.Logger.Fatal(echoServ.Start(cfg.Port))
}
