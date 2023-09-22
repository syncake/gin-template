package api

import (
	"gin-template/main/app/constant"
	"gin-template/main/app/controller"
	"github.com/gin-gonic/gin"
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
