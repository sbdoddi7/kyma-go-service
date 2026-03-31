package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/kyma-go-service/internal/handler"
)

func RegisterRoutes(r *gin.Engine) {
	h := handler.NewHandler()

	r.GET("/health", h.Health)
	r.GET("/ready", h.Ready)

	v1 := r.Group("v1")
	{
		v1.GET("hello", h.Hello)
	}

}
