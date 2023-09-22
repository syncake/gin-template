package main

import (
	"fmt"
	"gin-template/main/app/constant"
	"gin-template/main/bootstrap"
	"gin-template/main/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	bootstrap.Boot()
}

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	err := r.Run(fmt.Sprintf("%v:%v",
		constant.Config.GetString("serve.host"),
		constant.Config.GetString("serve.port"),
	))
	if err != nil {
		panic("serve failed")
	}
}
