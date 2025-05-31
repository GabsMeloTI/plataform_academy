package server

import (
	"gorm.io/gorm"
	"plataform_init/config/environment"
	"plataform_init/infra/database"
	"plataform_init/infra/token"
	"plataform_init/internal/user"
)

type ContainerDI struct {
	Config         environment.Config
	Conn           *gorm.DB
	UserRepository *user.Repository
	UserService    *user.Service
	UserHandler    *user.Handler
	JwtMaker       *token.JwtMaker
}

func NewContainerDI(config environment.Config) *ContainerDI {
	container := &ContainerDI{Config: config}

	container.db()
	container.buildRepository()
	container.buildService()
	container.buildHandlers()

	return container
}

func (c *ContainerDI) db() {
	dbConfig := database.Config{
		Host:     c.Config.DBHost,
		Port:     c.Config.DBPort,
		User:     c.Config.DBUser,
		Password: c.Config.DBPassword,
		Database: c.Config.DBDatabase,
		SSLMode:  c.Config.DBSSLMode,
		Driver:   c.Config.DBDriver,
	}

	var err error
	c.Conn, err = database.NewConnection(&dbConfig)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
}

func (c *ContainerDI) buildRepository() {
	c.UserRepository = user.NewUserRepository(c.Conn)
}

func (c *ContainerDI) buildService() {
	c.UserService = user.NewServiceUser(c.UserRepository, c.Config.SignatureString, c.Config.AwsBucketName)
}

func (c *ContainerDI) buildHandlers() {
	c.UserHandler = user.NewHandler(c.UserService)
}
