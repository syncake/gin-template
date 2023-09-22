package routes

import (
	"github.com/gin-gonic/gin"
	"go-chain/main/app/controller/api"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", api.Ping)
	r.GET("/fail", api.PingFailed)
}
