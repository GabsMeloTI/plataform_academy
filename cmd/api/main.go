package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os/signal"
	"plataform_init/config/environment"
	"plataform_init/config/server"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	e := echo.New()

	loadEnv := environment.NewConfig()
	container := server.NewContainerDI(loadEnv)

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := e.Shutdown(ctx); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	user := e.Group("/user")
	user.POST("/create", container.UserHandler.CreateUser)

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
