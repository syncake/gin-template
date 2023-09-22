package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chain/main/app/constant"
	"go-chain/main/bootstrap"
	"go-chain/main/routes"
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
