package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os/signal"
	"plataform_init/config/environment"
	"plataform_init/config/server"
	_mid "plataform_init/infra/middleware"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	e := echo.New()
	e.Use(middleware.Recover())

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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	user := e.Group("/user")
	user.POST("/create", container.UserHandler.CreateUser)
	user.POST("/login", container.UserHandler.Login)
	user.PUT("/upload", container.UserHandler.UpdateAvatar, _mid.CheckAuthorization)
	user.GET("/list", container.UserHandler.GetUserById, _mid.CheckAuthorization)
	//TODO: GET TRAZENDO SOMENTE OS USUARIO QUE SÃO ALUNOS
	//TODO: GET TRAZENDO SOMENTE OS USUARIO QUE SÃO PERSONAIS
	user.GET("/list/:role", container.UserHandler.GetUsersByRole)

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
