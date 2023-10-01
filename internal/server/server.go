package server

import (
	"github.com/gin-gonic/gin"
)

func NewServerHTTP() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	return r
}
