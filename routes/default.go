package routes

import (
	"gin-template/main/app/controller/api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", api.Ping)
	r.GET("/fail", api.PingFailed)
}
