package api

import (
	"github.com/gin-gonic/gin"
	"go-chain/main/app/constant"
	"go-chain/main/app/controller"
)

func Ping(c *gin.Context) {
	constant.Logger.Debugln("ping received")
	controller.Success(c, "pong", &map[string]interface{}{
		"name": "joker",
	})
}

func PingFailed(c *gin.Context) {
	controller.Fail(c, "", 0, nil)
}
