package main

import (
	"gin-mini-mall/internal/dao"
	"gin-mini-mall/internal/handler"
	"gin-mini-mall/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet()

var DaoSet = wire.NewSet(
	dao.NewDao,
	dao.NewUserDao,
)
var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)
var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func initApp(logger *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		DaoSet,
		ServiceSet,
		HandlerSet,
	))
}
