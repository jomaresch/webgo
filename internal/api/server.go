package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func (s *Server) Run() error {
	router := gin.Default()
	api := router.Group("api")
	api.GET("list", func(context *gin.Context) {
		context.JSON(200, gin.H{"data": []string{"test", "nice"}})
	})
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	return router.Run(":8080")
}
